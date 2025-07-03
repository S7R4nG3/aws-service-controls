# AWS Step Functions
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Entitlement | Implement strong identity verification and access controls for Step Functions execution and management | **IaC** - Use AWS IAM roles with least privilege principles for Step Functions execution roles and user access |  |
| **Critical** | EKM-01 | CSA Cloud Controls Matrix v4.0 | Encryption & Key Management | Implement proper encryption and key management for Step Functions state data | **IaC** - Configure KMS encryption for Step Functions state machines and use customer-managed keys |  |
| **High** | DSI-02 | CSA Cloud Controls Matrix v4.0 | Data Security & Information Lifecycle Management | Implement data classification and protection measures for Step Functions workflows | **User** - Classify data processed by Step Functions and implement appropriate security controls based on sensitivity |  |
| **High** | GRM-01 | CSA Cloud Controls Matrix v4.0 | Governance and Risk Management | Enable comprehensive logging for Step Functions execution and API calls | **Platform** - Enable CloudWatch Logs for Step Functions and configure CloudTrail for API logging |  |
| Medium | BCR-01 | CSA Cloud Controls Matrix v4.0 | Business Continuity Management & Operational Resilience | Implement business continuity and disaster recovery for Step Functions | **IaC** - Design Step Functions with cross-region failover capabilities and backup strategies |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage accounts with access to Step Functions resources | **Platform** - Use AWS Organizations and IAM to manage accounts and roles accessing Step Functions |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data in transit for Step Functions communications | **Platform** - Ensure all Step Functions communications use HTTPS/TLS encryption |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Generate audit records for Step Functions events | **IaC** - Configure CloudWatch Logs and CloudTrail to capture Step Functions execution events |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Step Functions for security events and anomalies | **Platform** - Use CloudWatch metrics and alarms to monitor Step Functions execution patterns |  |
| Medium | CP-2 | NIST 800-53 Rev 5 | Contingency Plan | Develop contingency plans for Step Functions failures | **User** - Create documented procedures for Step Functions failure scenarios and recovery |  |
| **Critical** | StepFunctions.1 | AWS Foundational Security Best Practices 1.0.0 | Step Functions state machines should have logging turned on | Step Functions state machines should have logging enabled to CloudWatch Logs | **IaC** - Configure loggingConfiguration in Step Functions state machine definition with CloudWatch Logs destination |  |
| **Critical** | IAM.1 | AWS Foundational Security Best Practices 1.0.0 | IAM policies should not allow full administrative privileges | Step Functions execution roles should follow least privilege principle | **IaC** - Create specific IAM policies for Step Functions with only required permissions |  |
| **High** | CloudTrail.1 | AWS Foundational Security Best Practices 1.0.0 | CloudTrail should be enabled and configured with at least one multi-Region trail | Enable CloudTrail to log Step Functions API calls | **Platform** - Configure CloudTrail to capture Step Functions API calls across all regions |  |
| **High** | KMS.1 | AWS Foundational Security Best Practices 1.0.0 | Customer managed keys should be rotated annually | Enable automatic key rotation for KMS keys used with Step Functions | **IaC** - Configure automatic key rotation for KMS keys used to encrypt Step Functions data |  |
| Medium | CloudWatch.1 | AWS Foundational Security Best Practices 1.0.0 | CloudWatch alarms should be configured for Step Functions metrics | Set up CloudWatch alarms for Step Functions execution failures and performance | **IaC** - Create CloudWatch alarms for ExecutionsFailed, ExecutionsTimedOut, and other critical metrics |  |
| **Critical** | StepFunctions.1 | AWS Security Hub 2024.1 | Step Functions state machines should have logging turned on | Step Functions state machines should have logging enabled to CloudWatch Logs | **IaC** - Configure loggingConfiguration in Step Functions state machine definition with CloudWatch Logs destination |  |
| **Critical** | IAM.1 | AWS Security Hub 2024.1 | IAM policies should not allow full administrative privileges | Step Functions execution roles should follow least privilege principle | **IaC** - Review and minimize IAM permissions for Step Functions execution roles regularly |  |
| **High** | CloudTrail.2 | AWS Security Hub 2024.1 | CloudTrail should have encryption at-rest enabled | Encrypt CloudTrail logs that capture Step Functions API calls | **IaC** - Configure CloudTrail with KMS encryption for Step Functions API audit logs |  |
| **High** | KMS.4 | AWS Security Hub 2024.1 | AWS KMS key rotation should be enabled | Enable automatic key rotation for KMS keys used with Step Functions | **IaC** - Configure automatic key rotation for KMS keys used to encrypt Step Functions data |  |
| Medium | CloudWatch.13 | AWS Security Hub 2024.1 | Route 53 public hosted zones should log DNS queries | Configure comprehensive monitoring for Step Functions execution and performance | **Platform** - Set up CloudWatch dashboards and SNS notifications for Step Functions monitoring |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-SF-01 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Use Express Workflows for high-volume, short-duration workloads | Implement Express Workflows instead of Standard Workflows for high-frequency, short-duration executions | **IaC** - Configure Step Functions with Express Workflow type for workloads with execution duration under 5 minutes and high frequency |  |
| **High** | COST-SF-02 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Optimize state transitions and reduce unnecessary steps | Minimize the number of state transitions to reduce per-transition costs | **User** - Review workflow definitions to combine states where possible and eliminate unnecessary Pass states |  |
| Medium | COST-SF-03 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Implement proper timeout configurations | Set appropriate timeouts to prevent long-running executions that increase costs | **IaC** - Configure TimeoutSeconds for tasks and overall execution to prevent runaway costs |  |
| Medium | COST-SF-04 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Use Map state efficiently for parallel processing | Optimize Map state configurations to balance performance and cost | **User** - Configure MaxConcurrency in Map states to control parallel execution costs |  |
| Medium | COST-SF-05 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Monitor and analyze Step Functions costs regularly | Implement cost monitoring and analysis for Step Functions usage | **Platform** - Use AWS Cost Explorer and CloudWatch metrics to track Step Functions costs and optimize based on usage patterns |  |
| Low | COST-SF-06 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Optimize data payload sizes | Minimize data passed between states to reduce storage and transfer costs | **User** - Use S3 or other storage services for large payloads and pass references instead of full data |  |
| Low | COST-SF-07 | AWS Step Functions Cost Optimization Best Practices 2024.1 | Implement execution history retention policies | Configure appropriate retention policies for Step Functions execution history | **Platform** - Use CloudWatch Logs retention settings to manage Step Functions log retention and associated costs |  |

