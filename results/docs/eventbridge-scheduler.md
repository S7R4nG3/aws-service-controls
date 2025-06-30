# AWS EventBridge Scheduler
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Identity and Access Management | IAM-01 | Implement strong identity and access management controls for EventBridge Scheduler resources | **IaC** - Configure IAM policies with least privilege access, implement role-based access control for schedule creation, modification, and deletion operations |  |
| HIGH | User Access Provisioning | IAM-02 | Ensure proper user access provisioning and deprovisioning for EventBridge Scheduler | **Platform** - Implement automated user lifecycle management with regular access reviews and time-based access controls |  |
| HIGH | Data Security and Information Lifecycle Management | DSI-01 | Protect schedule data and configuration information throughout its lifecycle | **IaC** - Enable encryption at rest and in transit for schedule data, implement data classification and retention policies |  |
| HIGH | Cryptography, Encryption & Key Management | CEK-01 | Implement proper encryption and key management for EventBridge Scheduler | **Platform** - Use AWS KMS for encryption keys, implement key rotation policies, and ensure proper key access controls |  |
| MEDIUM | Infrastructure & Virtualization Security | IVS-01 | Secure the underlying infrastructure supporting EventBridge Scheduler | **Platform** - Implement network security controls, VPC endpoints, and secure communication channels |  |
| MEDIUM | Security Testing & Audit | STA-01 | Regular security testing and auditing of EventBridge Scheduler configurations | **User** - Conduct regular security assessments, penetration testing, and configuration audits |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Account Management | AC-2 | Manage user accounts with access to EventBridge Scheduler | **Platform** - Implement centralized account management with automated provisioning and deprovisioning processes |  |
| HIGH | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to EventBridge Scheduler | **IaC** - Configure IAM policies with explicit deny statements and resource-based policies for granular access control |  |
| HIGH | Least Privilege | AC-6 | Employ the principle of least privilege for EventBridge Scheduler access | **IaC** - Implement minimal necessary permissions for schedule operations and regular privilege reviews |  |
| HIGH | Event Logging | AU-2 | Log security-relevant events for EventBridge Scheduler | **Platform** - Enable CloudTrail logging for all EventBridge Scheduler API calls and configuration changes |  |
| HIGH | Audit Record Review | AU-6 | Review and analyze audit records for EventBridge Scheduler | **User** - Implement automated log analysis and regular manual review of audit logs for anomalies |  |
| HIGH | Boundary Protection | SC-7 | Implement network boundary protection for EventBridge Scheduler | **IaC** - Configure VPC endpoints, security groups, and NACLs to control network access |  |
| HIGH | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Enforce TLS encryption for all API communications and data in transit |  |
| HIGH | Cryptographic Protection | SC-13 | Implement cryptographic protection for EventBridge Scheduler data | **IaC** - Configure AWS KMS encryption for schedule data and implement proper key management |  |
| MEDIUM | System Monitoring | SI-4 | Monitor the EventBridge Scheduler system for security events | **Platform** - Implement CloudWatch monitoring and alerting for suspicious activities and performance anomalies |  |
| MEDIUM | System Backup | CP-9 | Conduct backups of EventBridge Scheduler configurations | **User** - Implement automated backup procedures for schedule configurations and maintain recovery documentation |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | EventBridge custom bus should have a resource-based policy attached | EventBridge.1 | EventBridge custom buses should have resource-based policies to control access | **IaC** - Implement resource-based policies on EventBridge custom buses to restrict cross-account access and define allowed principals |  |
| MEDIUM | EventBridge custom bus should have a dead letter queue configured | EventBridge.2 | Configure dead letter queues for EventBridge Scheduler to handle failed invocations | **IaC** - Configure SQS dead letter queues for EventBridge Scheduler targets to capture failed invocations and enable retry mechanisms |  |
| HIGH | IAM policies should not allow full administrative privileges | IAM.1 | Avoid granting full administrative privileges for EventBridge Scheduler operations | **IaC** - Create specific IAM policies for EventBridge Scheduler operations with granular permissions instead of using wildcard actions |  |
| MEDIUM | IAM users should not have IAM policies attached | IAM.2 | Use IAM roles and groups instead of attaching policies directly to users | **Platform** - Implement role-based access management and group-based policy assignment for EventBridge Scheduler access |  |
| MEDIUM | IAM users' access keys should be rotated every 90 days or less | IAM.3 | Regularly rotate access keys used for EventBridge Scheduler API access | **User** - Implement automated access key rotation policies and use temporary credentials where possible |  |
| HIGH | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable CloudTrail to log EventBridge Scheduler API calls across all regions | **Platform** - Configure multi-region CloudTrail with S3 bucket logging and log file validation for comprehensive audit trails |  |
| LOW | CloudWatch should have a log group retention policy set | CloudWatch.1 | Set appropriate retention policies for EventBridge Scheduler CloudWatch logs | **IaC** - Configure log group retention policies based on compliance requirements and cost considerations |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | EventBridge Scheduler should use customer managed KMS keys | EventBridge.4 | Use customer managed KMS keys for EventBridge Scheduler encryption | **IaC** - Configure EventBridge Scheduler to use customer managed KMS keys with proper key policies and rotation |  |
| MEDIUM | EventBridge Scheduler should have cross-region replication enabled | EventBridge.5 | Enable cross-region replication for EventBridge Scheduler for disaster recovery | **IaC** - Implement cross-region schedule replication and failover mechanisms for business continuity |  |
| HIGH | AWS Config should be enabled | Config.1 | Enable AWS Config to track EventBridge Scheduler configuration changes | **Platform** - Configure AWS Config to monitor EventBridge Scheduler resources and maintain configuration history |  |
| HIGH | GuardDuty should be enabled | GuardDuty.1 | Enable GuardDuty to detect threats related to EventBridge Scheduler | **Platform** - Enable GuardDuty across all regions where EventBridge Scheduler is deployed |  |
| HIGH | Lambda functions should prohibit public access | Lambda.1 | Ensure Lambda functions invoked by EventBridge Scheduler are not publicly accessible | **IaC** - Configure Lambda function policies to restrict public access and implement proper resource-based policies |  |
| MEDIUM | Lambda functions should use supported runtimes | Lambda.2 | Use supported Lambda runtimes for functions invoked by EventBridge Scheduler | **User** - Regularly update Lambda functions to use supported runtimes and implement runtime upgrade policies |  |


