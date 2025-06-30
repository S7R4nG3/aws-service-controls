# AWS Aurora Serverless
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management | IAM-01 | Implement strong identity and access management controls for Aurora Serverless access | **IaC** - Configure IAM policies with least privilege principles, implement MFA, and use IAM database authentication |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Ensure data at rest and in transit encryption for Aurora Serverless | **IaC** - Enable encryption at rest using AWS KMS and enforce SSL/TLS for data in transit |  |
| **High** | Infrastructure and Virtualization Security | IVS-01 | Secure network infrastructure and virtual environments | **IaC** - Deploy Aurora Serverless in private subnets with VPC security groups and NACLs |  |
| Medium | Application and Interface Security | AIS-01 | Secure application interfaces and APIs | **User** - Implement secure connection strings and API authentication mechanisms |  |
| Medium | Business Continuity Management and Operational Resilience | BCR-01 | Ensure business continuity and disaster recovery capabilities | **IaC** - Configure automated backups, cross-region replication, and point-in-time recovery |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Aurora Serverless | **IaC** - Implement IAM policies and database-level access controls with role-based permissions |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable SSL/TLS encryption for all database connections and configure certificate validation |  |
| **High** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **IaC** - Enable Aurora encryption at rest using AWS KMS with customer-managed keys |  |
| Medium | Audit Generation | AU-12 | Generate audit records for events with defined auditable content | **IaC** - Enable Aurora database activity monitoring and CloudTrail logging |  |
| Medium | Information System Backup | CP-9 | Conduct backups of user-level information and system-level information | **IaC** - Configure automated backups with appropriate retention periods and cross-region backup copies |  |
| Medium | Information System Monitoring | SI-4 | Monitor the information system to detect attacks and indicators of potential attacks | **IaC** - Enable CloudWatch monitoring, Performance Insights, and GuardDuty for threat detection |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | RDS DB instances should have encryption at rest enabled | RDS.3 | Aurora Serverless clusters should have encryption at rest enabled | **IaC** - Enable encryption when creating Aurora Serverless cluster using AWS KMS keys |  |
| **High** | RDS cluster snapshots and database snapshots should be encrypted | RDS.4 | Ensure Aurora Serverless snapshots are encrypted | **IaC** - Enable snapshot encryption through cluster encryption configuration |  |
| **High** | RDS DB instances should be configured with multiple Availability Zones | RDS.5 | Configure Aurora Serverless for high availability across multiple AZs | **IaC** - Deploy Aurora Serverless cluster across multiple AZs within the VPC |  |
| Medium | Enhanced monitoring should be configured for RDS DB instances | RDS.6 | Enable enhanced monitoring for Aurora Serverless clusters | **IaC** - Configure Performance Insights and CloudWatch enhanced monitoring |  |
| Medium | IAM authentication should be configured for RDS clusters | RDS.12 | Enable IAM database authentication for Aurora Serverless | **IaC** - Enable IAM database authentication in cluster configuration and create database users that use IAM authentication |  |
| Low | RDS automatic minor version upgrades should be enabled | RDS.13 | Enable automatic minor version upgrades for Aurora Serverless | **IaC** - Set auto_minor_version_upgrade parameter to true in cluster configuration |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | RDS DB instances should prohibit public read access | RDS.2 | Ensure Aurora Serverless clusters are not publicly accessible | **IaC** - Set publicly_accessible parameter to false and deploy in private subnets |  |
| **Critical** | RDS snapshots should be private | RDS.1 | Ensure Aurora Serverless snapshots are not shared publicly | **Platform** - AWS ensures snapshots are private by default, verify no public sharing is configured |  |
| Medium | RDS clusters should have deletion protection enabled | RDS.7 | Enable deletion protection for Aurora Serverless clusters | **IaC** - Set deletion_protection parameter to true in cluster configuration |  |
| Medium | RDS DB instances should have deletion protection enabled | RDS.8 | Enable deletion protection for production Aurora Serverless clusters | **IaC** - Configure deletion protection through cluster settings and IaC templates |  |
| Medium | RDS instances should have logging enabled | RDS.11 | Enable appropriate logging for Aurora Serverless clusters | **IaC** - Enable audit logs, error logs, general logs, and slow query logs as appropriate |  |


## Operational Controls
---



## Cost Controls
---


### AWS Aurora Serverless Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size Aurora Serverless capacity settings | COST-01 | Configure appropriate minimum and maximum Aurora Capacity Units (ACUs) based on workload requirements | **IaC** - Set min_capacity and max_capacity parameters based on performance testing and monitoring data |
| **High** | Implement auto-pause functionality | COST-02 | Enable auto-pause for development and testing environments to reduce costs during inactivity | **IaC** - Configure auto_pause and seconds_until_auto_pause parameters for non-production clusters |
| Medium | Optimize backup retention policies | COST-03 | Set appropriate backup retention periods to balance cost and recovery requirements | **IaC** - Configure backup_retention_period based on business requirements, typically 7-35 days |
| Medium | Monitor and optimize Aurora Serverless usage patterns | COST-04 | Regularly review CloudWatch metrics to optimize capacity settings and identify cost savings opportunities | **User** - Set up CloudWatch dashboards and alarms to monitor ACU utilization and costs |
| Medium | Use appropriate storage optimization | COST-05 | Implement data lifecycle policies and remove unnecessary data to optimize storage costs | **User** - Implement database maintenance procedures to archive or delete old data and optimize storage usage |
| Low | Implement cost allocation and tagging | COST-06 | Use consistent tagging strategy for cost allocation and tracking | **IaC** - Apply standardized tags for environment, project, cost center, and owner to enable cost tracking and allocation |
| Low | Optimize cross-region backup strategy | COST-07 | Evaluate the need for cross-region backups and snapshots to optimize data transfer costs | **IaC** - Configure cross-region automated backups only when required for disaster recovery, considering data transfer costs |


