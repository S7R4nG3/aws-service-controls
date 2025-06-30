# AWS OpenSearch
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish comprehensive identity and access management policies for OpenSearch cluster access | **IaC** - Define IAM policies and roles with least privilege access, implement fine-grained access control using OpenSearch security plugin |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Implement encryption at rest and in transit for all OpenSearch data | **Platform** - Enable encryption at rest using AWS KMS keys and enforce HTTPS/TLS for all communications |  |
| **High** | Infrastructure and Virtualization Security | IVS-01 | Secure network infrastructure and implement network segmentation | **IaC** - Deploy OpenSearch in VPC with private subnets, configure security groups with minimal required ports |  |
| **High** | Logging and Monitoring | LOG-01 | Enable comprehensive logging and monitoring for security events | **Platform** - Enable CloudWatch logs, CloudTrail, and OpenSearch audit logs with appropriate retention policies |  |
| Medium | Business Continuity Management | BCR-01 | Implement backup and disaster recovery procedures | **IaC** - Configure automated snapshots to S3, implement multi-AZ deployment, define RTO/RPO requirements |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Account Management | AC-2 | Manage user accounts and access to OpenSearch resources | **User** - Implement regular account reviews, disable unused accounts, use IAM roles for service access |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect data in transit to and from OpenSearch | **Platform** - Enforce HTTPS/TLS 1.2+ for all client connections and inter-node communication |  |
| **Critical** | Protection of Information at Rest | SC-28 | Encrypt sensitive data stored in OpenSearch | **IaC** - Enable encryption at rest using customer-managed KMS keys with proper key rotation |  |
| **High** | Content of Audit Records | AU-3 | Generate comprehensive audit logs for security monitoring | **Platform** - Configure OpenSearch audit logging to capture authentication, authorization, and data access events |  |
| **High** | Information System Backup | CP-9 | Implement regular backup procedures for OpenSearch data | **IaC** - Configure automated daily snapshots with cross-region replication for disaster recovery |  |
| Medium | Information System Monitoring | SI-4 | Monitor OpenSearch for security events and anomalies | **Platform** - Implement CloudWatch monitoring with custom metrics and alerting for suspicious activities |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | OpenSearch domains should have encryption at rest enabled | OpenSearch.1 | Ensures OpenSearch domains encrypt data at rest to protect sensitive information | **IaC** - Set encryptionAtRestOptions.enabled to true in domain configuration with KMS key specification |  |
| **Critical** | OpenSearch domains should be in a VPC | OpenSearch.2 | Deploy OpenSearch domains within VPC for network isolation and security | **IaC** - Configure VPC options with private subnets and appropriate security groups in domain configuration |  |
| **Critical** | OpenSearch domains should encrypt data sent between nodes | OpenSearch.3 | Enable node-to-node encryption for inter-cluster communication security | **IaC** - Set nodeToNodeEncryptionOptions.enabled to true in domain configuration |  |
| **High** | OpenSearch domain error logging to CloudWatch Logs should be enabled | OpenSearch.4 | Enable error logging for troubleshooting and security monitoring | **IaC** - Configure logPublishingOptions for ERROR_LOGS with appropriate CloudWatch log group |  |
| **High** | OpenSearch domains should have audit logging enabled | OpenSearch.5 | Enable audit logging to track access and changes to OpenSearch resources | **IaC** - Configure logPublishingOptions for AUDIT_LOGS with CloudWatch log group and appropriate retention |  |
| Medium | OpenSearch domains should have at least three data nodes | OpenSearch.6 | Ensure high availability and fault tolerance with minimum three data nodes | **IaC** - Set clusterConfig.instanceCount to minimum of 3 with dedicatedMasterEnabled for production workloads |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | OpenSearch domains should have encryption at rest enabled | OpenSearch.1 | Detects OpenSearch domains without encryption at rest enabled | **Platform** - Enable automatic remediation through Security Hub or Config rules to enforce encryption settings |  |
| **Critical** | OpenSearch domains should be in a VPC | OpenSearch.2 | Identifies OpenSearch domains not deployed within a VPC | **IaC** - Ensure all OpenSearch domains are created with VPC configuration in CloudFormation/Terraform templates |  |
| **High** | OpenSearch domains should have fine-grained access control enabled | OpenSearch.7 | Ensures fine-grained access control is enabled for enhanced security | **IaC** - Enable advancedSecurityOptions with fine-grained access control and internal user database or SAML |  |
| **High** | OpenSearch domain connections should be encrypted using TLS 1.2 | OpenSearch.8 | Ensures all connections use minimum TLS 1.2 for secure communication | **IaC** - Configure domainEndpointOptions with TLS security policy set to Policy-Min-TLS-1-2-2019-07 or later |  |
| Medium | OpenSearch domains should have automatic snapshots enabled | OpenSearch.9 | Verifies automatic snapshot configuration for data backup and recovery | **IaC** - Configure snapshotOptions with automated snapshot hour setting in domain configuration |  |


## Operational Controls
---



## Cost Controls
---


### AWS OpenSearch Service Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Right-size OpenSearch instances | COST-01 | Monitor and optimize instance types and sizes based on actual usage patterns | **User** - Use CloudWatch metrics to analyze CPU, memory, and storage utilization; downsize or switch to more cost-effective instance types |
| **High** | Implement data lifecycle management | COST-02 | Use Index State Management policies to transition older data to cheaper storage tiers | **IaC** - Configure ISM policies to move data from hot to warm to cold storage, and delete old indices automatically |
| **High** | Optimize storage configuration | COST-03 | Select appropriate EBS volume types and sizes based on performance requirements | **IaC** - Use gp3 volumes instead of gp2 for better price-performance ratio, implement storage monitoring and alerting |
| Medium | Use Reserved Instances for predictable workloads | COST-04 | Purchase Reserved Instances for stable, long-running OpenSearch clusters | **User** - Analyze usage patterns and purchase 1-year or 3-year Reserved Instances for consistent workloads |
| Medium | Implement cluster scaling policies | COST-05 | Configure automatic scaling to match capacity with demand | **IaC** - Use Application Auto Scaling to automatically adjust cluster size based on metrics like CPU utilization or search latency |
| Medium | Optimize index and shard configuration | COST-06 | Configure optimal number of shards and replicas to minimize resource waste | **User** - Monitor shard sizes and distribution, use index templates to set appropriate shard counts for new indices |
| Medium | Monitor and alert on cost anomalies | COST-07 | Set up cost monitoring and alerting to detect unexpected spending increases | **Platform** - Configure AWS Cost Anomaly Detection and Budgets with notifications for OpenSearch service spending |
| Low | Use appropriate data compression | COST-08 | Enable compression to reduce storage costs and improve performance | **User** - Configure index compression settings and use appropriate field mappings to minimize storage footprint |


