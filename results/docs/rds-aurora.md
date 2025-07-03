# AWS Aurora
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Identity and Access Management | Implement strong identity and access management controls for Aurora database access | **Platform** - Configure IAM database authentication and use AWS IAM roles for Aurora access control |  |
| **Critical** | DSI-02 | CSA Cloud Controls Matrix v4.0 | Data Security and Information Lifecycle Management | Encrypt data at rest and in transit for Aurora databases | **IaC** - Enable encryption at rest using AWS KMS and enforce SSL/TLS connections |  |
| **High** | DSI-04 | CSA Cloud Controls Matrix v4.0 | Data Loss Prevention | Implement automated backup and point-in-time recovery for Aurora | **Platform** - Configure automated backups with appropriate retention periods and enable backtrack feature |  |
| **High** | CCM-09 | CSA Cloud Controls Matrix v4.0 | Change Control and Configuration Management | Manage Aurora configuration changes through controlled processes | **IaC** - Use CloudFormation or Terraform to manage Aurora cluster configurations |  |
| Medium | IVS-01 | CSA Cloud Controls Matrix v4.0 | Infrastructure and Virtualization Security | Secure Aurora network infrastructure and isolation | **IaC** - Deploy Aurora in private subnets with appropriate security groups and NACLs |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage Aurora database user accounts and access privileges | **User** - Implement least privilege access and regularly review database user accounts |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission to and from Aurora databases | **Platform** - Enforce SSL/TLS encryption for all Aurora connections |  |
| **Critical** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Encrypt Aurora data at rest using strong encryption | **IaC** - Enable Aurora encryption using AWS KMS with customer-managed keys |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Configure comprehensive logging for Aurora database activities | **Platform** - Enable Aurora audit logging and CloudTrail for API activities |  |
| **High** | CP-9 | NIST 800-53 Rev 5 | Information System Backup | Implement comprehensive backup strategies for Aurora | **Platform** - Configure automated backups, manual snapshots, and cross-region backup replication |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | Information System Monitoring | Monitor Aurora database performance and security events | **Platform** - Implement CloudWatch monitoring and Performance Insights for Aurora |  |
| Medium | CA-7 | NIST 800-53 Rev 5 | Continuous Monitoring | Establish continuous monitoring of Aurora security posture | **Platform** - Use AWS Config rules and Security Hub for continuous Aurora compliance monitoring |  |
| **Critical** | RDS.1 | AWS Foundational Security Best Practices 1.0.0 | RDS snapshots should be private | Ensure Aurora snapshots are not publicly accessible | **Platform** - Verify Aurora snapshots have private visibility and restrict sharing |  |
| **Critical** | RDS.2 | AWS Foundational Security Best Practices 1.0.0 | RDS DB instances should prohibit public read access | Ensure Aurora clusters are not publicly accessible | **IaC** - Set publicly_accessible to false in Aurora cluster configuration |  |
| **Critical** | RDS.3 | AWS Foundational Security Best Practices 1.0.0 | RDS DB instances should have encryption at rest enabled | Enable encryption at rest for Aurora clusters | **IaC** - Enable storage_encrypted parameter in Aurora cluster configuration |  |
| **High** | RDS.4 | AWS Foundational Security Best Practices 1.0.0 | RDS cluster snapshots and database snapshots should be encrypted at rest | Ensure Aurora snapshots are encrypted | **Platform** - Enable encryption for Aurora cluster snapshots using KMS keys |  |
| **High** | RDS.5 | AWS Foundational Security Best Practices 1.0.0 | RDS DB instances should be configured with multiple Availability Zones | Deploy Aurora clusters across multiple AZs for high availability | **IaC** - Configure Aurora cluster with multiple availability zones |  |
| Medium | RDS.6 | AWS Foundational Security Best Practices 1.0.0 | Enhanced monitoring should be configured for RDS DB instances | Enable enhanced monitoring for Aurora instances | **IaC** - Configure monitoring_interval parameter for Aurora instances |  |
| **High** | RDS.7 | AWS Security Hub 2023 | RDS clusters should have deletion protection enabled | Enable deletion protection for Aurora clusters | **IaC** - Set deletion_protection to true in Aurora cluster configuration |  |
| **High** | RDS.8 | AWS Security Hub 2023 | RDS DB instances should have deletion protection enabled | Enable deletion protection for Aurora instances | **IaC** - Set deletion_protection to true for Aurora instances |  |
| **High** | RDS.9 | AWS Security Hub 2023 | Database logging should be enabled | Enable comprehensive logging for Aurora databases | **Platform** - Enable Aurora audit logs, error logs, and slow query logs |  |
| Medium | RDS.10 | AWS Security Hub 2023 | IAM authentication should be configured for RDS instances | Enable IAM database authentication for Aurora | **IaC** - Enable iam_database_authentication_enabled parameter for Aurora cluster |  |
| Medium | RDS.11 | AWS Security Hub 2023 | RDS instances should have automatic backups enabled | Configure automatic backups for Aurora clusters | **IaC** - Set backup_retention_period to appropriate value for Aurora cluster |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-02 | AWS Aurora Cost Optimization Best Practices 2023 | Implement Aurora Serverless for variable workloads | Use Aurora Serverless v2 for unpredictable or intermittent workloads | **IaC** - Configure Aurora Serverless v2 with appropriate scaling configuration for variable workloads |  |
| **High** | COST-01 | AWS Aurora Cost Optimization Best Practices 2023 | Right-size Aurora instances | Monitor and adjust Aurora instance sizes based on actual workload requirements | **User** - Use CloudWatch metrics and Performance Insights to identify oversized instances and downsize appropriately |  |
| **High** | COST-03 | AWS Aurora Cost Optimization Best Practices 2023 | Optimize Aurora storage costs | Monitor storage growth and implement data lifecycle management | **User** - Regularly archive or delete old data and monitor storage metrics in CloudWatch |  |
| Medium | COST-04 | AWS Aurora Cost Optimization Best Practices 2023 | Use Reserved Instances for predictable workloads | Purchase Reserved Instances for long-term, predictable Aurora workloads | **User** - Analyze usage patterns and purchase 1-year or 3-year Reserved Instances for consistent workloads |  |
| Medium | COST-05 | AWS Aurora Cost Optimization Best Practices 2023 | Optimize backup retention and cross-region replication | Set appropriate backup retention periods and minimize unnecessary cross-region backups | **IaC** - Configure backup_retention_period based on business requirements and limit cross-region backup replication |  |
| Medium | COST-06 | AWS Aurora Cost Optimization Best Practices 2023 | Monitor and optimize data transfer costs | Minimize data transfer costs by optimizing application architecture and read replica placement | **User** - Place read replicas in same regions as applications and optimize query patterns to reduce data transfer |  |
| Medium | COST-07 | AWS Aurora Cost Optimization Best Practices 2023 | Implement automated scaling policies | Use Aurora Auto Scaling for read replicas to automatically adjust capacity | **IaC** - Configure Aurora Auto Scaling policies for read replicas based on CPU utilization or connection metrics |  |
| Low | COST-08 | AWS Aurora Cost Optimization Best Practices 2023 | Regular cost monitoring and alerting | Set up cost monitoring and alerts for Aurora spending | **Platform** - Configure AWS Cost Explorer alerts and budgets for Aurora service costs |  |

