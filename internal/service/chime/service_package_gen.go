// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package chime

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	chime_sdkv1 "github.com/aws/aws-sdk-go/service/chime"
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
			Factory:  ResourceVoiceConnector,
			TypeName: "aws_chime_voice_connector",
		},
		{
			Factory:  ResourceVoiceConnectorGroup,
			TypeName: "aws_chime_voice_connector_group",
		},
		{
			Factory:  ResourceVoiceConnectorLogging,
			TypeName: "aws_chime_voice_connector_logging",
		},
		{
			Factory:  ResourceVoiceConnectorOrigination,
			TypeName: "aws_chime_voice_connector_origination",
		},
		{
			Factory:  ResourceVoiceConnectorStreaming,
			TypeName: "aws_chime_voice_connector_streaming",
		},
		{
			Factory:  ResourceVoiceConnectorTermination,
			TypeName: "aws_chime_voice_connector_termination",
		},
		{
			Factory:  ResourceVoiceConnectorTerminationCredentials,
			TypeName: "aws_chime_voice_connector_termination_credentials",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Chime
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, sess *session_sdkv1.Session, endpoint string) *chime_sdkv1.Chime {
	return chime_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(endpoint)}))
}

var ServicePackage = &servicePackage{}
