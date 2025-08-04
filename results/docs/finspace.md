# AWS FinSpace
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA CCM v4.0 | Identity and Access Management Policy and Procedures | Establish comprehensive identity and access management policies for FinSpace environments and users | **Platform** - Configure FinSpace user management with proper role-based access controls and integration with AWS IAM |  |
| **Critical** | DSI-05 | CSA CCM v4.0 | Data Classification | Classify all financial data stored and processed within FinSpace according to sensitivity levels | **User** - Implement data classification tags and metadata within FinSpace datasets and organize data based on sensitivity |  |
| **Critical** | EKM-01 | CSA CCM v4.0 | Encryption Key Management | Implement proper encryption key management for data at rest and in transit within FinSpace | **IaC** - Configure AWS KMS keys for FinSpace encryption and establish key rotation policies through CloudFormation templates |  |
| **High** | IVS-06 | CSA CCM v4.0 | Network Security | Implement network segmentation and security controls for FinSpace environments | **IaC** - Deploy FinSpace within VPC with proper security groups, NACLs, and private subnets configuration |  |
| **High** | BCR-01 | CSA CCM v4.0 | Business Continuity Planning | Establish business continuity and disaster recovery procedures for FinSpace environments | **Platform** - Configure cross-region backup strategies and implement FinSpace environment replication procedures |  |
| Medium | AIS-01 | CSA CCM v4.0 | Audit Independent Reviews | Enable comprehensive logging and monitoring for all FinSpace activities and access patterns | **IaC** - Configure CloudTrail, CloudWatch, and FinSpace audit logging with appropriate retention policies |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage FinSpace user accounts with proper provisioning, modification, and deprovisioning procedures | **Platform** - Implement automated user lifecycle management through AWS SSO integration with FinSpace |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information in FinSpace | **Platform** - Ensure all FinSpace communications use TLS 1.2 or higher encryption protocols |  |
| **Critical** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect the confidentiality and integrity of information stored in FinSpace datasets | **IaC** - Configure FinSpace with customer-managed KMS keys for encryption at rest of all datasets and clusters |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Determine and audit events within FinSpace environments | **IaC** - Configure comprehensive audit logging for FinSpace API calls, user activities, and data access events |  |
| **High** | CP-9 | NIST 800-53 Rev 5 | System Backup | Conduct backups of FinSpace datasets and configurations | **Platform** - Implement automated backup procedures for FinSpace datasets using AWS Backup or custom solutions |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor FinSpace for security events and unauthorized activities | **IaC** - Deploy CloudWatch monitoring and AWS Security Hub integration for FinSpace security monitoring |  |
| **Critical** | FinSpace.1 | AWS Foundational Security Best Practices 1.0.0 | FinSpace environments should not be publicly accessible | Ensure FinSpace environments are deployed in private networks and not accessible from the public internet | **IaC** - Deploy FinSpace environments within VPC private subnets with no direct internet gateway access |  |
| **Critical** | FinSpace.2 | AWS Foundational Security Best Practices 1.0.0 | FinSpace clusters should be encrypted | Enable encryption for FinSpace clusters using AWS KMS keys | **IaC** - Configure FinSpace cluster creation with encryption enabled using customer-managed KMS keys |  |
| **High** | FinSpace.3 | AWS Foundational Security Best Practices 1.0.0 | FinSpace should have deletion protection enabled | Enable deletion protection for FinSpace environments to prevent accidental deletion | **Platform** - Configure FinSpace environment settings with deletion protection enabled and implement proper IAM policies |  |
| **High** | FinSpace.4 | AWS Foundational Security Best Practices 1.0.0 | FinSpace access logging should be enabled | Enable and configure access logging for FinSpace environments | **IaC** - Configure FinSpace audit logging and integrate with CloudWatch Logs for centralized log management |  |
| Medium | FinSpace.5 | AWS Foundational Security Best Practices 1.0.0 | FinSpace datasets should have appropriate permissions | Implement least privilege access controls for FinSpace datasets | **User** - Configure dataset permissions with role-based access controls and regular access reviews |  |
| **Critical** | finspace-environment-kms-key-configured | AWS Config 2023.1 | FinSpace environments should use customer managed keys | Use customer-managed KMS keys instead of AWS managed keys for enhanced control | **IaC** - Create and configure customer-managed KMS keys for FinSpace environment encryption |  |
| **High** | finspace-environment-audit-logging-enabled | AWS Config 2023.1 | FinSpace environments should have audit logging enabled | Ensure FinSpace environments have audit logging enabled for compliance and monitoring | **Platform** - Configure FinSpace audit logging and ensure CloudTrail captures all API calls |  |
| **High** | finspace-environment-in-vpc | AWS Config 2023.1 | FinSpace environments should be deployed in VPC | Deploy FinSpace environments within VPC for network isolation and security | **IaC** - Configure FinSpace environments to be deployed within VPC with appropriate subnet configuration |  |
| Medium | finspace-cluster-encryption-enabled | AWS Config 2023.1 | FinSpace clusters should have encryption enabled | Ensure FinSpace clusters are configured with encryption for data protection | **IaC** - Configure FinSpace clusters with encryption enabled using KMS keys |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS FinSpace Cost Optimization 2023 | Right-size FinSpace clusters | Monitor and optimize FinSpace cluster sizes based on actual usage patterns | **User** - Regularly review cluster utilization metrics and adjust cluster configurations to match workload requirements |  |
| **High** | COST-02 | AWS FinSpace Cost Optimization 2023 | Implement cluster scheduling | Schedule FinSpace clusters to run only during business hours when possible | **IaC** - Use AWS Lambda and EventBridge to automatically start and stop FinSpace clusters based on schedules |  |
| Medium | COST-03 | AWS FinSpace Cost Optimization 2023 | Optimize data storage | Implement data lifecycle policies to optimize storage costs for FinSpace datasets | **Platform** - Configure data retention policies and archive older datasets to lower-cost storage tiers |  |
| Medium | COST-04 | AWS FinSpace Cost Optimization 2023 | Monitor usage and costs | Implement cost monitoring and alerting for FinSpace usage | **IaC** - Configure AWS Cost Explorer, budgets, and billing alerts to track FinSpace spending patterns |  |
| Medium | COST-05 | AWS FinSpace Cost Optimization 2023 | Use appropriate instance types | Select cost-effective instance types for FinSpace clusters based on workload characteristics | **User** - Analyze workload requirements and select appropriate instance families and sizes for optimal cost-performance ratio |  |
| Low | COST-06 | AWS FinSpace Cost Optimization 2023 | Implement resource tagging | Tag all FinSpace resources for cost allocation and chargeback purposes | **IaC** - Apply consistent tagging strategies across all FinSpace resources for cost tracking and allocation |  |

