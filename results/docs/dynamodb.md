# AWS DynamoDB
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish formal identity and access management policies for DynamoDB resources including user provisioning, authentication, and authorization procedures. | **Platform/IaC** - Configure IAM policies, roles, and resource-based policies for DynamoDB tables. Implement least privilege access using IAM policy conditions and attribute-based access control. |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Implement encryption for data at rest and in transit, establish data classification schemes, and manage data lifecycle for DynamoDB tables. | **IaC** - Enable DynamoDB encryption at rest using AWS KMS customer-managed keys, configure encryption in transit using TLS, implement Point-in-Time Recovery and backup policies. |  |
| **High** | Governance and Risk Management | GRM-01 | Establish governance frameworks for DynamoDB usage including risk assessment, compliance monitoring, and security controls validation. | **Platform** - Use AWS Config rules to monitor DynamoDB compliance, implement AWS Security Hub for centralized security findings, establish tagging strategies for governance. |  |
| Medium | Interoperability and Portability | IVS-01 | Ensure secure data exchange and API security for DynamoDB integrations with other services and applications. | **User/IaC** - Implement VPC endpoints for DynamoDB, configure API Gateway with proper authentication for DynamoDB access, use secure SDK configurations. |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to DynamoDB resources in accordance with applicable access control policies. | **IaC** - Configure IAM policies with explicit deny statements, implement resource-based policies, use condition keys for fine-grained access control, enable DynamoDB resource-level permissions. |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information between clients and DynamoDB service. | **User/IaC** - Configure DynamoDB client SDKs to use HTTPS/TLS encryption, implement VPC endpoints to keep traffic within AWS network, validate SSL certificates in applications. |  |
| **Critical** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest in DynamoDB tables and backups. | **IaC** - Enable DynamoDB encryption at rest using AWS managed or customer-managed KMS keys, encrypt DynamoDB backups, implement proper key management policies. |  |
| **High** | Event Logging | AU-2 | Identify the types of events to be logged for DynamoDB operations and maintain audit logs. | **Platform/IaC** - Enable AWS CloudTrail for DynamoDB API logging, configure DynamoDB Streams for data-level events, implement CloudWatch logging for application-level events. |  |
| **High** | System Monitoring | SI-4 | Monitor DynamoDB for attacks and indicators of potential attacks in accordance with monitoring objectives. | **Platform/IaC** - Configure CloudWatch metrics and alarms for DynamoDB performance and security events, implement GuardDuty for threat detection, use AWS Security Hub for centralized monitoring. |  |
| **High** | System Backup | CP-9 | Conduct backups of DynamoDB data and system-level information contained in the system. | **IaC** - Enable DynamoDB Point-in-Time Recovery, configure automated backups, implement cross-region backup replication, establish backup retention policies. |  |
| Medium | Vulnerability Monitoring and Scanning | RA-5 | Monitor and scan for vulnerabilities in DynamoDB configurations and access patterns. | **Platform** - Use AWS Config rules for DynamoDB security configuration compliance, implement AWS Security Hub findings review, conduct regular access reviews and policy audits. |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | DynamoDB tables should automatically scale capacity with demand | DynamoDB.1 | DynamoDB tables should use on-demand billing mode or have auto scaling enabled to prevent throttling and potential denial of service. | **IaC** - Configure DynamoDB tables with on-demand billing mode or enable auto scaling for read and write capacity units with appropriate scaling policies. |  |
| **High** | DynamoDB tables should have point-in-time recovery enabled | DynamoDB.2 | Enable point-in-time recovery to protect against accidental writes or deletes to DynamoDB tables. | **IaC** - Enable Point-in-Time Recovery (PITR) on DynamoDB tables through CloudFormation, Terraform, or AWS CLI with appropriate retention settings. |  |
| **High** | DynamoDB Accelerator (DAX) clusters should be encrypted at rest | DynamoDB.3 | Enable encryption at rest for DynamoDB Accelerator clusters to protect cached data. | **IaC** - Configure DAX clusters with encryption at rest enabled using AWS managed or customer-managed KMS keys during cluster creation. |  |

### AWS Security Hub 2023
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | DynamoDB tables should be encrypted at rest | DynamoDB.4 | Ensure DynamoDB tables have encryption at rest enabled to protect sensitive data from unauthorized access. | **IaC** - Configure server-side encryption for DynamoDB tables using AWS managed keys (SSE-S3) or customer-managed KMS keys (SSE-KMS) based on data sensitivity requirements. |  |
| **High** | DynamoDB tables should have deletion protection enabled | DynamoDB.5 | Enable deletion protection on DynamoDB tables to prevent accidental deletion of critical data. | **IaC** - Set the DeletionProtectionEnabled attribute to true when creating or updating DynamoDB tables through infrastructure as code templates. |  |
| **High** | DynamoDB should have AWS CloudTrail logging configured | DynamoDB.6 | Configure CloudTrail to log DynamoDB API calls for security monitoring and compliance auditing. | **Platform** - Ensure CloudTrail is enabled with data events configured for DynamoDB tables, store logs in encrypted S3 buckets with appropriate retention policies. |  |


## Operational Controls
---



## Cost Controls
---


### AWS DynamoDB Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Choose the Right Billing Mode | COST-1 | Select between On-Demand and Provisioned billing modes based on traffic patterns to optimize costs. | **IaC** - Analyze traffic patterns and choose On-Demand for unpredictable workloads or Provisioned mode with auto-scaling for predictable workloads. Monitor and adjust based on CloudWatch metrics. |
| **High** | Implement Table and Index Design Optimization | COST-2 | Design efficient table structures and indexes to minimize read/write capacity requirements and storage costs. | **User** - Use single-table design patterns, optimize partition key distribution, minimize Global Secondary Index usage, and compress large attribute values. Regularly review access patterns. |
| Medium | Enable Auto Scaling for Provisioned Tables | COST-3 | Configure auto scaling to automatically adjust capacity based on actual usage patterns and avoid over-provisioning. | **IaC** - Enable DynamoDB auto scaling with appropriate target utilization percentages (70-80%), set minimum and maximum capacity limits, and configure scaling policies for tables and indexes. |
| Medium | Optimize Data Storage and Lifecycle | COST-4 | Implement data archiving strategies and use appropriate data types to minimize storage costs. | **User/IaC** - Use TTL (Time To Live) for automatic data expiration, archive old data to S3, compress large items, and use efficient data types. Implement data lifecycle policies. |
| Medium | Monitor and Analyze Cost Usage | COST-5 | Implement comprehensive cost monitoring and alerting to identify optimization opportunities. | **Platform** - Use AWS Cost Explorer, set up billing alerts, implement detailed tagging strategies, monitor DynamoDB Contributor Insights, and regularly review AWS Trusted Advisor recommendations. |
| Low | Use Reserved Capacity for Predictable Workloads | COST-6 | Purchase DynamoDB Reserved Capacity for stable, predictable workloads to achieve significant cost savings. | **Platform** - Analyze historical usage patterns and purchase one-year or three-year Reserved Capacity commitments for baseline capacity requirements. Monitor utilization to ensure optimal ROI. |
| Low | Optimize Global Tables and Cross-Region Replication | COST-7 | Carefully plan Global Tables deployment to minimize cross-region data transfer and replication costs. | **IaC** - Deploy Global Tables only in required regions, optimize data synchronization patterns, monitor cross-region transfer costs, and consider regional data locality requirements. |


