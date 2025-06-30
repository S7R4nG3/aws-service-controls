# AWS Step Functions
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management Policy and Procedures | IAM-01 | Establish and maintain identity and access management policies for Step Functions resources | **IaC** - Implement IAM policies with least privilege access, define roles for Step Functions execution, and establish proper resource-based policies |  |
| **High** | Data Classification | DCS-01 | Classify data processed by Step Functions workflows and apply appropriate protection measures | **User** - Identify sensitive data in state machines, implement data sanitization in workflow definitions, and ensure proper data handling procedures |  |
| **High** | Encryption Key Management | EKM-01 | Implement proper encryption key management for Step Functions data at rest and in transit | **IaC** - Configure KMS keys for encrypting Step Functions logs and state data, implement key rotation policies, and manage key access permissions |  |
| Medium | Audit Logging Policy and Procedures | LOG-01 | Enable comprehensive logging for Step Functions activities and state changes | **IaC** - Enable CloudWatch Logs for Step Functions, configure CloudTrail for API logging, and implement log retention policies |  |
| Medium | Network Security Policy | NET-01 | Implement network security controls for Step Functions VPC endpoints and service communications | **IaC** - Configure VPC endpoints for Step Functions, implement security groups and NACLs, and ensure encrypted communications |  |

### NIST SP 800-53 Rev. 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Step Functions resources | **IaC** - Implement resource-based policies, IAM roles with least privilege, and condition-based access controls for Step Functions |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information in Step Functions workflows | **Platform** - Leverage AWS native encryption in transit, ensure HTTPS endpoints for integrations, and validate SSL/TLS configurations |  |
| **High** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest in Step Functions | **IaC** - Enable encryption for Step Functions state data and logs using AWS KMS, configure appropriate encryption keys and access policies |  |
| Medium | Event Logging | AU-2 | Identify the types of events to be logged for Step Functions activities | **IaC** - Configure CloudWatch Logs for execution history, enable CloudTrail for API calls, and implement structured logging for workflow events |  |
| Medium | System Monitoring | SI-4 | Monitor Step Functions for attacks, intrusions, and unauthorized activities | **IaC** - Implement CloudWatch metrics and alarms, configure AWS Config rules for compliance monitoring, and set up security event detection |  |
| Low | System Backup | CP-9 | Conduct backups of Step Functions configurations and state machine definitions | **User** - Version control state machine definitions, backup IAM policies and configurations, and maintain disaster recovery procedures |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Step Functions state machines should have logging turned on | StepFunctions.1 | Step Functions state machines should have logging enabled to CloudWatch Logs for monitoring and troubleshooting | **IaC** - Configure loggingConfiguration in state machine definition with appropriate log level (ALL, ERROR, FATAL, or OFF) and CloudWatch Logs destination |  |
| **High** | Step Functions state machines should not have a dangling IAM role | StepFunctions.2 | Step Functions state machines should have valid IAM roles that exist and can be assumed by the service | **IaC** - Ensure IAM roles referenced in state machines exist, have proper trust policies, and include necessary permissions for integrated services |  |

### AWS Security Hub Latest
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Step Functions state machines should have logging turned on | StepFunctions.1 | This control checks whether an AWS Step Functions state machine has logging turned on | **IaC** - Enable logging in Step Functions state machine definition by configuring loggingConfiguration with appropriate CloudWatch Logs ARN and log level |  |
| **High** | Step Functions state machines should not have a dangling IAM role | StepFunctions.2 | This control checks whether Step Functions state machines have IAM roles that can be assumed by the service | **IaC** - Validate IAM role ARNs in state machine definitions, ensure roles exist and have proper trust relationships with states.amazonaws.com |  |


## Operational Controls
---



## Cost Controls
---


### AWS Step Functions Cost Optimization Best Practices Latest
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Choose appropriate Step Functions workflow type | COST-01 | Select Standard or Express workflows based on execution patterns and duration to optimize costs | **User** - Use Express workflows for high-volume, short-duration workloads (under 5 minutes) and Standard workflows for long-running processes with complex error handling |
| **High** | Optimize state transitions | COST-02 | Minimize the number of state transitions to reduce Step Functions charges | **User** - Combine multiple simple states into single states where possible, use parallel states efficiently, and avoid unnecessary pass states |
| Medium | Implement efficient error handling | COST-03 | Design error handling to minimize retry costs and failed state transitions | **User** - Configure appropriate retry intervals and maximum attempts, implement circuit breaker patterns, and use exponential backoff for retries |
| Medium | Monitor and analyze execution patterns | COST-04 | Use CloudWatch metrics to monitor execution costs and identify optimization opportunities | **IaC** - Set up CloudWatch dashboards for execution metrics, implement cost alerts for unexpected charges, and regularly review execution patterns |
| Medium | Optimize payload sizes | COST-05 | Minimize data passed between states to reduce execution costs and improve performance | **User** - Use ResultPath and Parameters to filter data between states, store large payloads in S3 and pass references, and implement data compression where appropriate |
| Low | Use appropriate timeout values | COST-06 | Set reasonable timeout values to prevent long-running executions that incur unnecessary costs | **User** - Configure TimeoutSeconds for states and workflows based on expected execution times, implement heartbeat timeouts for activities, and monitor timeout patterns |
| Low | Implement workflow versioning and cleanup | COST-07 | Manage state machine versions and clean up unused resources to avoid unnecessary storage costs | **User** - Regularly review and delete unused state machine versions, implement automated cleanup procedures, and archive old execution history data |


