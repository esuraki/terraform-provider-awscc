// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_view_version", viewVersionDataSource)
}

// viewVersionDataSource returns the Terraform awscc_connect_view_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::ViewVersion resource.
func viewVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the view.",
		//	  "type": "integer"
		//	}
		"version": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The version of the view.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description for the view version.",
		//	  "maxLength": 4096,
		//	  "minLength": 1,
		//	  "pattern": "^([\\p{L}\\p{N}_.:\\/=+\\-@,]+[\\p{L}\\p{Z}\\p{N}_.:\\/=+\\-@,]*)$",
		//	  "type": "string"
		//	}
		"version_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description for the view version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the view for which a version is being created.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/view/[-:a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"view_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the view for which a version is being created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewContentSha256
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The view content hash to be checked.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9]{64}$",
		//	  "type": "string"
		//	}
		"view_content_sha_256": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The view content hash to be checked.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ViewVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the created view version.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/view/[-:a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"view_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the created view version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::ViewVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::ViewVersion").WithTerraformTypeName("awscc_connect_view_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"version":              "Version",
		"version_description":  "VersionDescription",
		"view_arn":             "ViewArn",
		"view_content_sha_256": "ViewContentSha256",
		"view_version_arn":     "ViewVersionArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
