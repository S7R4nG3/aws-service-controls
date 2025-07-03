# AWS MQ
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| CRITICAL | MQ.1 | AWS Foundational Security Best Practices 1.0.0 | Amazon MQ brokers should not have public accessibility | MQ brokers should not be publicly accessible to prevent unauthorized access from the internet | **IaC** - Configure broker in private subnets with PubliclyAccessible set to false |  |
| HIGH | MQ.2 | AWS Foundational Security Best Practices 1.0.0 | Amazon MQ brokers should have audit logging enabled | Enable audit logging to track access and administrative activities | **IaC** - Configure Logs parameter with audit set to true in broker configuration |  |
| MEDIUM | MQ.3 | AWS Foundational Security Best Practices 1.0.0 | Amazon MQ brokers should have general logging enabled | Enable general logging to capture broker operational information | **IaC** - Configure Logs parameter with general set to true in broker configuration |  |
| MEDIUM | MQ.4 | AWS Foundational Security Best Practices 1.0.0 | Amazon MQ brokers should have automatic minor version upgrades enabled | Enable automatic minor version upgrades to receive security patches and bug fixes | **IaC** - Set AutoMinorVersionUpgrade to true during broker creation |  |
| MEDIUM | MQ.5 | AWS Foundational Security Best Practices 1.0.0 | Amazon MQ brokers should use encryption in transit | Enable encryption in transit to protect data while in motion | **IaC** - Configure SSL/TLS encryption for client connections |  |
| CRITICAL | AC-2 | NIST 800-53 Rev 5 | Account Management | Implement proper account management for MQ broker users including creation, modification, and deletion | **User** - Establish user accounts with appropriate permissions, implement regular access reviews, and remove unused accounts |  |
| CRITICAL | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to MQ resources | **IaC** - Configure IAM policies and MQ broker user permissions to enforce least privilege access |  |
| CRITICAL | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable SSL/TLS encryption for client connections and configure security groups to allow only secure protocols |  |
| HIGH | AU-2 | NIST 800-53 Rev 5 | Event Logging | Ensure that auditable events are defined and logged by the system | **IaC** - Enable CloudTrail for API logging and configure MQ audit logging to capture broker events |  |
| HIGH | SC-7 | NIST 800-53 Rev 5 | Boundary Protection | Monitor and control communications at external boundaries and key internal boundaries | **IaC** - Deploy MQ brokers in private subnets with appropriate security groups and NACLs |  |
| HIGH | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of system-level information and user-level information | **Platform** - AWS MQ automatically performs daily backups; configure retention period based on requirements |  |
| MEDIUM | IA-5 | NIST 800-53 Rev 5 | Authenticator Management | Manage system authenticators including passwords and certificates | **User** - Implement strong password policies for MQ users and rotate credentials regularly |  |
| MEDIUM | CM-2 | NIST 800-53 Rev 5 | Baseline Configuration | Develop, document, and maintain baseline configurations | **IaC** - Define and maintain standardized MQ broker configurations using Infrastructure as Code |  |
| CRITICAL | IAM-01 | CSA Cloud Controls Matrix v4.0 | Identity and Access Management Policy and Procedures | Establish policies and procedures for identity and access management | **User** - Develop and implement IAM policies specific to MQ access management and user provisioning |  |
| CRITICAL | DSI-01 | CSA Cloud Controls Matrix v4.0 | Data Security and Information Lifecycle Management | Implement controls for data security throughout its lifecycle | **IaC** - Configure encryption at rest and in transit for MQ messages and implement data classification policies |  |
| HIGH | IVS-01 | CSA Cloud Controls Matrix v4.0 | Infrastructure and Virtualization Security | Implement security controls for infrastructure and virtualization layers | **IaC** - Deploy MQ brokers with proper network segmentation and security group configurations |  |
| HIGH | AIS-01 | CSA Cloud Controls Matrix v4.0 | Application and Interface Security | Establish comprehensive logging and monitoring capabilities | **IaC** - Enable CloudWatch logging, CloudTrail, and MQ audit logs with appropriate retention policies |  |
| HIGH | BCR-01 | CSA Cloud Controls Matrix v4.0 | Business Continuity Management and Operational Resilience | Implement business continuity and disaster recovery capabilities | **IaC** - Configure multi-AZ deployment with automatic failover and backup strategies |  |
| MEDIUM | EKM-01 | CSA Cloud Controls Matrix v4.0 | Encryption and Key Management | Implement appropriate encryption and key management practices | **IaC** - Use AWS KMS for encryption key management and enable encryption for MQ broker storage |  |
| MEDIUM | GRM-01 | CSA Cloud Controls Matrix v4.0 | Governance and Risk Management | Establish governance framework and risk management processes | **User** - Implement regular security assessments and compliance monitoring for MQ deployments |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| HIGH | COST-01 | AWS MQ Cost Optimization Best Practices 2024 | Right-size broker instances | Select appropriate instance types based on actual workload requirements to avoid over-provisioning | **User** - Monitor CPU and memory utilization metrics and resize instances based on actual usage patterns |  |
| HIGH | COST-02 | AWS MQ Cost Optimization Best Practices 2024 | Optimize deployment mode selection | Choose single-instance deployment for development/testing and multi-AZ only when high availability is required | **IaC** - Use single-instance deployment for non-production environments and active/standby multi-AZ only for production workloads requiring HA |  |
| MEDIUM | COST-03 | AWS MQ Cost Optimization Best Practices 2024 | Implement automated start/stop scheduling | Automatically stop non-production brokers during off-hours to reduce costs | **IaC** - Use AWS Lambda functions with EventBridge rules to automatically stop and start brokers based on schedules |  |
| MEDIUM | COST-04 | AWS MQ Cost Optimization Best Practices 2024 | Monitor and optimize storage usage | Regularly monitor storage consumption and implement message retention policies | **User** - Configure appropriate message time-to-live settings and monitor storage utilization through CloudWatch metrics |  |
| MEDIUM | COST-05 | AWS MQ Cost Optimization Best Practices 2024 | Use AWS Cost Explorer and budgets | Set up cost monitoring and alerting to track MQ spending | **User** - Create AWS budgets for MQ services and use Cost Explorer to analyze spending patterns and identify optimization opportunities |  |
| LOW | COST-06 | AWS MQ Cost Optimization Best Practices 2024 | Optimize network data transfer costs | Minimize cross-AZ and cross-region data transfer by co-locating applications with MQ brokers | **IaC** - Deploy MQ brokers in the same AZ as client applications when possible and use VPC endpoints to reduce data transfer costs |  |
| LOW | COST-07 | AWS MQ Cost Optimization Best Practices 2024 | Regular cost review and optimization | Conduct periodic reviews of MQ usage and costs to identify optimization opportunities | **User** - Establish monthly cost review processes and use AWS Trusted Advisor recommendations for MQ optimization |  |

