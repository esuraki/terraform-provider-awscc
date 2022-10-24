// Code generated by generators/plural-data-source/main.go; DO NOT EDIT.

package greengrassv2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSGreengrassV2DeploymentsDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::GreengrassV2::Deployment", "awscc_greengrassv2_deployments", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyDataSourceConfig(),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(fmt.Sprintf("data.%s", td.ResourceName), "ids.#"),
			),
		},
	})
}