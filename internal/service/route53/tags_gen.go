// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package route53

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// ListTags lists route53 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn route53iface.Route53API, identifier, resourceType string) (tftags.KeyValueTags, error) {
	input := &route53.ListTagsForResourceInput{
		ResourceId:   aws.String(identifier),
		ResourceType: aws.String(resourceType),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.ResourceTagSet.Tags), nil
}

// ListTags lists route53 service tags and set them in Context.
// It is called from outside this package.
func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier, resourceType string) error {
	tags, err := ListTags(ctx, meta.(*conns.AWSClient).Route53Conn(), identifier, resourceType)

	if err != nil {
		return err
	}

	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(tags)
	}

	return nil
}

// []*SERVICE.Tag handling

// Tags returns route53 service tags.
func Tags(tags tftags.KeyValueTags) []*route53.Tag {
	result := make([]*route53.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &route53.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from route53 service tags.
func KeyValueTags(ctx context.Context, tags []*route53.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// GetTagsIn returns route53 service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) []*route53.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets route53 service tags in Context.
func SetTagsOut(ctx context.Context, tags []*route53.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// createTags creates route53 service tags for new resources.
func createTags(ctx context.Context, conn route53iface.Route53API, identifier, resourceType string, tags []*route53.Tag) error {
	if len(tags) == 0 {
		return nil
	}

	return UpdateTags(ctx, conn, identifier, resourceType, nil, KeyValueTags(ctx, tags))
}

// UpdateTags updates route53 service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(ctx context.Context, conn route53iface.Route53API, identifier, resourceType string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)
	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.Route53)
	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.Route53)

	// Ensure we do not send empty requests.
	if len(removedTags) == 0 && len(updatedTags) == 0 {
		return nil
	}

	input := &route53.ChangeTagsForResourceInput{
		ResourceId:   aws.String(identifier),
		ResourceType: aws.String(resourceType),
	}

	if len(updatedTags) > 0 {
		input.AddTags = Tags(updatedTags)
	}

	if len(removedTags) > 0 {
		input.RemoveTagKeys = aws.StringSlice(removedTags.Keys())
	}

	_, err := conn.ChangeTagsForResourceWithContext(ctx, input)

	if err != nil {
		return fmt.Errorf("tagging resource (%s): %w", identifier, err)
	}

	return nil
}

// UpdateTags updates route53 service tags.
// It is called from outside this package.
func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier, resourceType string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).Route53Conn(), identifier, resourceType, oldTags, newTags)
}
