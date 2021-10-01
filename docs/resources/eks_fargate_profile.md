---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_eks_fargate_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema for AWS::EKS::FargateProfile
---

# awscc_eks_fargate_profile (Resource)

Resource Schema for AWS::EKS::FargateProfile



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cluster_name** (String) Name of the Cluster
- **pod_execution_role_arn** (String) The IAM policy arn for pods
- **selectors** (Attributes List) (see [below for nested schema](#nestedatt--selectors))

### Optional

- **fargate_profile_name** (String) Name of FargateProfile
- **subnets** (List of String)
- **tags** (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **arn** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--selectors"></a>
### Nested Schema for `selectors`

Required:

- **labels** (Attributes List) (see [below for nested schema](#nestedatt--selectors--labels))
- **namespace** (String)

<a id="nestedatt--selectors--labels"></a>
### Nested Schema for `selectors.labels`

Required:

- **key** (String) The key name of the label.
- **value** (String) The value for the label.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- **value** (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_eks_fargate_profile.example <resource ID>
```