# AWS EventBridge
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-02 | CSA Cloud Controls Matrix v4.0.6 | User Access Management | Establish and maintain proper identity and access management controls for EventBridge resources | **IaC** - Implement IAM policies with least privilege principles, define resource-based policies for event buses, and establish proper role-based access controls |  |
| **High** | DSI-02 | CSA Cloud Controls Matrix v4.0.6 | Data Security and Encryption | Protect sensitive data within EventBridge events and implement proper data classification | **Platform** - Enable encryption at rest and in transit, implement event content filtering, and establish data retention policies |  |
| **High** | IVS-01 | CSA Cloud Controls Matrix v4.0.6 | Network Security | Secure the underlying infrastructure and network communications for EventBridge | **IaC** - Configure VPC endpoints for private connectivity, implement network ACLs, and establish secure communication channels |  |
| Medium | BCR-02 | CSA Cloud Controls Matrix v4.0.6 | Business Continuity Planning | Ensure EventBridge service availability and implement disaster recovery procedures | **Platform** - Configure cross-region replication for critical event buses, implement dead letter queues, and establish backup procedures |  |
| Medium | STA-01 | CSA Cloud Controls Matrix v4.0.6 | Supply Chain Management | Maintain visibility and control over EventBridge integrations and third-party connections | **User** - Document all event sources and targets, maintain inventory of integrations, and establish approval processes for new connections |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage EventBridge access through proper account management procedures | **IaC** - Implement automated account provisioning/deprovisioning, regular access reviews, and principle of least privilege for EventBridge permissions |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information in EventBridge | **Platform** - Enable TLS encryption for all EventBridge communications, implement message authentication codes, and use encrypted channels |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Implement comprehensive auditing for EventBridge activities | **Platform** - Enable CloudTrail logging for EventBridge API calls, configure detailed event logging, and implement log integrity protection |  |
| **High** | SC-7 | NIST 800-53 Rev 5 | Boundary Protection | Implement network boundary protection for EventBridge communications | **IaC** - Configure VPC endpoints, implement security groups with restrictive rules, and establish network segmentation |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Implement backup procedures for EventBridge configurations and critical data | **IaC** - Backup event bus configurations, rules, and targets using Infrastructure as Code, implement cross-region backup strategies |  |
| **Critical** | EventBridge.1 | AWS Foundational Security Standard 1.0.0 | EventBridge custom event buses should have resource-based policy attached | Custom EventBridge event buses should have resource-based policies to control access | **IaC** - Attach resource-based policies to custom event buses defining who can publish events, create rules, or manage the bus |  |
| **High** | EventBridge.2 | AWS Foundational Security Standard 1.0.0 | EventBridge rules should have targets configured | EventBridge rules should have at least one target to ensure events are properly processed | **IaC** - Configure appropriate targets for all EventBridge rules and implement dead letter queues for failed event processing |  |
| **High** | EventBridge.3 | AWS Foundational Security Standard 1.0.0 | EventBridge should use encryption at rest | EventBridge event buses should be encrypted at rest using AWS KMS | **Platform** - Enable server-side encryption for EventBridge using AWS managed keys or customer managed KMS keys |  |
| Medium | EventBridge.4 | AWS Foundational Security Standard 1.0.0 | EventBridge API calls should be logged | All EventBridge API calls should be logged for security monitoring and compliance | **Platform** - Enable AWS CloudTrail to log all EventBridge API calls and configure log analysis for security events |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-EB-01 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Optimize Event Filtering | Implement precise event filtering to reduce unnecessary event processing and associated costs | **User** - Use specific event patterns in rules to filter events at the source, reducing downstream processing costs and improving efficiency |  |
| **High** | COST-EB-02 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Monitor and Optimize Cross-Region Data Transfer | Minimize cross-region data transfer costs by optimizing event routing and replication strategies | **IaC** - Design event architectures to minimize cross-region transfers, use regional event buses where possible, and implement cost-effective replication strategies |  |
| Medium | COST-EB-03 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Implement Event Batching | Use event batching where appropriate to reduce the number of individual event invocations and associated costs | **User** - Configure targets like Lambda functions and SQS queues to handle batched events when possible to reduce per-invocation costs |  |
| Medium | COST-EB-04 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Right-size Target Resources | Ensure target resources are appropriately sized for the event volume and processing requirements | **IaC** - Monitor target resource utilization and adjust Lambda function memory, ECS task sizes, and other compute resources based on actual event processing needs |  |
| Medium | COST-EB-05 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Implement Dead Letter Queues Efficiently | Configure dead letter queues to minimize storage costs while maintaining error handling capabilities | **IaC** - Set appropriate message retention periods for dead letter queues, implement automated cleanup processes, and monitor DLQ usage to optimize storage costs |  |
| Low | COST-EB-06 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Use Custom Event Buses Strategically | Evaluate the cost-benefit of custom event buses versus using the default event bus | **User** - Assess whether custom event buses provide sufficient value for the additional management overhead and costs, consolidate where appropriate |  |
| Low | COST-EB-07 | AWS EventBridge Cost Optimization Best Practices 2024.1 | Monitor Event Volume and Costs | Implement comprehensive monitoring of EventBridge usage and costs to identify optimization opportunities | **Platform** - Set up CloudWatch metrics and billing alerts for EventBridge usage, create dashboards to track event volumes and costs over time |  |

