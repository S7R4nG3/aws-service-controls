# AWS Elastic File System (EFS)
---


### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure EFS file system policies and POSIX permissions to restrict access based on IAM roles and resource-based policies |  |
| HIGH | Least Privilege | AC-6 | Employ the principle of least privilege for specific security functions and privileged accounts | **IaC** - Grant minimal required EFS permissions through IAM policies and file system access points with specific path restrictions |  |
| HIGH | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable encryption in transit for EFS mount targets using TLS and configure EFS utils with encryption parameters |  |
| HIGH | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **IaC** - Enable EFS encryption at rest using AWS KMS keys and configure appropriate key policies for access control |  |
| MEDIUM | Event Logging | AU-2 | Identify the types of events that the system is capable of logging | **IaC** - Enable CloudTrail logging for EFS API calls and configure VPC Flow Logs for network traffic monitoring |  |
| MEDIUM | Audit Record Review, Analysis, and Reporting | AU-6 | Review and analyze system audit records for indications of inappropriate or unusual activity | **Platform** - Configure CloudWatch alarms and AWS Config rules to monitor EFS access patterns and configuration changes |  |
| MEDIUM | Boundary Protection | SC-7 | Monitor and control communications at the external interfaces to the system | **IaC** - Configure VPC security groups and NACLs to restrict EFS mount target access to authorized subnets and ports |  |
| MEDIUM | System Backup | CP-9 | Conduct backups of user-level and system-level information | **IaC** - Enable EFS automatic backups using AWS Backup service and configure backup retention policies |  |

### CSA CCM v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Identity and Access Management Policy and Procedures | IAM-01 | Policies and procedures shall be established for identity and access management | **User** - Establish and document EFS access management procedures including role-based access controls and regular access reviews |  |
| HIGH | Data Security and Information Lifecycle Management | DSI-01 | Policies and procedures shall be established for data security and information lifecycle management | **User** - Implement data classification policies for EFS stored data and establish lifecycle management procedures |  |
| HIGH | Encryption and Key Management | EKM-01 | Policies and procedures shall be established for encryption and key management | **IaC** - Implement AWS KMS key management for EFS encryption with proper key rotation and access policies |  |
| MEDIUM | Infrastructure and Virtualization Security | IVS-01 | Policies and procedures shall be established for infrastructure and virtualization security | **IaC** - Configure EFS mount targets within secure VPC subnets with appropriate network segmentation |  |
| MEDIUM | Business Continuity Management and Operational Resilience | BCR-01 | Policies and procedures shall be established for business continuity management | **User** - Establish business continuity plans that include EFS backup and recovery procedures across multiple AZs |  |
| MEDIUM | Audit Logging / Intrusion Detection | LOG-01 | Policies and procedures shall be established for audit logging and intrusion detection | **IaC** - Configure comprehensive logging for EFS access through CloudTrail and implement monitoring for suspicious activities |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| MEDIUM | Amazon EFS should be configured to encrypt file data at rest | EFS.1 | EFS file systems should have encryption at rest enabled to protect sensitive data | **IaC** - Enable encryption at rest during EFS creation or modify existing file systems to enable encryption |  |
| MEDIUM | Amazon EFS volumes should be in backup plans | EFS.2 | EFS file systems should be included in AWS Backup plans for data protection | **IaC** - Create AWS Backup plans that include EFS file systems and configure appropriate backup schedules |  |
| MEDIUM | EFS access points should enforce a user identity | EFS.3 | EFS access points should be configured with POSIX user and group information | **IaC** - Configure EFS access points with CreationInfo to enforce specific POSIX user and group identities |  |
| MEDIUM | EFS access points should enforce a root directory | EFS.4 | EFS access points should specify a root directory to limit file system access | **IaC** - Configure EFS access points with RootDirectory to restrict access to specific file system paths |  |

### AWS Security Hub Latest
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| MEDIUM | EFS file systems should encrypt data at rest | EFS.1 | Checks whether Amazon EFS file systems are encrypted at rest | **IaC** - Enable encryption at rest for EFS file systems using AWS managed or customer managed KMS keys |  |
| MEDIUM | EFS file systems should be in a backup plan | EFS.2 | Checks whether EFS file systems are added to AWS Backup plans | **IaC** - Include EFS file systems in AWS Backup plans with appropriate backup frequency and retention |  |
| LOW | EFS access points should enforce a user identity | EFS.3 | Checks whether EFS access points are configured to enforce a user identity | **IaC** - Configure EFS access points with POSIX user information to enforce specific user identities |  |
| LOW | EFS access points should enforce a root directory | EFS.4 | Checks whether EFS access points enforce a root directory | **IaC** - Configure EFS access points with root directory settings to limit file system access scope |  |


## Operational Controls
---



## Cost Controls
---


### AWS EFS Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| HIGH | Implement EFS Intelligent Tiering | COST-1 | Enable Intelligent Tiering to automatically move files to lower-cost storage classes | **IaC** - Configure EFS lifecycle policies to automatically transition files to Infrequent Access (IA) storage class based on access patterns |
| HIGH | Use Provisioned Throughput Only When Necessary | COST-2 | Avoid provisioned throughput mode unless burst throughput is insufficient | **User** - Monitor EFS performance metrics and only enable provisioned throughput when burst mode limits are consistently exceeded |
| MEDIUM | Monitor and Optimize EFS Usage | COST-3 | Regularly monitor EFS usage patterns and optimize storage allocation | **Platform** - Use CloudWatch metrics to track EFS usage, identify unused files, and implement data cleanup policies |
| MEDIUM | Right-size EFS Performance Mode | COST-4 | Choose appropriate performance mode based on actual requirements | **IaC** - Use General Purpose mode for most workloads and only use Max I/O mode when higher IOPS are required |
| MEDIUM | Implement Data Lifecycle Management | COST-5 | Establish policies for data retention and deletion to minimize storage costs | **User** - Create automated processes to identify and remove obsolete data from EFS file systems |
| LOW | Use Regional Storage Classes Appropriately | COST-6 | Choose between Standard and One Zone storage classes based on availability requirements | **IaC** - Use EFS One Zone storage class for workloads that don't require multi-AZ resilience to reduce costs |
| LOW | Optimize Mount Target Configuration | COST-7 | Configure mount targets efficiently to minimize network costs | **IaC** - Place mount targets in the same AZ as compute resources to minimize cross-AZ data transfer charges |


