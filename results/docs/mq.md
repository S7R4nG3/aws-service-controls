# AWS MQ
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy | IAM-01 | Ensure proper identity and access management policies are established for AWS MQ resources | **IaC** - Define IAM policies with least privilege access for AWS MQ brokers, queues, and administrative functions using CloudFormation or Terraform |  |
| **Critical** | Encryption Key Management | EKM-01 | Implement proper encryption key management for data at rest and in transit | **Platform** - Enable AWS KMS encryption for MQ brokers and configure customer-managed keys for enhanced control |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and lifecycle management for message data | **User** - Classify message data based on sensitivity and implement appropriate retention policies |  |
| **High** | Infrastructure and Virtualization Security | IVS-01 | Secure network infrastructure and virtualization components | **IaC** - Deploy MQ brokers in private subnets with proper VPC configuration and security groups |  |
| Medium | Secure Technology Architecture | STA-01 | Implement secure architecture patterns for message queuing systems | **IaC** - Design multi-AZ deployments with proper network segmentation and security controls |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to AWS MQ resources | **Platform** - Configure MQ broker access control using AWS IAM policies and resource-based policies |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Enable TLS encryption for all MQ connections and configure SSL/TLS certificates |  |
| **Critical** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **Platform** - Enable encryption at rest for MQ brokers using AWS KMS keys |  |
| **High** | Audit Generation | AU-12 | Generate audit records for security-relevant events | **Platform** - Enable CloudTrail logging for MQ API calls and configure broker logging |  |
| **High** | Information System Monitoring | SI-4 | Monitor the information system to detect attacks and indicators of potential attacks | **Platform** - Configure CloudWatch monitoring and alerting for MQ brokers and set up security event notifications |  |
| Medium | Information System Backup | CP-9 | Conduct backups of user-level and system-level information | **Platform** - Configure automated backups for MQ broker configurations and implement message persistence |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | MQ brokers should not be publicly accessible | MQ.1 | Ensure MQ brokers are not accessible from the internet | **IaC** - Deploy MQ brokers with PubliclyAccessible set to false and use VPC endpoints for private connectivity |  |
| **High** | MQ brokers should use active/standby deployment mode | MQ.2 | Configure MQ brokers for high availability using active/standby mode | **IaC** - Set deployment mode to ACTIVE_STANDBY_MULTI_AZ for production workloads |  |
| **High** | MQ brokers should have encryption in transit enabled | MQ.3 | Ensure all data in transit is encrypted using TLS | **Platform** - Configure SSL/TLS encryption for all broker connections and disable non-encrypted protocols |  |
| **High** | MQ brokers should have encryption at rest enabled | MQ.4 | Ensure data at rest is encrypted using AWS KMS | **Platform** - Enable encryption at rest using AWS KMS customer-managed keys |  |
| Medium | MQ brokers should have logging enabled | MQ.5 | Enable comprehensive logging for security monitoring | **Platform** - Enable general and audit logging for MQ brokers and configure log retention policies |  |

### AWS Security Hub 2023.04.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | ActiveMQ brokers should have automatic minor version upgrade enabled | MQ.6 | Ensure automatic minor version upgrades are enabled for security patches | **Platform** - Enable AutoMinorVersionUpgrade for MQ brokers to receive security updates automatically |  |
| **High** | MQ brokers should restrict access using security groups | MQ.7 | Configure restrictive security group rules for MQ brokers | **IaC** - Create security groups with minimal required ports and source IP restrictions |  |
| Medium | MQ brokers should have multi-factor authentication enabled | MQ.8 | Implement MFA for administrative access to MQ brokers | **User** - Configure MFA for IAM users and roles with MQ administrative permissions |  |
| Medium | MQ brokers should be deployed across multiple availability zones | MQ.9 | Deploy MQ brokers across multiple AZs for resilience | **IaC** - Configure subnet groups spanning multiple availability zones for broker deployment |  |


## Operational Controls
---



## Cost Controls
---


### AWS MQ Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size MQ broker instances | COST-01 | Select appropriate instance types based on actual workload requirements | **User** - Monitor CPU, memory, and network utilization to determine optimal instance sizes and downsize overprovisioned brokers |
| **High** | Use single-instance deployment for non-production workloads | COST-02 | Deploy single-instance brokers for development and testing environments | **IaC** - Configure SINGLE_INSTANCE deployment mode for non-production environments to reduce costs by 50% |
| Medium | Implement automated broker lifecycle management | COST-03 | Automatically start and stop brokers based on usage patterns | **IaC** - Use Lambda functions with CloudWatch Events to automatically stop/start brokers during off-hours |
| Medium | Optimize message retention policies | COST-04 | Configure appropriate message retention periods to minimize storage costs | **User** - Set message TTL and dead letter queue policies to prevent unnecessary message accumulation |
| Medium | Monitor and optimize data transfer costs | COST-05 | Minimize cross-AZ and cross-region data transfer charges | **IaC** - Deploy applications and MQ brokers in the same AZ when possible and use VPC endpoints to avoid internet gateway charges |
| Low | Implement cost allocation tags | COST-06 | Tag MQ resources for cost tracking and optimization | **IaC** - Apply consistent tagging strategy to track costs by environment, application, and team |
| Low | Regular cost review and optimization | COST-07 | Establish regular cost review processes for MQ resources | **User** - Set up monthly cost reports and reviews using AWS Cost Explorer and Trusted Advisor recommendations |


