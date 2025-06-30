# AWS Application Auto Scaling
---


### Cloud Security Alliance (CSA) Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | User Access Management | IAM-02 | Ensure proper user access management and authentication for Application Auto Scaling resources | **Platform** - Configure IAM policies with least privilege principles, implement MFA for administrative access, and use IAM roles for service-to-service authentication |  |
| **Critical** | Privileged User Management | IAM-03 | Control and monitor privileged access to Auto Scaling configurations and policies | **IaC** - Define IAM policies in IaC templates that restrict Auto Scaling administrative actions to specific roles and implement approval workflows for policy changes |  |
| **High** | Anti-Malware | AIS-01 | Ensure scaled instances are protected against malware | **User** - Configure AMIs with anti-malware solutions and ensure Auto Scaling launches instances from secure, hardened base images |  |
| **High** | Data Loss Prevention | DSP-03 | Implement data protection measures for instances managed by Auto Scaling | **IaC** - Configure Auto Scaling to use encrypted EBS volumes and implement data classification policies in launch templates |  |
| **High** | Network Security | IVS-06 | Ensure network security controls are applied to Auto Scaling groups | **IaC** - Configure security groups and NACLs in Auto Scaling launch configurations to restrict network access and implement VPC isolation |  |
| Medium | Audit Logging / Intrusion Detection | STA-09 | Enable comprehensive logging and monitoring for Auto Scaling activities | **Platform** - Enable CloudTrail logging for Auto Scaling API calls and configure CloudWatch alarms for scaling events and policy violations |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Account Management | AC-2 | Manage accounts with access to Application Auto Scaling resources and configurations | **Platform** - Implement automated account provisioning and deprovisioning processes, regularly review account permissions, and maintain account inventory for Auto Scaling access |  |
| **Critical** | Access Enforcement | AC-3 | Enforce authorized access to Auto Scaling resources and prevent unauthorized modifications | **IaC** - Implement resource-based policies and IAM conditions to enforce access controls based on time, location, and request context for Auto Scaling operations |  |
| **High** | Boundary Protection | SC-7 | Implement network boundary protection for Auto Scaling groups and instances | **IaC** - Configure VPC boundaries, security groups, and NACLs in Auto Scaling launch templates to control network traffic flow |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect data transmission for Auto Scaling communications and scaled applications | **User** - Configure TLS/SSL encryption for application communications and ensure Auto Scaling health checks use encrypted protocols where possible |  |
| **High** | Protection of Information at Rest | SC-28 | Ensure data encryption at rest for instances managed by Auto Scaling | **IaC** - Configure encrypted EBS volumes and encrypted snapshots in Auto Scaling launch templates using AWS KMS keys |  |
| Medium | Event Logging | AU-2 | Log Auto Scaling events and configuration changes for audit purposes | **Platform** - Enable CloudTrail for Auto Scaling API logging and configure CloudWatch Events to capture scaling activities and policy changes |  |
| Medium | System Monitoring | SI-4 | Monitor Auto Scaling system performance and security events | **Platform** - Implement CloudWatch monitoring for Auto Scaling metrics and configure automated alerting for anomalous scaling behavior |  |

### AWS Foundational Security Best Practices 1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Auto Scaling groups associated with a load balancer should use load balancer health checks | AutoScaling.1 | Auto Scaling groups should use load balancer health checks to ensure unhealthy instances are replaced | **IaC** - Configure Auto Scaling groups to use ELB health checks in addition to EC2 health checks to ensure proper instance health monitoring |  |
| **High** | Amazon EC2 Auto Scaling group should cover multiple Availability Zones | AutoScaling.2 | Auto Scaling groups should span multiple AZs for high availability and fault tolerance | **IaC** - Configure Auto Scaling groups to deploy instances across at least two Availability Zones to ensure resilience |  |
| **High** | Auto Scaling group launch configurations should not have public IP addresses | AutoScaling.3 | Launch configurations should not automatically assign public IP addresses unless required | **IaC** - Configure launch templates to disable automatic public IP assignment and use NAT gateways or load balancers for internet access |  |
| Medium | Auto Scaling group launch configuration should not have metadata response hop limit greater than 1 | AutoScaling.4 | Limit metadata service response hop limit to prevent potential security issues | **IaC** - Set metadata response hop limit to 1 in launch templates to restrict metadata service access to the instance itself |  |
| Medium | Auto Scaling groups should use EC2 launch templates | AutoScaling.5 | Use launch templates instead of launch configurations for better feature support and management | **IaC** - Migrate from launch configurations to launch templates to access newer EC2 features and improved versioning capabilities |  |
| Medium | Auto Scaling groups should use multiple instance types in multiple Availability Zones | AutoScaling.6 | Use multiple instance types to improve availability and optimize costs | **IaC** - Configure Auto Scaling groups with multiple instance types and use mixed instances policies to improve fault tolerance and cost optimization |  |

