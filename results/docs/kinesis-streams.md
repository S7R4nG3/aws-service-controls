# AWS Kinesis Data Streams
---


### CSA CCM v5 5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| high | Identity and Access Management | IAM-01 | Implement proper identity and access management controls for Kinesis Data Streams access | **IaC** - Configure IAM policies with least privilege access using specific Kinesis actions and resource ARNs in CloudFormation/Terraform templates |  |
| high | Data Security and Information Lifecycle Management | DSI-01 | Ensure data encryption at rest and in transit for streaming data | **IaC** - Enable server-side encryption using AWS KMS keys and enforce HTTPS/TLS for data producers and consumers |  |
| medium | Data Classification | DSI-02 | Classify and tag streaming data based on sensitivity levels | **user** - Implement data classification tags on Kinesis streams and establish data handling procedures based on classification |  |
| medium | Audit Logging and Monitoring | LOG-01 | Enable comprehensive logging and monitoring of Kinesis Data Streams activities | **IaC** - Configure CloudTrail, CloudWatch metrics, and VPC Flow Logs for comprehensive monitoring of stream access and operations |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| high | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Kinesis streams | **IaC** - Implement resource-based policies and IAM roles with condition statements to control access to specific streams and operations |  |
| high | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **platform** - AWS provides TLS encryption for data in transit by default; ensure client applications use HTTPS endpoints |  |
| high | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **IaC** - Configure server-side encryption with customer-managed KMS keys for data at rest in Kinesis streams |  |
| medium | Event Logging | AU-2 | Ensure auditable events are defined and logging is configured | **IaC** - Enable CloudTrail logging for all Kinesis API calls and configure CloudWatch for operational metrics |  |
| medium | Information System Backup | CP-9 | Conduct backups of user-level information contained in the information system | **user** - Implement application-level data archival to S3 or other durable storage for critical streaming data |  |
| medium | Information System Monitoring | SI-4 | Monitor the information system to detect attacks and indicators of potential attacks | **IaC** - Configure CloudWatch alarms for anomalous activity, failed API calls, and unusual data patterns in streams |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| high | Kinesis Data Streams should be encrypted at rest | Kinesis.1 | Kinesis Data Streams should have server-side encryption enabled | **IaC** - Enable server-side encryption using AWS managed keys or customer managed KMS keys in stream configuration |  |
| high | IAM policies should not allow full administrative privileges | IAM.1 | Avoid overly permissive IAM policies for Kinesis access | **IaC** - Create specific IAM policies with minimal required permissions for Kinesis operations, avoiding wildcard permissions |  |
| medium | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable CloudTrail to log Kinesis API activities across regions | **platform** - Ensure CloudTrail is enabled at the account level to capture all Kinesis API calls and management events |  |
| medium | CloudWatch should be configured to monitor Kinesis streams | CloudWatch.1 | Monitor Kinesis streams performance and security metrics | **IaC** - Configure CloudWatch dashboards and alarms for key Kinesis metrics including throughput, errors, and iterator age |  |

### AWS Security Hub Current
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| high | Kinesis Data Streams should be encrypted at rest with AWS KMS | KDS.1 | Ensure Kinesis streams are encrypted using AWS KMS for data protection | **IaC** - Configure StreamEncryption property with KMS key ARN in infrastructure templates |  |
| medium | VPC endpoints should be used for Kinesis access from private subnets | VPC.1 | Use VPC endpoints to access Kinesis without internet gateway | **IaC** - Create VPC endpoints for Kinesis service and configure route tables to use private connectivity |  |
| medium | Cross-account access should use roles not users | IAM.2 | Use IAM roles for cross-account Kinesis access instead of IAM users | **IaC** - Create cross-account IAM roles with trust policies for external account access to Kinesis streams |  |
| low | API logging should be enabled for security monitoring | LOG.1 | Enable comprehensive API logging for security analysis | **platform** - Ensure CloudTrail data events are enabled for Kinesis streams to capture detailed API activity |  |


## Operational Controls
---



## Cost Controls
---


### AWS Kinesis Data Streams Cost Optimization Current
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| high | Right-size shard capacity | COST-01 | Monitor and adjust the number of shards based on actual throughput requirements | **user** - Use CloudWatch metrics to monitor IncomingRecords, OutgoingRecords, and IteratorAgeMilliseconds to determine optimal shard count and implement auto-scaling policies |
| high | Optimize data retention period | COST-02 | Set appropriate retention periods to balance operational needs with storage costs | **IaC** - Configure RetentionPeriodHours parameter based on business requirements, default is 24 hours but can be extended up to 365 days |
| medium | Use on-demand capacity mode judiciously | COST-03 | Evaluate on-demand vs provisioned capacity modes based on usage patterns | **user** - Use on-demand mode for unpredictable workloads and provisioned mode for steady, predictable traffic to optimize costs |
| medium | Implement efficient consumer patterns | COST-04 | Optimize consumer applications to reduce unnecessary API calls and data transfer | **user** - Use enhanced fan-out for multiple consumers, implement proper checkpointing, and batch record processing to reduce costs |
| medium | Monitor and alert on cost metrics | COST-05 | Set up cost monitoring and budgets for Kinesis usage | **IaC** - Configure AWS Budgets and Cost Explorer alerts for Kinesis spend, and use cost allocation tags for tracking |
| low | Optimize cross-region data transfer | COST-06 | Minimize cross-region data transfer costs by co-locating resources | **user** - Deploy Kinesis streams, producers, and consumers in the same AWS region when possible to avoid data transfer charges |
| low | Archive historical data efficiently | COST-07 | Use cost-effective storage for long-term data archival instead of extended retention | **user** - Implement Kinesis Data Firehose or custom consumers to archive data to S3 with appropriate storage classes for long-term retention |


