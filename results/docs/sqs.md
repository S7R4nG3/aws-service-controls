# AWS Simple Queue Service (SQS)
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v5.0 | Identity and Access Management | Implement proper identity and access management controls for SQS resources | **IaC/Platform** - Configure IAM policies with least privilege principles, use resource-based policies for queue access, and implement cross-account access controls |  |
| **Critical** | EKM-01 | CSA Cloud Controls Matrix v5.0 | Encryption and Key Management | Implement proper encryption key management for SQS message encryption | **IaC/Platform** - Use AWS KMS for server-side encryption, implement customer-managed keys with proper rotation policies |  |
| **High** | DSI-01 | CSA Cloud Controls Matrix v5.0 | Data Security and Information Lifecycle Management | Protect data in transit and at rest for SQS messages | **Platform/User** - Enable server-side encryption, use HTTPS endpoints, implement message retention policies |  |
| **High** | BCR-01 | CSA Cloud Controls Matrix v5.0 | Business Continuity Management and Operational Resilience | Ensure business continuity and disaster recovery capabilities | **IaC** - Deploy queues across multiple availability zones, implement cross-region replication strategies |  |
| **High** | AIS-01 | CSA Cloud Controls Matrix v5.0 | Application and Interface Security | Implement comprehensive logging and monitoring for SQS operations | **Platform** - Enable CloudTrail logging, configure CloudWatch metrics and alarms, implement access logging |  |
| Medium | GRM-01 | CSA Cloud Controls Matrix v5.0 | Governance and Risk Management | Implement governance controls for SQS resource management | **Platform** - Use resource tagging, implement service control policies, establish naming conventions |  |
| **Critical** | AC-3 | NIST SP 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to SQS resources | **IaC/Platform** - Implement IAM policies with principle of least privilege, use condition keys for fine-grained access control |  |
| **Critical** | SC-13 | NIST SP 800-53 Rev 5 | Cryptographic Protection | Implement cryptographic mechanisms to protect SQS message confidentiality and integrity | **Platform** - Enable server-side encryption with KMS, use HTTPS for all API calls, implement client-side encryption if required |  |
| **High** | AU-2 | NIST SP 800-53 Rev 5 | Event Logging | Ensure SQS events are captured and logged for security monitoring | **Platform** - Enable CloudTrail for API logging, configure CloudWatch for operational metrics, implement custom logging for application events |  |
| **High** | CP-9 | NIST SP 800-53 Rev 5 | System Backup | Implement backup and recovery procedures for critical SQS configurations | **IaC** - Use Infrastructure as Code for queue configuration backup, implement cross-region deployment strategies |  |
| **High** | SI-4 | NIST SP 800-53 Rev 5 | System Monitoring | Monitor SQS for unauthorized access and anomalous behavior | **Platform** - Configure CloudWatch alarms for queue metrics, implement GuardDuty for threat detection, set up custom monitoring dashboards |  |
| Medium | SC-7 | NIST SP 800-53 Rev 5 | Boundary Protection | Implement network boundary protection for SQS access | **IaC/Platform** - Use VPC endpoints for private network access, configure security groups and NACLs, implement IP-based access restrictions |  |
| **Critical** | SQS.1 | AWS Foundational Technical Review 2024.04.01 | Amazon SQS queues should be encrypted at rest | SQS queues should have server-side encryption enabled to protect message data at rest | **IaC/Platform** - Enable server-side encryption using AWS KMS keys when creating or modifying SQS queues |  |
| **High** | CloudTrail.1 | AWS Foundational Technical Review 2024.04.01 | CloudTrail should be enabled and configured with at least one multi-Region trail | Ensure CloudTrail is capturing SQS API calls across all regions | **Platform** - Enable CloudTrail with multi-region configuration to capture all SQS API activities |  |
| **High** | Config.1 | AWS Foundational Technical Review 2024.04.01 | AWS Config should be enabled | Enable AWS Config to track SQS configuration changes | **Platform** - Configure AWS Config to monitor SQS queue configuration changes and compliance status |  |
| **Critical** | SQS.1 | AWS Security Hub 2024.04.01 | Amazon SQS queues should be encrypted at rest | Ensure all SQS queues have server-side encryption enabled | **IaC/Platform** - Configure server-side encryption for all SQS queues using AWS managed or customer managed KMS keys |  |
| **High** | CloudTrail.2 | AWS Security Hub 2024.04.01 | CloudTrail should have encryption at rest enabled | Ensure CloudTrail logs capturing SQS events are encrypted | **Platform** - Enable CloudTrail log file encryption using KMS keys to protect audit logs |  |
| **High** | IAM.1 | AWS Security Hub 2024.04.01 | IAM policies should not allow full administrative privileges | Restrict IAM policies to prevent overly broad SQS permissions | **IaC/Platform** - Review and refactor IAM policies to grant minimum necessary permissions for SQS operations |  |
| Medium | CloudWatch.1 | AWS Security Hub 2024.04.01 | A log metric filter and alarm should exist for usage of root user | Monitor for root user access to SQS resources | **Platform** - Create CloudWatch metric filters and alarms to detect root user activity on SQS resources |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-SQS-001 | AWS SQS Cost Optimization Best Practices 2023 | Optimize Message Polling Strategy | Implement long polling to reduce the number of empty receive requests and minimize costs | **User** - Configure ReceiveMessageWaitTimeSeconds to 20 seconds for long polling, reducing API calls and associated costs |  |
| **High** | COST-SQS-002 | AWS SQS Cost Optimization Best Practices 2023 | Right-size Message Retention Period | Set appropriate message retention periods to avoid unnecessary storage costs | **IaC** - Configure MessageRetentionPeriod based on business requirements, default is 4 days but can be reduced for temporary messages |  |
| **High** | COST-SQS-003 | AWS SQS Cost Optimization Best Practices 2023 | Use FIFO Queues Only When Necessary | FIFO queues cost more than standard queues, use only when ordering is required | **IaC** - Evaluate business requirements and use standard queues when message ordering is not critical |  |
| Medium | COST-SQS-004 | AWS SQS Cost Optimization Best Practices 2023 | Implement Message Batching | Use batch operations to reduce the number of API calls and associated costs | **User** - Use SendMessageBatch and ReceiveMessage with MaxNumberOfMessages parameter to process multiple messages per API call |  |
| Medium | COST-SQS-005 | AWS SQS Cost Optimization Best Practices 2023 | Monitor and Clean Up Dead Letter Queues | Regularly clean up dead letter queues to avoid accumulating storage costs | **User/IaC** - Implement automated cleanup processes for dead letter queues and set appropriate retention periods |  |
| Medium | COST-SQS-006 | AWS SQS Cost Optimization Best Practices 2023 | Optimize Message Size | Keep message sizes under 64KB to avoid extended client library charges | **User** - Design messages to stay within the 64KB limit, use S3 for larger payloads with message pointers |  |
| Medium | COST-SQS-007 | AWS SQS Cost Optimization Best Practices 2023 | Use VPC Endpoints for Internal Traffic | Implement VPC endpoints to avoid data transfer charges for internal communication | **IaC** - Create VPC endpoints for SQS to eliminate data transfer costs for traffic within the VPC |  |
| Low | COST-SQS-008 | AWS SQS Cost Optimization Best Practices 2023 | Implement Queue Resource Tagging | Use consistent tagging strategy for cost allocation and monitoring | **IaC** - Implement comprehensive tagging strategy including cost center, environment, and project tags for accurate cost allocation |  |
| Low | COST-SQS-009 | AWS SQS Cost Optimization Best Practices 2023 | Monitor Usage with CloudWatch | Set up CloudWatch alarms to monitor unexpected usage spikes | **Platform** - Configure CloudWatch alarms for NumberOfMessagesSent, NumberOfMessagesReceived, and ApproximateNumberOfMessages metrics |  |

