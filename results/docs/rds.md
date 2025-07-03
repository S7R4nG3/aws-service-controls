# AWS RDS
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage information system accounts, including establishing, activating, modifying, reviewing, disabling, and removing accounts | **User** - Implement database user account management with least privilege principles, regular access reviews, and automated provisioning/deprovisioning |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure IAM policies, database security groups, and resource-based policies to enforce access controls |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable SSL/TLS encryption for database connections and configure force SSL parameters |  |
| **High** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect the confidentiality and integrity of information at rest | **IaC** - Enable RDS encryption at rest using AWS KMS keys for database instances and snapshots |  |
| Medium | AU-12 | NIST 800-53 Rev 5 | Audit Record Generation | Provide audit record generation capability for the events defined in AU-2 a at the information system components | **IaC** - Enable database audit logging and configure log exports to CloudWatch Logs |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of user-level information contained in the system | **IaC** - Configure automated backups, backup retention periods, and point-in-time recovery settings |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor events on the system and detect attack attempts | **Platform** - Enable Amazon GuardDuty for RDS Protection and configure CloudWatch monitoring with appropriate alarms |  |
| **High** | DSI-01 | CSA CCM v4.0.1 | Data Classification | Data and objects containing data shall be assigned a classification by a data owner based on data type, value, sensitivity, and criticality to the organization | **User** - Implement data classification policies and use AWS tags to classify RDS instances based on data sensitivity |  |
| **High** | DSI-02 | CSA CCM v4.0.1 | Data Handling/Labeling/Security Policy | Procedures shall be established, documented, approved, communicated, applied, evaluated and maintained for labeling, handling and the security of data and objects containing data | **User** - Establish data handling procedures specific to RDS environments including backup retention and deletion policies |  |
| **High** | EKM-01 | CSA CCM v4.0.1 | Entitlement | Encryption keys shall be managed, stored, and provisioned throughout their lifecycle with proper access, authorization, and use | **IaC** - Implement AWS KMS key management with appropriate key rotation and access policies for RDS encryption |  |
| Medium | IVS-01 | CSA CCM v4.0.1 | Clock Synchronization | System clocks shall be synchronized to a standard time source across all relevant information processing systems | **Platform** - Ensure RDS instances use synchronized time sources through AWS managed infrastructure |  |
| Medium | LOG-01 | CSA CCM v4.0.1 | Audit Logs / Audit Trails | Policies and procedures shall be established to maintain audit logs and audit trails | **IaC** - Configure comprehensive logging policies for RDS including slow query logs, error logs, and general logs |  |
| Medium | TVM-01 | CSA CCM v4.0.1 | Vulnerability / Patch Management | Policies and procedures shall be established to ensure timely and effective identification, analysis, prioritization, and resolution of vulnerabilities | **Platform** - Enable auto minor version upgrades and establish maintenance windows for RDS instances |  |
| **Critical** | RDS.1 | AWS Foundational Security Standard 1.0.0 | RDS snapshots should not be public | RDS snapshots should not be publicly accessible to prevent unauthorized access to database data | **IaC** - Ensure all RDS snapshots have public access disabled and implement snapshot sharing controls |  |
| **Critical** | RDS.2 | AWS Foundational Security Standard 1.0.0 | RDS DB instances should prohibit public access | RDS instances should not be publicly accessible unless specifically required | **IaC** - Configure RDS instances with PubliclyAccessible set to false and place in private subnets |  |
| **High** | RDS.3 | AWS Foundational Security Standard 1.0.0 | RDS DB instances should have encryption at rest enabled | RDS instances should be encrypted to protect sensitive data at rest | **IaC** - Enable encryption at rest for RDS instances using AWS KMS customer managed keys |  |
| **High** | RDS.4 | AWS Foundational Security Standard 1.0.0 | RDS cluster snapshots and database snapshots should be encrypted at rest | RDS snapshots should be encrypted to protect backup data | **IaC** - Ensure automated and manual snapshots are encrypted using KMS encryption |  |
| Medium | RDS.5 | AWS Foundational Security Standard 1.0.0 | RDS DB instances should be configured with multiple Availability Zones | RDS instances should use Multi-AZ deployments for high availability | **IaC** - Configure Multi-AZ deployment for production RDS instances to ensure high availability |  |
| Medium | RDS.8 | AWS Foundational Security Standard 1.0.0 | RDS DB instances should have deletion protection enabled | Deletion protection helps prevent accidental deletion of RDS instances | **IaC** - Enable deletion protection for production RDS instances to prevent accidental deletion |  |
| Low | RDS.6 | AWS Foundational Security Standard 1.0.0 | Enhanced monitoring should be configured for RDS DB instances | Enhanced monitoring provides real-time operating system metrics for RDS instances | **IaC** - Enable enhanced monitoring with appropriate granularity and IAM role configuration |  |
| **High** | RDS.10 | AWS Security Hub 2023.11 | RDS instances should be deployed in a VPC | RDS instances should be deployed within a VPC for network isolation | **IaC** - Deploy RDS instances within VPC with proper subnet groups and security group configurations |  |
| Medium | RDS.11 | AWS Security Hub 2023.11 | RDS instances should have logging enabled | RDS instances should have appropriate logging enabled for security monitoring | **IaC** - Enable and configure database engine specific logs and export to CloudWatch Logs |  |
| Medium | RDS.12 | AWS Security Hub 2023.11 | RDS instances should use a custom administrator username | RDS instances should not use default administrator usernames | **User** - Configure custom administrator usernames instead of default names like 'admin' or 'root' |  |
| Medium | RDS.14 | AWS Security Hub 2023.11 | Amazon Aurora clusters should have backtracking enabled | Aurora clusters should enable backtracking for point-in-time recovery capabilities | **IaC** - Enable backtracking for Aurora clusters with appropriate retention period configuration |  |
| Medium | RDS.15 | AWS Security Hub 2023.11 | RDS DB clusters should be configured for multiple Availability Zones | RDS clusters should span multiple AZs for high availability | **IaC** - Configure RDS clusters across multiple Availability Zones for fault tolerance |  |
| Low | RDS.13 | AWS Security Hub 2023.11 | RDS automatic minor version upgrades should be enabled | Automatic minor version upgrades help ensure instances receive security patches | **IaC** - Enable auto minor version upgrades and configure appropriate maintenance windows |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS RDS Cost Optimization 2024 | Right-size RDS instances | Regularly review and optimize RDS instance types and sizes based on actual usage patterns | **User** - Use CloudWatch metrics and AWS Trusted Advisor to identify underutilized instances and resize appropriately |  |
| **High** | COST-02 | AWS RDS Cost Optimization 2024 | Utilize Reserved Instances | Purchase Reserved Instances for predictable workloads to achieve significant cost savings | **User** - Analyze usage patterns and purchase 1-year or 3-year Reserved Instances for steady-state workloads |  |
| Medium | COST-03 | AWS RDS Cost Optimization 2024 | Implement automated start/stop scheduling | Automatically stop non-production RDS instances during off-hours to reduce costs | **IaC** - Use AWS Lambda functions with EventBridge rules to automate RDS instance start/stop schedules |  |
| Medium | COST-04 | AWS RDS Cost Optimization 2024 | Optimize storage configuration | Choose appropriate storage types and configure storage autoscaling to optimize costs | **IaC** - Use GP3 storage for cost optimization and enable storage autoscaling to avoid over-provisioning |  |
| Medium | COST-05 | AWS RDS Cost Optimization 2024 | Manage backup retention | Optimize backup retention periods to balance data protection needs with storage costs | **IaC** - Set appropriate backup retention periods and use snapshot lifecycle policies for long-term archival |  |
| Low | COST-06 | AWS RDS Cost Optimization 2024 | Monitor and optimize data transfer costs | Minimize data transfer costs by optimizing cross-AZ and cross-region data movement | **User** - Use VPC endpoints where possible and optimize application architecture to reduce data transfer |  |
| Low | COST-07 | AWS RDS Cost Optimization 2024 | Implement resource tagging for cost allocation | Use consistent tagging strategies to track and allocate RDS costs across teams and projects | **IaC** - Implement comprehensive tagging strategy with cost center, environment, and project tags |  |

