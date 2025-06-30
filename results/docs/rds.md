# AWS Relational Database Service (RDS)
---


### Cloud Security Alliance (CSA) Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish comprehensive identity and access management policies for RDS instances and related resources | **Platform/IaC** - Configure IAM policies, roles, and database users through AWS IAM and RDS parameter groups with least privilege access principles |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and lifecycle management for RDS databases | **User/IaC** - Configure automated backups, retention policies, and data classification tags through RDS backup settings and AWS tags |  |
| **Critical** | Encryption and Key Management | EKM-01 | Implement encryption at rest and in transit for RDS instances | **IaC** - Enable RDS encryption using AWS KMS keys and enforce SSL/TLS connections through parameter groups |  |
| **High** | Business Continuity Management and Operational Resilience | BCR-01 | Establish business continuity and disaster recovery procedures for RDS | **Platform/IaC** - Configure Multi-AZ deployments, read replicas, and automated backup strategies with cross-region replication |  |
| **High** | Audit Logging / Intrusion Detection | LOG-01 | Enable comprehensive logging and monitoring for RDS instances | **Platform/IaC** - Enable RDS Performance Insights, CloudWatch monitoring, and database audit logs with integration to CloudTrail |  |
| Medium | Vulnerability Management | IVS-01 | Implement vulnerability scanning and patch management for RDS instances | **Platform** - Enable automatic minor version upgrades and configure maintenance windows for RDS instances |  |

### NIST Special Publication 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to RDS resources | **IaC** - Implement IAM policies with resource-based permissions and database user management with role-based access control |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information to and from RDS | **IaC** - Enforce SSL/TLS encryption for all database connections and configure security groups to restrict network access |  |
| **Critical** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest in RDS | **IaC** - Enable RDS encryption at rest using AWS KMS customer managed keys with appropriate key rotation policies |  |
| **High** | Event Logging | AU-2 | Identify the types of events to be logged by RDS and supporting infrastructure | **Platform/IaC** - Configure database audit logging, CloudWatch logs, and CloudTrail for comprehensive event capture |  |
| **High** | System Backup | CP-9 | Conduct backups of user-level and system-level information contained in RDS | **Platform/IaC** - Configure automated backups with appropriate retention periods and cross-region backup replication |  |
| **High** | System Monitoring | SI-4 | Monitor RDS instances to detect attacks and indicators of potential attacks | **Platform/IaC** - Enable CloudWatch monitoring, RDS Performance Insights, and integrate with AWS GuardDuty for threat detection |  |
| Medium | Vulnerability Monitoring and Scanning | RA-5 | Monitor and scan for vulnerabilities in RDS instances and supporting infrastructure | **Platform** - Enable AWS Security Hub findings and implement regular security assessments of RDS configurations |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | RDS snapshots should be private | RDS.1 | RDS snapshots should not be public to prevent unauthorized access to sensitive data | **IaC** - Ensure RDS snapshot sharing is restricted and automated checks prevent public snapshots |  |
| **Critical** | RDS DB instances should have encryption at rest enabled | RDS.3 | RDS DB instances should be encrypted at rest to protect sensitive data | **IaC** - Enable encryption at rest for all RDS instances using AWS KMS keys |  |
| **Critical** | RDS cluster snapshots and database snapshots should be encrypted at rest | RDS.4 | RDS snapshots should be encrypted to protect backup data | **IaC** - Ensure all RDS snapshots inherit encryption from the source database or are explicitly encrypted |  |
| **High** | RDS DB instances should be configured with multiple Availability Zones | RDS.5 | RDS instances should use Multi-AZ for high availability and disaster recovery | **IaC** - Configure Multi-AZ deployment for production RDS instances to ensure high availability |  |
| **High** | Enhanced monitoring should be configured for RDS DB instances | RDS.6 | Enhanced monitoring provides real-time metrics for the operating system of RDS instances | **IaC** - Enable Enhanced Monitoring for RDS instances with appropriate granularity and retention |  |
| **High** | RDS DB instances should have deletion protection enabled | RDS.8 | Deletion protection prevents accidental deletion of RDS instances | **IaC** - Enable deletion protection for production RDS instances to prevent accidental deletion |  |
| Medium | IAM authentication should be configured for RDS instances | RDS.12 | IAM database authentication provides secure access without embedding credentials | **IaC** - Enable IAM database authentication where supported and configure appropriate IAM policies |  |
| Medium | RDS automatic minor version upgrades should be enabled | RDS.13 | Automatic minor version upgrades ensure security patches are applied | **Platform/IaC** - Enable automatic minor version upgrades and configure appropriate maintenance windows |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | RDS DB instances should prohibit public read access | RDS.2 | RDS instances should not be publicly accessible to prevent unauthorized access | **IaC** - Configure RDS instances with PubliclyAccessible set to false and place in private subnets |  |
| **High** | RDS clusters should have deletion protection enabled | RDS.7 | RDS clusters should have deletion protection to prevent accidental deletion | **IaC** - Enable deletion protection for RDS clusters in production environments |  |
| **High** | Database logging should be enabled | RDS.9 | RDS instances should have appropriate database logs enabled for security monitoring | **IaC** - Enable relevant database logs (error, general, slow query) and export to CloudWatch Logs |  |
| Medium | IAM authentication should be configured for RDS instances | RDS.10 | RDS instances should use IAM authentication where possible | **IaC** - Enable IAM database authentication for supported engine types and configure IAM policies |  |
| Medium | RDS instances should have automatic backups enabled | RDS.11 | RDS instances should have automated backups configured for data protection | **IaC** - Configure automated backups with appropriate backup retention periods and backup windows |  |


## Operational Controls
---



## Cost Controls
---


### AWS RDS Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size RDS instances based on utilization | COST-RDS-1 | Monitor and adjust RDS instance sizes based on actual CPU, memory, and I/O utilization | **Platform** - Use CloudWatch metrics and RDS Performance Insights to analyze utilization and resize instances appropriately |
| **High** | Implement Reserved Instance purchasing strategy | COST-RDS-2 | Purchase Reserved Instances for predictable workloads to achieve significant cost savings | **Platform** - Analyze usage patterns and purchase 1-year or 3-year Reserved Instances for stable workloads |
| **High** | Optimize storage type and provisioned IOPS | COST-RDS-3 | Select appropriate storage types and avoid over-provisioning IOPS | **IaC** - Use gp3 storage where appropriate and monitor IOPS utilization to avoid over-provisioning |
| Medium | Implement automated start/stop scheduling for non-production environments | COST-RDS-4 | Schedule RDS instances to stop during non-business hours for development and testing environments | **Platform/IaC** - Use AWS Lambda and EventBridge to automate start/stop schedules for non-production RDS instances |
| Medium | Optimize backup retention and cross-region replication | COST-RDS-5 | Set appropriate backup retention periods and minimize unnecessary cross-region backup copies | **IaC** - Configure backup retention based on compliance requirements and evaluate cross-region backup necessity |
| Medium | Monitor and optimize data transfer costs | COST-RDS-6 | Minimize data transfer costs by optimizing read replica placement and application architecture | **Platform/IaC** - Place read replicas in the same region as applications and optimize query patterns to reduce data transfer |
| Medium | Implement database connection pooling | COST-RDS-7 | Use RDS Proxy or application-level connection pooling to reduce the need for larger instance sizes | **IaC** - Configure RDS Proxy to pool database connections and reduce connection overhead |
| Low | Regular cleanup of unused snapshots and automated backups | COST-RDS-8 | Implement automated cleanup policies for old snapshots and manual backups | **Platform** - Create automated processes to delete old manual snapshots and optimize automated backup retention |


