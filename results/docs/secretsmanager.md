# AWS Secrets Manager
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | IAM-01 | CSA Cloud Controls Matrix v4.0.10 | Identity and Access Management | Ensure proper identity and access management controls are implemented for Secrets Manager access | **IaC** - Implement IAM policies with least privilege access, enforce resource-based policies, and use IAM roles for cross-account access |  |
| **High** | EKM-01 | CSA Cloud Controls Matrix v4.0.10 | Encryption Key Management | Implement proper encryption key management for secrets at rest and in transit | **Platform** - Configure AWS KMS customer managed keys for secret encryption, enable automatic key rotation, and implement key access policies |  |
| **High** | DCS-01 | CSA Cloud Controls Matrix v4.0.10 | Data Classification and Handling | Classify and handle secrets according to data sensitivity levels | **User** - Tag secrets with appropriate classification levels, implement naming conventions, and establish handling procedures |  |
| Medium | BCR-01 | CSA Cloud Controls Matrix v4.0.10 | Business Continuity Management | Ensure secrets are available for business continuity and disaster recovery | **IaC** - Enable cross-region replication for critical secrets, implement backup strategies, and document recovery procedures |  |
| Medium | AIS-01 | CSA Cloud Controls Matrix v4.0.10 | Audit Independent Systems | Implement comprehensive logging and auditing for all Secrets Manager operations | **Platform** - Enable CloudTrail logging for Secrets Manager API calls, configure log retention, and integrate with SIEM systems |  |
| **High** | AC-2 | NIST SP 800-53 Rev 5 | Account Management | Manage accounts with access to secrets through proper provisioning and deprovisioning | **IaC** - Implement automated account lifecycle management, regular access reviews, and role-based access controls |  |
| **High** | AC-3 | NIST SP 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to Secrets Manager | **IaC** - Configure resource-based policies, implement condition-based access, and use VPC endpoints for network isolation |  |
| **High** | SC-13 | NIST SP 800-53 Rev 5 | Cryptographic Protection | Implement cryptographic mechanisms to protect secrets | **Platform** - Use AWS KMS encryption with customer managed keys, implement encryption in transit with TLS 1.2+, and enable automatic rotation |  |
| Medium | AU-2 | NIST SP 800-53 Rev 5 | Event Logging | Log events related to secret access and management | **Platform** - Enable comprehensive CloudTrail logging, configure CloudWatch for monitoring, and implement log aggregation |  |
| Medium | CP-9 | NIST SP 800-53 Rev 5 | Information System Backup | Conduct backups of secrets and related configuration information | **IaC** - Implement cross-region replication, export secrets for backup purposes, and maintain disaster recovery procedures |  |
| Medium | SI-4 | NIST SP 800-53 Rev 5 | System Monitoring | Monitor the system for unauthorized access and usage | **Platform** - Configure CloudWatch alarms for unusual access patterns, implement GuardDuty integration, and set up notification systems |  |
| Medium | SecretsManager.1 | AWS Foundational Security Best Practices 1.0.0 | Secrets Manager secrets should have automatic rotation enabled | Secrets should be automatically rotated to reduce the risk of compromise | **IaC** - Configure automatic rotation schedules, implement Lambda rotation functions, and test rotation procedures |  |
| Medium | SecretsManager.2 | AWS Foundational Security Best Practices 1.0.0 | Secrets Manager secrets configured with automatic rotation should rotate successfully | Ensure that automatic rotation completes successfully without errors | **Platform** - Monitor rotation status, configure alerting for failed rotations, and implement retry mechanisms |  |
| Medium | SecretsManager.4 | AWS Foundational Security Best Practices 1.0.0 | Secrets Manager secrets should be rotated within a specified number of days | Secrets should be rotated within organization-defined timeframes | **IaC** - Configure rotation schedules based on secret sensitivity, implement compliance monitoring, and automate rotation enforcement |  |
| Low | SecretsManager.3 | AWS Foundational Security Best Practices 1.0.0 | Remove unused Secrets Manager secrets | Unused secrets should be removed to reduce attack surface | **User** - Regularly audit secret usage, implement automated cleanup processes, and maintain secret inventory |  |
| Medium | SecretsManager.1 | AWS Security Hub 2024.1.0 | Secrets Manager secrets should have automatic rotation enabled | This control checks whether a secret stored in AWS Secrets Manager is configured with automatic rotation | **IaC** - Enable automatic rotation in Secrets Manager configuration, create rotation Lambda functions, and schedule regular rotations |  |
| Medium | SecretsManager.2 | AWS Security Hub 2024.1.0 | Secrets Manager secrets configured with automatic rotation should rotate successfully | This control checks whether an AWS Secrets Manager secret rotated successfully based on the rotation schedule | **Platform** - Monitor CloudWatch metrics for rotation success, configure SNS notifications for failures, and implement automated remediation |  |
| Medium | SecretsManager.4 | AWS Security Hub 2024.1.0 | Secrets Manager secrets should be rotated within a specified number of days | This control checks whether your secrets have been rotated within the specified number of days | **IaC** - Configure rotation policies with appropriate timeframes, implement compliance monitoring, and automate rotation scheduling |  |
| Low | SecretsManager.3 | AWS Security Hub 2024.1.0 | Remove unused Secrets Manager secrets | This control checks whether your secrets have been accessed within a specified number of days | **User** - Implement secret usage tracking, create automated cleanup policies, and maintain documentation of active secrets |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Secrets Manager Cost Optimization 2023 | Remove Unused Secrets | Regularly identify and delete secrets that are no longer being used to avoid unnecessary charges | **User** - Implement automated monitoring of secret usage patterns, set up alerts for unused secrets, and establish regular cleanup processes |  |
| Medium | COST-02 | AWS Secrets Manager Cost Optimization 2023 | Optimize Secret Storage | Minimize the size of stored secrets to reduce storage costs | **User** - Store only essential secret data, avoid storing large configuration files as secrets, and use references where appropriate |  |
| Medium | COST-03 | AWS Secrets Manager Cost Optimization 2023 | Efficient API Usage | Optimize API calls to reduce per-request charges through caching and batching | **User** - Implement client-side caching for frequently accessed secrets, use batch operations where possible, and avoid unnecessary API calls |  |
| Medium | COST-04 | AWS Secrets Manager Cost Optimization 2023 | Right-size Rotation Frequency | Balance security requirements with cost by optimizing rotation frequency | **IaC** - Configure rotation schedules based on secret sensitivity and compliance requirements, avoid over-rotation of low-risk secrets |  |
| Low | COST-05 | AWS Secrets Manager Cost Optimization 2023 | Monitor and Budget | Implement cost monitoring and budgeting for Secrets Manager usage | **Platform** - Set up CloudWatch billing alarms, use AWS Cost Explorer to track Secrets Manager costs, and implement budget controls |  |
| Low | COST-06 | AWS Secrets Manager Cost Optimization 2023 | Regional Optimization | Optimize cross-region replication based on actual disaster recovery needs | **IaC** - Evaluate cross-region replication requirements, replicate only critical secrets, and choose cost-effective regions |  |

