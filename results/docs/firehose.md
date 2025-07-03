# AWS Data Firehose
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Identity and Access Management Policy | Implement comprehensive identity and access management policies for Firehose streams and related resources | **IaC** - Define IAM policies with least privilege principles using CloudFormation or Terraform templates for Firehose access |  |
| **Critical** | DSI-01 | CSA Cloud Controls Matrix v4.0 | Data Security and Information Lifecycle Management | Ensure data encryption in transit and at rest for all Firehose delivery streams | **IaC** - Configure server-side encryption using AWS KMS keys and enforce HTTPS/TLS for data transmission in Firehose configuration |  |
| **High** | AIS-01 | CSA Cloud Controls Matrix v4.0 | Audit Logging / Intrusion Detection | Enable comprehensive logging and monitoring for Firehose activities | **IaC** - Configure CloudTrail logging for Firehose API calls and enable CloudWatch monitoring for delivery metrics |  |
| **High** | DSI-05 | CSA Cloud Controls Matrix v4.0 | Data Loss Prevention | Implement data loss prevention mechanisms to protect sensitive data in Firehose streams | **Platform** - Enable error record handling and configure backup S3 bucket for failed deliveries |  |
| Medium | TVM-01 | CSA Cloud Controls Matrix v4.0 | Vulnerability Management | Regularly assess and remediate vulnerabilities in Firehose configurations | **User** - Conduct periodic security assessments of Firehose configurations and access patterns |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage accounts with access to Firehose resources with proper provisioning and deprovisioning procedures | **User** - Implement automated account lifecycle management for users accessing Firehose resources through IAM roles and policies |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission confidentiality and integrity for Firehose streams | **IaC** - Configure TLS encryption for all data sources and destinations, and enable server-side encryption for delivery streams |  |
| **Critical** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect information at rest in Firehose temporary storage and destination systems | **IaC** - Enable server-side encryption using AWS KMS customer-managed keys for all Firehose delivery streams |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Log auditable events for Firehose service usage and access | **Platform** - Enable CloudTrail for Firehose API logging and configure CloudWatch Logs for delivery stream events |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Firehose systems for security-relevant events and anomalies | **IaC** - Set up CloudWatch alarms for Firehose delivery failures, throttling, and unusual access patterns |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Implement backup procedures for Firehose configuration and data recovery | **IaC** - Configure error record processing with S3 backup bucket and implement cross-region replication for critical streams |  |
| **Critical** | KinesisFirehose.1 | AWS Foundational Security Standard 1.0 | Kinesis Data Firehose delivery streams should be encrypted at rest | Firehose delivery streams should have server-side encryption enabled to protect data at rest | **IaC** - Configure SSEDescription parameter with KMS encryption in Firehose delivery stream configuration |  |
| **High** | CloudTrail.1 | AWS Foundational Security Standard 1.0 | CloudTrail should be enabled and configured with at least one multi-Region trail | Enable CloudTrail to log Firehose API calls for security monitoring and compliance | **Platform** - Ensure CloudTrail is enabled with multi-region configuration to capture all Firehose API activities |  |
| **High** | IAM.1 | AWS Foundational Security Standard 1.0 | IAM policies should not allow full '*:*' administrative privileges | Restrict IAM policies for Firehose to follow least privilege principle | **IaC** - Create specific IAM policies for Firehose operations without wildcard permissions, scoped to specific resources |  |
| Medium | CloudWatch.1 | AWS Foundational Security Standard 1.0 | CloudWatch alarms should be configured for critical metrics | Configure CloudWatch alarms for Firehose delivery failures and performance metrics | **IaC** - Set up CloudWatch alarms for DeliveryToS3.Success, DeliveryToS3.Records, and other critical Firehose metrics |  |
| **Critical** | KinesisFirehose.1 | AWS Security Hub 2023 | Kinesis Data Firehose delivery streams should be encrypted at rest | Kinesis Data Firehose delivery streams should be encrypted at rest to protect sensitive data | **IaC** - Enable server-side encryption using AWS KMS keys in Firehose delivery stream configuration |  |
| **High** | Config.1 | AWS Security Hub 2023 | AWS Config should be enabled | Enable AWS Config to track configuration changes for Firehose resources | **Platform** - Configure AWS Config rules to monitor Firehose delivery stream configurations and compliance status |  |
| **High** | GuardDuty.1 | AWS Security Hub 2023 | GuardDuty should be enabled | Enable GuardDuty to detect malicious activity related to Firehose resources | **Platform** - Enable GuardDuty in all regions where Firehose is deployed for threat detection |  |
| Medium | AccessAnalyzer.1 | AWS Security Hub 2023 | IAM Access Analyzer should be enabled | Use IAM Access Analyzer to review and validate Firehose resource access policies | **Platform** - Enable IAM Access Analyzer to identify unintended access to Firehose resources from external entities |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Data Firehose Cost Optimization Best Practices 2024 | Optimize Data Compression | Enable data compression to reduce storage costs and improve delivery performance | **IaC** - Configure compression settings (GZIP, Snappy, or ZIP) in Firehose delivery stream to reduce data size and associated costs |  |
| **High** | COST-02 | AWS Data Firehose Cost Optimization Best Practices 2024 | Right-size Buffer Intervals and Sizes | Optimize buffer size and interval settings to minimize per-record costs while meeting latency requirements | **IaC** - Configure appropriate BufferSize (1-128 MB) and BufferInterval (60-900 seconds) based on data volume patterns to optimize cost per GB |  |
| Medium | COST-03 | AWS Data Firehose Cost Optimization Best Practices 2024 | Implement Data Format Conversion | Convert data to columnar formats like Parquet or ORC to reduce storage costs and improve query performance | **IaC** - Enable format conversion to Parquet or ORC in Firehose configuration to optimize downstream storage and analytics costs |  |
| Medium | COST-04 | AWS Data Firehose Cost Optimization Best Practices 2024 | Monitor and Optimize Data Processing | Optimize Lambda data transformation functions to reduce processing costs and improve efficiency | **User** - Review and optimize Lambda function code for data transformation, right-size memory allocation, and minimize execution time |  |
| Medium | COST-05 | AWS Data Firehose Cost Optimization Best Practices 2024 | Implement Lifecycle Policies for Destinations | Configure lifecycle policies for S3 destinations to automatically transition data to lower-cost storage classes | **IaC** - Set up S3 lifecycle policies to transition delivered data to IA, Glacier, or Deep Archive storage classes based on access patterns |  |
| Low | COST-06 | AWS Data Firehose Cost Optimization Best Practices 2024 | Monitor Usage with Cost Allocation Tags | Implement comprehensive tagging strategy for cost tracking and allocation | **IaC** - Apply consistent tags to Firehose resources including environment, project, team, and cost center for detailed cost analysis |  |
| Low | COST-07 | AWS Data Firehose Cost Optimization Best Practices 2024 | Optimize Error Record Handling | Configure efficient error record processing to minimize retry costs and failed delivery charges | **IaC** - Set appropriate retry duration and configure error record processing to S3 bucket with lifecycle policies for cost optimization |  |

