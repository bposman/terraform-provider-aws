// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package shield

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	shield_sdkv1 "github.com/aws/aws-sdk-go/service/shield"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceProtection,
			TypeName: "aws_shield_protection",
			Name:     "Protection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceProtectionGroup,
			TypeName: "aws_shield_protection_group",
			Name:     "Protection Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "protection_group_arn",
			},
		},
		{
			Factory:  ResourceProtectionHealthCheckAssociation,
			TypeName: "aws_shield_protection_health_check_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Shield
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *shield_sdkv1.Shield {
	return shield_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
