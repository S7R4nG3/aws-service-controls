# AWS Neptune
---


### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Account Management | AC-2 | Manage information system accounts, group memberships, privileges, workflow, notifications, deactivations, and authorizations | **IaC** - Configure IAM roles and policies for Neptune access with least privilege principles using CloudFormation or Terraform |  |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Implement IAM policies and resource-based policies to control access to Neptune clusters and databases |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable SSL/TLS encryption for all Neptune connections and configure security groups to allow only encrypted traffic |  |
| **High** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **IaC** - Enable encryption at rest for Neptune clusters using AWS KMS keys in CloudFormation templates |  |
| Medium | Audit Generation | AU-12 | Provide audit record generation capability for defined auditable events | **IaC** - Configure Neptune audit logging and integrate with CloudTrail for comprehensive audit trail |  |
| Medium | Boundary Protection | SC-7 | Monitor and control communications at the external boundary and key internal boundaries | **IaC** - Deploy Neptune in private subnets with proper VPC configuration and security group rules |  |
| Medium | Baseline Configuration | CM-2 | Develop, document, and maintain a current baseline configuration | **IaC** - Use Infrastructure as Code to maintain consistent Neptune cluster configurations and parameter groups |  |
| Medium | Information System Monitoring | SI-4 | Monitor the information system to detect attacks and indicators of potential attacks | **Platform** - Configure CloudWatch monitoring and alarms for Neptune performance metrics and security events |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Neptune DB clusters should be encrypted | Neptune.1 | Neptune DB clusters should have encryption enabled to protect data at rest | **IaC** - Set StorageEncrypted parameter to true and specify KmsKeyId in Neptune cluster configuration |  |
| Medium | Neptune DB clusters should publish audit logs to CloudWatch Logs | Neptune.2 | Neptune clusters should be configured to export audit logs to CloudWatch | **IaC** - Configure EnableCloudwatchLogsExports parameter with audit log types in cluster template |  |
| Medium | Neptune DB cluster snapshots should be encrypted | Neptune.3 | Neptune cluster snapshots should be encrypted to protect backup data | **IaC** - Ensure KmsKeyId is specified when creating manual snapshots and verify automatic snapshots inherit encryption |  |
| Low | Neptune DB clusters should have deletion protection enabled | Neptune.4 | Enable deletion protection to prevent accidental cluster deletion | **IaC** - Set DeletionProtection parameter to true in Neptune cluster configuration |  |
| Medium | Neptune DB clusters should have automated backups enabled | Neptune.5 | Ensure automated backups are enabled with appropriate retention period | **IaC** - Configure BackupRetentionPeriod with minimum 7 days and set PreferredBackupWindow |  |

### CSA Cloud Security Alliance 5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management | IAM-01 | Implement strong identity verification and access management controls | **User** - Configure multi-factor authentication for IAM users accessing Neptune and implement role-based access control |  |
| **High** | Encryption Key Management | EKM-02 | Implement proper encryption key management and rotation policies | **IaC** - Use AWS KMS with customer-managed keys for Neptune encryption and implement automatic key rotation |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and protection throughout its lifecycle | **User** - Classify Neptune data based on sensitivity and implement appropriate access controls and data retention policies |  |
| Medium | Infrastructure and Virtualization Security | IVS-01 | Secure the underlying infrastructure and virtualization layers | **IaC** - Deploy Neptune in isolated VPC with private subnets and implement network segmentation |  |
| Medium | Security Threat Assessment | STA-01 | Continuously assess and monitor for security threats | **Platform** - Enable AWS GuardDuty and Security Hub to monitor Neptune-related security events and threats |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Neptune clusters should be encrypted | Neptune.1 | This control checks whether Neptune clusters are encrypted | **IaC** - Enable encryption by setting StorageEncrypted to true in Neptune cluster resource definition |  |
| Medium | Neptune clusters should have audit logging enabled | Neptune.2 | This control checks whether Neptune clusters have audit logging enabled | **IaC** - Configure EnableCloudwatchLogsExports parameter to include audit logs in cluster template |  |
| Medium | Neptune clusters should be deployed across multiple Availability Zones | Neptune.3 | This control checks whether Neptune clusters are configured for high availability | **IaC** - Configure DB subnet group to span multiple AZs and deploy read replicas across different AZs |  |
| **High** | Neptune cluster snapshots should not be public | Neptune.4 | This control checks whether Neptune cluster snapshots are publicly accessible | **User** - Ensure manual snapshots are not shared publicly and implement IAM policies to prevent public sharing |  |
| Low | Neptune clusters should have automatic minor version upgrade enabled | Neptune.5 | This control checks whether automatic minor version upgrades are enabled | **IaC** - Set AutoMinorVersionUpgrade parameter to true in Neptune cluster configuration |  |


## Operational Controls
---



## Cost Controls
---


### AWS Neptune Cost Optimization 2023.1
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size Neptune instances | COST-01 | Monitor and adjust Neptune instance types based on actual workload requirements | **User** - Use CloudWatch metrics to monitor CPU, memory, and storage utilization and resize instances accordingly |
| Medium | Optimize backup retention periods | COST-02 | Set appropriate backup retention periods to balance data protection and storage costs | **IaC** - Configure BackupRetentionPeriod based on business requirements, typically 7-35 days |
| Medium | Use read replicas efficiently | COST-03 | Deploy read replicas only when needed and monitor their utilization | **User** - Monitor read replica utilization and remove underutilized replicas or scale down instance types |
| Medium | Implement automated start/stop for development environments | COST-04 | Automatically stop Neptune clusters during non-business hours for development and testing | **IaC** - Use Lambda functions with CloudWatch Events to automatically start/stop Neptune clusters on schedule |
| Medium | Monitor and optimize data transfer costs | COST-05 | Minimize cross-AZ and cross-region data transfer charges | **IaC** - Deploy applications in same AZ as Neptune primary instance and use VPC endpoints to reduce data transfer |
| Low | Clean up unused snapshots | COST-06 | Regularly review and delete unnecessary manual snapshots | **User** - Implement automated snapshot lifecycle management to delete old snapshots based on retention policies |
| Medium | Use appropriate instance families | COST-07 | Select cost-effective instance families based on workload characteristics | **IaC** - Choose T3 instances for variable workloads, R5 for memory-intensive, or customize based on performance requirements |
| Low | Implement resource tagging for cost allocation | COST-08 | Tag Neptune resources for accurate cost tracking and allocation | **IaC** - Apply consistent tagging strategy including cost center, environment, and project tags to all Neptune resources |


