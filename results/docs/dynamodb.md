# AWS DynamoDB
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v4.0.10 | Entitlement | Policies and procedures shall be established to govern the entitlement or access rights granted to customer/tenant access to data and functions within the cloud service | **Platform** - Implement IAM policies with least privilege access, use IAM roles for applications, and enforce MFA for administrative access |  |
| **Critical** | EKM-02 | CSA Cloud Controls Matrix v4.0.10 | Key Generation | Encryption keys shall be generated using approved algorithms and key sizes as defined in recognized standards | **IaC** - Configure DynamoDB encryption at rest using AWS KMS keys, implement customer-managed keys for sensitive data |  |
| **High** | AIS-01 | CSA Cloud Controls Matrix v4.0.10 | Anti-Virus / Malicious Software | Comprehensive logging and monitoring for security events and malicious activities | **IaC** - Enable AWS CloudTrail for API logging, configure DynamoDB Streams for data change monitoring |  |
| **High** | DSI-01 | CSA Cloud Controls Matrix v4.0.10 | Classification | Data and objects containing data shall be assigned a classification by the data owner based on data type, jurisdiction, applicable regulation and sensitivity | **User** - Classify DynamoDB data based on sensitivity, implement data retention policies using TTL features |  |
| Medium | IVS-01 | CSA Cloud Controls Matrix v4.0.10 | Baseline Requirements | Baseline security requirements shall be applied to networks based on their classification and taking into account zone concept implementations | **IaC** - Configure VPC endpoints for private access, implement security groups and NACLs |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage DynamoDB access accounts and credentials with proper lifecycle management | **Platform** - Use IAM service accounts and roles, implement automated account provisioning and deprovisioning |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission to and from DynamoDB using encryption | **IaC** - Enforce HTTPS/TLS 1.2+ for all DynamoDB connections, configure SDK with SSL/TLS settings |  |
| **Critical** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Encrypt sensitive data stored in DynamoDB tables | **IaC** - Enable DynamoDB encryption at rest using AWS managed or customer managed KMS keys |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Generate audit logs for all DynamoDB operations and access attempts | **IaC** - Configure CloudTrail data events for DynamoDB, enable VPC Flow Logs for network monitoring |  |
| **High** | CP-9 | NIST 800-53 Rev 5 | Information System Backup | Implement backup and recovery procedures for DynamoDB data | **IaC** - Enable DynamoDB point-in-time recovery, configure automated backups and cross-region replication |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | Information System Monitoring | Monitor DynamoDB for unauthorized access and anomalous activity | **Platform** - Configure CloudWatch alarms for DynamoDB metrics, implement GuardDuty for threat detection |  |
| Medium | DynamoDB.1 | AWS Foundational Security Best Practices 1.0.0 | DynamoDB tables should automatically scale capacity with demand | Configure DynamoDB tables with on-demand billing or auto-scaling to handle traffic spikes securely | **IaC** - Enable DynamoDB on-demand billing mode or configure auto-scaling for provisioned capacity |  |
| Medium | DynamoDB.2 | AWS Foundational Security Best Practices 1.0.0 | DynamoDB tables should have point-in-time recovery enabled | Enable point-in-time recovery to protect against accidental data loss or corruption | **IaC** - Enable point-in-time recovery on all DynamoDB tables containing critical business data |  |
| Medium | DynamoDB.3 | AWS Foundational Security Best Practices 1.0.0 | DynamoDB Accelerator (DAX) clusters should be encrypted at rest | Encrypt DAX clusters to protect cached data at rest | **IaC** - Configure DAX clusters with encryption at rest using AWS managed keys |  |
| Low | DynamoDB.4 | AWS Foundational Security Best Practices 1.0.0 | DynamoDB tables should be present in a backup plan | Include DynamoDB tables in AWS Backup plans for centralized backup management | **IaC** - Create AWS Backup plans that include DynamoDB tables with appropriate retention policies |  |
| Medium | DynamoDB.1 | AWS Security Hub 2023.04 | DynamoDB tables should automatically scale capacity with demand | Ensure DynamoDB tables can handle unexpected load without service disruption | **IaC** - Configure auto-scaling policies or use on-demand billing mode for DynamoDB tables |  |
| Medium | DynamoDB.2 | AWS Security Hub 2023.04 | DynamoDB tables should have point-in-time recovery enabled | Enable continuous backups to recover from data corruption or accidental deletion | **IaC** - Enable PITR on production DynamoDB tables and configure appropriate retention periods |  |
| Medium | DynamoDB.3 | AWS Security Hub 2023.04 | DynamoDB Accelerator clusters should be encrypted at rest | Protect sensitive data cached in DAX clusters through encryption | **IaC** - Enable server-side encryption for DAX clusters using AWS KMS keys |  |
| Low | DynamoDB.6 | AWS Security Hub 2023.04 | DynamoDB tables should have deletion protection enabled | Prevent accidental deletion of critical DynamoDB tables | **IaC** - Enable deletion protection on production DynamoDB tables through table configuration |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-001 | AWS DynamoDB Cost Optimization Best Practices 2023 | Choose Appropriate Billing Mode | Select the most cost-effective billing mode based on traffic patterns | **IaC** - Analyze traffic patterns and choose between on-demand and provisioned capacity modes, switch to on-demand for unpredictable workloads |  |
| **High** | COST-002 | AWS DynamoDB Cost Optimization Best Practices 2023 | Implement Data Lifecycle Management | Use TTL to automatically delete expired data and reduce storage costs | **User** - Configure Time To Live (TTL) attributes on DynamoDB items to automatically expire and delete old data |  |
| **High** | COST-003 | AWS DynamoDB Cost Optimization Best Practices 2023 | Optimize Table Design for Cost Efficiency | Design tables with efficient partition keys and minimal secondary indexes | **User** - Use composite partition keys for even data distribution, minimize Global Secondary Indexes, and avoid hot partitions |  |
| Medium | COST-004 | AWS DynamoDB Cost Optimization Best Practices 2023 | Enable Auto Scaling | Configure auto-scaling to automatically adjust capacity based on demand | **IaC** - Enable DynamoDB auto-scaling with appropriate target utilization percentages and scaling policies |  |
| Medium | COST-005 | AWS DynamoDB Cost Optimization Best Practices 2023 | Monitor and Optimize Reserved Capacity | Purchase reserved capacity for predictable workloads to reduce costs | **Platform** - Analyze usage patterns and purchase DynamoDB reserved capacity for consistent workloads to achieve up to 76% cost savings |  |
| Medium | COST-006 | AWS DynamoDB Cost Optimization Best Practices 2023 | Implement Efficient Query Patterns | Use efficient query and scan operations to minimize consumed capacity units | **User** - Use Query instead of Scan operations, implement pagination, use projection expressions to limit returned attributes |  |
| Medium | COST-007 | AWS DynamoDB Cost Optimization Best Practices 2023 | Archive Cold Data | Move infrequently accessed data to more cost-effective storage solutions | **User** - Export old data to S3 using DynamoDB export features, implement data archiving strategies for cold data |  |
| Low | COST-008 | AWS DynamoDB Cost Optimization Best Practices 2023 | Monitor Cost and Usage | Implement comprehensive cost monitoring and alerting for DynamoDB usage | **Platform** - Set up CloudWatch billing alarms, use AWS Cost Explorer to analyze DynamoDB costs, implement cost allocation tags |  |

