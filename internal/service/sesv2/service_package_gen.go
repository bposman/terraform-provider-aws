// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sesv2

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	sesv2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
		},
		{
			Factory:  DataSourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceConfigurationSet,
			TypeName: "aws_sesv2_configuration_set",
			Name:     "Configuration Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConfigurationSetEventDestination,
			TypeName: "aws_sesv2_configuration_set_event_destination",
		},
		{
			Factory:  ResourceContactList,
			TypeName: "aws_sesv2_contact_list",
			Name:     "Contact List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDedicatedIPAssignment,
			TypeName: "aws_sesv2_dedicated_ip_assignment",
		},
		{
			Factory:  ResourceDedicatedIPPool,
			TypeName: "aws_sesv2_dedicated_ip_pool",
			Name:     "Dedicated IP Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceEmailIdentity,
			TypeName: "aws_sesv2_email_identity",
			Name:     "Email Identity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceEmailIdentityFeedbackAttributes,
			TypeName: "aws_sesv2_email_identity_feedback_attributes",
		},
		{
			Factory:  ResourceEmailIdentityMailFromAttributes,
			TypeName: "aws_sesv2_email_identity_mail_from_attributes",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SESV2
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, cfg aws_sdkv2.Config, endpoint string) *sesv2_sdkv2.Client {
	return sesv2_sdkv2.NewFromConfig(cfg, func(o *sesv2_sdkv2.Options) {
		if endpoint != "" {
			o.EndpointResolver = sesv2_sdkv2.EndpointResolverFromURL(endpoint)
		}
	})
}

var ServicePackage = &servicePackage{}
