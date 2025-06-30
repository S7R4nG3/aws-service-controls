# AWS VPC Security Group
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy | IAM-01 | Establish comprehensive identity and access management policies for security group management | **Platform** - Configure IAM policies with principle of least privilege for security group operations, implement resource-based policies and conditions |  |
| **High** | User Access Provisioning | IAM-02 | Implement proper user access provisioning and deprovisioning procedures | **Platform** - Use IAM roles and groups for security group access, implement automated provisioning workflows |  |
| **Critical** | Network Architecture | IVS-01 | Design secure network architecture with proper segmentation | **IaC** - Define security groups in IaC templates with proper rule definitions and network segmentation patterns |  |
| **High** | Network Controls | IVS-02 | Implement network security controls including firewalls and access controls | **IaC** - Configure security group rules to deny by default, allow only necessary traffic, implement layered security |  |
| **High** | Logging and Monitoring | LOG-01 | Implement comprehensive logging and monitoring for security events | **Platform** - Enable CloudTrail for security group changes, implement VPC Flow Logs, configure monitoring dashboards |  |
| Medium | Business Continuity Management | BCR-01 | Ensure business continuity through proper backup and recovery procedures | **IaC** - Implement security group configurations in version-controlled IaC templates for disaster recovery |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **Platform** - Configure IAM policies to enforce least privilege access to security group management functions |  |
| **Critical** | Information Flow Enforcement | AC-4 | Control information flows within the system and between interconnected systems | **IaC** - Design security group rules to enforce information flow policies between network segments |  |
| **High** | Event Logging | AU-2 | Identify the types of events that the system is capable of logging | **Platform** - Enable CloudTrail logging for all security group API calls and configuration changes |  |
| **High** | Audit Record Review, Analysis, and Reporting | AU-6 | Review and analyze information system audit records for indications of inappropriate or unusual activity | **User** - Implement regular review processes for security group changes and access patterns using CloudWatch and Security Hub |  |
| **High** | Baseline Configuration | CM-2 | Develop, document, and maintain a current baseline configuration | **IaC** - Maintain baseline security group configurations in version-controlled IaC templates |  |
| Medium | Configuration Settings | CM-6 | Establish and document configuration settings for system components | **IaC** - Document and standardize security group configuration settings in organizational templates |  |
| Medium | System Backup | CP-9 | Conduct backups of user-level information contained in the system | **Platform** - Implement automated backup of security group configurations using AWS Config and IaC state management |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Amazon EBS snapshots should not be public | EC2.1 | Ensure EBS snapshots are not publicly accessible | **Platform** - Configure security groups to restrict access to EBS resources and implement proper snapshot policies |  |
| **Critical** | VPC default security group should not allow inbound and outbound traffic | EC2.2 | Default security groups should restrict all traffic | **IaC** - Remove all inbound and outbound rules from default security groups, create custom security groups for specific needs |  |
| **Critical** | Security groups should not allow unrestricted access to ports with high risk | EC2.19 | Restrict access to high-risk ports like SSH, RDP, and database ports | **IaC** - Configure security group rules to allow access only from specific IP ranges or security groups for high-risk ports |  |
| **High** | Both VPC default security groups should not allow inbound and outbound traffic | EC2.20 | Ensure both default security groups in VPC restrict all traffic | **IaC** - Systematically review and remove all rules from default security groups across all VPCs |  |
| **High** | Network ACLs should not allow ingress from 0.0.0.0/0 to port 22 or 3389 | EC2.21 | Restrict SSH and RDP access in network ACLs | **IaC** - Configure security groups to complement NACL restrictions for administrative access ports |  |
| Medium | AWS Config should be enabled | Config.1 | Enable AWS Config to track security group configuration changes | **Platform** - Enable AWS Config in all regions to monitor security group configuration compliance |  |

### AWS Security Hub Current
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Security groups should not allow unrestricted inbound traffic | SecurityHub.EC2.1 | Security groups should not have inbound rules that allow unrestricted traffic | **IaC** - Review and restrict security group rules to allow only necessary traffic from specific sources |  |
| **High** | Security groups should not allow unrestricted outbound traffic | SecurityHub.EC2.2 | Security groups should have restrictive outbound rules | **IaC** - Configure explicit outbound rules instead of allowing all outbound traffic |  |
| **High** | CloudTrail should be enabled | SecurityHub.CloudTrail.1 | Enable CloudTrail to log security group API calls | **Platform** - Enable CloudTrail across all regions to capture security group management activities |  |
| Medium | AWS Config should be enabled | SecurityHub.Config.1 | Enable AWS Config for compliance monitoring | **Platform** - Configure AWS Config rules to monitor security group compliance with organizational policies |  |
| **High** | IAM policies should not allow full administrative privileges | SecurityHub.IAM.1 | Restrict IAM policies for security group management to necessary permissions only | **Platform** - Create specific IAM policies for security group management with least privilege principles |  |


## Operational Controls
---



## Cost Controls
---


### AWS VPC Security Group Cost Optimization Best Practices
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize Security Group Usage | COST-01 | Eliminate unused and redundant security groups to reduce management overhead | **User** - Regularly audit and remove unused security groups, consolidate similar security groups where appropriate |
| Medium | Implement Efficient Rule Management | COST-02 | Optimize security group rules to reduce complexity and management costs | **IaC** - Use security group references instead of IP addresses where possible, implement rule consolidation strategies |
| Medium | Automate Security Group Management | COST-03 | Reduce operational costs through automation of security group lifecycle management | **Platform** - Implement automated security group provisioning and deprovisioning using IaC and AWS Lambda functions |
| Low | Monitor Security Group Utilization | COST-04 | Track security group usage to identify cost optimization opportunities | **Platform** - Use AWS Cost Explorer and custom CloudWatch metrics to monitor security group-related costs and usage patterns |
| Low | Implement Tagging Strategy | COST-05 | Use consistent tagging for cost allocation and management | **IaC** - Apply consistent cost allocation tags to security groups for better cost tracking and chargeback processes |


