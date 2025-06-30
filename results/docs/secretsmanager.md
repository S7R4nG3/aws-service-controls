# AWS Secrets Manager
---


### CSA CCM v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | User Access Management | IAM-02 | Ensure proper identity and access management for users accessing secrets | **IaC** - Implement least privilege IAM policies using resource-based policies and condition keys for Secrets Manager access |  |
| **High** | Entitlement Establishment | EKM-01 | Establish proper entitlements for secret access based on roles and responsibilities | **IaC** - Create role-based IAM policies with specific secret ARN restrictions and condition-based access controls |  |
| **High** | Encryption and Key Management | DSI-07 | Ensure proper encryption of secrets at rest and in transit | **Platform** - Configure AWS KMS customer-managed keys for secrets encryption and enable encryption in transit using TLS |  |
| Medium | Network Monitoring | IVS-06 | Monitor network access to secrets management service | **Platform** - Enable VPC Flow Logs and configure CloudTrail for API call monitoring |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Account Management | AC-2 | Manage user accounts with access to secrets | **User** - Implement centralized account management using AWS IAM with regular access reviews and automated provisioning/deprovisioning |  |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for accessing secrets | **IaC** - Configure resource-based policies and IAM policies with explicit deny statements and condition keys |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect confidentiality and integrity of transmitted secrets | **Platform** - Use HTTPS/TLS for all API communications and implement VPC endpoints for private connectivity |  |
| **High** | Content of Audit Records | AU-3 | Generate audit records containing information for security investigations | **Platform** - Enable CloudTrail logging with detailed event information including source IP, user identity, and timestamps |  |
| Medium | Information System Backup | CP-9 | Conduct backups of secrets and recovery information | **IaC** - Configure cross-region secret replication and implement backup strategies using AWS Backup |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Secrets Manager secrets should have automatic rotation enabled | SecretsManager.1 | Ensure secrets are automatically rotated to reduce security risks | **IaC** - Configure automatic rotation using Lambda functions and specify rotation schedules in CloudFormation/Terraform |  |
| **High** | Secrets Manager secrets configured with automatic rotation should rotate successfully | SecretsManager.2 | Verify that configured automatic rotation completes successfully | **User** - Monitor rotation status through CloudWatch metrics and set up alerts for failed rotations |  |
| Medium | Remove unused Secrets Manager secrets | SecretsManager.3 | Identify and remove secrets that are no longer in use | **User** - Implement regular auditing using AWS Config rules and CloudTrail analysis to identify unused secrets |  |
| Medium | Secrets Manager secrets should be rotated within a specified number of days | SecretsManager.4 | Ensure secrets are rotated within organization-defined timeframes | **IaC** - Configure rotation schedules and monitoring using AWS Config rules and CloudWatch alarms |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Secrets Manager secrets should have automatic rotation enabled | SecretsManager.1 | Automatically rotate secrets to maintain security posture | **IaC** - Enable automatic rotation in secret configuration with appropriate Lambda rotation functions |  |
| **Critical** | IAM policies should not allow full '*' administrative privileges | IAM.1 | Restrict overly permissive policies for Secrets Manager access | **IaC** - Create specific IAM policies with explicit resource ARNs and required actions only |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable comprehensive logging for Secrets Manager API calls | **Platform** - Configure multi-region CloudTrail with S3 bucket logging and log file validation |  |
| **High** | IAM customer managed policies should not allow decryption actions on all KMS keys | KMS.1 | Restrict KMS key access for secret decryption | **IaC** - Implement specific KMS key policies with resource-level permissions for secret encryption keys |  |
| Medium | AWS Config should be enabled | Config.1 | Monitor configuration changes to Secrets Manager resources | **Platform** - Enable AWS Config with Secrets Manager resource recording and compliance rules |  |


## Operational Controls
---



## Cost Controls
---


### AWS Secrets Manager Cost Optimization 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Remove Unused Secrets | COST-01 | Regularly identify and delete secrets that are no longer in use to avoid ongoing charges | **User** - Implement automated discovery using CloudTrail logs to identify secrets not accessed within defined timeframes and schedule for deletion |
| **High** | Optimize Rotation Frequency | COST-02 | Balance security requirements with cost by optimizing rotation schedules | **IaC** - Configure rotation schedules based on risk assessment rather than default frequencies, using longer intervals where security permits |
| Medium | Consolidate Secrets | COST-03 | Combine related secrets into single secret entries where security boundaries allow | **User** - Review secret architecture to identify opportunities for consolidation, such as storing multiple database credentials in JSON format |
| Medium | Monitor Secret Usage | COST-04 | Implement monitoring to track secret usage patterns and identify cost optimization opportunities | **Platform** - Set up CloudWatch dashboards and Cost Explorer analysis to track Secrets Manager spending and usage metrics |
| Medium | Use Appropriate Replication Strategy | COST-05 | Implement cross-region replication only where required for disaster recovery | **IaC** - Configure replication based on RTO/RPO requirements and avoid unnecessary cross-region replication charges |
| Low | Implement Secret Lifecycle Management | COST-06 | Establish processes for secret lifecycle management to prevent secret sprawl | **User** - Create governance processes with regular reviews, automated cleanup, and approval workflows for new secrets |


