# AWS Elastic Load Balancer
---


### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure IAM policies and roles with least privilege access to ELB resources using resource-based policies and condition statements |  |
| HIGH | Boundary Protection | SC-7 | Monitor and control communications at the external boundary of the system | **IaC** - Configure security groups and NACLs to restrict traffic, implement WAF rules, and use internal load balancers for private traffic |  |
| HIGH | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Enable SSL/TLS termination on ALB/NLB with strong cipher suites and redirect HTTP to HTTPS |  |
| HIGH | Event Logging | AU-2 | Identify the types of events that the system is capable of logging | **IaC** - Enable access logging to S3 bucket and configure CloudTrail for ELB API calls |  |
| MEDIUM | System Monitoring | SI-4 | Monitor the system to detect attacks and indicators of potential attacks | **IaC** - Configure CloudWatch metrics and alarms for ELB performance and security events |  |
| MEDIUM | Contingency Plan | CP-2 | Develop a contingency plan for the system that identifies essential missions and business functions | **User** - Deploy ELB across multiple AZs and implement cross-region disaster recovery procedures |  |

### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Identity and Access Management Policy and Procedures | IAM-01 | Establish and maintain identity and access management policies and procedures | **User** - Implement IAM policies with least privilege for ELB management and establish regular access reviews |  |
| HIGH | Clock Synchronization | IVS-01 | Ensure system clocks are synchronized across all relevant systems | **Platform** - AWS manages clock synchronization for ELB service infrastructure |  |
| HIGH | Data Classification | DSI-01 | Data and objects are classified according to organizational policies | **User** - Tag ELB resources with appropriate data classification labels and implement corresponding security controls |  |
| HIGH | Encryption Key Management | EKM-01 | Encryption keys are managed throughout their lifecycle | **IaC** - Use AWS Certificate Manager for SSL/TLS certificates and implement proper key rotation policies |  |
| MEDIUM | Policy Enforcement Point | IPY-01 | Implement policy enforcement points for network traffic | **IaC** - Configure listener rules and target group health checks to enforce traffic policies |  |
| MEDIUM | Threat and Vulnerability Management | TVM-01 | Implement threat and vulnerability management processes | **User** - Regularly assess ELB configurations using AWS Config rules and Security Hub findings |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| MEDIUM | Classic Load Balancers with SSL/HTTPS listeners should use predefined security policies | ELB.2 | Classic Load Balancers should use predefined security policies that have strong ciphers | **IaC** - Configure CLB with predefined security policies like ELBSecurityPolicy-TLS-1-2-2017-01 or newer |  |
| HIGH | Application Load Balancer should be configured to redirect HTTP requests to HTTPS | ELB.3 | ALB listeners should redirect HTTP requests to HTTPS to ensure secure communication | **IaC** - Configure ALB listener rules to redirect HTTP traffic to HTTPS using redirect actions |  |
| MEDIUM | Application Load Balancers should be configured to drop HTTP headers | ELB.4 | ALB should be configured to drop invalid HTTP headers to prevent potential attacks | **IaC** - Enable routing.http.drop_invalid_header_fields.enabled attribute on ALB |  |
| MEDIUM | Application and Network Load Balancers should have logging enabled | ELB.5 | Load balancers should have access logging enabled to track requests | **IaC** - Enable access logging to S3 bucket with appropriate bucket policies and lifecycle management |  |
| LOW | Application Load Balancer deletion protection should be enabled | ELB.6 | Enable deletion protection to prevent accidental deletion of load balancers | **IaC** - Set deletion_protection attribute to true in ALB configuration |  |
| LOW | Classic Load Balancers should have connection draining enabled | ELB.7 | Enable connection draining to complete in-flight requests during instance deregistration | **IaC** - Configure connection draining with appropriate timeout values on CLB |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Classic Load Balancer listeners should be configured with HTTPS or TLS termination | ELB.1 | CLB should use secure listeners to encrypt data in transit | **IaC** - Configure HTTPS/TLS listeners on CLB with valid SSL certificates from ACM |  |
| MEDIUM | Classic Load Balancers with SSL/HTTPS listeners should use predefined security policies | ELB.8 | Use predefined security policies with strong encryption algorithms | **IaC** - Apply security policies like ELBSecurityPolicy-TLS-1-2-2017-01 or ELBSecurityPolicy-FS-1-2-Res-2020-10 |  |
| HIGH | Network Load Balancers should have TLS listeners for encrypted connections | ELB.9 | NLB should use TLS listeners to encrypt traffic between clients and load balancer | **IaC** - Configure TLS listeners on NLB with SSL certificates and appropriate security policies |  |
| HIGH | Load balancers should span multiple Availability Zones | ELB.10 | Load balancers should be deployed across multiple AZs for high availability | **IaC** - Configure load balancer subnets across at least two different Availability Zones |  |
| MEDIUM | Load balancer target groups should have health check enabled | ELB.11 | Target groups should have health checks configured to ensure traffic is routed to healthy targets | **IaC** - Configure health check settings including path, interval, timeout, and healthy/unhealthy thresholds |  |
| LOW | Application Load Balancers should be configured with defensive or strictest desync mitigation mode | ELB.12 | ALB should be configured to handle HTTP desync attacks | **IaC** - Set routing.http.desync_mitigation_mode to defensive or strictest on ALB |  |


## Operational Controls
---



## Cost Controls
---


### AWS ELB Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| HIGH | Right-size Load Balancer Type | COST-01 | Choose the appropriate load balancer type based on actual requirements to avoid unnecessary costs | **User** - Evaluate traffic patterns and requirements to select between ALB, NLB, and CLB based on feature needs and cost structure |
| HIGH | Optimize Load Balancer Capacity Units (LCU) | COST-02 | Monitor and optimize LCU usage to reduce costs for ALB and NLB | **User** - Monitor CloudWatch metrics for processed bytes, active connections, new connections, and rule evaluations to optimize LCU consumption |
| MEDIUM | Consolidate Load Balancers | COST-03 | Consolidate multiple load balancers where possible to reduce fixed hourly costs | **IaC** - Use host-based and path-based routing rules to serve multiple applications from a single ALB instead of deploying separate load balancers |
| MEDIUM | Optimize Cross-Zone Load Balancing | COST-04 | Evaluate cross-zone load balancing settings to optimize data transfer costs | **IaC** - Disable cross-zone load balancing for NLB if traffic distribution allows, and consider regional traffic patterns for ALB |
| LOW | Implement Efficient Health Checks | COST-05 | Optimize health check frequency and target to reduce unnecessary load and potential costs | **IaC** - Configure appropriate health check intervals and paths to minimize resource consumption while maintaining reliability |
| MEDIUM | Monitor and Optimize Idle Load Balancers | COST-06 | Identify and remove unused or underutilized load balancers | **User** - Set up CloudWatch alarms and regular reviews to identify load balancers with minimal traffic and evaluate their necessity |
| LOW | Optimize Target Group Configuration | COST-07 | Configure target groups efficiently to minimize processing overhead | **IaC** - Optimize target group size, deregistration delay, and connection draining settings to reduce unnecessary resource consumption |
| LOW | Leverage Reserved Capacity | COST-08 | Consider reserved capacity options for predictable, long-term load balancer usage | **User** - Analyze usage patterns and consider AWS Savings Plans or other commitment-based pricing models for consistent workloads |


