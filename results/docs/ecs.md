# AWS Elastic Container Service (ECS)
---


### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Amazon ECS task definitions should have secure networking modes and user definitions | ECS.1 | ECS task definitions should not have elevated privileges and should use secure networking modes to prevent unauthorized access | **IaC** - Configure task definitions with non-root users, avoid privileged mode, and use awsvpc network mode for better isolation |  |
| HIGH | ECS services should not have public IP addresses assigned automatically | ECS.2 | ECS services should not automatically assign public IP addresses to tasks to reduce attack surface | **IaC** - Set assignPublicIp to DISABLED in service configuration and use NAT Gateway or VPC endpoints for outbound connectivity |  |
| HIGH | ECS task definitions should not share the host's process namespace | ECS.3 | Task definitions should not share the host's process namespace to maintain container isolation | **IaC** - Set pidMode to task or omit pidMode parameter in task definition to avoid sharing host process namespace |  |
| MEDIUM | ECS containers should run as non-privileged | ECS.4 | Containers should not run with privileged access to the host system | **IaC** - Set privileged to false in container definition and avoid running containers as root user |  |
| MEDIUM | ECS containers should be limited in their access to host resources | ECS.5 | Container access to host resources should be restricted to minimize security risks | **IaC** - Limit or avoid host port mappings, restrict volume mounts, and configure appropriate ulimits |  |
| MEDIUM | Secrets should not be passed as container environment variables | ECS.8 | Sensitive information should not be passed through environment variables | **IaC** - Use AWS Systems Manager Parameter Store or AWS Secrets Manager with secrets retrieval in task definitions |  |
| MEDIUM | ECS Fargate services should run on the latest Fargate platform version | ECS.10 | Fargate services should use the latest platform version for security patches and improvements | **Platform** - Set platformVersion to LATEST or specify the most recent platform version in service configuration |  |
| LOW | ECS clusters should use Container Insights | ECS.12 | Container Insights should be enabled for monitoring and observability | **IaC** - Enable Container Insights in cluster configuration for enhanced monitoring and logging capabilities |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Account Management | AC-2 | Manage user accounts, group memberships, privileges, and access authorizations | **User** - Implement IAM roles and policies for ECS access, use least privilege principle, and regularly audit access permissions |  |
| HIGH | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure IAM policies, resource-based policies, and security groups to enforce access controls |  |
| HIGH | Boundary Protection | SC-7 | Monitor and control communications at the external boundary and key internal boundaries | **IaC** - Implement VPC security groups, NACLs, and service mesh for network boundary protection |  |
| MEDIUM | Event Logging | AU-2 | Identify the types of events to be logged by the system | **IaC** - Enable CloudTrail for API logging, configure CloudWatch Logs for container logs, and enable VPC Flow Logs |  |
| MEDIUM | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Use TLS encryption for all communication, implement Application Load Balancer with SSL/TLS termination |  |
| MEDIUM | Baseline Configuration | CM-2 | Develop, document, and maintain baseline configurations | **IaC** - Use Infrastructure as Code templates, maintain version control for configurations, and implement configuration drift detection |  |
| MEDIUM | System Monitoring | SI-4 | Monitor the system to detect attacks and indicators of potential attacks | **Platform** - Enable GuardDuty, configure CloudWatch alarms, and implement AWS Security Hub for centralized monitoring |  |
| LOW | System Backup | CP-9 | Conduct backups of user-level and system-level information | **IaC** - Implement automated backups for ECS configurations, use AWS Backup for persistent volumes |  |

### CSA Cloud Controls Matrix v4.0.10
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Identity and Access Management Policy and Procedures | IAM-01 | Policies and procedures shall be established for identity and access management | **User** - Establish IAM policies for ECS resource access, implement role-based access control, and document access procedures |  |
| HIGH | User Access Provisioning | IAM-02 | User access shall be provisioned based on job function and least privilege | **User** - Use IAM roles for ECS tasks, implement least privilege access, and regularly review access permissions |  |
| HIGH | Data Security and Information Lifecycle Management | DSI-01 | Policies and procedures shall be established for data security and information lifecycle management | **IaC** - Implement encryption at rest and in transit, use AWS KMS for key management, and establish data retention policies |  |
| MEDIUM | Infrastructure and Virtualization Security Policy and Procedures | IVS-01 | Policies and procedures shall be established for infrastructure and virtualization security | **IaC** - Implement container security best practices, use security groups for network isolation, and maintain secure base images |  |
| MEDIUM | Application and Interface Security | AIS-01 | Policies and procedures shall be established for application and interface security | **IaC** - Implement secure coding practices, use WAF for web applications, and configure API security measures |  |
| MEDIUM | Business Continuity Management and Operational Resilience | BCR-01 | Policies and procedures shall be established for business continuity and operational resilience | **IaC** - Implement multi-AZ deployments, configure auto-scaling, and establish disaster recovery procedures |  |
| MEDIUM | Change Control and Configuration Management | CCC-01 | Policies and procedures shall be established for change control and configuration management | **IaC** - Use version control for infrastructure code, implement CI/CD pipelines, and maintain configuration baselines |  |
| LOW | Data Center Security | DCS-01 | Policies and procedures shall be established for data center security | **Platform** - Leverage AWS physical security controls, implement network segmentation, and use dedicated tenancy when required |  |

