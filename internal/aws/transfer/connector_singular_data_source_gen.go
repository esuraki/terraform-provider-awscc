// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_transfer_connector", connectorDataSource)
}

// connectorDataSource returns the Terraform awscc_transfer_connector data source.
// This Terraform data source corresponds to the CloudFormation AWS::Transfer::Connector resource.
func connectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_role": {
			// Property: AccessRole
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the access role for the connector.",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "arn:.*role/.*",
			//	  "type": "string"
			//	}
			Description: "Specifies the access role for the connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			//	  "maxLength": 1600,
			//	  "minLength": 20,
			//	  "pattern": "arn:.*",
			//	  "type": "string"
			//	}
			Description: "Specifies the unique Amazon Resource Name (ARN) for the workflow.",
			Type:        types.StringType,
			Computed:    true,
		},
		"as_2_config": {
			// Property: As2Config
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Configuration for an AS2 connector.",
			//	  "properties": {
			//	    "Compression": {
			//	      "description": "Compression setting for this AS2 connector configuration.",
			//	      "enum": [
			//	        "ZLIB",
			//	        "DISABLED"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "EncryptionAlgorithm": {
			//	      "description": "Encryption algorithm for this AS2 connector configuration.",
			//	      "enum": [
			//	        "AES128_CBC",
			//	        "AES192_CBC",
			//	        "AES256_CBC",
			//	        "NONE"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "LocalProfileId": {
			//	      "description": "A unique identifier for the local profile.",
			//	      "maxLength": 19,
			//	      "minLength": 19,
			//	      "pattern": "^p-([0-9a-f]{17})$",
			//	      "type": "string"
			//	    },
			//	    "MdnResponse": {
			//	      "description": "MDN Response setting for this AS2 connector configuration.",
			//	      "enum": [
			//	        "SYNC",
			//	        "NONE"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "MdnSigningAlgorithm": {
			//	      "description": "MDN Signing algorithm for this AS2 connector configuration.",
			//	      "enum": [
			//	        "SHA256",
			//	        "SHA384",
			//	        "SHA512",
			//	        "SHA1",
			//	        "NONE",
			//	        "DEFAULT"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "MessageSubject": {
			//	      "description": "The message subject for this AS2 connector configuration.",
			//	      "maxLength": 1024,
			//	      "minLength": 1,
			//	      "pattern": "",
			//	      "type": "string"
			//	    },
			//	    "PartnerProfileId": {
			//	      "description": "A unique identifier for the partner profile.",
			//	      "maxLength": 19,
			//	      "minLength": 19,
			//	      "pattern": "^p-([0-9a-f]{17})$",
			//	      "type": "string"
			//	    },
			//	    "SigningAlgorithm": {
			//	      "description": "Signing algorithm for this AS2 connector configuration.",
			//	      "enum": [
			//	        "SHA256",
			//	        "SHA384",
			//	        "SHA512",
			//	        "SHA1",
			//	        "NONE"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Configuration for an AS2 connector.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"compression": {
						// Property: Compression
						Description: "Compression setting for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"encryption_algorithm": {
						// Property: EncryptionAlgorithm
						Description: "Encryption algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"local_profile_id": {
						// Property: LocalProfileId
						Description: "A unique identifier for the local profile.",
						Type:        types.StringType,
						Computed:    true,
					},
					"mdn_response": {
						// Property: MdnResponse
						Description: "MDN Response setting for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"mdn_signing_algorithm": {
						// Property: MdnSigningAlgorithm
						Description: "MDN Signing algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"message_subject": {
						// Property: MessageSubject
						Description: "The message subject for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"partner_profile_id": {
						// Property: PartnerProfileId
						Description: "A unique identifier for the partner profile.",
						Type:        types.StringType,
						Computed:    true,
					},
					"signing_algorithm": {
						// Property: SigningAlgorithm
						Description: "Signing algorithm for this AS2 connector configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"connector_id": {
			// Property: ConnectorId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A unique identifier for the connector.",
			//	  "maxLength": 19,
			//	  "minLength": 19,
			//	  "pattern": "^c-([0-9a-f]{17})$",
			//	  "type": "string"
			//	}
			Description: "A unique identifier for the connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"logging_role": {
			// Property: LoggingRole
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the logging role for the connector.",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "arn:.*role/.*",
			//	  "type": "string"
			//	}
			Description: "Specifies the logging role for the connector.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "Creates a key-value pair for a specific resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The name assigned to the tag that you create.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "Contains one or more values that you assigned to the key name you create.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "Key-value pairs that can be used to group and search for workflows. Tags are metadata attached to workflows for any purpose.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name assigned to the tag that you create.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "Contains one or more values that you assigned to the key name you create.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"url": {
			// Property: Url
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "URL for Connector",
			//	  "maxLength": 255,
			//	  "type": "string"
			//	}
			Description: "URL for Connector",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Transfer::Connector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Connector").WithTerraformTypeName("awscc_transfer_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role":           "AccessRole",
		"arn":                   "Arn",
		"as_2_config":           "As2Config",
		"compression":           "Compression",
		"connector_id":          "ConnectorId",
		"encryption_algorithm":  "EncryptionAlgorithm",
		"key":                   "Key",
		"local_profile_id":      "LocalProfileId",
		"logging_role":          "LoggingRole",
		"mdn_response":          "MdnResponse",
		"mdn_signing_algorithm": "MdnSigningAlgorithm",
		"message_subject":       "MessageSubject",
		"partner_profile_id":    "PartnerProfileId",
		"signing_algorithm":     "SigningAlgorithm",
		"tags":                  "Tags",
		"url":                   "Url",
		"value":                 "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
