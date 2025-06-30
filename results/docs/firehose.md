# AWS Data Firehose
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management | IAM-01 | Implement proper identity and access management controls for Firehose delivery streams | **IaC** - Define IAM roles and policies with least privilege access to Firehose delivery streams using CloudFormation or Terraform |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Ensure data encryption in transit and at rest for all Firehose delivery streams | **IaC** - Configure server-side encryption using AWS KMS keys and enable SSL/TLS encryption for data in transit |  |
| **High** | Governance and Risk Management | GRM-01 | Establish governance framework for Firehose stream management and data handling | **Platform** - Implement AWS Config rules to monitor Firehose configuration compliance and AWS Organizations policies |  |
| **High** | Data Classification | DSI-02 | Classify data flowing through Firehose streams and apply appropriate protection measures | **User** - Tag Firehose delivery streams with data classification labels and implement data loss prevention measures |  |
| Medium | Application and Interface Security | AIS-01 | Secure APIs and interfaces used to interact with Firehose | **IaC** - Implement VPC endpoints for Firehose API calls and configure API throttling limits |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Firehose resources | **IaC** - Implement resource-based policies and IAM policies with explicit deny statements for unauthorized access |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information through Firehose | **IaC** - Enable SSL/TLS encryption for all data transmission and configure encryption parameters in delivery stream settings |  |
| **High** | Event Logging | AU-2 | Ensure Firehose activities are logged and auditable | **Platform** - Enable CloudTrail logging for all Firehose API calls and configure CloudWatch logging for delivery stream events |  |
| **High** | Protection of Information at Rest | SC-28 | Protect information at rest in Firehose buffers and destination storage | **IaC** - Configure KMS encryption for Firehose delivery streams and ensure destination storage encryption |  |
| **High** | System Monitoring | SI-4 | Monitor Firehose delivery streams for security events and anomalies | **Platform** - Configure CloudWatch alarms for delivery failures, data transformation errors, and unusual traffic patterns |  |
| Medium | System Backup | CP-9 | Ensure backup capabilities for critical Firehose configurations | **IaC** - Version control Firehose configurations in Infrastructure as Code templates and implement cross-region replication |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Firehose delivery streams should be encrypted at rest | Firehose.1 | Ensure all Firehose delivery streams have server-side encryption enabled | **IaC** - Configure SSEConfiguration with KMS encryption in delivery stream definition |  |
| **High** | Firehose delivery streams should have error logging enabled | Firehose.2 | Enable error logging for Firehose delivery streams to CloudWatch Logs | **IaC** - Configure CloudWatchLoggingOptions in the delivery stream configuration |  |
| **High** | Firehose delivery streams should use VPC endpoints | Firehose.3 | Use VPC endpoints for Firehose to keep traffic within AWS network | **IaC** - Create VPC endpoints for Firehose service and configure route tables appropriately |  |
| Medium | Firehose should have proper IAM roles and policies | Firehose.4 | Configure least privilege IAM roles for Firehose service operations | **IaC** - Create dedicated IAM roles with minimal required permissions for Firehose delivery stream operations |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Kinesis Data Firehose delivery streams should be encrypted at rest | KinesisFirehose.1 | This control checks whether Amazon Kinesis Data Firehose delivery streams are encrypted at rest | **IaC** - Enable server-side encryption using AWS KMS customer managed keys in delivery stream configuration |  |
| **High** | Kinesis Data Firehose delivery streams should have server access logging enabled | KinesisFirehose.2 | Enable server access logging for S3 destinations of Firehose delivery streams | **IaC** - Configure S3 bucket logging for destination buckets and enable CloudWatch Logs for delivery stream errors |  |
| **High** | Kinesis Data Firehose delivery streams should not be publicly accessible | KinesisFirehose.3 | Ensure Firehose delivery streams are not accessible from public networks | **Platform** - Configure VPC endpoints and security groups to restrict access to authorized networks only |  |
| Medium | Kinesis Data Firehose should have monitoring enabled | KinesisFirehose.4 | Enable comprehensive monitoring for Firehose delivery streams | **Platform** - Configure CloudWatch metrics and alarms for delivery success rates, error rates, and data freshness |  |


## Operational Controls
---



## Cost Controls
---


### AWS Kinesis Data Firehose Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize buffer settings | COST-001 | Configure appropriate buffer size and interval to minimize costs while meeting delivery requirements | **IaC** - Set buffer size to maximum (128 MB) and buffer interval to maximum (900 seconds) when latency requirements allow |
| **High** | Implement data compression | COST-002 | Enable data compression to reduce data transfer and storage costs | **IaC** - Configure GZIP compression in the delivery stream configuration for supported destinations |
| Medium | Optimize data transformation | COST-003 | Minimize data transformation overhead and Lambda invocation costs | **User** - Use efficient Lambda functions for data transformation and consider batch processing to reduce invocation frequency |
| Medium | Right-size delivery stream configuration | COST-004 | Monitor and adjust delivery stream settings based on actual usage patterns | **Platform** - Use CloudWatch metrics to analyze throughput patterns and adjust buffer settings and scaling accordingly |
| Medium | Implement data lifecycle policies | COST-005 | Configure appropriate storage classes and lifecycle policies for destination storage | **IaC** - Configure S3 Intelligent Tiering or lifecycle transitions for long-term data storage cost optimization |
| Low | Monitor and alert on cost anomalies | COST-006 | Implement cost monitoring and alerting for Firehose usage | **Platform** - Set up AWS Budgets and Cost Anomaly Detection for Firehose service costs with appropriate alert thresholds |
| Low | Optimize error handling and retry logic | COST-007 | Implement efficient error handling to minimize retry costs and failed delivery charges | **IaC** - Configure appropriate retry duration and error output prefix to minimize processing costs for failed records |