## Operational Controls
---



## Cost Controls
---


### AWS EventBridge Scheduler Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| HIGH | Optimize Schedule Frequency | COST-01 | Reduce unnecessary schedule invocations to minimize costs | **User** - Review schedule frequencies and consolidate similar schedules, use appropriate time intervals based on business requirements |
| HIGH | Implement Schedule Lifecycle Management | COST-02 | Automatically delete or disable unused schedules to reduce costs | **IaC** - Implement automated lifecycle policies to identify and remove inactive schedules, use tagging for schedule management |
| MEDIUM | Use Flexible Time Windows | COST-03 | Implement flexible time windows for non-critical schedules to optimize resource usage | **User** - Configure flexible time windows for schedules that don't require exact timing to allow for cost optimization |
| MEDIUM | Monitor and Alert on Usage Patterns | COST-04 | Implement monitoring and alerting for EventBridge Scheduler usage and costs | **Platform** - Set up CloudWatch dashboards and billing alerts to monitor EventBridge Scheduler usage patterns and costs |
| MEDIUM | Optimize Target Selection | COST-05 | Choose cost-effective target services for EventBridge Scheduler invocations | **User** - Evaluate target service costs and choose appropriate compute options (Lambda vs ECS vs Step Functions) based on workload requirements |
| MEDIUM | Implement Batch Processing | COST-06 | Batch similar operations to reduce the number of individual schedule invocations | **User** - Group related tasks into batch operations and reduce schedule frequency by processing multiple items per invocation |
| LOW | Use Appropriate Retry Policies | COST-07 | Configure optimal retry policies to avoid unnecessary costs from excessive retries | **IaC** - Configure appropriate retry counts and backoff strategies to minimize costs while maintaining reliability |
| LOW | Implement Resource Tagging | COST-08 | Use consistent tagging for cost allocation and tracking of EventBridge Scheduler resources | **IaC** - Implement comprehensive tagging strategy for schedules and related resources to enable cost tracking and optimization |


