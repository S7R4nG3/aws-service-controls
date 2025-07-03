# AWS Elastic Load Balancer
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | IAM-02 | CSA CCM v4.0 | User Access Management | Implement proper access controls for ELB management operations | **Platform** - Configure IAM policies with least privilege access for ELB operations, use resource-based policies for cross-account access |  |
| **High** | IAM-03 | CSA CCM v4.0 | Privileged User Management | Restrict administrative access to load balancer configuration | **Platform** - Create dedicated IAM roles for ELB administration with MFA requirements and time-based access controls |  |
| **High** | IVS-01 | CSA CCM v4.0 | Network Security | Implement network segmentation and security group controls | **IaC** - Configure security groups to restrict inbound/outbound traffic, implement VPC flow logs for network monitoring |  |
| **High** | DSI-04 | CSA CCM v4.0 | Data Protection in Transit | Encrypt data in transit using TLS/SSL | **IaC** - Configure SSL/TLS listeners with strong cipher suites, implement SSL termination at load balancer level |  |
| Medium | LOG-02 | CSA CCM v4.0 | Audit Logging | Enable comprehensive logging for load balancer activities | **IaC** - Enable access logging to S3, configure CloudWatch metrics and alarms for monitoring |  |
| **High** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage accounts with access to ELB resources | **Platform** - Implement IAM account lifecycle management with regular access reviews and automated provisioning/deprovisioning |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for ELB operations | **Platform** - Use IAM policies and resource tags to enforce access controls based on business requirements |  |
| **High** | SC-7 | NIST 800-53 Rev 5 | Boundary Protection | Monitor and control communications at external boundaries | **IaC** - Configure security groups and NACLs to control traffic flow, implement WAF integration where applicable |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protection of transmitted information | **IaC** - Configure HTTPS listeners with TLS 1.2+ and strong cipher suites, implement SSL redirect policies |  |
| Medium | AU-2 | NIST 800-53 Rev 5 | Event Logging | Determine auditable events for load balancer | **IaC** - Enable CloudTrail for API logging, configure access logs with appropriate attributes |  |
| Medium | CP-9 | NIST 800-53 Rev 5 | System Backup | Backup ELB configuration and associated data | **IaC** - Implement Infrastructure as Code for configuration backup and version control |  |
| Medium | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor ELB for attacks and indicators of potential attacks | **IaC** - Configure CloudWatch alarms for health checks, response times, and error rates |  |
| **High** | ELB.3 | AWS Foundational Security Best Practices 1.0 | Classic Load Balancer listeners should be configured with HTTPS or TLS termination | Ensure encrypted connections | **IaC** - Configure HTTPS/TLS listeners and redirect HTTP traffic to HTTPS |  |
| **High** | ELB.5 | AWS Foundational Security Best Practices 1.0 | Application and Network Load Balancers should span multiple Availability Zones | Ensure high availability and fault tolerance | **IaC** - Configure load balancer with subnets in multiple AZs for redundancy |  |
| Medium | ELB.2 | AWS Foundational Security Best Practices 1.0 | Classic Load Balancers with SSL/HTTPS listeners should use predefined security policies | Ensure secure SSL policies are applied | **IaC** - Configure ELB with AWS predefined SSL security policies that include strong cipher suites |  |
| Medium | ELB.4 | AWS Foundational Security Best Practices 1.0 | Application Load Balancers should be configured to drop HTTP headers | Remove unnecessary HTTP headers for security | **IaC** - Enable deletion of HTTP headers attribute on ALB configuration |  |
| Medium | ELB.6 | AWS Foundational Security Best Practices 1.0 | Application Load Balancer deletion protection should be enabled | Prevent accidental deletion of load balancers | **IaC** - Enable deletion protection attribute in ALB configuration |  |
| Low | ELB.7 | AWS Foundational Security Best Practices 1.0 | Classic Load Balancers should have connection draining enabled | Gracefully handle instance deregistration | **IaC** - Configure connection draining with appropriate timeout values |  |
| **High** | ELB.10 | AWS Security Hub Current | Classic Load Balancers should span multiple Availability Zones | Ensure load balancer resilience across AZs | **IaC** - Configure ELB with instances registered in multiple availability zones |  |
| **High** | ELB.13 | AWS Security Hub Current | Application, Network, and Gateway Load Balancers should span multiple Availability Zones | Ensure multi-AZ deployment for resilience | **IaC** - Deploy load balancer across minimum of two availability zones with appropriate subnet configuration |  |
| Medium | ELB.8 | AWS Security Hub Current | Classic Load Balancers with SSL/HTTPS listeners should use predefined security policies | Use AWS predefined security policies for SSL configuration | **IaC** - Apply ELBSecurityPolicy with latest TLS versions and secure cipher suites |  |
| Medium | ELB.12 | AWS Security Hub Current | Application Load Balancers should be configured with defensive or strictest desync mitigation mode | Protect against HTTP desync attacks | **IaC** - Configure desync mitigation mode to defensive or strictest in ALB attributes |  |
| Medium | ELB.14 | AWS Security Hub Current | Classic Load Balancers should be configured with defensive or strictest desync mitigation mode | Enable desync attack protection | **IaC** - Set desync mitigation mode attribute to defensive or strictest for Classic Load Balancers |  |
| Low | ELB.9 | AWS Security Hub Current | Classic Load Balancers should have cross-zone load balancing enabled | Distribute traffic evenly across availability zones | **IaC** - Enable cross-zone load balancing attribute in ELB configuration |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-ELB-01 | AWS ELB Cost Optimization Best Practices Current | Right-size Load Balancer Type | Choose appropriate load balancer type based on requirements | **User** - Evaluate whether ALB, NLB, CLB, or GWLB best fits your use case to avoid over-provisioning costs |  |
| **High** | COST-ELB-02 | AWS ELB Cost Optimization Best Practices Current | Optimize Load Balancer Capacity Units (LCUs) | Monitor and optimize LCU consumption for ALB and NLB | **Platform** - Use CloudWatch metrics to monitor LCU usage and optimize rules, connections, and bandwidth to reduce costs |  |
| Medium | COST-ELB-03 | AWS ELB Cost Optimization Best Practices Current | Consolidate Load Balancers | Use host-based and path-based routing to reduce number of load balancers | **IaC** - Configure multiple target groups and routing rules on single ALB instead of deploying multiple load balancers |  |
| Medium | COST-ELB-04 | AWS ELB Cost Optimization Best Practices Current | Remove Unused Load Balancers | Identify and decommission idle or unused load balancers | **Platform** - Implement automated monitoring to identify load balancers with zero or minimal traffic and evaluate for removal |  |
| Medium | COST-ELB-05 | AWS ELB Cost Optimization Best Practices Current | Optimize Data Transfer Costs | Minimize cross-AZ data transfer charges | **IaC** - Enable cross-zone load balancing judiciously and consider target placement to minimize inter-AZ traffic |  |
| Medium | COST-ELB-07 | AWS ELB Cost Optimization Best Practices Current | Implement Cost Monitoring and Alerting | Set up cost monitoring for ELB resources | **Platform** - Configure AWS Cost Explorer and CloudWatch billing alarms to monitor ELB costs and usage trends |  |
| Low | COST-ELB-06 | AWS ELB Cost Optimization Best Practices Current | Use Reserved Capacity for Predictable Workloads | Consider reserved pricing for steady-state load balancer usage | **Platform** - Analyze usage patterns and purchase reserved capacity where applicable to reduce costs |  |
| Low | COST-ELB-08 | AWS ELB Cost Optimization Best Practices Current | Optimize SSL Certificate Management | Use AWS Certificate Manager to avoid third-party certificate costs | **IaC** - Utilize AWS ACM for SSL certificates instead of purchasing third-party certificates |  |

