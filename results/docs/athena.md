# AWS Athena
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage information system accounts, including establishing, activating, modifying, reviewing, disabling, and removing accounts | **Platform/IaC** - Implement IAM policies and roles with least privilege access to Athena resources, regularly review and audit user accounts |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure IAM policies to control access to specific databases, tables, and S3 locations used by Athena |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **Platform** - Enable TLS encryption for all Athena API calls and query results in transit |  |
| **High** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect the confidentiality and integrity of information at rest | **IaC** - Enable S3 encryption for Athena query results and source data using KMS keys |  |
| Medium | AU-2 | NIST 800-53 Rev 5 | Audit Events | Identify the types of events that the system is capable of logging in support of the audit function | **IaC** - Enable CloudTrail logging for all Athena API calls and configure VPC Flow Logs if using VPC endpoints |  |
| Medium | AU-12 | NIST 800-53 Rev 5 | Audit Record Generation | Provide audit record generation capability for defined auditable events | **IaC** - Configure CloudWatch logging for Athena query execution and performance metrics |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor the system to detect attacks and indicators of potential attacks | **IaC** - Set up CloudWatch alarms for unusual query patterns, failed queries, and resource consumption |  |
| **High** | IAM-01 | CSA CCM v4.0 | Identity and Access Management Policy and Procedures | Establish and maintain identity and access management policies and procedures | **User/IaC** - Document and implement IAM policies for Athena access with role-based permissions and regular reviews |  |
| **High** | IAM-02 | CSA CCM v4.0 | User Access Provisioning | User access is provisioned, de-provisioned, and reviewed | **User** - Implement automated provisioning and de-provisioning of Athena access through IAM roles and groups |  |
| **High** | EKM-02 | CSA CCM v4.0 | Encryption Key Management | Encryption keys are managed throughout their lifecycle | **IaC** - Use AWS KMS for managing encryption keys for Athena query results and implement key rotation policies |  |
| Medium | DSI-01 | CSA CCM v4.0 | Data Security and Information Lifecycle Management | Data and objects are protected throughout their lifecycle | **IaC** - Implement data classification and retention policies for Athena query results and underlying S3 data |  |
| Medium | DSI-07 | CSA CCM v4.0 | Secure Disposal | Secure disposal of data when no longer required | **IaC** - Configure S3 lifecycle policies for Athena result buckets to automatically delete old query results |  |
| **High** | IAM.1 | AWS Foundational Security Best Practices 1.0.0 | IAM policies should not allow full administrative privileges | This control checks whether IAM policies grant full administrative privileges | **IaC** - Apply principle of least privilege for Athena IAM policies, avoiding wildcard permissions |  |
| **High** | S3.1 | AWS Foundational Security Best Practices 1.0.0 | S3 Block Public Access setting should be enabled | This control checks whether S3 Block Public Access settings are configured for S3 buckets used by Athena | **IaC** - Enable S3 Block Public Access settings on buckets containing Athena data sources and results |  |
| Medium | Athena.1 | AWS Foundational Security Best Practices 1.0.0 | Athena workgroups should encrypt query results | This control checks whether an Amazon Athena workgroup encrypts query results | **IaC** - Configure workgroup settings to enforce encryption of query results using S3 server-side encryption |  |
| **High** | CloudTrail.1 | AWS Security Hub Current | CloudTrail should be enabled and configured with at least one multi-Region trail | This control checks that there is at least one multi-region CloudTrail trail | **Platform** - Enable multi-region CloudTrail to capture all Athena API calls across regions |  |
| Medium | CloudWatch.1 | AWS Security Hub Current | A log metric filter and alarm should exist for usage of root user | This control checks whether a log metric filter and alarm exist for root user usage | **IaC** - Create CloudWatch alarms for root user access to Athena resources |  |
| Medium | EC2.2 | AWS Security Hub Current | VPC default security groups should not allow inbound or outbound traffic | This control checks that the default security group restricts all traffic | **IaC** - Configure restrictive security groups for VPC endpoints used with Athena |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Athena Cost Optimization Best Practices Current | Optimize Data Storage Format | Use columnar storage formats like Parquet or ORC to reduce data scanned and improve query performance | **User** - Convert data to Parquet or ORC format and partition data appropriately to minimize data scanned per query |  |
| **High** | COST-02 | AWS Athena Cost Optimization Best Practices Current | Implement Data Partitioning | Partition data by commonly queried fields to reduce the amount of data scanned | **User** - Design partition schemes based on query patterns and use partition projection where applicable |  |
| **High** | COST-03 | AWS Athena Cost Optimization Best Practices Current | Use Data Compression | Compress data to reduce storage costs and improve query performance | **User** - Apply appropriate compression algorithms (GZIP, Snappy, LZO) based on data type and access patterns |  |
| Medium | COST-04 | AWS Athena Cost Optimization Best Practices Current | Configure Query Result Location Lifecycle | Set up S3 lifecycle policies for query result locations to automatically delete old results | **IaC** - Configure S3 lifecycle rules to transition old query results to cheaper storage classes or delete them |  |
| Medium | COST-05 | AWS Athena Cost Optimization Best Practices Current | Implement Workgroup Query Controls | Use workgroups to set query limits and control resource usage | **IaC** - Configure workgroups with data usage controls, query timeout limits, and result encryption requirements |  |
| Medium | COST-07 | AWS Athena Cost Optimization Best Practices Current | Monitor and Alert on Usage | Set up monitoring and alerting for Athena usage and costs | **IaC** - Create CloudWatch alarms for data processed, query costs, and set up AWS Budgets for cost control |  |
| Medium | COST-06 | AWS Athena Cost Optimization Best Practices Current | Optimize Query Performance | Write efficient queries to minimize data processing and reduce costs | **User** - Use LIMIT clauses, specific column selection, predicate pushdown, and avoid SELECT * statements |  |
| Low | COST-08 | AWS Athena Cost Optimization Best Practices Current | Use Approximation Functions | Use approximation functions for large datasets when exact results are not required | **User** - Utilize functions like approx_distinct() and approx_percentile() for faster, cost-effective analytics |  |

