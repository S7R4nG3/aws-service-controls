# AWS Batch
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management | IAM-01 | Implement strong identity and access management controls for AWS Batch resources | **IaC** - Define IAM roles and policies in CloudFormation/Terraform templates with least privilege access to Batch resources |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Ensure proper data classification and protection for data processed in Batch jobs | **User** - Implement data encryption at rest and in transit, classify data sensitivity levels, and apply appropriate security controls |  |
| **High** | Encryption and Key Management | EKM-01 | Implement proper encryption key management for Batch resources | **Platform** - Use AWS KMS for encryption key management and enable encryption for EBS volumes and S3 buckets used by Batch |  |
| Medium | Infrastructure and Virtualization Security | IVS-01 | Secure the underlying infrastructure and compute environments for Batch | **IaC** - Configure secure compute environments with proper security groups, subnets, and instance types in infrastructure templates |  |
| Medium | Logging and Monitoring | LOG-01 | Implement comprehensive logging and monitoring for Batch activities | **Platform** - Enable CloudTrail logging, CloudWatch monitoring, and configure log aggregation for Batch job execution |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Account Management | AC-2 | Manage AWS Batch service accounts and access permissions | **IaC** - Implement automated account provisioning and deprovisioning through IAM roles and policies defined in infrastructure code |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect data transmission to and from Batch compute environments | **Platform** - Enable TLS/SSL encryption for all API communications and data transfers to/from Batch resources |  |
| **High** | Protection of Information at Rest | SC-28 | Encrypt data at rest in Batch storage systems | **IaC** - Configure EBS volume encryption and S3 bucket encryption in infrastructure templates for Batch storage |  |
| Medium | Event Logging | AU-2 | Log security-relevant events in AWS Batch | **Platform** - Enable CloudTrail for API logging and CloudWatch for operational logging of Batch activities |  |
| Medium | System Monitoring | SI-4 | Monitor AWS Batch systems for security events and anomalies | **Platform** - Implement CloudWatch monitoring, AWS Security Hub, and GuardDuty for continuous security monitoring |  |
| Low | Baseline Configuration | CM-2 | Maintain baseline configurations for Batch compute environments | **IaC** - Define and maintain standardized compute environment configurations using infrastructure as code templates |  |

### AWS Foundational Security Standard v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Batch compute environments should use EC2 instances with IMDSv2 | Batch.1 | Configure Batch compute environments to use EC2 instances with Instance Metadata Service version 2 | **IaC** - Set httpTokens to 'required' and httpPutResponseHopLimit appropriately in launch templates used by Batch compute environments |  |
| **High** | Batch job definitions should not have privileged mode enabled | Batch.2 | Ensure Batch job definitions do not run containers in privileged mode | **User** - Set privileged parameter to false in containerProperties of job definitions |  |
| Medium | Batch compute environments should be in a VPC | Batch.3 | Deploy Batch compute environments within a VPC for network isolation | **IaC** - Configure subnets parameter in compute environment definitions to deploy instances in private subnets |  |
| Medium | Batch job queues should have logging enabled | Batch.4 | Enable logging for Batch job queues to track job execution | **Platform** - Configure CloudWatch Logs integration and enable container logging in job definitions |  |
| Low | Batch compute environments should use managed scaling | Batch.5 | Use AWS managed scaling for Batch compute environments | **IaC** - Set type to 'MANAGED' in compute environment configuration for automatic scaling management |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Batch job definitions should not have elevated privileges | Batch.1 | Job definitions should not grant unnecessary elevated privileges to containers | **User** - Configure job definitions with minimal required privileges and avoid privileged mode, readonlyRootFilesystem should be true |  |
| Medium | Batch compute environments should have detailed monitoring enabled | Batch.2 | Enable detailed monitoring for Batch compute environments | **IaC** - Set instanceRole and enable detailed monitoring in compute environment launch templates |  |
| Medium | Batch job definitions should specify resource limits | Batch.3 | Define appropriate CPU and memory limits for Batch jobs | **User** - Specify vcpus and memory parameters in job definition containerProperties to prevent resource exhaustion |  |
| Low | Batch compute environments should use encryption for EBS volumes | Batch.4 | Enable EBS encryption for compute environment storage | **IaC** - Configure encrypted EBS volumes in launch templates used by Batch compute environments |  |


## Operational Controls
---



## Cost Controls
---


### AWS Batch Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Use Spot Instances for Fault-Tolerant Workloads | COST-01 | Leverage EC2 Spot Instances in Batch compute environments for significant cost savings on fault-tolerant workloads | **IaC** - Configure bidPercentage and allocationStrategy in compute environment to use Spot instances with diversified instance types |
| **High** | Implement Auto Scaling and Right-Sizing | COST-02 | Use managed compute environments with auto scaling to automatically adjust capacity based on job queue demand | **IaC** - Set appropriate minvCpus, maxvCpus, and desiredvCpus in managed compute environments for optimal scaling |
| Medium | Optimize Instance Types and Families | COST-03 | Select appropriate EC2 instance types based on workload characteristics to optimize cost-performance ratio | **User** - Analyze job resource requirements and specify optimal instance types in compute environment instanceTypes parameter |
| Medium | Implement Job Queue Prioritization | COST-04 | Use job queue priorities to ensure critical jobs run first and optimize resource utilization | **User** - Configure priority values for job queues and implement job scheduling strategies to maximize throughput |
| Medium | Enable Cost Allocation Tags | COST-05 | Implement comprehensive tagging strategy for cost tracking and allocation across different teams or projects | **IaC** - Define consistent tagging strategy in infrastructure templates for all Batch resources including compute environments and job queues |
| Low | Optimize Storage Costs | COST-06 | Use appropriate storage types and lifecycle policies for temporary and persistent data | **Platform** - Configure S3 storage classes, lifecycle policies, and use appropriate EBS volume types for different data access patterns |
| Low | Implement Resource Monitoring and Alerting | COST-07 | Set up monitoring and alerting for cost anomalies and resource utilization | **Platform** - Configure CloudWatch billing alarms, AWS Cost Explorer, and AWS Budgets to monitor and alert on Batch costs |