### AWS Security Hub 2.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| HIGH | Amazon ECS task definitions should have secure networking modes and user definitions | ECS.1 | This control checks whether an Amazon ECS task definition with host networking mode has privileged or user container definitions | **IaC** - Configure task definitions without privileged mode and avoid host networking mode unless absolutely necessary |  |
| HIGH | ECS services should not have public IP addresses assigned automatically | ECS.2 | This control checks whether Amazon ECS services are configured to automatically assign public IP addresses | **IaC** - Configure ECS services to not automatically assign public IP addresses and use load balancers for external access |  |
| MEDIUM | ECS task definitions should not share the host's process namespace | ECS.3 | This control checks if Amazon ECS task definitions are configured to share a host's process namespace with its containers | **IaC** - Ensure task definitions do not set pidMode to 'host' to maintain container isolation |  |
| MEDIUM | ECS containers should run as non-privileged | ECS.4 | This control checks if the privileged parameter in the container definition of Amazon ECS Task Definitions is set to true | **IaC** - Set privileged parameter to false in container definitions and run containers with minimal required privileges |  |
| MEDIUM | ECS containers should be limited in their access to host resources | ECS.5 | This control checks if Amazon ECS containers are limited in their access to host resources | **IaC** - Configure appropriate resource limits and avoid mounting sensitive host directories |  |
| MEDIUM | Secrets should not be passed as container environment variables | ECS.8 | This control checks whether the key value of any variables in the environment file of Amazon ECS task definitions contains secrets | **IaC** - Use AWS Secrets Manager or Systems Manager Parameter Store instead of environment variables for sensitive data |  |
| MEDIUM | ECS Fargate services should run on the latest Fargate platform version | ECS.10 | This control checks if Amazon ECS Fargate services are running the latest Fargate platform version | **Platform** - Regularly update Fargate platform version to the latest available version for security patches |  |
| LOW | ECS clusters should use Container Insights | ECS.12 | This control checks if Amazon ECS clusters use Container Insights | **IaC** - Enable Container Insights on ECS clusters for enhanced monitoring and observability |  |


## Operational Controls
---



## Cost Controls
---


### AWS ECS Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| HIGH | Right-size ECS Tasks and Services | COST-01 | Optimize CPU and memory allocation for ECS tasks to avoid over-provisioning | **IaC** - Monitor resource utilization using CloudWatch metrics and adjust task definitions to match actual resource requirements |
| HIGH | Use Fargate Spot for Fault-Tolerant Workloads | COST-02 | Leverage Fargate Spot pricing for cost-effective execution of fault-tolerant workloads | **IaC** - Configure ECS services to use Fargate Spot capacity providers for workloads that can handle interruptions |
| HIGH | Implement Auto Scaling | COST-03 | Configure auto scaling to automatically adjust the number of tasks based on demand | **IaC** - Set up Application Auto Scaling policies based on CPU utilization, memory utilization, or custom metrics |
| MEDIUM | Use EC2 Spot Instances for ECS Clusters | COST-04 | Leverage EC2 Spot Instances in ECS clusters for significant cost savings | **IaC** - Configure ECS cluster with mixed instance types including Spot instances using capacity providers |
| MEDIUM | Optimize Container Images | COST-05 | Reduce container image sizes to minimize storage costs and improve startup times | **User** - Use multi-stage builds, minimal base images, and remove unnecessary packages from container images |
| MEDIUM | Implement Resource-Based Scheduling | COST-06 | Use placement strategies to optimize resource utilization across cluster instances | **IaC** - Configure binpack placement strategy to maximize resource utilization and minimize instance count |
| MEDIUM | Use Reserved Instances for Predictable Workloads | COST-07 | Purchase Reserved Instances for predictable workloads to reduce costs | **User** - Analyze usage patterns and purchase Reserved Instances for consistent workloads running on EC2-backed ECS clusters |
| LOW | Implement Cost Monitoring and Alerting | COST-08 | Set up cost monitoring and alerting to track ECS spending | **Platform** - Use AWS Cost Explorer, set up billing alerts, and implement cost allocation tags for ECS resources |
| LOW | Optimize Load Balancer Usage | COST-09 | Right-size and optimize load balancer configurations to reduce costs | **IaC** - Use Application Load Balancer efficiently, consider shared load balancers for multiple services, and optimize health check intervals |
| LOW | Schedule Non-Critical Workloads | COST-10 | Schedule non-critical workloads to run during off-peak hours | **IaC** - Use scheduled scaling or EventBridge rules to run batch workloads during off-peak hours when costs are lower |