### AWS Security Hub 2023
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Auto Scaling launch configuration public IP | ASG-1 | Auto Scaling launch configurations should not assign public IP addresses | **IaC** - Configure launch configurations and templates to disable automatic public IP assignment for instances |  |
| **High** | Auto Scaling group ELB health check | ASG-2 | Auto Scaling groups with load balancers should use ELB health checks | **IaC** - Enable ELB health checks for Auto Scaling groups that are associated with load balancers |  |
| Medium | Auto Scaling group multi-AZ | ASG-3 | Auto Scaling groups should be configured to use multiple Availability Zones | **IaC** - Configure Auto Scaling groups to span multiple Availability Zones for improved fault tolerance |  |
| Medium | Auto Scaling group launch template | ASG-4 | Auto Scaling groups should use launch templates instead of launch configurations | **IaC** - Migrate Auto Scaling groups from launch configurations to launch templates for enhanced functionality |  |


## Operational Controls
---



## Cost Controls
---


### AWS Application Auto Scaling Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Use Target Tracking Scaling Policies | COST-1 | Implement target tracking scaling policies for more efficient resource utilization | **IaC** - Configure target tracking scaling policies based on appropriate metrics like CPU utilization, request count, or custom metrics to maintain optimal performance while minimizing costs |
| **High** | Implement Mixed Instance Types | COST-2 | Use multiple instance types and purchasing options to optimize costs | **IaC** - Configure Auto Scaling groups with mixed instance types policy to combine On-Demand and Spot instances across different instance families |
| **High** | Optimize Scaling Policies | COST-3 | Fine-tune scaling policies to prevent over-provisioning and under-provisioning | **User** - Regularly review and adjust scaling policy thresholds, cooldown periods, and step scaling configurations based on application performance patterns |
| Medium | Use Predictive Scaling | COST-4 | Implement predictive scaling for workloads with predictable patterns | **User** - Enable predictive scaling for applications with recurring traffic patterns to proactively scale resources and reduce over-provisioning |
| Medium | Set Appropriate Cooldown Periods | COST-5 | Configure proper cooldown periods to prevent rapid scaling oscillations | **IaC** - Set cooldown periods based on application startup time and metric stabilization periods to avoid unnecessary scaling actions |
| Medium | Monitor and Optimize Instance Sizes | COST-6 | Regularly review instance utilization and right-size instances | **User** - Use CloudWatch metrics and AWS Cost Explorer to identify over-provisioned instances and adjust Auto Scaling configurations accordingly |
| Medium | Implement Scheduled Scaling | COST-7 | Use scheduled scaling for predictable workload patterns | **IaC** - Configure scheduled scaling actions for known traffic patterns such as business hours, maintenance windows, or batch processing schedules |
| Low | Use Application-Specific Metrics | COST-8 | Implement custom metrics for more accurate scaling decisions | **User** - Create custom CloudWatch metrics specific to application performance indicators rather than relying solely on infrastructure metrics |
| Low | Regular Cost Analysis | COST-9 | Perform regular cost analysis of Auto Scaling activities | **User** - Use AWS Cost and Usage Reports and CloudWatch Insights to analyze Auto Scaling costs and identify optimization opportunities |


