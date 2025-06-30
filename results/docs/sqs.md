# AWS Simple Queue Service (SQS)
---


### CSA Cloud Controls Matrix v5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | User Access Authorization | IAM-02 | Verify user access authorization and user access reviews for entitlements | **IaC** - Implement IAM policies with least privilege access to SQS queues, use resource-based policies to control queue access |  |
| HIGH | Data Classification | DSI-01 | Data and objects must be classified based on organizational requirements | **User** - Classify data before sending to SQS queues, implement appropriate tagging strategy for queues based on data sensitivity |  |
| HIGH | Data Handling/Labeling/Security Policy | DSI-02 | Procedures for data handling, labeling, and security policies must be established | **User** - Establish data handling procedures for SQS messages, implement message labeling standards for sensitive data |  |
| HIGH | Encryption Key Management | EKM-01 | Encryption keys must be managed throughout their lifecycle | **IaC** - Configure SQS with server-side encryption using AWS KMS, implement key rotation policies |  |
| MEDIUM | Network Segmentation | IVS-01 | Network environments must be segregated and traffic filtered | **IaC** - Configure VPC endpoints for SQS to keep traffic within VPC, implement proper security group rules |  |
| MEDIUM | Audit Logging | LOG-01 | Audit logs must be generated for all security events | **Platform** - Enable CloudTrail logging for SQS API calls, configure CloudWatch metrics and alarms |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Implement IAM policies with condition keys to enforce access controls based on time, IP address, and MFA |  |
| HIGH | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Use HTTPS/TLS for all SQS communications, enable server-side encryption for messages at rest |  |
| HIGH | Event Logging | AU-2 | Identify the types of events that the system is capable of logging | **Platform** - Configure CloudTrail to log all SQS API calls, enable CloudWatch Logs for application-level logging |  |
| HIGH | Identification and Authentication | IA-2 | Uniquely identify and authenticate organizational users | **IaC** - Require MFA for sensitive SQS operations, use IAM roles for service-to-service authentication |  |
| MEDIUM | Boundary Protection | SC-7 | Monitor, control, and protect communications at the external boundary | **IaC** - Implement VPC endpoints for SQS, configure security groups and NACLs appropriately |  |
| MEDIUM | System Backup | CP-9 | Conduct backups of user-level and system-level information | **User** - Implement message backup strategies using dead letter queues, configure message retention policies |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Amazon SQS queues should be encrypted at rest | SQS.1 | SQS queues should use server-side encryption to protect sensitive data | **IaC** - Configure SQS queues with KmsMasterKeyId parameter to enable encryption at rest |  |
| HIGH | Amazon SQS queues should not be publicly accessible | SQS.2 | SQS queue policies should not allow public access | **IaC** - Review and restrict SQS queue policies to prevent public access, use IAM policies for access control |  |
| MEDIUM | Amazon SQS dead letter queues should be configured | SQS.3 | Configure dead letter queues to handle message processing failures | **IaC** - Configure dead letter queues with appropriate maxReceiveCount to handle failed message processing |  |

### AWS Security Hub Current
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | SQS queues should be encrypted at rest | SQS.1 | This control checks whether SQS queues are encrypted at rest | **IaC** - Enable server-side encryption using AWS managed keys or customer managed keys |  |
| HIGH | SQS queues should not be publicly accessible | SQS.2 | This control checks whether SQS queues have resource-based policies that allow public access | **IaC** - Review queue policies to ensure they don't contain wildcard principals or public access statements |  |
| MEDIUM | SQS dead letter queues should be configured | SQS.3 | This control checks whether SQS queues have dead letter queues configured | **IaC** - Configure RedrivePolicy with dead letter queue ARN and maxReceiveCount |  |
| HIGH | CloudTrail should be enabled | CloudTrail.1 | This control checks whether CloudTrail is enabled for SQS API logging | **Platform** - Enable CloudTrail in all regions to capture SQS API calls and management events |  |


## Operational Controls
---



## Cost Controls
---


### AWS SQS Cost Optimization Best Practices Current
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| HIGH | Use Long Polling | COST-01 | Reduce the number of empty receives and lower costs by using long polling | **IaC** - Set ReceiveMessageWaitTimeSeconds to 1-20 seconds to enable long polling and reduce API calls |
| HIGH | Optimize Message Batching | COST-02 | Use batch operations to reduce the number of API calls and lower costs | **User** - Implement SendMessageBatch and DeleteMessageBatch operations to process up to 10 messages per API call |
| MEDIUM | Set Appropriate Message Retention | COST-03 | Configure message retention period based on business requirements to avoid unnecessary storage costs | **IaC** - Set MessageRetentionPeriod to the minimum required duration (60 seconds to 14 days) |
| MEDIUM | Monitor Queue Metrics | COST-04 | Monitor queue metrics to identify cost optimization opportunities | **Platform** - Set up CloudWatch metrics monitoring for NumberOfEmptyReceives, ApproximateNumberOfMessages, and other cost-related metrics |
| MEDIUM | Use Standard Queues When Possible | COST-05 | Use standard queues instead of FIFO queues when strict ordering is not required | **IaC** - Deploy standard queues (lower cost) unless FIFO ordering and deduplication are business requirements |
| LOW | Implement Queue Cleanup Policies | COST-06 | Regularly clean up unused queues to avoid unnecessary charges | **User** - Implement automated processes to identify and delete unused queues, set up tagging for queue lifecycle management |
| LOW | Optimize Visibility Timeout | COST-07 | Set appropriate visibility timeout to reduce duplicate processing and associated costs | **IaC** - Configure VisibilityTimeoutSeconds based on message processing time to minimize reprocessing costs |


