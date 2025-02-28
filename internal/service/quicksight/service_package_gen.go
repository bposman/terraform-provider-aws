// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceFolderMembership,
			Name:    "Folder Membership",
		},
		{
			Factory: newResourceIAMPolicyAssignment,
			Name:    "IAM Policy Assignment",
		},
		{
			Factory: newResourceIngestion,
			Name:    "Ingestion",
		},
		{
			Factory: newResourceNamespace,
			Name:    "Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
		},
		{
			Factory:  DataSourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  DataSourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccountSubscription,
			TypeName: "aws_quicksight_account_subscription",
			Name:     "Account Subscription",
		},
		{
			Factory:  ResourceDataSet,
			TypeName: "aws_quicksight_data_set",
			Name:     "Data Set",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_quicksight_data_source",
			Name:     "Data Source",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceFolder,
			TypeName: "aws_quicksight_folder",
			Name:     "Folder",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceGroup,
			TypeName: "aws_quicksight_group",
			Name:     "Group",
		},
		{
			Factory:  ResourceGroupMembership,
			TypeName: "aws_quicksight_group_membership",
			Name:     "Group Membership",
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_quicksight_user",
			Name:     "User",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.QuickSight
}

var ServicePackage = &servicePackage{}
