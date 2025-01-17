// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lex_resource_policy", resourcePolicyResource)
}

// resourcePolicyResource returns the Terraform awscc_lex_resource_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lex::ResourcePolicy resource.
func resourcePolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Physical ID of the resource policy.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Physical ID of the resource policy.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Policy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
		//	  "type": "object"
		//	}
		"policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1011),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: RevisionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9]+$",
		//	  "type": "string"
		//	}
		"revision_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "A resource policy with specified policy statements that attaches to a Lex bot or bot alias.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::ResourcePolicy").WithTerraformTypeName("awscc_lex_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":           "Id",
		"policy":       "Policy",
		"resource_arn": "ResourceArn",
		"revision_id":  "RevisionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
