---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_lambda_version Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Lambda::Version
---

# awscc_lambda_version (Data Source)

Data Source schema for AWS::Lambda::Version



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `code_sha_256` (String) Only publish a version if the hash value matches the value that's specified. Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.
- `description` (String) A description for the version to override the description in the function configuration. Updates are not supported for this property.
- `function_arn` (String) The ARN of the version.
- `function_name` (String) The name of the Lambda function.
- `provisioned_concurrency_config` (Attributes) Specifies a provisioned concurrency configuration for a function's version. Updates are not supported for this property. (see [below for nested schema](#nestedatt--provisioned_concurrency_config))
- `runtime_policy` (Attributes) Specifies the runtime management configuration of a function. Displays runtimeVersionArn only for Manual. (see [below for nested schema](#nestedatt--runtime_policy))
- `version` (String) The version number.

<a id="nestedatt--provisioned_concurrency_config"></a>
### Nested Schema for `provisioned_concurrency_config`

Read-Only:

- `provisioned_concurrent_executions` (Number) The amount of provisioned concurrency to allocate for the version.


<a id="nestedatt--runtime_policy"></a>
### Nested Schema for `runtime_policy`

Read-Only:

- `runtime_version_arn` (String) The ARN of the runtime the function is configured to use. If the runtime update mode is manual, the ARN is returned, otherwise null is returned.
- `update_runtime_on` (String) The runtime update mode.