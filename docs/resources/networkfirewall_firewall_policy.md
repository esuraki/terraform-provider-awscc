---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_networkfirewall_firewall_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource type definition for AWS::NetworkFirewall::FirewallPolicy
---

# awscc_networkfirewall_firewall_policy (Resource)

Resource type definition for AWS::NetworkFirewall::FirewallPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **firewall_policy** (Attributes) (see [below for nested schema](#nestedatt--firewall_policy))
- **firewall_policy_name** (String)

### Optional

- **description** (String)
- **tags** (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **firewall_policy_arn** (String) A resource ARN.
- **firewall_policy_id** (String)
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--firewall_policy"></a>
### Nested Schema for `firewall_policy`

Required:

- **stateful_rule_group_references** (Attributes Set) (see [below for nested schema](#nestedatt--firewall_policy--stateful_rule_group_references))
- **stateless_custom_actions** (Attributes Set) (see [below for nested schema](#nestedatt--firewall_policy--stateless_custom_actions))
- **stateless_default_actions** (Set of String)
- **stateless_fragment_default_actions** (Set of String)
- **stateless_rule_group_references** (Attributes Set) (see [below for nested schema](#nestedatt--firewall_policy--stateless_rule_group_references))

<a id="nestedatt--firewall_policy--stateful_rule_group_references"></a>
### Nested Schema for `firewall_policy.stateful_rule_group_references`

Required:

- **resource_arn** (String) A resource ARN.


<a id="nestedatt--firewall_policy--stateless_custom_actions"></a>
### Nested Schema for `firewall_policy.stateless_custom_actions`

Required:

- **action_definition** (Attributes) (see [below for nested schema](#nestedatt--firewall_policy--stateless_custom_actions--action_definition))
- **action_name** (String)

<a id="nestedatt--firewall_policy--stateless_custom_actions--action_definition"></a>
### Nested Schema for `firewall_policy.stateless_custom_actions.action_definition`

Required:

- **publish_metric_action** (Attributes) (see [below for nested schema](#nestedatt--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action))

<a id="nestedatt--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action"></a>
### Nested Schema for `firewall_policy.stateless_custom_actions.action_definition.publish_metric_action`

Required:

- **dimensions** (Attributes Set) (see [below for nested schema](#nestedatt--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action--dimensions))

<a id="nestedatt--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action--dimensions"></a>
### Nested Schema for `firewall_policy.stateless_custom_actions.action_definition.publish_metric_action.dimensions`

Required:

- **value** (String)





<a id="nestedatt--firewall_policy--stateless_rule_group_references"></a>
### Nested Schema for `firewall_policy.stateless_rule_group_references`

Required:

- **priority** (Number)
- **resource_arn** (String) A resource ARN.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_networkfirewall_firewall_policy.example <resource ID>
```