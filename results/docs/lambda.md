# AWS Lambda
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Entitlement | Establish comprehensive identity and access management policies for Lambda functions including least privilege access | **IaC** - Define IAM roles and policies in CloudFormation/Terraform templates with minimal required permissions for Lambda execution |  |
| **High** | EKM-03 | CSA Cloud Controls Matrix v4.0 | Sensitive Data Protection | Implement proper encryption key management for Lambda environment variables and data at rest | **Platform** - Configure AWS KMS keys for encrypting Lambda environment variables and enable encryption at rest |  |
| **High** | DSI-01 | CSA Cloud Controls Matrix v4.0 | Classification | Classify and protect sensitive data processed by Lambda functions according to data classification policies | **User** - Implement data classification tagging and ensure sensitive data is encrypted in transit and at rest within Lambda code |  |
| Medium | AIS-01 | CSA Cloud Controls Matrix v4.0 | Application Security | Secure Lambda function interfaces and APIs with proper authentication and authorization mechanisms | **IaC** - Configure API Gateway with appropriate authentication methods and implement function-level security controls |  |
| Medium | IVS-01 | CSA Cloud Controls Matrix v4.0 | Network Security | Ensure Lambda functions are deployed in secure network configurations with proper VPC settings | **IaC** - Configure VPC settings, security groups, and subnets for Lambda functions requiring network isolation |  |
| Medium | BCR-01 | CSA Cloud Controls Matrix v4.0 | Business Continuity Planning | Implement business continuity and disaster recovery procedures for Lambda-based applications | **Platform** - Deploy Lambda functions across multiple availability zones and implement cross-region backup strategies |  |
| Low | SEF-01 | CSA Cloud Controls Matrix v4.0 | Contact / Personnel Security | Enable comprehensive logging and monitoring for Lambda functions to support incident response | **Platform** - Configure CloudWatch Logs, X-Ray tracing, and CloudTrail for comprehensive Lambda function monitoring |  |
| **High** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage Lambda function service accounts and execution roles with proper lifecycle management | **IaC** - Implement automated IAM role provisioning and deprovisioning using infrastructure as code with regular access reviews |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for Lambda function execution and resource access | **Platform** - Configure resource-based policies and IAM policies to enforce least privilege access for Lambda functions |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information to and from Lambda functions | **User** - Implement TLS encryption for all Lambda function communications and validate certificates in function code |  |
| Medium | AU-2 | NIST 800-53 Rev 5 | Event Logging | Enable comprehensive audit logging for Lambda function execution and API calls | **Platform** - Configure CloudWatch Logs retention policies and enable detailed monitoring for all Lambda functions |  |
| Medium | CM-2 | NIST 800-53 Rev 5 | Baseline Configuration | Establish and maintain baseline configurations for Lambda function deployments | **IaC** - Use infrastructure as code templates to maintain consistent Lambda function configurations and deployment baselines |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Lambda functions for security events and performance anomalies | **Platform** - Configure CloudWatch alarms, AWS X-Ray, and GuardDuty for comprehensive Lambda function monitoring |  |
| Low | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of Lambda function code and configuration information | **User** - Implement version control for Lambda function code and automated backup of function configurations to S3 |  |
| **High** | Lambda.1 | AWS Foundational Security Best Practices 1.0.0 | Lambda functions should prohibit public access by other accounts | Ensure Lambda functions cannot be invoked by unauthorized external accounts | **Platform** - Configure resource-based policies to restrict function invocation to authorized accounts and services only |  |
| **High** | Lambda.2 | AWS Foundational Security Best Practices 1.0.0 | Lambda functions should use supported runtimes | Ensure Lambda functions use supported and up-to-date runtime versions | **User** - Regularly update Lambda function runtime versions and implement automated scanning for deprecated runtimes |  |
| Medium | Lambda.3 | AWS Foundational Security Best Practices 1.0.0 | Lambda functions should use dead letter queues | Configure dead letter queues for Lambda functions to handle failed executions | **IaC** - Configure SQS or SNS dead letter queues in Lambda function configuration for proper error handling |  |
| Medium | Lambda.5 | AWS Foundational Security Best Practices 1.0.0 | VPC Lambda functions should operate in multiple Availability Zones | Deploy Lambda functions across multiple AZs for high availability when using VPC configuration | **IaC** - Configure Lambda VPC settings to include subnets from multiple availability zones in infrastructure templates |  |
| Low | Lambda.4 | AWS Foundational Security Best Practices 1.0.0 | Lambda functions should have a dead letter queue configured | Ensure proper error handling and monitoring for failed Lambda function executions | **Platform** - Configure CloudWatch alarms to monitor dead letter queue metrics and implement automated remediation |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Lambda Cost Optimization Best Practices 2024 | Right-size Lambda Function Memory | Optimize Lambda function memory allocation to balance performance and cost | **User** - Use AWS Lambda Power Tuning tool to determine optimal memory configuration and implement automated memory optimization |  |
| **High** | COST-02 | AWS Lambda Cost Optimization Best Practices 2024 | Implement Function Timeout Optimization | Set appropriate timeout values to prevent unnecessary charges from long-running functions | **IaC** - Configure optimal timeout values in infrastructure templates based on function execution patterns and requirements |  |
| Medium | COST-03 | AWS Lambda Cost Optimization Best Practices 2024 | Use Provisioned Concurrency Judiciously | Apply provisioned concurrency only where cold start latency is critical to avoid unnecessary costs | **Platform** - Implement application auto-scaling for provisioned concurrency based on usage patterns and schedule-based scaling |  |
| Medium | COST-04 | AWS Lambda Cost Optimization Best Practices 2024 | Optimize Code for Efficiency | Write efficient code to minimize execution time and memory usage | **User** - Implement code optimization techniques including connection reuse, efficient libraries, and minimal package sizes |  |
| Medium | COST-05 | AWS Lambda Cost Optimization Best Practices 2024 | Implement Cost Monitoring and Alerting | Monitor Lambda costs and set up alerts for unusual spending patterns | **Platform** - Configure AWS Cost Explorer, CloudWatch metrics, and billing alerts to monitor Lambda function costs and usage patterns |  |
| Low | COST-06 | AWS Lambda Cost Optimization Best Practices 2024 | Use ARM-based Graviton Processors | Leverage ARM-based processors for better price-performance ratio where compatible | **IaC** - Configure Lambda functions to use arm64 architecture in infrastructure templates for supported workloads |  |
| Low | COST-07 | AWS Lambda Cost Optimization Best Practices 2024 | Implement Function Cleanup and Lifecycle Management | Remove unused Lambda functions and versions to avoid unnecessary storage costs | **Platform** - Implement automated cleanup of unused function versions and aliases using lifecycle policies and AWS Config rules |  |

