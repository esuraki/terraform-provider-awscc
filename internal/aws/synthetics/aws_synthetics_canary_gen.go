// Code generated by generators/resource/main.go; DO NOT EDIT.

package synthetics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_synthetics_canary", awsSyntheticsCanary)
}

// awsSyntheticsCanary returns the Terraform aws_synthetics_canary resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Synthetics::Canary resource type.
func awsSyntheticsCanary(ctx context.Context) (tfsdk.ResourceType, error) {
	// Subproperty definitions.

	// Definition: Code
	// Property: Handler
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr791084952 := schema.Attribute{}
	attr791084952.Type = types.StringType
	attr791084952.Optional = true

	// Definition: Code
	// Property: S3Bucket
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr130402388 := schema.Attribute{}
	attr130402388.Type = types.StringType
	attr130402388.Optional = true

	// Definition: Code
	// Property: S3Key
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr713593033 := schema.Attribute{}
	attr713593033.Type = types.StringType
	attr713593033.Optional = true

	// Definition: Code
	// Property: S3ObjectVersion
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1933333155 := schema.Attribute{}
	attr1933333155.Type = types.StringType
	attr1933333155.Optional = true

	// Definition: Code
	// Property: Script
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1604953967 := schema.Attribute{}
	attr1604953967.Type = types.StringType
	attr1604953967.Optional = true

	// Property references for Code:
	attr2036185364 := make(map[string]schema.Attribute, 5)
	attr2036185364["handler"] = attr791084952
	attr2036185364["s_3_bucket"] = attr130402388
	attr2036185364["s_3_key"] = attr713593033
	attr2036185364["s_3_object_version"] = attr1933333155
	attr2036185364["script"] = attr1604953967

	// Definition: RunConfig
	// Property: ActiveTracing
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Enable active tracing if set to true",
	     "type": "boolean"
	   }
	*/
	attr1943719100 := schema.Attribute{}
	attr1943719100.Type = types.BoolType
	attr1943719100.Optional = true
	attr1943719100.Description = `Enable active tracing if set to true`

	// Definition: RunConfig
	// Property: EnvironmentVariables
	// CloudFormation resource type schema:
	/*
	   {
	     "additionalProperties": false,
	     "description": "Environment variable key-value pairs.",
	     "patternProperties": {
	       "[a-zA-Z][a-zA-Z0-9_]+": {
	         "type": "string"
	       }
	     },
	     "type": "object"
	   }
	*/
	attr3358771378 := schema.Attribute{}
	// Pattern: "[a-zA-Z][a-zA-Z0-9_]+"
	attr3358771378.Type = types.MapType{ElemType: types.StringType}
	attr3358771378.Optional = true
	attr3358771378.Description = `Environment variable key-value pairs.`

	// Definition: RunConfig
	// Property: MemoryInMB
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide maximum memory available for canary in MB",
	     "type": "integer"
	   }
	*/
	attr781052099 := schema.Attribute{}
	attr781052099.Type = types.NumberType
	attr781052099.Optional = true
	attr781052099.Description = `Provide maximum memory available for canary in MB`

	// Definition: RunConfig
	// Property: TimeoutInSeconds
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide maximum canary timeout per run in seconds",
	     "type": "integer"
	   }
	*/
	attr500709785 := schema.Attribute{}
	attr500709785.Type = types.NumberType
	attr500709785.Optional = true
	attr500709785.Description = `Provide maximum canary timeout per run in seconds`

	// Property references for RunConfig:
	attr3375960452 := make(map[string]schema.Attribute, 4)
	attr3375960452["active_tracing"] = attr1943719100
	attr3375960452["environment_variables"] = attr3358771378
	attr3375960452["memory_in_mb"] = attr781052099
	attr3375960452["timeout_in_seconds"] = attr500709785

	// Definition: Schedule
	// Property: DurationInSeconds
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr438806832 := schema.Attribute{}
	attr438806832.Type = types.StringType
	attr438806832.Optional = true

	// Definition: Schedule
	// Property: Expression
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr982960382 := schema.Attribute{}
	attr982960382.Type = types.StringType
	attr982960382.Optional = true

	// Property references for Schedule:
	attr2751052524 := make(map[string]schema.Attribute, 2)
	attr2751052524["duration_in_seconds"] = attr438806832
	attr2751052524["expression"] = attr982960382

	// Definition: Tag
	// Property: Key
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
	     "maxLength": 128,
	     "minLength": 1,
	     "type": "string"
	   }
	*/
	attr2859957898 := schema.Attribute{}
	attr2859957898.Type = types.StringType
	attr2859957898.Optional = true
	attr2859957898.Description = `The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. `

	// Definition: Tag
	// Property: Value
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
	     "maxLength": 256,
	     "minLength": 0,
	     "type": "string"
	   }
	*/
	attr3708964976 := schema.Attribute{}
	attr3708964976.Type = types.StringType
	attr3708964976.Optional = true
	attr3708964976.Description = `The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. `

	// Property references for Tag:
	attr4169356339 := make(map[string]schema.Attribute, 2)
	attr4169356339["key"] = attr2859957898
	attr4169356339["value"] = attr3708964976

	// Definition: VPCConfig
	// Property: SecurityGroupIds
	// CloudFormation resource type schema:
	/*
	   {
	     "items": {
	       "type": "string"
	     },
	     "type": "array"
	   }
	*/
	attr634697343 := schema.Attribute{}
	attr634697343.Type = types.ListType{ElemType: types.StringType}
	attr634697343.Optional = true

	// Definition: VPCConfig
	// Property: SubnetIds
	// CloudFormation resource type schema:
	/*
	   {
	     "items": {
	       "type": "string"
	     },
	     "type": "array"
	   }
	*/
	attr2512506947 := schema.Attribute{}
	attr2512506947.Type = types.ListType{ElemType: types.StringType}
	attr2512506947.Optional = true

	// Definition: VPCConfig
	// Property: VpcId
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr4286326268 := schema.Attribute{}
	attr4286326268.Type = types.StringType
	attr4286326268.Optional = true

	// Property references for VPCConfig:
	attr2512402886 := make(map[string]schema.Attribute, 3)
	attr2512402886["security_group_ids"] = attr634697343
	attr2512402886["subnet_ids"] = attr2512506947
	attr2512402886["vpc_id"] = attr4286326268

	// Root property definition.

	// Definition: (Root)
	// Property: ArtifactS3Location
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide the s3 bucket output location for test results",
	     "pattern": "^(s3|S3)://",
	     "type": "string"
	   }
	*/
	attr3370581841 := schema.Attribute{}
	attr3370581841.Type = types.StringType
	attr3370581841.Required = true
	attr3370581841.Description = `Provide the s3 bucket output location for test results`

	// Definition: (Root)
	// Property: Code
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide the canary script source",
	     "$ref": "#/definitions/Code"
	   }
	*/
	attr3077982955 := schema.Attribute{}
	attr3077982955.Attributes = schema.SingleNestedAttributes(attr2036185364)
	attr3077982955.Required = true
	attr3077982955.Description = `Provide the canary script source`

	// Definition: (Root)
	// Property: ExecutionRoleArn
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Lambda Execution role used to run your canaries",
	     "type": "string"
	   }
	*/
	attr3105035707 := schema.Attribute{}
	attr3105035707.Type = types.StringType
	attr3105035707.Required = true
	attr3105035707.Description = `Lambda Execution role used to run your canaries`

	// Definition: (Root)
	// Property: FailureRetentionPeriod
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Retention period of failed canary runs represented in number of days",
	     "type": "integer"
	   }
	*/
	attr2772195751 := schema.Attribute{}
	attr2772195751.Type = types.NumberType
	attr2772195751.Optional = true
	attr2772195751.Description = `Retention period of failed canary runs represented in number of days`

	// Definition: (Root)
	// Property: Id
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Id of the canary",
	     "type": "string"
	   }
	*/
	attr2710565547 := schema.Attribute{}
	attr2710565547.Type = types.StringType
	attr2710565547.Computed = true
	attr2710565547.Description = `Id of the canary`

	// Definition: (Root)
	// Property: Name
	// PrimaryIdentifier: true
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Name of the canary.",
	     "pattern": "^[0-9a-z_\\-]{1,21}$",
	     "type": "string"
	   }
	*/
	attr1435824245 := schema.Attribute{}
	attr1435824245.Type = types.StringType
	attr1435824245.Required = true
	attr1435824245.Description = `Name of the canary.`

	// Definition: (Root)
	// Property: RunConfig
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide canary run configuration",
	     "$ref": "#/definitions/RunConfig"
	   }
	*/
	attr2654279805 := schema.Attribute{}
	attr2654279805.Attributes = schema.SingleNestedAttributes(attr3375960452)
	attr2654279805.Optional = true
	attr2654279805.Description = `Provide canary run configuration`

	// Definition: (Root)
	// Property: RuntimeVersion
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Runtime version of Synthetics Library",
	     "type": "string"
	   }
	*/
	attr2015302114 := schema.Attribute{}
	attr2015302114.Type = types.StringType
	attr2015302114.Required = true
	attr2015302114.Description = `Runtime version of Synthetics Library`

	// Definition: (Root)
	// Property: Schedule
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Frequency to run your canaries",
	     "$ref": "#/definitions/Schedule"
	   }
	*/
	attr4051156835 := schema.Attribute{}
	attr4051156835.Attributes = schema.SingleNestedAttributes(attr2751052524)
	attr4051156835.Required = true
	attr4051156835.Description = `Frequency to run your canaries`

	// Definition: (Root)
	// Property: StartCanaryAfterCreation
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Runs canary if set to True. Default is False",
	     "type": "boolean"
	   }
	*/
	attr4092632865 := schema.Attribute{}
	attr4092632865.Type = types.BoolType
	attr4092632865.Required = true
	attr4092632865.Description = `Runs canary if set to True. Default is False`

	// Definition: (Root)
	// Property: State
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "State of the canary",
	     "type": "string"
	   }
	*/
	attr902619179 := schema.Attribute{}
	attr902619179.Type = types.StringType
	attr902619179.Computed = true
	attr902619179.Description = `State of the canary`

	// Definition: (Root)
	// Property: SuccessRetentionPeriod
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Retention period of successful canary runs represented in number of days",
	     "type": "integer"
	   }
	*/
	attr3841818316 := schema.Attribute{}
	attr3841818316.Type = types.NumberType
	attr3841818316.Optional = true
	attr3841818316.Description = `Retention period of successful canary runs represented in number of days`

	// Definition: (Root)
	// Property: Tags
	// CloudFormation resource type schema:
	/*
	   {
	     "items": {
	       "$ref": "#/definitions/Tag"
	     },
	     "type": "array",
	     "uniqueItems": false
	   }
	*/
	attr2151425999 := schema.Attribute{}
	attr2151425999Options := schema.ListNestedAttributesOptions{}
	attr2151425999.Attributes = schema.ListNestedAttributes(attr4169356339, attr2151425999Options)
	attr2151425999.Optional = true

	// Definition: (Root)
	// Property: VPCConfig
	// CloudFormation resource type schema:
	/*
	   {
	     "description": "Provide VPC Configuration if enabled.",
	     "$ref": "#/definitions/VPCConfig"
	   }
	*/
	attr259913979 := schema.Attribute{}
	attr259913979.Attributes = schema.SingleNestedAttributes(attr2512402886)
	attr259913979.Optional = true
	attr259913979.Description = `Provide VPC Configuration if enabled.`

	// Property references for (Root):
	attr1863345718 := make(map[string]schema.Attribute, 14)
	attr1863345718["artifact_s_3_location"] = attr3370581841
	attr1863345718["code"] = attr3077982955
	attr1863345718["execution_role_arn"] = attr3105035707
	attr1863345718["failure_retention_period"] = attr2772195751
	attr1863345718["id"] = attr2710565547
	attr1863345718["name"] = attr1435824245
	attr1863345718["run_config"] = attr2654279805
	attr1863345718["runtime_version"] = attr2015302114
	attr1863345718["schedule"] = attr4051156835
	attr1863345718["start_canary_after_creation"] = attr4092632865
	attr1863345718["state"] = attr902619179
	attr1863345718["success_retention_period"] = attr3841818316
	attr1863345718["tags"] = attr2151425999
	attr1863345718["vpc_config"] = attr259913979

	// Resource instance unique identifier.
	attr1863345718["identifier"] = schema.Attribute{
		Type:        types.StringType,
		Computed:    true,
		Description: "The resource instance's unique identifier.",
	}

	schema := schema.Schema{
		Version:    1,
		Attributes: attr1863345718,
	}

	resourceType := generic.NewResourceType(
		"AWS::Synthetics::Canary",
		"aws_synthetics_canary",
		schema,
	)

	tflog.Debug(ctx, "Generated schema for %s:\n\n%v", "aws_synthetics_canary", schema)

	return resourceType, nil
}