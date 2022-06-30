// Code generated by generators/resource/main.go; DO NOT EDIT.

package emrserverless

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_emrserverless_application", applicationResourceType)
}

// applicationResourceType returns the Terraform awscc_emrserverless_application resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EMRServerless::Application resource type.
func applicationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_id": {
			// Property: ApplicationId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the EMR Serverless Application.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The ID of the EMR Serverless Application.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the EMR Serverless Application.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the EMR Serverless Application.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"auto_start_configuration": {
			// Property: AutoStartConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration for Auto Start of Application.",
			//   "properties": {
			//     "Enabled": {
			//       "default": true,
			//       "description": "If set to true, the Application will automatically start. Defaults to true.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configuration for Auto Start of Application.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enabled": {
						// Property: Enabled
						Description: "If set to true, the Application will automatically start. Defaults to true.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Bool{Value: true}),
							tfsdk.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
		},
		"auto_stop_configuration": {
			// Property: AutoStopConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration for Auto Stop of Application.",
			//   "properties": {
			//     "Enabled": {
			//       "default": true,
			//       "description": "If set to true, the Application will automatically stop after being idle. Defaults to true.",
			//       "type": "boolean"
			//     },
			//     "IdleTimeoutMinutes": {
			//       "description": "The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Configuration for Auto Stop of Application.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"enabled": {
						// Property: Enabled
						Description: "If set to true, the Application will automatically stop after being idle. Defaults to true.",
						Type:        types.BoolType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Bool{Value: true}),
							tfsdk.UseStateForUnknown(),
						},
					},
					"idle_timeout_minutes": {
						// Property: IdleTimeoutMinutes
						Description: "The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.",
						Type:        types.Int64Type,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"initial_capacity": {
			// Property: InitialCapacity
			// CloudFormation resource type schema:
			// {
			//   "description": "Initial capacity initialized when an Application is started.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "Worker type for an analytics framework.",
			//         "maxLength": 50,
			//         "minLength": 1,
			//         "pattern": "^[a-zA-Z]+[-_]*[a-zA-Z]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "WorkerConfiguration": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Cpu": {
			//                 "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
			//                 "maxLength": 15,
			//                 "minLength": 1,
			//                 "pattern": "^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$",
			//                 "type": "string"
			//               },
			//               "Disk": {
			//                 "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
			//                 "maxLength": 15,
			//                 "minLength": 1,
			//                 "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$",
			//                 "type": "string"
			//               },
			//               "Memory": {
			//                 "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
			//                 "maxLength": 15,
			//                 "minLength": 1,
			//                 "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "Cpu",
			//               "Memory"
			//             ],
			//             "type": "object"
			//           },
			//           "WorkerCount": {
			//             "description": "Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped",
			//             "format": "int64",
			//             "maximum": 1000000,
			//             "minimum": 1,
			//             "type": "integer"
			//           }
			//         },
			//         "required": [
			//           "WorkerCount",
			//           "WorkerConfiguration"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Initial capacity initialized when an Application is started.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "Worker type for an analytics framework.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 50),
							validate.StringMatch(regexp.MustCompile("^[a-zA-Z]+[-_]*[a-zA-Z]+$"), ""),
						},
					},
					"value": {
						// Property: Value
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"worker_configuration": {
									// Property: WorkerConfiguration
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cpu": {
												// Property: Cpu
												Description: "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 15),
													validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$"), ""),
												},
											},
											"disk": {
												// Property: Disk
												Description: "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
												Type:        types.StringType,
												Optional:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 15),
													validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"), ""),
												},
											},
											"memory": {
												// Property: Memory
												Description: "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 15),
													validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$"), ""),
												},
											},
										},
									),
									Required: true,
								},
								"worker_count": {
									// Property: WorkerCount
									Description: "Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped",
									Type:        types.Int64Type,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(1, 1000000),
									},
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"maximum_capacity": {
			// Property: MaximumCapacity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.",
			//   "properties": {
			//     "Cpu": {
			//       "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
			//       "maxLength": 15,
			//       "minLength": 1,
			//       "pattern": "^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$",
			//       "type": "string"
			//     },
			//     "Disk": {
			//       "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
			//       "maxLength": 15,
			//       "minLength": 1,
			//       "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$",
			//       "type": "string"
			//     },
			//     "Memory": {
			//       "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
			//       "maxLength": 15,
			//       "minLength": 1,
			//       "pattern": "^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Cpu",
			//     "Memory"
			//   ],
			//   "type": "object"
			// }
			Description: "Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"cpu": {
						// Property: Cpu
						Description: "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 15),
							validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(vCPU|vcpu|VCPU)?$"), ""),
						},
					},
					"disk": {
						// Property: Disk
						Description: "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 15),
							validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)$"), ""),
						},
					},
					"memory": {
						// Property: Memory
						Description: "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 15),
							validate.StringMatch(regexp.MustCompile("^[1-9][0-9]*(\\s)?(GB|gb|gB|Gb)?$"), ""),
						},
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "User friendly Application name.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9._\\/#-]+$",
			//   "type": "string"
			// }
			Description: "User friendly Application name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9._\\/#-]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"network_configuration": {
			// Property: NetworkConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Network Configuration for customer VPC connectivity.",
			//   "properties": {
			//     "SecurityGroupIds": {
			//       "description": "The ID of the security groups in the VPC to which you want to connect your job or application.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "Identifier of a security group",
			//         "maxLength": 32,
			//         "minLength": 1,
			//         "pattern": "^[-0-9a-zA-Z]+",
			//         "type": "string"
			//       },
			//       "maxItems": 5,
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "SubnetIds": {
			//       "description": "The ID of the subnets in the VPC to which you want to connect your job or application.",
			//       "insertionOrder": false,
			//       "items": {
			//         "description": "Identifier of a subnet",
			//         "maxLength": 32,
			//         "minLength": 1,
			//         "pattern": "^[-0-9a-zA-Z]+",
			//         "type": "string"
			//       },
			//       "maxItems": 16,
			//       "minItems": 1,
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Network Configuration for customer VPC connectivity.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"security_group_ids": {
						// Property: SecurityGroupIds
						Description: "The ID of the security groups in the VPC to which you want to connect your job or application.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 5),
							validate.ArrayForEach(validate.StringLenBetween(1, 32)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^[-0-9a-zA-Z]+"), "")),
						},
					},
					"subnet_ids": {
						// Property: SubnetIds
						Description: "The ID of the subnets in the VPC to which you want to connect your job or application.",
						Type:        types.SetType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 16),
							validate.ArrayForEach(validate.StringLenBetween(1, 32)),
							validate.ArrayForEach(validate.StringMatch(regexp.MustCompile("^[-0-9a-zA-Z]+"), "")),
						},
					},
				},
			),
			Optional: true,
		},
		"release_label": {
			// Property: ReleaseLabel
			// CloudFormation resource type schema:
			// {
			//   "description": "EMR release label.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "^[A-Za-z0-9._/-]+$",
			//   "type": "string"
			// }
			Description: "EMR release label.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9._/-]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Tag map with key and value",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "^[A-Za-z0-9 /_.:=+@-]+$",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "pattern": "^[A-Za-z0-9 /_.:=+@-]*$",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Tag map with key and value",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The value for the tag. You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
							validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9 /_.:=+@-]+$"), ""),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
							validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9 /_.:=+@-]*$"), ""),
						},
					},
				},
			),
			Optional: true,
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of the application",
			//   "type": "string"
			// }
			Description: "The type of the application",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::EMRServerless::Application Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EMRServerless::Application").WithTerraformTypeName("awscc_emrserverless_application")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":           "ApplicationId",
		"arn":                      "Arn",
		"auto_start_configuration": "AutoStartConfiguration",
		"auto_stop_configuration":  "AutoStopConfiguration",
		"cpu":                      "Cpu",
		"disk":                     "Disk",
		"enabled":                  "Enabled",
		"idle_timeout_minutes":     "IdleTimeoutMinutes",
		"initial_capacity":         "InitialCapacity",
		"key":                      "Key",
		"maximum_capacity":         "MaximumCapacity",
		"memory":                   "Memory",
		"name":                     "Name",
		"network_configuration":    "NetworkConfiguration",
		"release_label":            "ReleaseLabel",
		"security_group_ids":       "SecurityGroupIds",
		"subnet_ids":               "SubnetIds",
		"tags":                     "Tags",
		"type":                     "Type",
		"value":                    "Value",
		"worker_configuration":     "WorkerConfiguration",
		"worker_count":             "WorkerCount",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}