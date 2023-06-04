// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package directconnect

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	directconnect_sdkv1 "github.com/aws/aws-sdk-go/service/directconnect"
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
			Factory:  DataSourceConnection,
			TypeName: "aws_dx_connection",
		},
		{
			Factory:  DataSourceGateway,
			TypeName: "aws_dx_gateway",
		},
		{
			Factory:  DataSourceLocation,
			TypeName: "aws_dx_location",
		},
		{
			Factory:  DataSourceLocations,
			TypeName: "aws_dx_locations",
		},
		{
			Factory:  DataSourceRouterConfiguration,
			TypeName: "aws_dx_router_configuration",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceBGPPeer,
			TypeName: "aws_dx_bgp_peer",
		},
		{
			Factory:  ResourceConnection,
			TypeName: "aws_dx_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConnectionAssociation,
			TypeName: "aws_dx_connection_association",
		},
		{
			Factory:  ResourceConnectionConfirmation,
			TypeName: "aws_dx_connection_confirmation",
		},
		{
			Factory:  ResourceGateway,
			TypeName: "aws_dx_gateway",
		},
		{
			Factory:  ResourceGatewayAssociation,
			TypeName: "aws_dx_gateway_association",
		},
		{
			Factory:  ResourceGatewayAssociationProposal,
			TypeName: "aws_dx_gateway_association_proposal",
		},
		{
			Factory:  ResourceHostedConnection,
			TypeName: "aws_dx_hosted_connection",
		},
		{
			Factory:  ResourceHostedPrivateVirtualInterface,
			TypeName: "aws_dx_hosted_private_virtual_interface",
		},
		{
			Factory:  ResourceHostedPrivateVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_private_virtual_interface_accepter",
			Name:     "Hosted Private Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceHostedPublicVirtualInterface,
			TypeName: "aws_dx_hosted_public_virtual_interface",
		},
		{
			Factory:  ResourceHostedPublicVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_public_virtual_interface_accepter",
			Name:     "Hosted Public Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceHostedTransitVirtualInterface,
			TypeName: "aws_dx_hosted_transit_virtual_interface",
		},
		{
			Factory:  ResourceHostedTransitVirtualInterfaceAccepter,
			TypeName: "aws_dx_hosted_transit_virtual_interface_accepter",
			Name:     "Hosted Transit Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceLag,
			TypeName: "aws_dx_lag",
			Name:     "LAG",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceMacSecKeyAssociation,
			TypeName: "aws_dx_macsec_key_association",
		},
		{
			Factory:  ResourcePrivateVirtualInterface,
			TypeName: "aws_dx_private_virtual_interface",
			Name:     "Private Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourcePublicVirtualInterface,
			TypeName: "aws_dx_public_virtual_interface",
			Name:     "Public Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTransitVirtualInterface,
			TypeName: "aws_dx_transit_virtual_interface",
			Name:     "Transit Virtual Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.DirectConnect
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *directconnect_sdkv1.DirectConnect {
	return directconnect_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
