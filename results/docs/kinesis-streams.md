# AWS Kinesis Data Streams
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| Medium | Kinesis.2 | AWS Foundational Security Best Practices 1.0.0 | Kinesis Data Streams should have server-side encryption enabled | This control checks whether server-side encryption is enabled for Amazon Kinesis Data Streams | **IaC** - Configure server-side encryption using AWS KMS keys in Terraform or CloudFormation templates with encryption_type parameter |  |
| Medium | Kinesis.1 | AWS Foundational Security Best Practices 1.0.0 | Kinesis Data Streams should be encrypted at rest | This control checks whether Amazon Kinesis Data Streams are encrypted at rest with server-side encryption | **IaC** - Define IAM roles and policies with minimal required permissions for producers and consumers in IaC templates |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage user accounts and access to Kinesis Data Streams with proper authentication and authorization | **Platform** - Configure AWS IAM with MFA enforcement and regular access reviews for Kinesis stream access |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data in transit to and from Kinesis Data Streams using encryption | **User** - Use HTTPS/TLS for all API calls and configure client-side encryption for sensitive data |  |
| **High** | SC-13 | NIST 800-53 Rev 5 | Cryptographic Protection | Use FIPS 140-2 validated cryptographic modules for data protection | **IaC** - Configure AWS KMS with customer-managed keys and enable encryption for all streams |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Enable comprehensive logging of all Kinesis Data Streams activities | **IaC** - Enable CloudTrail logging for all Kinesis API calls and configure CloudWatch for stream metrics |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Kinesis Data Streams for unauthorized access and anomalous activities | **Platform** - Configure CloudWatch alarms for unusual access patterns and integrate with AWS Security Hub |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Implement backup and recovery procedures for critical stream configurations | **IaC** - Version control all stream configurations and implement automated backup of stream metadata |  |
| **Critical** | IAM-01 | CSA CCM v4.0.10 | Identity and Access Management | Implement strong identity and access management controls for Kinesis Data Streams | **Platform** - Configure AWS IAM with role-based access control, MFA, and regular access certifications |  |
| **Critical** | DSI-01 | CSA CCM v4.0.10 | Data Security and Information Lifecycle Management | Classify and protect data throughout its lifecycle in Kinesis Data Streams | **User** - Implement data classification tags and configure appropriate retention policies based on data sensitivity |  |
| **High** | EKM-01 | CSA CCM v4.0.10 | Encryption and Key Management | Implement proper encryption and key management for Kinesis Data Streams | **IaC** - Use AWS KMS with customer-managed keys and implement key rotation policies |  |
| **High** | IVS-01 | CSA CCM v4.0.10 | Infrastructure and Virtualization Security | Secure the underlying infrastructure supporting Kinesis Data Streams | **IaC** - Deploy streams in private subnets with proper security groups and NACLs |  |
| Medium | BCR-01 | CSA CCM v4.0.10 | Business Continuity Management and Operational Resilience | Ensure business continuity and operational resilience of Kinesis Data Streams | **IaC** - Configure cross-region replication and implement disaster recovery procedures |  |
| Medium | Kinesis.1 | AWS Security Hub Latest | Kinesis Data Streams should be encrypted at rest | This control checks whether Amazon Kinesis Data Streams are encrypted at rest with server-side encryption | **IaC** - Enable server-side encryption using AWS KMS in stream configuration templates |  |
| **High** | CloudTrail.1 | AWS Security Hub Latest | CloudTrail should be enabled and configured with at least one multi-Region trail that includes read and write management events | AWS CloudTrail can help in non-repudiation by recording AWS Management Console actions and API calls | **Platform** - Configure CloudTrail with S3 bucket logging and log file validation enabled |  |
| Medium | Config.1 | AWS Security Hub Latest | AWS Config should be enabled | This control checks whether AWS Config is enabled in the account for the local Region and is recording all resources | **Platform** - Enable AWS Config rules to monitor Kinesis stream configuration compliance |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Kinesis Data Streams Cost Optimization Latest | Right-size shard capacity based on throughput requirements | Monitor and adjust shard count to match actual data throughput to avoid over-provisioning | **Platform** - Use CloudWatch metrics to monitor shard utilization and implement auto-scaling or manual adjustments |  |
| **High** | COST-02 | AWS Kinesis Data Streams Cost Optimization Latest | Optimize data retention period | Set appropriate retention periods to balance data availability with storage costs | **IaC** - Configure retention period parameter in stream configuration based on business requirements |  |
| Medium | COST-03 | AWS Kinesis Data Streams Cost Optimization Latest | Use On-Demand capacity for variable workloads | Consider On-Demand capacity mode for streams with unpredictable or variable traffic patterns | **IaC** - Configure stream mode as ON_DEMAND in templates for variable workloads |  |
| Medium | COST-04 | AWS Kinesis Data Streams Cost Optimization Latest | Implement efficient data compression | Compress data before sending to Kinesis to reduce payload size and associated costs | **User** - Implement data compression in producer applications before putting records to streams |  |
| Medium | COST-05 | AWS Kinesis Data Streams Cost Optimization Latest | Monitor and optimize PUT requests | Batch multiple records in single PUT requests to reduce API call costs | **User** - Use PutRecords API instead of PutRecord for batch operations in producer applications |  |
| Low | COST-06 | AWS Kinesis Data Streams Cost Optimization Latest | Implement proper tagging for cost allocation | Tag Kinesis streams appropriately for cost tracking and allocation across teams or projects | **IaC** - Define consistent tagging strategy in IaC templates for cost center, environment, and project identification |  |
| Low | COST-07 | AWS Kinesis Data Streams Cost Optimization Latest | Regular cost review and optimization | Regularly review Kinesis costs and usage patterns to identify optimization opportunities | **Platform** - Set up AWS Cost Explorer reports and budgets to monitor Kinesis spending trends |  |

