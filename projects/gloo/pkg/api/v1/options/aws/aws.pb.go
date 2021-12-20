// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/aws/aws.proto

package aws

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DestinationSpec_InvocationStyle int32

const (
	DestinationSpec_SYNC  DestinationSpec_InvocationStyle = 0
	DestinationSpec_ASYNC DestinationSpec_InvocationStyle = 1
)

// Enum value maps for DestinationSpec_InvocationStyle.
var (
	DestinationSpec_InvocationStyle_name = map[int32]string{
		0: "SYNC",
		1: "ASYNC",
	}
	DestinationSpec_InvocationStyle_value = map[string]int32{
		"SYNC":  0,
		"ASYNC": 1,
	}
)

func (x DestinationSpec_InvocationStyle) Enum() *DestinationSpec_InvocationStyle {
	p := new(DestinationSpec_InvocationStyle)
	*p = x
	return p
}

func (x DestinationSpec_InvocationStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DestinationSpec_InvocationStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_enumTypes[0].Descriptor()
}

func (DestinationSpec_InvocationStyle) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_enumTypes[0]
}

func (x DestinationSpec_InvocationStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DestinationSpec_InvocationStyle.Descriptor instead.
func (DestinationSpec_InvocationStyle) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescGZIP(), []int{2, 0}
}

// Upstream Spec for AWS Lambda Upstreams
// AWS Upstreams represent a collection of Lambda Functions for a particular AWS Account (IAM Role or User account)
// in a particular region
type UpstreamSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The AWS Region where the desired Lambda Functions exist
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// A [Gloo Secret Ref](https://gloo.solo.io/introduction/concepts/#Secrets) to an AWS Secret
	// AWS Secrets can be created with `glooctl secret create aws ...`
	// If the secret is created manually, it must conform to the following structure:
	//  ```
	//  access_key: <aws access key>
	//  secret_key: <aws secret key>
	//  session_token: <(optional) aws session token>
	//  ```
	SecretRef *core.ResourceRef `protobuf:"bytes,2,opt,name=secret_ref,json=secretRef,proto3" json:"secret_ref,omitempty"`
	// The list of Lambda Functions contained within this region.
	// This list will be automatically populated by Gloo if discovery is enabled for AWS Lambda Functions
	LambdaFunctions []*LambdaFunctionSpec `protobuf:"bytes,3,rep,name=lambda_functions,json=lambdaFunctions,proto3" json:"lambda_functions,omitempty"`
	// (Optional): role_arn to use when assuming a role for a given request via STS.
	// If set this role_arn will override the value found in AWS_ROLE_ARN
	// This option will only be respected if STS credentials are enabled.
	// To enable STS credential fetching see Settings.Gloo.AwsOptions in settings.proto.
	RoleArn string `protobuf:"bytes,4,opt,name=role_arn,json=roleArn,proto3" json:"role_arn,omitempty"`
}

func (x *UpstreamSpec) Reset() {
	*x = UpstreamSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamSpec) ProtoMessage() {}

func (x *UpstreamSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamSpec.ProtoReflect.Descriptor instead.
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *UpstreamSpec) GetSecretRef() *core.ResourceRef {
	if x != nil {
		return x.SecretRef
	}
	return nil
}

func (x *UpstreamSpec) GetLambdaFunctions() []*LambdaFunctionSpec {
	if x != nil {
		return x.LambdaFunctions
	}
	return nil
}

func (x *UpstreamSpec) GetRoleArn() string {
	if x != nil {
		return x.RoleArn
	}
	return ""
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions:
// - name of the function
// - qualifier for the function
type LambdaFunctionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the logical name gloo should associate with this function. if left empty, it will default to
	// lambda_function_name+qualifier
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// The Name of the Lambda Function as it appears in the AWS Lambda Portal
	LambdaFunctionName string `protobuf:"bytes,2,opt,name=lambda_function_name,json=lambdaFunctionName,proto3" json:"lambda_function_name,omitempty"`
	// The Qualifier for the Lambda Function. Qualifiers act as a kind of version
	// for Lambda Functions. See https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html for more info.
	Qualifier string `protobuf:"bytes,3,opt,name=qualifier,proto3" json:"qualifier,omitempty"`
}

