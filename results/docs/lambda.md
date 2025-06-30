# AWS Lambda
---


### CSA Cloud Controls Matrix 4.0.10
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish comprehensive identity and access management policies for Lambda function access | **Platform/IaC** - Implement IAM policies using least privilege principle, define execution roles with minimal required permissions, and establish service-linked roles for Lambda functions |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and protection mechanisms for data processed by Lambda functions | **User/IaC** - Classify data processed by functions, implement encryption at rest and in transit, use AWS KMS for key management, and establish data retention policies |  |
| **Critical** | Encryption and Key Management | EKM-01 | Ensure proper encryption and key management for Lambda environment variables and data | **User/IaC** - Enable encryption for environment variables using AWS KMS, implement proper key rotation policies, and use customer-managed keys for sensitive data |  |
| **High** | Infrastructure and Virtualization Security | IVS-01 | Implement network security controls and VPC configuration for Lambda functions | **IaC** - Configure Lambda functions within VPC when accessing private resources, implement security groups and NACLs, and use VPC endpoints for AWS service communication |  |
| **High** | Threat and Vulnerability Management | TVM-01 | Implement vulnerability scanning and threat detection for Lambda functions | **Platform/User** - Enable AWS GuardDuty, implement dependency scanning for function code, use AWS Inspector for container images, and establish vulnerability remediation processes |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Account Management | AC-2 | Implement proper account management for Lambda function access and execution roles | **Platform/IaC** - Create dedicated execution roles for each Lambda function, implement role-based access control, and establish account lifecycle management procedures |  |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Lambda functions and resources | **Platform/IaC** - Implement resource-based policies, use API Gateway for function access control, and configure cross-account access policies with proper conditions |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **User/IaC** - Use HTTPS/TLS for all API communications, implement certificate pinning where applicable, and encrypt data in transit between Lambda and other services |  |
| **High** | Event Logging | AU-2 | Determine auditable events and ensure Lambda function activities are logged | **Platform/IaC** - Enable CloudTrail for API calls, configure CloudWatch Logs for function execution logs, and implement structured logging within function code |  |
| **High** | Boundary Protection | SC-7 | Monitor and control communications at external boundaries and key internal boundaries | **IaC** - Configure VPC boundaries for Lambda functions, implement security groups as firewalls, and use AWS WAF for API Gateway protection |  |
| **High** | System Monitoring | SI-4 | Monitor Lambda functions for security events and performance anomalies | **Platform/User** - Configure CloudWatch alarms for function metrics, implement AWS X-Ray for distributed tracing, and use AWS Security Hub for centralized monitoring |  |
| Medium | System Backup | CP-9 | Conduct backups of Lambda function code and configuration | **User/IaC** - Store function code in version control systems, implement function versioning and aliases, and backup deployment packages to S3 |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Lambda functions should prohibit public access by other accounts | Lambda.1 | Ensure Lambda functions are not publicly accessible unless specifically required | **IaC** - Review and restrict resource-based policies, avoid wildcard principals in policies, and implement condition statements for cross-account access |  |
| **Critical** | Lambda functions should use supported runtimes | Lambda.2 | Ensure Lambda functions use currently supported runtime versions | **User** - Regularly update function runtimes to supported versions, establish runtime deprecation tracking, and test functions with new runtime versions |  |
| **High** | Lambda functions should be in a VPC | Lambda.3 | Configure Lambda functions within a VPC when accessing private resources | **IaC** - Place functions requiring private resource access in VPC, configure appropriate subnets and security groups, and use VPC endpoints for AWS services |  |
| Medium | Lambda functions should have a dead letter queue configured | Lambda.4 | Configure dead letter queues for asynchronous Lambda function invocations | **IaC** - Configure SQS or SNS as dead letter queue destination, set appropriate retry policies, and implement monitoring for failed invocations |  |
| Medium | VPC Lambda functions should operate in multiple Availability Zones | Lambda.5 | Ensure VPC-configured Lambda functions can operate across multiple AZs | **IaC** - Configure function subnets in multiple availability zones, ensure proper subnet routing, and validate cross-AZ connectivity |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Lambda functions should have reserved concurrency configured | Lambda.ConcurrentExecutions | Configure reserved concurrency to prevent resource exhaustion attacks | **IaC** - Set appropriate reserved concurrency limits, monitor concurrency usage, and implement throttling controls |  |
| **High** | Lambda environment variables should be encrypted | Lambda.EnvironmentEncryption | Encrypt sensitive data in Lambda environment variables | **User/IaC** - Enable KMS encryption for environment variables, use AWS Systems Manager Parameter Store for sensitive configuration, and implement proper key management |  |
| Medium | Lambda functions should have tracing enabled | Lambda.Tracing | Enable AWS X-Ray tracing for Lambda functions to improve observability | **IaC** - Enable X-Ray tracing mode, configure sampling rules, and implement distributed tracing across services |  |
| Medium | Lambda functions should use code signing | Lambda.CodeSigning | Implement code signing to ensure function integrity | **User/IaC** - Create signing profiles, configure code signing policies, and establish code integrity verification processes |  |


## Operational Controls
---



## Cost Controls
---


### AWS Lambda Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size Lambda function memory allocation | COST-001 | Optimize memory allocation to balance performance and cost | **User** - Use AWS Lambda Power Tuning tool to find optimal memory configuration, monitor memory utilization metrics, and adjust memory based on actual usage patterns |
| **High** | Minimize Lambda function execution time | COST-002 | Optimize function code to reduce execution duration and associated costs | **User** - Optimize code performance, minimize cold start times, use connection pooling, implement efficient algorithms, and reduce package size |
| **High** | Use appropriate Lambda pricing model | COST-003 | Choose between on-demand and provisioned concurrency based on usage patterns | **IaC/User** - Analyze traffic patterns, use provisioned concurrency only for predictable workloads, implement auto-scaling for provisioned concurrency, and monitor cost impact |
| Medium | Optimize Lambda function architecture | COST-004 | Design efficient architectures to minimize Lambda invocations and data transfer | **User** - Batch process records where possible, use appropriate event sources, minimize cross-region data transfer, and implement efficient retry mechanisms |
| Medium | Monitor and alert on Lambda costs | COST-005 | Implement cost monitoring and alerting for Lambda usage | **Platform/IaC** - Set up AWS Cost Explorer reports, configure billing alerts, use AWS Budgets for cost thresholds, and implement cost allocation tags |
| Medium | Clean up unused Lambda resources | COST-006 | Regularly review and remove unused Lambda functions and versions | **User** - Implement automated cleanup of old function versions, remove unused functions, clean up associated resources like layers and triggers, and establish resource lifecycle policies |
| Low | Optimize Lambda layers usage | COST-007 | Use Lambda layers efficiently to reduce deployment package size and improve reusability | **User/IaC** - Share common libraries through layers, optimize layer sizes, version layers appropriately, and avoid duplicate dependencies across layers |


