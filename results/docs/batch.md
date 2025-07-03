# AWS Batch
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | Cloud Security Alliance (CSA) Cloud Controls Matrix 4.0.10 | Entitlement | Implement strong identity and access management controls for AWS Batch resources including compute environments, job queues, and job definitions | **IaC** - Define IAM roles, policies, and service-linked roles in CloudFormation/Terraform templates with least privilege access principles |  |
| **Critical** | DSI-07 | Cloud Security Alliance (CSA) Cloud Controls Matrix 4.0.10 | Secure Disposal or Re-use of Equipment | Ensure secure disposal and data sanitization of compute resources used by AWS Batch when decommissioned | **Platform** - Leverage AWS managed disposal processes and implement data encryption to ensure secure decommissioning of Batch compute resources |  |
| **High** | BCR-01 | Cloud Security Alliance (CSA) Cloud Controls Matrix 4.0.10 | Business Continuity Planning | Ensure Batch workloads can recover from failures and maintain operational continuity | **IaC** - Deploy Batch resources across multiple AZs and implement automated retry mechanisms in job definitions |  |
| **High** | AIS-01 | Cloud Security Alliance (CSA) Cloud Controls Matrix 4.0.10 | Audit Logging / Intrusion Detection | Enable comprehensive logging and monitoring for all Batch operations and API calls | **Platform** - Configure CloudTrail, CloudWatch Logs, and AWS Config to monitor Batch API calls and resource changes |  |
| Medium | TVM-01 | Cloud Security Alliance (CSA) Cloud Controls Matrix 4.0.10 | Vulnerability Management | Regularly scan and update container images and underlying infrastructure for security vulnerabilities | **User** - Implement container image scanning in CI/CD pipelines and regularly update AMIs used in compute environments |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage AWS accounts and user access for Batch services with proper provisioning and deprovisioning procedures | **Platform** - Implement AWS Organizations, SSO, and automated account lifecycle management for Batch access |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect data transmission between Batch components and external systems using encryption | **IaC** - Configure VPC endpoints, security groups, and TLS encryption for all Batch communications |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to Batch resources and job execution | **IaC** - Implement resource-based policies and IAM conditions to restrict Batch resource access |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Generate audit logs for Batch job submissions, state changes, and administrative actions | **Platform** - Enable CloudWatch Events for Batch state changes and CloudTrail for API logging |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor Batch infrastructure and workloads for security events and anomalous behavior | **Platform** - Deploy GuardDuty, Security Hub, and custom CloudWatch alarms for Batch monitoring |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Implement backup procedures for Batch job definitions, configurations, and critical data | **IaC** - Use AWS Backup or custom solutions to backup Batch configurations and associated data |  |
| **Critical** | Batch.1 | AWS Foundational Security Standard 1.0 | AWS Batch compute environments should use EC2 launch templates | Ensure Batch compute environments use EC2 launch templates for consistent and secure instance configuration | **IaC** - Define EC2 launch templates with security groups, IAM instance profiles, and encrypted EBS volumes in infrastructure code |  |
| **Critical** | Batch.2 | AWS Foundational Security Standard 1.0 | AWS Batch job definitions should not have privileged or root access | Prevent Batch job definitions from running with privileged access or as root user | **IaC** - Configure job definitions with non-privileged user contexts and avoid privileged flags in container parameters |  |
| **High** | Batch.3 | AWS Foundational Security Standard 1.0 | AWS Batch compute environments should be placed in private subnets | Deploy Batch compute environments in private subnets to reduce attack surface | **IaC** - Configure compute environments to use private subnet IDs and ensure NAT Gateway for outbound connectivity |  |
| **High** | Batch.4 | AWS Foundational Security Standard 1.0 | AWS Batch should use encrypted EBS volumes | Enable EBS encryption for all volumes used by Batch compute environments | **IaC** - Configure launch templates with encrypted EBS volumes using AWS KMS keys |  |
| Medium | Batch.5 | AWS Foundational Security Standard 1.0 | AWS Batch job queues should have logging enabled | Enable logging for Batch job queues to track job execution and performance | **Platform** - Configure CloudWatch Logs integration and enable job queue event logging |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS Batch Cost Optimization 2023 | Use Spot Instances for Fault-Tolerant Workloads | Leverage EC2 Spot Instances in Batch compute environments to reduce costs for fault-tolerant workloads | **IaC** - Configure compute environments with SPOT allocation strategy and appropriate bid prices in infrastructure templates |  |
| **High** | COST-02 | AWS Batch Cost Optimization 2023 | Implement Auto Scaling for Compute Environments | Configure appropriate scaling policies to automatically adjust compute capacity based on job queue demand | **IaC** - Set minvCpus to 0 and configure appropriate maxvCpus and desiredvCpus in compute environment definitions |  |
| **High** | COST-03 | AWS Batch Cost Optimization 2023 | Right-Size Instance Types | Select appropriate EC2 instance types based on workload requirements to optimize cost-performance ratio | **IaC** - Define instance type arrays with cost-optimized instances and use allocation strategy DIVERSIFIED |  |
| Medium | COST-04 | AWS Batch Cost Optimization 2023 | Optimize Job Resource Requirements | Define accurate resource requirements in job definitions to prevent over-provisioning | **User** - Monitor job resource utilization and adjust vCpus, memory, and GPU requirements in job definitions accordingly |  |
| Medium | COST-05 | AWS Batch Cost Optimization 2023 | Use Multi-AZ Deployment Strategically | Deploy across multiple AZs to take advantage of Spot Instance availability and pricing variations | **IaC** - Configure compute environments with subnets from multiple availability zones to optimize Spot pricing |  |
| Medium | COST-06 | AWS Batch Cost Optimization 2023 | Implement Job Queue Prioritization | Use job queue priorities to ensure critical jobs run first and optimize resource allocation | **IaC** - Configure job queues with appropriate priority levels and associate with cost-optimized compute environments |  |
| Medium | COST-07 | AWS Batch Cost Optimization 2023 | Monitor and Alert on Cost Metrics | Implement cost monitoring and alerting to track Batch spending and identify optimization opportunities | **Platform** - Set up AWS Cost Explorer, budgets, and CloudWatch alarms to monitor Batch-related EC2 and EBS costs |  |
| Low | COST-08 | AWS Batch Cost Optimization 2023 | Use Efficient Container Images | Optimize container images to reduce startup time and storage costs | **User** - Use minimal base images, implement multi-stage builds, and optimize layer caching in container image creation |  |
| Low | COST-09 | AWS Batch Cost Optimization 2023 | Implement Lifecycle Policies for Logs and Data | Configure lifecycle policies for CloudWatch Logs and associated S3 storage to manage long-term costs | **IaC** - Set log retention periods and S3 lifecycle transitions to cheaper storage classes in infrastructure templates |  |