func (x *LambdaFunctionSpec) Reset() {
	*x = LambdaFunctionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LambdaFunctionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LambdaFunctionSpec) ProtoMessage() {}

func (x *LambdaFunctionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LambdaFunctionSpec.ProtoReflect.Descriptor instead.
func (*LambdaFunctionSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescGZIP(), []int{1}
}

func (x *LambdaFunctionSpec) GetLogicalName() string {
	if x != nil {
		return x.LogicalName
	}
	return ""
}

func (x *LambdaFunctionSpec) GetLambdaFunctionName() string {
	if x != nil {
		return x.LambdaFunctionName
	}
	return ""
}

func (x *LambdaFunctionSpec) GetQualifier() string {
	if x != nil {
		return x.Qualifier
	}
	return ""
}

// Each Lambda Function Spec contains data necessary for Gloo to invoke Lambda functions
type DestinationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Logical Name of the LambdaFunctionSpec to be invoked.
	LogicalName string `protobuf:"bytes,1,opt,name=logical_name,json=logicalName,proto3" json:"logical_name,omitempty"`
	// Can be either Sync or Async.
	InvocationStyle DestinationSpec_InvocationStyle `protobuf:"varint,2,opt,name=invocation_style,json=invocationStyle,proto3,enum=aws.options.gloo.solo.io.DestinationSpec_InvocationStyle" json:"invocation_style,omitempty"`
	// de-jsonify response bodies returned from aws lambda
	ResponseTransformation bool `protobuf:"varint,5,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
}

func (x *DestinationSpec) Reset() {
	*x = DestinationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationSpec) ProtoMessage() {}

func (x *DestinationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationSpec.ProtoReflect.Descriptor instead.
func (*DestinationSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescGZIP(), []int{2}
}

func (x *DestinationSpec) GetLogicalName() string {
	if x != nil {
		return x.LogicalName
	}
	return ""
}

func (x *DestinationSpec) GetInvocationStyle() DestinationSpec_InvocationStyle {
	if x != nil {
		return x.InvocationStyle
	}
	return DestinationSpec_SYNC
}

func (x *DestinationSpec) GetResponseTransformation() bool {
	if x != nil {
		return x.ResponseTransformation
	}
	return false
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x77, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12,
	0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x6b, 0x69, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x12, 0x57, 0x0a, 0x10, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f, 0x6c, 0x61, 0x6d,
	0x62, 0x64, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x4c, 0x61, 0x6d, 0x62,
	0x64, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x39, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x49, 0x6e,
	0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x0f, 0x69,
	0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x37,
	0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x59,
	0x4e, 0x43, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x01, 0x42,
	0x4a, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x77, 0x73, 0xc0,
	0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xd0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_goTypes = []interface{}{
	(DestinationSpec_InvocationStyle)(0), // 0: aws.options.gloo.solo.io.DestinationSpec.InvocationStyle
	(*UpstreamSpec)(nil),                 // 1: aws.options.gloo.solo.io.UpstreamSpec
	(*LambdaFunctionSpec)(nil),           // 2: aws.options.gloo.solo.io.LambdaFunctionSpec
	(*DestinationSpec)(nil),              // 3: aws.options.gloo.solo.io.DestinationSpec
	(*core.ResourceRef)(nil),             // 4: core.solo.io.ResourceRef
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_depIdxs = []int32{
	4, // 0: aws.options.gloo.solo.io.UpstreamSpec.secret_ref:type_name -> core.solo.io.ResourceRef
	2, // 1: aws.options.gloo.solo.io.UpstreamSpec.lambda_functions:type_name -> aws.options.gloo.solo.io.LambdaFunctionSpec
	0, // 2: aws.options.gloo.solo.io.DestinationSpec.invocation_style:type_name -> aws.options.gloo.solo.io.DestinationSpec.InvocationStyle
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LambdaFunctionSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_options_aws_aws_proto_depIdxs = nil
}
