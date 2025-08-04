# AWS DocumentDB
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | DSI-01 | CSA Cloud Controls Matrix v4.0.10 | Data Security and Information Lifecycle Management | Implement data classification and protection mechanisms for DocumentDB data | **Platform** - Enable encryption at rest and in transit for DocumentDB clusters using AWS KMS keys |  |
| **High** | EKM-01 | CSA Cloud Controls Matrix v4.0.10 | Encryption and Key Management | Manage encryption keys for DocumentDB data protection | **IaC** - Configure customer-managed KMS keys for DocumentDB encryption with proper key rotation policies |  |
| **High** | IAM-01 | CSA Cloud Controls Matrix v4.0.10 | Identity and Access Management Policy and Procedures | Establish comprehensive identity and access management policies for DocumentDB resources | **IaC** - Implement IAM policies and roles through CloudFormation or Terraform templates with least privilege access to DocumentDB clusters |  |
| Medium | IVS-01 | CSA Cloud Controls Matrix v4.0.10 | Vulnerability Management | Implement vulnerability assessments for DocumentDB infrastructure | **Platform** - Enable AWS Security Hub and Inspector for automated vulnerability scanning of DocumentDB resources |  |
| Medium | SEF-01 | CSA Cloud Controls Matrix v4.0.10 | Security Incident Management, E-Discovery & Cloud Forensics | Implement comprehensive logging for DocumentDB activities | **IaC** - Configure CloudTrail, CloudWatch, and DocumentDB profiler logging through infrastructure templates |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to DocumentDB resources | **IaC** - Implement RBAC through DocumentDB database users and roles with VPC-based network access controls |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted DocumentDB information | **Platform** - Enable TLS encryption for all client connections to DocumentDB clusters |  |
| **High** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect confidentiality and integrity of DocumentDB data at rest | **IaC** - Configure encryption at rest for DocumentDB clusters using AWS KMS customer-managed keys |  |
| Medium | AU-2 | NIST 800-53 Rev 5 | Event Logging | Identify events that are to be logged for DocumentDB operations | **IaC** - Enable DocumentDB audit logging and profiler to capture database operations and performance metrics |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of DocumentDB data and system-level information | **IaC** - Configure automated backups with appropriate retention periods and cross-region backup replication |  |
| Medium | SC-7 | NIST 800-53 Rev 5 | Boundary Protection | Monitor and control communications at external boundaries and key internal boundaries | **IaC** - Deploy DocumentDB clusters in private subnets with VPC security groups and NACLs for network segmentation |  |
| **High** | DocumentDB.1 | AWS Foundational Security Standard 1.0.0 | DocumentDB clusters should be encrypted | DocumentDB clusters should have encryption at rest enabled | **IaC** - Enable storage encryption parameter when creating DocumentDB clusters in infrastructure templates |  |
| Medium | DocumentDB.2 | AWS Foundational Security Standard 1.0.0 | DocumentDB clusters should have backup retention period greater than or equal to 7 days | DocumentDB clusters should be configured with adequate backup retention | **IaC** - Set backup retention period to minimum 7 days in DocumentDB cluster configuration |  |
| Medium | DocumentDB.3 | AWS Foundational Security Standard 1.0.0 | DocumentDB clusters should have an adequate number of log exports enabled | DocumentDB clusters should export audit and profiler logs to CloudWatch | **IaC** - Enable audit and profiler log exports in DocumentDB cluster configuration |  |
| Low | DocumentDB.4 | AWS Foundational Security Standard 1.0.0 | DocumentDB clusters should be configured to copy tags to snapshots | DocumentDB clusters should copy all tags to snapshots for proper resource management | **IaC** - Enable copy-tags-to-snapshot parameter in DocumentDB cluster configuration |  |
| Low | DocumentDB.5 | AWS Foundational Security Standard 1.0.0 | DocumentDB clusters should have deletion protection enabled | DocumentDB clusters should be protected against accidental deletion | **IaC** - Enable deletion protection parameter in DocumentDB cluster configuration |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS DocumentDB Cost Optimization Best Practices 2023 | Right-size DocumentDB instances | Regularly monitor and adjust DocumentDB instance types based on actual resource utilization | **User** - Use CloudWatch metrics to analyze CPU, memory, and I/O utilization and resize instances accordingly |  |
| **High** | COST-02 | AWS DocumentDB Cost Optimization Best Practices 2023 | Optimize storage costs | Monitor storage growth and implement data lifecycle policies to manage storage costs | **User** - Regularly review storage utilization patterns and archive or delete unnecessary data |  |
| Medium | COST-03 | AWS DocumentDB Cost Optimization Best Practices 2023 | Configure appropriate backup retention | Set backup retention periods based on business requirements to avoid unnecessary storage costs | **IaC** - Configure backup retention period to match business recovery requirements without excessive retention |  |
| Medium | COST-04 | AWS DocumentDB Cost Optimization Best Practices 2023 | Use Reserved Instances for predictable workloads | Purchase DocumentDB Reserved Instances for steady-state workloads to reduce costs | **User** - Analyze usage patterns and purchase 1-year or 3-year Reserved Instances for consistent workloads |  |
| Medium | COST-05 | AWS DocumentDB Cost Optimization Best Practices 2023 | Implement resource tagging for cost allocation | Use consistent tagging strategy to track and allocate DocumentDB costs across teams and projects | **IaC** - Implement mandatory tags for cost center, environment, and project in DocumentDB resources |  |
| Medium | COST-06 | AWS DocumentDB Cost Optimization Best Practices 2023 | Monitor and alert on cost anomalies | Set up cost monitoring and alerts to detect unexpected DocumentDB spending | **Platform** - Configure AWS Cost Anomaly Detection and Budgets to monitor DocumentDB spending patterns |  |
| Low | COST-07 | AWS DocumentDB Cost Optimization Best Practices 2023 | Optimize cross-region data transfer costs | Minimize cross-region data transfer by co-locating applications and DocumentDB clusters | **IaC** - Deploy DocumentDB clusters in the same region as consuming applications to reduce data transfer costs |  |
| Low | COST-08 | AWS DocumentDB Cost Optimization Best Practices 2023 | Review and optimize read replica usage | Regularly assess the necessity and utilization of DocumentDB read replicas | **User** - Monitor read replica utilization and remove underutilized replicas or redistribute read traffic |  |

