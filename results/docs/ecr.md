# AWS Elastic Container Registry
---


### Cloud Security Alliance (CSA) Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Identity and Access Management Policy and Procedures | IAM-01 | Establish and maintain identity and access management policies for ECR resources | **Platform + IaC** - Configure ECR repository policies, IAM roles and policies for ECR access using IaC templates with least privilege principles |  |
| **Critical** | Data Security and Information Lifecycle Management | DSI-01 | Implement data protection controls for container images and sensitive data | **Platform + User** - Enable ECR image scanning, implement encryption at rest and in transit, establish image lifecycle policies |  |
| **Critical** | Encryption and Key Management | EKM-01 | Implement proper encryption and key management for ECR repositories | **Platform + IaC** - Configure ECR repositories with KMS encryption using customer-managed keys, implement key rotation policies |  |
| **High** | Audit Logging / Intrusion Detection | LOG-01 | Enable comprehensive logging and monitoring for ECR activities | **Platform + IaC** - Enable CloudTrail logging for ECR API calls, configure VPC Flow Logs for ECR endpoints, set up CloudWatch monitoring |  |
| **High** | Vulnerability / Patch Management | IVS-01 | Implement vulnerability scanning and patch management for container images | **Platform + User** - Enable ECR enhanced scanning, implement automated vulnerability remediation workflows, establish image promotion policies |  |
| **High** | Network Security | NET-01 | Implement network security controls for ECR access | **Platform + IaC** - Configure VPC endpoints for ECR, implement security groups and NACLs, enable private repository access |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to ECR resources | **Platform + IaC** - Implement resource-based policies on ECR repositories, configure cross-account access controls, use IAM conditions for fine-grained access |  |
| **Critical** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **Platform** - Ensure HTTPS/TLS encryption for all ECR API communications, configure Docker clients for secure registry communication |  |
| **Critical** | Protection of Information at Rest | SC-28 | Protect the confidentiality and integrity of information at rest | **Platform + IaC** - Enable ECR encryption at rest using AWS KMS, implement server-side encryption for all repositories |  |
| **High** | Event Logging | AU-2 | Ensure that auditable events are being logged | **Platform + IaC** - Configure CloudTrail to log ECR API events, enable ECR repository access logging, implement log aggregation |  |
| **High** | Malicious Code Protection | SI-3 | Implement malicious code protection mechanisms | **Platform + User** - Enable ECR image scanning for malware detection, implement admission controllers for Kubernetes clusters |  |
| **High** | System Monitoring | SI-4 | Monitor the system to detect attacks and indicators of potential attacks | **Platform + IaC** - Configure CloudWatch metrics for ECR, set up GuardDuty for threat detection, implement automated alerting |  |
| Medium | Vulnerability Monitoring and Scanning | RA-5 | Monitor and scan for vulnerabilities and remediate them | **Platform + User** - Enable ECR enhanced scanning with Inspector integration, implement vulnerability management workflows |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | ECR private repositories should have image scanning configured | ECR.1 | ECR repositories should have image scanning enabled to identify software vulnerabilities | **Platform + IaC** - Enable scan on push for ECR repositories, configure enhanced scanning for comprehensive vulnerability detection |  |
| Medium | ECR private repositories should have tag immutability configured | ECR.2 | ECR repositories should have tag immutability enabled to prevent image tags from being overwritten | **Platform + IaC** - Configure imageTagMutability to IMMUTABLE in ECR repository configuration |  |
| Medium | ECR repositories should have at least one lifecycle policy configured | ECR.3 | ECR repositories should have lifecycle policies to manage image retention and reduce storage costs | **Platform + IaC** - Implement lifecycle policies to automatically delete untagged images and old image versions |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | ECR private repositories should have image scanning configured | ECR.1 | This control checks whether ECR private repositories have image scanning enabled | **Platform + IaC** - Enable automatic scanning on image push and configure enhanced scanning for detailed vulnerability reports |  |
| Medium | ECR private repositories should have tag immutability configured | ECR.2 | This control checks whether ECR private repositories have tag immutability enabled | **Platform + IaC** - Set repository imageTagMutability to IMMUTABLE to prevent accidental overwrites of image tags |  |
| Low | ECR repositories should have at least one lifecycle policy configured | ECR.3 | This control checks whether ECR repositories have lifecycle policies configured | **Platform + IaC** - Create and apply lifecycle policies to automatically clean up old or unused container images |  |


## Operational Controls
---



## Cost Controls
---


### AWS ECR Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Implement Repository Lifecycle Policies | COST-01 | Configure lifecycle policies to automatically delete unused and old container images to reduce storage costs | **Platform + IaC** - Set up lifecycle rules to delete untagged images after 1 day and keep only the latest 10 tagged images per repository |
| **High** | Optimize Image Size and Layers | COST-02 | Minimize container image size and optimize layers to reduce storage and data transfer costs | **User** - Use multi-stage builds, minimize base images, combine RUN commands, and remove unnecessary packages and files |
| Medium | Monitor and Analyze Repository Usage | COST-03 | Regularly monitor repository usage and costs to identify optimization opportunities | **Platform + User** - Use CloudWatch metrics and AWS Cost Explorer to track ECR usage patterns and storage costs |
| Medium | Use Regional Repositories | COST-04 | Deploy ECR repositories in the same region as compute resources to minimize data transfer costs | **Platform + IaC** - Create ECR repositories in regions where container workloads are deployed to avoid cross-region data transfer charges |
| Medium | Implement Image Replication Strategy | COST-05 | Use ECR replication rules strategically to balance availability and cost | **Platform + IaC** - Configure selective replication rules based on image tags or repository filters rather than replicating all images |
| Low | Optimize Scanning Costs | COST-06 | Balance security scanning needs with cost by optimizing scanning frequency and scope | **Platform + User** - Enable scan-on-push for production images only, use basic scanning for development repositories unless enhanced scanning is required |


