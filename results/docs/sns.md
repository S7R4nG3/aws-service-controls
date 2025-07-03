# AWS Simple Notification Service
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| HIGH | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage information system accounts including account creation, enabling, modification, disabling, removal, and auditing | **Platform** - Implement IAM policies with least privilege access for SNS topics and subscriptions. Use IAM roles for service-to-service communication and regularly audit access permissions. |  |
| HIGH | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to information and system resources | **IaC** - Define SNS topic policies and subscription filters in CloudFormation/Terraform templates to enforce granular access controls and message filtering. |  |
| HIGH | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable HTTPS for SNS endpoints and implement message signing/verification for critical notifications using AWS SDK encryption features. |  |
| MEDIUM | SC-13 | NIST 800-53 Rev 5 | Cryptographic Protection | Determine the cryptographic uses and implement cryptographic controls | **IaC** - Enable server-side encryption for SNS topics using AWS KMS keys and implement message payload encryption for sensitive data. |  |
| MEDIUM | AU-2 | NIST 800-53 Rev 5 | Event Logging | Identify the types of events that the system is capable of logging in support of the audit function | **Platform** - Enable CloudTrail logging for SNS API calls and CloudWatch metrics for message delivery tracking and failure analysis. |  |
| LOW | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of user-level information contained in the system | **User** - Implement backup strategies for SNS topic configurations and subscription lists using AWS Config or custom automation scripts. |  |
| HIGH | IAM-01 | CSA Cloud Controls Matrix v4.0 | Identity and Access Management Capability | Policies and procedures shall be established to maintain baseline requirements for identity and access management capability | **Platform** - Implement multi-factor authentication for SNS administrative access and use federated identity providers for user authentication. |  |
| HIGH | EKM-01 | CSA Cloud Controls Matrix v4.0 | Encryption and Key Management | Encryption keys shall be managed throughout their lifecycle | **IaC** - Use AWS KMS customer-managed keys for SNS topic encryption and implement key rotation policies in infrastructure templates. |  |
| HIGH | STA-01 | CSA Cloud Controls Matrix v4.0 | Secure Transmission | Data in transit shall be protected using secure communications protocols | **Platform** - Configure SNS to use TLS 1.2+ for all HTTPS endpoints and validate SSL certificates for webhook subscriptions. |  |
| MEDIUM | LOG-01 | CSA Cloud Controls Matrix v4.0 | Logging and Monitoring | Comprehensive logging records shall be maintained and monitored for anomalous or suspicious activity | **IaC** - Deploy CloudWatch alarms for SNS delivery failures and implement centralized logging for message processing metrics. |  |
| MEDIUM | GRC-01 | CSA Cloud Controls Matrix v4.0 | Governance, Risk and Compliance Framework | Governance and enterprise risk management framework shall be established and maintained | **User** - Establish SNS usage policies, regular access reviews, and compliance reporting mechanisms for message handling. |  |
| HIGH | SNS.1 | AWS Foundational Security Best Practices 1.0.0 | SNS topics should be encrypted at rest using AWS KMS | Ensure SNS topics are encrypted at rest to protect sensitive message content | **IaC** - Configure KmsMasterKeyId parameter in SNS topic definitions and use customer-managed KMS keys for enhanced security control. |  |
| MEDIUM | SNS.2 | AWS Foundational Security Best Practices 1.0.0 | Delivery status logging should be enabled for notification messages sent to a platform endpoint | Enable delivery status logging to track message delivery success and failures | **Platform** - Configure delivery status attributes for each platform endpoint and integrate with CloudWatch for monitoring and alerting. |  |
| HIGH | SNS.1 | AWS Security Hub 2023.1 | SNS topics should be encrypted at rest using AWS KMS | This control checks whether Amazon SNS topics are encrypted at rest using AWS KMS | **IaC** - Enable KMS encryption in SNS topic resource definitions and ensure proper key policies are configured for topic access. |  |
| MEDIUM | SNS.2 | AWS Security Hub 2023.1 | Delivery status logging should be enabled for notification messages sent to a platform endpoint | This control checks whether delivery status logging is enabled for Amazon SNS topic subscriptions to platform endpoints | **Platform** - Enable delivery status logging attributes and configure appropriate IAM roles for SNS to write to CloudWatch Logs. |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| HIGH | COST-01 | AWS SNS Cost Optimization Best Practices 2023 | Optimize Message Filtering | Use message filtering to reduce unnecessary message deliveries and associated costs | **IaC** - Implement message filtering policies at the subscription level to ensure only relevant messages are delivered to endpoints, reducing billable message deliveries. |  |
| HIGH | COST-02 | AWS SNS Cost Optimization Best Practices 2023 | Choose Appropriate Delivery Protocols | Select cost-effective delivery protocols based on use case requirements | **User** - Use SQS subscriptions instead of HTTP/HTTPS for internal communications and batch processing to reduce per-message costs. |  |
| MEDIUM | COST-03 | AWS SNS Cost Optimization Best Practices 2023 | Implement Dead Letter Queues | Use dead letter queues to prevent repeated delivery attempts for failed messages | **IaC** - Configure dead letter queues for SNS subscriptions to avoid excessive retry charges and implement proper error handling workflows. |  |
| MEDIUM | COST-04 | AWS SNS Cost Optimization Best Practices 2023 | Monitor and Optimize Message Size | Optimize message payload sizes to reduce data transfer and processing costs | **User** - Implement message compression and avoid sending large payloads through SNS; use S3 references for large data objects instead. |  |
| MEDIUM | COST-05 | AWS SNS Cost Optimization Best Practices 2023 | Regular Subscription Audit | Regularly audit and clean up unused or inactive subscriptions | **Platform** - Implement automated monitoring to identify and remove unused subscriptions and topics to avoid unnecessary charges for inactive resources. |  |
| LOW | COST-06 | AWS SNS Cost Optimization Best Practices 2023 | Leverage Free Tier Appropriately | Understand and optimize usage within AWS Free Tier limits for development and testing | **User** - Monitor monthly usage against Free Tier limits (1 million requests, 100,000 HTTP deliveries, 1,000 email deliveries) and implement usage controls for non-production environments. |  |

