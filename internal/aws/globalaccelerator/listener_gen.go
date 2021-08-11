// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_globalaccelerator_listener", listenerResourceType)
}

// listenerResourceType returns the Terraform aws_globalaccelerator_listener resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GlobalAccelerator::Listener resource type.
func listenerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"accelerator_arn": {
			// Property: AcceleratorArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the accelerator.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Type:        types.StringType,
			Required:    true,
			// AcceleratorArn is a force-new attribute.
		},
		"client_affinity": {
			// Property: ClientAffinity
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
			     "enum": [
			       "NONE",
			       "SOURCE_IP"
			     ],
			     "type": "string"
			   }
			*/
			Description: "Client affinity lets you direct all requests from a user to the same endpoint.",
			Type:        types.StringType,
			Optional:    true,
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Amazon Resource Name (ARN) of the listener.",
			     "type": "string"
			   }
			*/
			Description: "The Amazon Resource Name (ARN) of the listener.",
			Type:        types.StringType,
			Computed:    true,
		},
		"port_ranges": {
			// Property: PortRanges
			// CloudFormation resource type schema:
			/*
			   {
			     "items": {
			       "description": "A port range to support for connections from  clients to your accelerator.",
			       "properties": {
			         "FromPort": {
			           "description": "A network port number",
			           "$ref": "#/definitions/Port",
			           "type": "integer"
			         },
			         "ToPort": {
			           "description": "A network port number",
			           "$ref": "#/definitions/Port",
			           "type": "integer"
			         }
			       },
			       "$ref": "#/definitions/PortRange",
			       "required": [
			         "FromPort",
			         "ToPort"
			       ],
			       "type": "object"
			     },
			     "type": "array"
			   }
			*/
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"from_port": {
						// Property: FromPort
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A network port number",
						     "$ref": "#/definitions/Port",
						     "type": "integer"
						   }
						*/
						Description: "A network port number",
						Type:        types.NumberType,
						Required:    true,
					},
					"to_port": {
						// Property: ToPort
						// CloudFormation resource type schema:
						/*
						   {
						     "description": "A network port number",
						     "$ref": "#/definitions/Port",
						     "type": "integer"
						   }
						*/
						Description: "A network port number",
						Type:        types.NumberType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The protocol for the listener.",
			     "enum": [
			       "TCP",
			       "UDP"
			     ],
			     "type": "string"
			   }
			*/
			Description: "The protocol for the listener.",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::GlobalAccelerator::Listener",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Listener").WithTerraformTypeName("aws_globalaccelerator_listener").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_globalaccelerator_listener", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}