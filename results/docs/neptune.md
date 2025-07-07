# AWS Neptune
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v5 | Identity and Access Management | Implement robust identity and access management controls for Neptune cluster access | **IaC** - Configure IAM roles and policies with least privilege access, implement database authentication using IAM database authentication |  |
| **Critical** | DSI-01 | CSA Cloud Controls Matrix v5 | Data Security and Information Lifecycle Management | Ensure data encryption at rest and in transit for Neptune databases | **IaC** - Enable encryption at rest using KMS keys and enforce SSL/TLS for all connections |  |
| **High** | IVS-01 | CSA Cloud Controls Matrix v5 | Infrastructure and Virtualization Security | Implement network isolation and security groups for Neptune instances | **IaC** - Deploy Neptune in private subnets with restrictive security groups and NACLs |  |
| **High** | LOG-01 | CSA Cloud Controls Matrix v5 | Logging and Monitoring | Enable comprehensive logging and monitoring for Neptune activities | **IaC** - Enable CloudTrail, CloudWatch logs, and Neptune audit logs with appropriate retention policies |  |
| Medium | BCR-01 | CSA Cloud Controls Matrix v5 | Business Continuity Management and Operational Resilience | Implement backup and disaster recovery procedures for Neptune | **IaC** - Configure automated backups, point-in-time recovery, and cross-region replication |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage Neptune database accounts and access credentials | **User** - Implement regular account reviews, disable unused accounts, and use IAM database authentication |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission to and from Neptune | **IaC** - Enforce SSL/TLS encryption for all Neptune connections and disable unencrypted endpoints |  |
| **Critical** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Encrypt Neptune data at rest | **IaC** - Enable encryption at rest using customer-managed KMS keys with proper key rotation |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Log Neptune database events and administrative actions | **IaC** - Enable Neptune audit logs and CloudTrail logging for API calls |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Neptune performance and security events | **IaC** - Configure CloudWatch monitoring, set up alarms for anomalous activities |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Implement backup procedures for Neptune data | **IaC** - Configure automated backups with appropriate retention periods and test restore procedures |  |
| **Critical** | Neptune.1 | AWS Foundational Security Best Practices 1.0 | Neptune DB clusters should be encrypted | Neptune clusters must have encryption at rest enabled | **IaC** - Set StorageEncrypted parameter to true when creating Neptune clusters |  |
| **High** | Neptune.2 | AWS Foundational Security Best Practices 1.0 | Neptune DB clusters should have deletion protection enabled | Enable deletion protection to prevent accidental cluster deletion | **IaC** - Set DeletionProtection parameter to true for Neptune clusters |  |
| **High** | Neptune.3 | AWS Foundational Security Best Practices 1.0 | Neptune DB clusters should have backup retention configured | Configure appropriate backup retention period for Neptune clusters | **IaC** - Set BackupRetentionPeriod to minimum 7 days for production workloads |  |
| **High** | Neptune.4 | AWS Foundational Security Best Practices 1.0 | Neptune DB clusters should not be publicly accessible | Ensure Neptune clusters are not accessible from the internet | **IaC** - Deploy Neptune clusters in private subnets and configure security groups to restrict access |  |
| Medium | Neptune.5 | AWS Foundational Security Best Practices 1.0 | Neptune DB clusters should have IAM database authentication enabled | Enable IAM database authentication for Neptune clusters | **IaC** - Set EnableIAMDatabaseAuthentication parameter to true |  |
| **High** | Neptune.6 | AWS Security Hub 2023.1 | Neptune clusters should have logging enabled | Enable audit logging for Neptune clusters | **IaC** - Configure EnableCloudwatchLogsExports parameter to include audit logs |  |
| **High** | Neptune.7 | AWS Security Hub 2023.1 | Neptune clusters should be deployed across multiple Availability Zones | Deploy Neptune clusters with multi-AZ configuration for high availability | **IaC** - Configure DB subnet group with subnets in multiple AZs and enable multi-AZ deployment |  |
| Medium | Neptune.8 | AWS Security Hub 2023.1 | Neptune clusters should have automated minor version upgrades enabled | Enable automatic minor version upgrades for Neptune clusters | **IaC** - Set AutoMinorVersionUpgrade parameter to true for Neptune instances |  |
| Medium | Neptune.9 | AWS Security Hub 2023.1 | Neptune clusters should have appropriate maintenance windows configured | Configure maintenance windows for Neptune clusters to minimize business impact | **IaC** - Set PreferredMaintenanceWindow parameter to off-peak hours |  |
| Low | Neptune.10 | AWS Security Hub 2023.1 | Neptune clusters should have tags applied | Apply appropriate tags to Neptune clusters for resource management | **IaC** - Define and apply consistent tagging strategy including cost center, environment, and owner tags |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-001 | AWS Neptune Cost Optimization 2023 | Right-size Neptune instances | Select appropriate instance types based on workload requirements | **User** - Monitor CPU, memory, and I/O utilization to select optimal instance types and avoid over-provisioning |  |
| **High** | COST-002 | AWS Neptune Cost Optimization 2023 | Implement automated scaling | Use Neptune Auto Scaling to automatically adjust read replica count based on demand | **IaC** - Configure Auto Scaling policies for read replicas to scale based on CPU utilization or connection count |  |
| Medium | COST-003 | AWS Neptune Cost Optimization 2023 | Optimize backup retention | Set appropriate backup retention periods to balance cost and compliance requirements | **IaC** - Configure backup retention period based on business requirements, avoiding excessive retention periods |  |
| Medium | COST-004 | AWS Neptune Cost Optimization 2023 | Monitor and optimize data transfer costs | Minimize data transfer costs by optimizing query patterns and using appropriate endpoints | **User** - Use cluster endpoints for write operations and reader endpoints for read operations to minimize cross-AZ traffic |  |
| Medium | COST-005 | AWS Neptune Cost Optimization 2023 | Implement cost allocation tags | Apply consistent tagging strategy for cost allocation and tracking | **IaC** - Define and apply cost allocation tags including project, department, and environment for detailed cost tracking |  |
| Medium | COST-006 | AWS Neptune Cost Optimization 2023 | Schedule non-production environments | Stop or scale down non-production Neptune clusters during off-hours | **IaC** - Implement Lambda functions with EventBridge rules to automatically stop/start development and test environments |  |
| Low | COST-007 | AWS Neptune Cost Optimization 2023 | Monitor unused resources | Identify and remove unused Neptune clusters and read replicas | **User** - Regularly review CloudWatch metrics to identify low-utilization clusters and unused read replicas |  |
| Low | COST-008 | AWS Neptune Cost Optimization 2023 | Optimize storage usage | Monitor and optimize Neptune storage usage | **User** - Regular data lifecycle management and archiving of historical data to reduce storage costs |  |

