# AWS CloudWatch Synthetics
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-02 | CSA Cloud Controls Matrix v4.0 | User Access Management | Implement proper identity and access management controls for CloudWatch Synthetics resources | **IaC** - Use AWS IAM policies to restrict access to CloudWatch Synthetics canaries based on principle of least privilege |  |
| **High** | DSI-01 | CSA Cloud Controls Matrix v4.0 | Data Security and Information Lifecycle Management | Protect sensitive data within Synthetics canary configurations and execution results | **Platform** - Enable encryption in transit and at rest for CloudWatch Synthetics data using AWS KMS |  |
| **High** | AIS-02 | CSA Cloud Controls Matrix v4.0 | Application Security | Secure application interfaces and APIs monitored by Synthetics canaries | **User** - Configure canaries to use secure authentication methods and avoid hardcoded credentials |  |
| Medium | BCR-03 | CSA Cloud Controls Matrix v4.0 | Business Continuity Planning | Ensure business continuity for monitoring capabilities | **IaC** - Deploy canaries across multiple regions and availability zones for redundancy |  |
| **Critical** | AC-2 | NIST SP 800-53 Rev 5 | Account Management | Manage accounts that have access to CloudWatch Synthetics resources | **Platform** - Use AWS IAM to create and manage user accounts with appropriate permissions for Synthetics operations |  |
| **Critical** | AC-3 | NIST SP 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to CloudWatch Synthetics | **IaC** - Implement resource-based policies and IAM policies to control access to specific canaries and operations |  |
| **High** | SC-8 | NIST SP 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **Platform** - Ensure all API calls and data transmission use TLS 1.2 or higher encryption |  |
| **High** | AU-2 | NIST SP 800-53 Rev 5 | Event Logging | Identify the types of events to be logged | **Platform** - Enable AWS CloudTrail logging for all CloudWatch Synthetics API calls and operations |  |
| Medium | CP-2 | NIST SP 800-53 Rev 5 | Contingency Plan | Develop a contingency plan for the system | **User** - Create backup monitoring strategies and failover procedures for critical canaries |  |
| **Critical** | IAM.1 | AWS Foundational Security Best Practices 1.0.0 | IAM policies should not allow full administrative privileges | Avoid overly permissive IAM policies for CloudWatch Synthetics access | **IaC** - Create granular IAM policies that grant only necessary permissions for Synthetics operations |  |
| **High** | CloudWatch.1 | AWS Foundational Security Best Practices 1.0.0 | CloudWatch log groups should be encrypted | CloudWatch Synthetics log groups should be encrypted at rest | **IaC** - Configure CloudWatch log groups used by Synthetics canaries with KMS encryption keys |  |
| Medium | SecretsManager.1 | AWS Foundational Security Best Practices 1.0.0 | Secrets Manager secrets should have automatic rotation enabled | Rotate credentials used by Synthetics canaries regularly | **Platform** - Store canary credentials in AWS Secrets Manager with automatic rotation enabled |  |
| **High** | CloudWatch.12 | AWS Security Hub 2024.1 | CloudWatch should not have overly permissive policies | CloudWatch Synthetics should not have overly permissive resource policies | **IaC** - Review and restrict CloudWatch Synthetics resource policies to follow least privilege principle |  |
| **High** | EC2.2 | AWS Security Hub 2024.1 | VPC default security groups should not allow inbound or outbound traffic | Synthetics canaries running in VPC should use custom security groups | **IaC** - Configure custom security groups for VPC-enabled canaries with minimal required network access |  |
| Medium | KMS.1 | AWS Security Hub 2024.1 | IAM customer managed policies should not allow decryption actions on all KMS keys | Limit KMS key access for CloudWatch Synthetics encryption | **IaC** - Create specific KMS key policies for CloudWatch Synthetics with restricted decrypt permissions |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Optimize Canary Execution Frequency | Right-size canary execution frequency based on business requirements to reduce unnecessary runs | **User** - Configure canary schedules to run only as frequently as needed for effective monitoring, avoiding over-monitoring |  |
| **High** | COST-02 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Implement Canary Lifecycle Management | Regularly review and decommission unused or redundant canaries | **User** - Establish processes to identify and remove canaries that are no longer needed or are duplicating monitoring coverage |  |
| Medium | COST-03 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Optimize CloudWatch Logs Retention | Set appropriate log retention periods for Synthetics canary logs to control storage costs | **IaC** - Configure CloudWatch log group retention periods based on compliance and operational requirements rather than using indefinite retention |  |
| Medium | COST-04 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Use Appropriate Canary Runtime Versions | Select the most cost-effective runtime version that meets functional requirements | **User** - Choose canary runtime versions that provide required functionality without unnecessary overhead or premium features |  |
| Medium | COST-05 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Monitor and Alert on Canary Costs | Implement cost monitoring and alerting for CloudWatch Synthetics usage | **Platform** - Set up AWS Cost Explorer and CloudWatch billing alarms to track Synthetics costs and receive notifications for unexpected increases |  |
| Low | COST-06 | AWS CloudWatch Synthetics Cost Optimization Best Practices 2024 | Optimize Screenshot and HAR File Storage | Configure appropriate retention and compression settings for canary artifacts | **IaC** - Set S3 lifecycle policies for canary artifact buckets to transition older screenshots and HAR files to cheaper storage classes or delete them after specified periods |  |

