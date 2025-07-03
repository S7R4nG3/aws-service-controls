# AWS EventBridge Scheduler
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| HIGH | IAM-02 | CSA Cloud Controls Matrix v4.0 | User ID Credentials | Ensure proper identity and access management controls for EventBridge Scheduler resources | **IaC** - Implement IAM policies with least privilege access using resource-based policies and condition keys for EventBridge Scheduler |  |
| HIGH | EKM-01 | CSA Cloud Controls Matrix v4.0 | Encryption & Key Management | Ensure encryption of EventBridge Scheduler data at rest and in transit | **IaC** - Configure KMS encryption for EventBridge Scheduler using customer-managed keys and enable encryption for all schedule data |  |
| MEDIUM | AIS-01 | CSA Cloud Controls Matrix v4.0 | Application & Interface Security | Secure application interfaces and APIs for EventBridge Scheduler | **User** - Implement API authentication and authorization controls, use HTTPS endpoints, and validate all inputs to EventBridge Scheduler APIs |  |
| MEDIUM | DSI-01 | CSA Cloud Controls Matrix v4.0 | Data Security & Information Lifecycle Management | Implement data classification and lifecycle management for EventBridge Scheduler | **Platform** - Configure data retention policies and implement data classification tags for EventBridge Scheduler schedules and metadata |  |
| LOW | SEF-01 | CSA Cloud Controls Matrix v4.0 | Security Incident Management, E-Discovery & Cloud Forensics | Enable comprehensive logging and monitoring for EventBridge Scheduler | **IaC** - Enable CloudTrail logging, CloudWatch monitoring, and configure log retention for EventBridge Scheduler activities |  |
| HIGH | AC-2 | NIST SP 800-53 Rev 5 | Account Management | Implement proper account management for EventBridge Scheduler access | **IaC** - Create dedicated service accounts with minimal permissions and implement automated account lifecycle management for EventBridge Scheduler |  |
| HIGH | SC-8 | NIST SP 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission to and from EventBridge Scheduler | **Platform** - Ensure all EventBridge Scheduler API calls use TLS 1.2 or higher and implement VPC endpoints for secure communication |  |
| MEDIUM | AU-2 | NIST SP 800-53 Rev 5 | Event Logging | Configure comprehensive event logging for EventBridge Scheduler | **IaC** - Enable CloudTrail data events for EventBridge Scheduler and configure CloudWatch Logs for all scheduler activities |  |
| MEDIUM | SI-4 | NIST SP 800-53 Rev 5 | System Monitoring | Implement continuous monitoring for EventBridge Scheduler | **IaC** - Configure CloudWatch alarms for EventBridge Scheduler metrics and implement automated response to security events |  |
| LOW | CP-9 | NIST SP 800-53 Rev 5 | System Backup | Implement backup and recovery procedures for EventBridge Scheduler configurations | **User** - Implement automated backup of EventBridge Scheduler configurations and maintain recovery procedures documentation |  |
| HIGH | CloudTrail.1 | AWS Foundational Security Standard v1.0.0 | CloudTrail should be enabled and configured with at least one multi-Region trail | Enable CloudTrail logging for EventBridge Scheduler API calls | **Platform** - Configure multi-region CloudTrail to capture all EventBridge Scheduler API calls and management events |  |
| MEDIUM | IAM.3 | AWS Foundational Security Standard v1.0.0 | IAM users' access keys should be rotated every 90 days or less | Implement access key rotation for EventBridge Scheduler service accounts | **User** - Implement automated access key rotation for service accounts accessing EventBridge Scheduler and use IAM roles where possible |  |
| MEDIUM | KMS.1 | AWS Foundational Security Standard v1.0.0 | IAM customer managed policies should not allow decryption actions on all KMS keys | Restrict KMS key access for EventBridge Scheduler encryption | **IaC** - Create specific KMS key policies for EventBridge Scheduler encryption and limit decrypt permissions to specific keys |  |
| LOW | Config.1 | AWS Foundational Security Standard v1.0.0 | AWS Config should be enabled | Enable AWS Config to track EventBridge Scheduler configuration changes | **Platform** - Enable AWS Config in all regions to track configuration changes to EventBridge Scheduler resources |  |
| HIGH | EventBridge.4 | AWS Security Hub v2.0 | EventBridge should be encrypted | EventBridge Scheduler data should be encrypted at rest and in transit | **IaC** - Configure KMS encryption for EventBridge Scheduler and ensure all data is encrypted using customer-managed keys |  |
| MEDIUM | CloudWatch.1 | AWS Security Hub v2.0 | CloudWatch should have log group retention set | Configure log retention for EventBridge Scheduler CloudWatch logs | **IaC** - Set appropriate log retention periods for EventBridge Scheduler CloudWatch log groups based on compliance requirements |  |
| LOW | Lambda.2 | AWS Security Hub v2.0 | Lambda functions should have a dead letter queue configured | Configure dead letter queues for Lambda functions triggered by EventBridge Scheduler | **IaC** - Configure DLQ for Lambda functions invoked by EventBridge Scheduler to handle failed executions |  |
| LOW | SNS.1 | AWS Security Hub v2.0 | SNS topics should be encrypted at rest | Encrypt SNS topics used with EventBridge Scheduler | **IaC** - Enable KMS encryption for SNS topics that receive notifications from EventBridge Scheduler |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| HIGH | COST-01 | AWS EventBridge Scheduler Cost Optimization 2024 | Optimize Schedule Frequency | Right-size schedule frequency to avoid unnecessary invocations | **User** - Review and optimize schedule frequencies, use cron expressions efficiently, and avoid over-scheduling to minimize request charges |  |
| HIGH | COST-02 | AWS EventBridge Scheduler Cost Optimization 2024 | Implement Schedule Lifecycle Management | Remove unused or obsolete schedules to reduce ongoing costs | **IaC** - Implement automated cleanup of unused schedules using tags and lifecycle policies, and regularly audit active schedules |  |
| MEDIUM | COST-03 | AWS EventBridge Scheduler Cost Optimization 2024 | Use Flexible Time Windows | Implement flexible time windows to optimize resource utilization | **User** - Configure flexible time windows for non-critical schedules to allow AWS to optimize execution timing and reduce costs |  |
| MEDIUM | COST-04 | AWS EventBridge Scheduler Cost Optimization 2024 | Monitor and Alert on Usage | Implement cost monitoring and alerting for EventBridge Scheduler usage | **IaC** - Set up CloudWatch billing alarms and AWS Budgets to monitor EventBridge Scheduler costs and usage patterns |  |
| LOW | COST-05 | AWS EventBridge Scheduler Cost Optimization 2024 | Optimize Target Service Costs | Optimize costs of services invoked by EventBridge Scheduler | **User** - Review and optimize the cost of target services (Lambda, SNS, SQS) invoked by EventBridge Scheduler schedules |  |

