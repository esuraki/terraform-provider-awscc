// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lambda_event_source_mapping", eventSourceMappingDataSource)
}

// eventSourceMappingDataSource returns the Terraform awscc_lambda_event_source_mapping data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lambda::EventSourceMapping resource.
func eventSourceMappingDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"amazon_managed_kafka_event_source_config": {
			// Property: AmazonManagedKafkaEventSourceConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specific configuration settings for an MSK event source.",
			//	  "properties": {
			//	    "ConsumerGroupId": {
			//	      "description": "The identifier for the Kafka Consumer Group to join.",
			//	      "maxLength": 200,
			//	      "minLength": 1,
			//	      "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Specific configuration settings for an MSK event source.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"consumer_group_id": {
						// Property: ConsumerGroupId
						Description: "The identifier for the Kafka Consumer Group to join.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"batch_size": {
			// Property: BatchSize
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The maximum number of items to retrieve in a single batch.",
			//	  "maximum": 10000,
			//	  "minimum": 1,
			//	  "type": "integer"
			//	}
			Description: "The maximum number of items to retrieve in a single batch.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"bisect_batch_on_function_error": {
			// Property: BisectBatchOnFunctionError
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) If the function returns an error, split the batch in two and retry.",
			//	  "type": "boolean"
			//	}
			Description: "(Streams) If the function returns an error, split the batch in two and retry.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"destination_config": {
			// Property: DestinationConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
			//	  "properties": {
			//	    "OnFailure": {
			//	      "additionalProperties": false,
			//	      "description": "The destination configuration for failed invocations.",
			//	      "properties": {
			//	        "Destination": {
			//	          "description": "The Amazon Resource Name (ARN) of the destination resource.",
			//	          "maxLength": 1024,
			//	          "minLength": 12,
			//	          "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "(Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"on_failure": {
						// Property: OnFailure
						Description: "The destination configuration for failed invocations.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"destination": {
									// Property: Destination
									Description: "The Amazon Resource Name (ARN) of the destination resource.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Disables the event source mapping to pause polling and invocation.",
			//	  "type": "boolean"
			//	}
			Description: "Disables the event source mapping to pause polling and invocation.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"event_source_arn": {
			// Property: EventSourceArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the event source.",
			//	  "maxLength": 1024,
			//	  "minLength": 12,
			//	  "pattern": "arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}(-gov)?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the event source.",
			Type:        types.StringType,
			Computed:    true,
		},
		"filter_criteria": {
			// Property: FilterCriteria
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The filter criteria to control event filtering.",
			//	  "properties": {
			//	    "Filters": {
			//	      "description": "List of filters of this FilterCriteria",
			//	      "items": {
			//	        "additionalProperties": false,
			//	        "description": "The filter object that defines parameters for ESM filtering.",
			//	        "properties": {
			//	          "Pattern": {
			//	            "description": "The filter pattern that defines which events should be passed for invocations.",
			//	            "maxLength": 4096,
			//	            "minLength": 0,
			//	            "pattern": ".*",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "maxItems": 20,
			//	      "minItems": 1,
			//	      "type": "array",
			//	      "uniqueItems": true
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The filter criteria to control event filtering.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"filters": {
						// Property: Filters
						Description: "List of filters of this FilterCriteria",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"pattern": {
									// Property: Pattern
									Description: "The filter pattern that defines which events should be passed for invocations.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"function_name": {
			// Property: FunctionName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the Lambda function.",
			//	  "maxLength": 140,
			//	  "minLength": 1,
			//	  "pattern": "(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}(-gov)?-[a-z]+-\\d{1}:)?(\\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\\$LATEST|[a-zA-Z0-9-_]+))?",
			//	  "type": "string"
			//	}
			Description: "The name of the Lambda function.",
			Type:        types.StringType,
			Computed:    true,
		},
		"function_response_types": {
			// Property: FunctionResponseTypes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) A list of response types supported by the function.",
			//	  "items": {
			//	    "enum": [
			//	      "ReportBatchItemFailures"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "maxLength": 1,
			//	  "minLength": 0,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "(Streams) A list of response types supported by the function.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Event Source Mapping Identifier UUID.",
			//	  "maxLength": 36,
			//	  "minLength": 36,
			//	  "pattern": "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "Event Source Mapping Identifier UUID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"maximum_batching_window_in_seconds": {
			// Property: MaximumBatchingWindowInSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
			//	  "maximum": 300,
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Description: "(Streams) The maximum amount of time to gather records before invoking the function, in seconds.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"maximum_record_age_in_seconds": {
			// Property: MaximumRecordAgeInSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
			//	  "maximum": 604800,
			//	  "minimum": -1,
			//	  "type": "integer"
			//	}
			Description: "(Streams) The maximum age of a record that Lambda sends to a function for processing.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"maximum_retry_attempts": {
			// Property: MaximumRetryAttempts
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) The maximum number of times to retry when the function returns an error.",
			//	  "maximum": 10000,
			//	  "minimum": -1,
			//	  "type": "integer"
			//	}
			Description: "(Streams) The maximum number of times to retry when the function returns an error.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"parallelization_factor": {
			// Property: ParallelizationFactor
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) The number of batches to process from each shard concurrently.",
			//	  "maximum": 10,
			//	  "minimum": 1,
			//	  "type": "integer"
			//	}
			Description: "(Streams) The number of batches to process from each shard concurrently.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"queues": {
			// Property: Queues
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(ActiveMQ) A list of ActiveMQ queues.",
			//	  "items": {
			//	    "maxLength": 1000,
			//	    "minLength": 1,
			//	    "pattern": "[\\s\\S]*",
			//	    "type": "string"
			//	  },
			//	  "maxItems": 1,
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "(ActiveMQ) A list of ActiveMQ queues.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"scaling_config": {
			// Property: ScalingConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The scaling configuration for the event source.",
			//	  "properties": {
			//	    "MaximumConcurrency": {
			//	      "description": "The maximum number of concurrent functions that the event source can invoke.",
			//	      "maximum": 1000,
			//	      "minimum": 2,
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The scaling configuration for the event source.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"maximum_concurrency": {
						// Property: MaximumConcurrency
						Description: "The maximum number of concurrent functions that the event source can invoke.",
						Type:        types.Int64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"self_managed_event_source": {
			// Property: SelfManagedEventSource
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Self-managed event source endpoints.",
			//	  "properties": {
			//	    "Endpoints": {
			//	      "additionalProperties": false,
			//	      "description": "The endpoints for a self-managed event source.",
			//	      "properties": {
			//	        "KafkaBootstrapServers": {
			//	          "description": "A list of Kafka server endpoints.",
			//	          "items": {
			//	            "description": "The URL of a Kafka server.",
			//	            "maxLength": 300,
			//	            "minLength": 1,
			//	            "pattern": "^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9]):[0-9]{1,5}",
			//	            "type": "string"
			//	          },
			//	          "maxItems": 10,
			//	          "minItems": 1,
			//	          "type": "array",
			//	          "uniqueItems": true
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Self-managed event source endpoints.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoints": {
						// Property: Endpoints
						Description: "The endpoints for a self-managed event source.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kafka_bootstrap_servers": {
									// Property: KafkaBootstrapServers
									Description: "A list of Kafka server endpoints.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"self_managed_kafka_event_source_config": {
			// Property: SelfManagedKafkaEventSourceConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specific configuration settings for a Self-Managed Apache Kafka event source.",
			//	  "properties": {
			//	    "ConsumerGroupId": {
			//	      "description": "The identifier for the Kafka Consumer Group to join.",
			//	      "maxLength": 200,
			//	      "minLength": 1,
			//	      "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Specific configuration settings for a Self-Managed Apache Kafka event source.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"consumer_group_id": {
						// Property: ConsumerGroupId
						Description: "The identifier for the Kafka Consumer Group to join.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"source_access_configurations": {
			// Property: SourceAccessConfigurations
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of SourceAccessConfiguration.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "The configuration used by AWS Lambda to access event source",
			//	    "properties": {
			//	      "Type": {
			//	        "description": "The type of source access configuration.",
			//	        "enum": [
			//	          "BASIC_AUTH",
			//	          "VPC_SUBNET",
			//	          "VPC_SECURITY_GROUP",
			//	          "SASL_SCRAM_512_AUTH",
			//	          "SASL_SCRAM_256_AUTH",
			//	          "VIRTUAL_HOST",
			//	          "CLIENT_CERTIFICATE_TLS_AUTH",
			//	          "SERVER_ROOT_CA_CERTIFICATE"
			//	        ],
			//	        "type": "string"
			//	      },
			//	      "URI": {
			//	        "description": "The URI for the source access configuration resource.",
			//	        "maxLength": 200,
			//	        "minLength": 1,
			//	        "pattern": "[a-zA-Z0-9-\\/*:_+=.@-]*",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "maxItems": 22,
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "A list of SourceAccessConfiguration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Description: "The type of source access configuration.",
						Type:        types.StringType,
						Computed:    true,
					},
					"uri": {
						// Property: URI
						Description: "The URI for the source access configuration resource.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"starting_position": {
			// Property: StartingPosition
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
			//	  "maxLength": 12,
			//	  "minLength": 6,
			//	  "pattern": "(LATEST|TRIM_HORIZON|AT_TIMESTAMP)+",
			//	  "type": "string"
			//	}
			Description: "The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.",
			Type:        types.StringType,
			Computed:    true,
		},
		"starting_position_timestamp": {
			// Property: StartingPositionTimestamp
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
			//	  "type": "number"
			//	}
			Description: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.",
			Type:        types.Float64Type,
			Computed:    true,
		},
		"topics": {
			// Property: Topics
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Kafka) A list of Kafka topics.",
			//	  "items": {
			//	    "maxLength": 249,
			//	    "minLength": 1,
			//	    "pattern": "^[^.]([a-zA-Z0-9\\-_.]+)",
			//	    "type": "string"
			//	  },
			//	  "maxItems": 1,
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "(Kafka) A list of Kafka topics.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"tumbling_window_in_seconds": {
			// Property: TumblingWindowInSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
			//	  "maximum": 900,
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Description: "(Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.",
			Type:        types.Int64Type,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Lambda::EventSourceMapping",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::EventSourceMapping").WithTerraformTypeName("awscc_lambda_event_source_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"amazon_managed_kafka_event_source_config": "AmazonManagedKafkaEventSourceConfig",
		"batch_size":                             "BatchSize",
		"bisect_batch_on_function_error":         "BisectBatchOnFunctionError",
		"consumer_group_id":                      "ConsumerGroupId",
		"destination":                            "Destination",
		"destination_config":                     "DestinationConfig",
		"enabled":                                "Enabled",
		"endpoints":                              "Endpoints",
		"event_source_arn":                       "EventSourceArn",
		"filter_criteria":                        "FilterCriteria",
		"filters":                                "Filters",
		"function_name":                          "FunctionName",
		"function_response_types":                "FunctionResponseTypes",
		"id":                                     "Id",
		"kafka_bootstrap_servers":                "KafkaBootstrapServers",
		"maximum_batching_window_in_seconds":     "MaximumBatchingWindowInSeconds",
		"maximum_concurrency":                    "MaximumConcurrency",
		"maximum_record_age_in_seconds":          "MaximumRecordAgeInSeconds",
		"maximum_retry_attempts":                 "MaximumRetryAttempts",
		"on_failure":                             "OnFailure",
		"parallelization_factor":                 "ParallelizationFactor",
		"pattern":                                "Pattern",
		"queues":                                 "Queues",
		"scaling_config":                         "ScalingConfig",
		"self_managed_event_source":              "SelfManagedEventSource",
		"self_managed_kafka_event_source_config": "SelfManagedKafkaEventSourceConfig",
		"source_access_configurations":           "SourceAccessConfigurations",
		"starting_position":                      "StartingPosition",
		"starting_position_timestamp":            "StartingPositionTimestamp",
		"topics":                                 "Topics",
		"tumbling_window_in_seconds":             "TumblingWindowInSeconds",
		"type":                                   "Type",
		"uri":                                    "URI",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
