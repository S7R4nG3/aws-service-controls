# AWS EventBridge
---


### Cloud Security Alliance Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management | IAM-01 | Implement strong identity and access management controls for EventBridge resources | **IaC** - Configure IAM policies with least privilege principles, use resource-based policies for event buses, and implement cross-account access controls |  |
| **High** | Encryption Key Management | EKM-01 | Implement proper encryption key management for EventBridge data protection | **IaC** - Configure customer-managed KMS keys for EventBridge encryption at rest and in transit |  |
| **High** | Audit Logging | LOG-01 | Enable comprehensive audit logging for all EventBridge operations | **Platform** - Enable CloudTrail logging for EventBridge API calls and configure log retention policies |  |
| Medium | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and lifecycle management for event data | **User** - Classify event data based on sensitivity and implement appropriate retention policies |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to EventBridge resources | **IaC** - Implement resource-based policies and IAM roles with condition-based access controls |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Ensure all EventBridge communications use TLS encryption and configure VPC endpoints where required |  |
| **High** | Event Logging | AU-2 | Determine which events are to be logged by EventBridge | **IaC** - Configure comprehensive CloudTrail logging and enable EventBridge rule execution logging |  |
| **High** | Boundary Protection | SC-7 | Monitor and control communications at external boundaries | **IaC** - Implement VPC endpoints for EventBridge and configure security groups for target services |  |
| Medium | System Backup | CP-9 | Conduct backups of EventBridge configurations and rules | **User** - Implement automated backup of EventBridge rules and configurations using infrastructure as code |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | EventBridge custom event buses should have resource-based policy attached | EventBridge.1 | Custom event buses should have resource-based policies to control access | **IaC** - Attach resource-based policies to custom event buses defining allowed principals and actions |  |
| **High** | EventBridge rules should have targets configured | EventBridge.2 | EventBridge rules without targets may indicate misconfigurations | **User** - Ensure all EventBridge rules have appropriate targets configured and remove unused rules |  |
| **High** | EventBridge custom event buses should be encrypted | EventBridge.3 | Custom event buses should use encryption at rest | **IaC** - Configure KMS encryption for custom event buses using customer-managed keys |  |
| Medium | EventBridge should not allow public access | EventBridge.4 | EventBridge resources should not be publicly accessible | **IaC** - Review and restrict resource-based policies to prevent public access |  |

### AWS Security Hub 2023.04.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | EventBridge custom event buses should have a resource-based policy attached | EventBridge.1 | This control checks whether EventBridge custom event buses have resource-based policies attached | **IaC** - Create and attach resource-based policies to custom event buses with appropriate access controls |  |
| **High** | EventBridge rules should have targets | EventBridge.2 | This control checks whether EventBridge rules have targets configured | **User** - Configure appropriate targets for all EventBridge rules and remove orphaned rules |  |
| **High** | EventBridge custom event buses should be encrypted | EventBridge.3 | This control checks whether EventBridge custom event buses are encrypted | **IaC** - Enable encryption for custom event buses using AWS KMS customer-managed keys |  |
| Medium | EventBridge global endpoints should be encrypted | EventBridge.4 | This control checks whether EventBridge global endpoints use encryption | **Platform** - Ensure EventBridge global endpoints are configured with encryption enabled |  |


## Operational Controls
---



## Cost Controls
---


### AWS EventBridge Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize Event Pattern Matching | COST-01 | Use specific event patterns to reduce unnecessary rule evaluations and processing costs | **User** - Design event patterns with specific filters to minimize rule evaluations and avoid broad pattern matching |
| **High** | Implement Event Deduplication | COST-02 | Prevent duplicate events to reduce processing costs and downstream resource consumption | **User** - Implement event deduplication logic and use idempotency tokens where applicable |
| Medium | Monitor and Remove Unused Rules | COST-03 | Regularly audit and remove unused EventBridge rules to reduce costs | **User** - Implement automated monitoring to identify and remove unused rules and event buses |
| Medium | Optimize Cross-Region Replication | COST-04 | Minimize unnecessary cross-region event replication to reduce data transfer costs | **IaC** - Configure replication rules to replicate only necessary events across regions |
| Medium | Use Event Batching | COST-05 | Implement event batching for high-volume scenarios to reduce per-event costs | **User** - Configure target services to process events in batches where supported |
| Low | Implement Event Filtering at Source | COST-06 | Filter events at the source to reduce the number of events sent to EventBridge | **User** - Configure event sources to filter events before sending to EventBridge |
| Low | Monitor Event Processing Costs | COST-07 | Implement cost monitoring and alerting for EventBridge usage | **Platform** - Set up CloudWatch billing alarms and cost budgets for EventBridge service usage |


