# AWS Athena
---


### CSA CCM v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish comprehensive identity and access management policies for Athena service access | **Platform** - Configure IAM policies with least privilege access to Athena resources, implement role-based access control (RBAC) for query execution and data access |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-07 | Implement data classification and lifecycle management for data accessed through Athena | **IaC** - Define data classification schemas in S3 bucket policies and implement lifecycle rules for query results stored in result buckets |  |
| **High** | Encryption and Key Management | EKM-01 | Ensure encryption at rest and in transit for all Athena operations | **IaC** - Configure S3 bucket encryption for data sources and result locations, enable SSL/TLS for Athena API calls and JDBC/ODBC connections |  |
| **High** | Information Security Monitoring | IVS-01 | Implement comprehensive monitoring and logging for Athena service usage | **Platform** - Enable CloudTrail logging for Athena API calls, configure CloudWatch metrics for query performance and cost monitoring |  |
| Medium | Threat and Vulnerability Management | TVM-01 | Implement vulnerability scanning and threat detection for Athena workloads | **Platform** - Use AWS Security Hub and GuardDuty to monitor for suspicious query patterns and unauthorized access attempts |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Athena resources | **IaC** - Implement fine-grained IAM policies with resource-level permissions for databases, tables, and workgroups in Athena |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Enforce HTTPS/TLS 1.2+ for all Athena API communications and client connections |  |
| **High** | Event Logging | AU-2 | Ensure auditable events are defined and logged appropriately | **IaC** - Configure CloudTrail to log all Athena API calls including StartQueryExecution, GetQueryResults, and workgroup management activities |  |
| **High** | Boundary Protection | SC-7 | Monitor and control communications at external boundaries of the system | **IaC** - Configure VPC endpoints for Athena service access and implement network ACLs to restrict access to authorized networks |  |
| Medium | System Backup | CP-9 | Conduct backups of system-level information contained in the system | **User** - Implement backup strategies for Athena metadata catalogs and ensure S3 data sources have appropriate backup and versioning enabled |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Athena workgroups should be encrypted at rest | Athena.1 | Athena workgroups should have encryption at rest enabled for query results | **IaC** - Configure workgroups with ResultConfiguration encryption using SSE-S3, SSE-KMS, or CSE-KMS encryption |  |
| **Critical** | S3 Block Public Access setting should be enabled | S3.1 | S3 buckets used by Athena should have Block Public Access enabled | **IaC** - Enable S3 Block Public Access settings on all buckets used for Athena data sources and query results |  |
| **High** | IAM policies should not allow full administrative privileges | IAM.1 | IAM policies for Athena access should follow principle of least privilege | **IaC** - Create specific IAM policies for Athena that grant only necessary permissions for query execution and data access |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable CloudTrail logging to capture Athena API calls across all regions | **Platform** - Configure multi-region CloudTrail with S3 bucket logging and log file validation for Athena service events |  |
| Medium | AWS Config should be enabled | Config.1 | AWS Config should be enabled to track configuration changes to Athena resources | **Platform** - Enable AWS Config to monitor changes to Athena workgroups, data catalogs, and associated IAM policies |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Athena workgroups should encrypt query results | Athena.1 | Athena workgroups must encrypt query results to protect sensitive data | **IaC** - Configure workgroup ResultConfiguration with EncryptionConfiguration using appropriate KMS keys |  |
| **Critical** | S3 buckets should block public access | S3.8 | S3 buckets used with Athena should have public access blocked | **IaC** - Apply bucket policies and Block Public Access settings to prevent unauthorized access to Athena data sources |  |
| **High** | IAM customer managed policies should not allow decryption actions on all KMS keys | IAM.21 | Restrict KMS key access in IAM policies used for Athena encryption | **IaC** - Create specific IAM policies that grant decrypt permissions only for KMS keys used by Athena workgroups |  |
| **High** | A log metric filter and alarm should exist for usage of root user | CloudWatch.1 | Monitor root user access to Athena services through CloudWatch alarms | **Platform** - Create CloudWatch metric filters and alarms to detect root user access to Athena resources |  |
| Medium | VPC flow logs should be enabled | VPC.7 | Enable VPC flow logs to monitor network traffic to Athena VPC endpoints | **IaC** - Configure VPC flow logs for subnets containing Athena VPC endpoints to monitor network access patterns |  |


## Operational Controls
---



## Cost Controls
---


### AWS Athena Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Implement Query Result Caching | COST-001 | Enable query result caching to avoid re-running identical queries and reduce data scanning costs | **Platform** - Configure workgroups to enable query result reuse and set appropriate cache TTL values based on data refresh patterns |
| **High** | Optimize Data Storage Format | COST-002 | Use columnar storage formats like Parquet or ORC to reduce data scanning and improve query performance | **User** - Convert data sources to Parquet or ORC format with appropriate compression algorithms (SNAPPY, GZIP) to minimize storage costs and query execution time |
| **High** | Implement Data Partitioning | COST-003 | Partition data to limit the amount of data scanned during query execution | **User** - Design partition schemes based on common query patterns (date, region, department) and ensure queries include partition predicates to limit scanning |
| Medium | Configure Workgroup Query Limits | COST-004 | Set data usage controls and query limits to prevent runaway queries and unexpected costs | **IaC** - Configure workgroup settings with BytesScannedCutoffPerQuery limits and enable query result location enforcement to control costs |
| Medium | Optimize Query Performance | COST-005 | Write efficient queries to minimize data scanning and execution time | **User** - Use specific column selection instead of SELECT *, implement appropriate WHERE clauses, and utilize LIMIT clauses for exploratory queries |
| Medium | Implement S3 Lifecycle Policies | COST-006 | Configure lifecycle policies for query results and temporary data to reduce storage costs | **IaC** - Set up S3 lifecycle rules to automatically transition query results to cheaper storage classes (IA, Glacier) or delete them after defined retention periods |
| Medium | Monitor and Alert on Query Costs | COST-007 | Implement cost monitoring and alerting to track Athena usage and spending patterns | **Platform** - Configure CloudWatch metrics and billing alerts for Athena data scanning costs, and use AWS Cost Explorer to analyze usage patterns by workgroup and user |
| Low | Use Approximate Functions When Appropriate | COST-008 | Utilize approximate aggregate functions for large datasets when exact precision is not required | **User** - Replace exact functions like COUNT DISTINCT with approximate functions like approx_distinct() for large dataset analysis to reduce query execution time and costs |


