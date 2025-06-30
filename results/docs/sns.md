# AWS Simple Notification Service (SNS)
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy | IAM-01 | Establish and maintain identity and access management policies for SNS resources | **IaC** - Implement IAM policies using CloudFormation or Terraform to define least-privilege access to SNS topics and subscriptions |  |
| **High** | User Access Provisioning | IAM-02 | Implement proper user provisioning and access controls for SNS resources | **Platform** - Use AWS IAM to create roles and policies that grant minimal necessary permissions for SNS operations |  |
| **Critical** | Encryption Key Management | EKM-01 | Implement proper encryption key management for SNS message encryption | **IaC** - Configure AWS KMS keys for SNS topic encryption and implement key rotation policies |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and lifecycle management for SNS messages | **User** - Classify message content and implement appropriate retention policies and data handling procedures |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to SNS resources | **IaC** - Implement resource-based policies and IAM policies to enforce access controls on SNS topics and subscriptions |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted SNS messages | **Platform** - Enable HTTPS endpoints for SNS subscriptions and use TLS encryption for message delivery |  |
| **High** | Event Logging | AU-2 | Identify events to be logged for SNS activities | **IaC** - Configure CloudTrail to log SNS API calls and enable SNS message delivery status logging |  |
| **High** | Protection of Information at Rest | SC-28 | Protect information at rest in SNS message storage | **IaC** - Enable server-side encryption for SNS topics using AWS KMS customer-managed keys |  |
| Medium | System Backup | CP-9 | Conduct backups of SNS configurations and subscriptions | **IaC** - Implement Infrastructure as Code to maintain SNS topic and subscription configurations as versioned backups |  |

### AWS Foundational Security Standard v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | SNS topics should be encrypted at rest using AWS KMS | SNS.1 | SNS topics should be encrypted at rest to protect sensitive message content | **IaC** - Configure KmsMasterKeyId parameter when creating SNS topics to enable encryption at rest |  |
| **High** | Delivery status logging should be enabled for notification messages sent to a topic | SNS.2 | Enable delivery status logging to monitor message delivery success and failures | **IaC** - Configure delivery status logging attributes for SNS topics to track message delivery to various endpoints |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable CloudTrail logging for SNS API calls for audit and compliance purposes | **Platform** - Configure CloudTrail to capture SNS API events across all regions where SNS is used |  |
| Medium | IAM policies should not allow full administrative privileges | IAM.1 | Implement least privilege access for SNS resources | **IaC** - Create specific IAM policies that grant only necessary SNS permissions rather than using wildcard permissions |  |

### AWS Security Hub Current
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | SNS topics should be encrypted at rest using AWS KMS | SNS.1 | Ensure SNS topics use KMS encryption for data protection at rest | **IaC** - Apply KMS encryption to SNS topics using customer-managed keys with appropriate key policies |  |
| **High** | Delivery status logging should be enabled for notification messages sent to a topic | SNS.2 | Enable message delivery status logging for operational visibility | **IaC** - Configure SNS topic attributes to enable delivery status logging for HTTP/HTTPS, Lambda, SQS, and other endpoints |  |
| **High** | AWS Config should be enabled | Config.1 | Enable AWS Config to monitor SNS resource configuration changes | **Platform** - Configure AWS Config rules to monitor SNS topic configuration compliance and changes |  |
| Medium | A log metric filter and alarm should exist for usage of root user | CloudWatch.1 | Monitor root user access to SNS resources | **IaC** - Create CloudWatch alarms and metric filters to detect root user SNS API usage |  |


## Operational Controls
---



## Cost Controls
---


### AWS SNS Cost Optimization Best Practices Current
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize Message Batching | COST-01 | Use message batching where possible to reduce the number of API calls and associated costs | **User** - Implement batch publishing using PublishBatch API to send up to 10 messages in a single request |
| **High** | Right-size Message Payloads | COST-02 | Optimize message size to avoid unnecessary charges for oversized messages | **User** - Keep message payloads under 256KB and use message attributes efficiently to minimize billing increments |
| Medium | Implement Dead Letter Queues | COST-03 | Use dead letter queues to avoid repeated delivery attempts and associated costs | **IaC** - Configure dead letter queues for SNS subscriptions to prevent costly retry cycles |
| Medium | Monitor and Optimize Cross-Region Data Transfer | COST-04 | Minimize cross-region data transfer costs by optimizing SNS topic and subscription placement | **IaC** - Deploy SNS topics in the same region as subscribers to avoid cross-region data transfer charges |
| Medium | Use Message Filtering | COST-05 | Implement message filtering to reduce unnecessary message deliveries and associated costs | **IaC** - Configure subscription filter policies to deliver only relevant messages to subscribers |
| Low | Optimize SMS and Email Delivery | COST-06 | Optimize SMS and email delivery patterns to minimize premium messaging costs | **User** - Use appropriate message prioritization and avoid sending SMS to premium numbers unnecessarily |
| Low | Implement Cost Monitoring and Alerting | COST-07 | Set up cost monitoring and alerting for SNS usage to prevent unexpected charges | **Platform** - Configure AWS Cost Explorer and CloudWatch billing alarms to monitor SNS costs and usage patterns |


