# AWS Elastic Container Registry
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-02 | CSA CCM v4.0 | User Access Authorization | Ensure proper authentication and authorization controls are implemented for ECR access | **Platform** - Configure IAM policies and roles with least privilege access to ECR repositories, implement resource-based policies for repository access control |  |
| **High** | DSI-01 | CSA CCM v4.0 | Data Classification | Implement encryption for data at rest and in transit for container images | **IaC** - Enable encryption at rest using AWS KMS keys and ensure HTTPS/TLS for data in transit through ECR API endpoints |  |
| **High** | DSI-07 | CSA CCM v4.0 | Secure Disposal | Implement secure deletion of container images and associated metadata | **IaC** - Configure lifecycle policies to automatically delete unused images and implement secure deletion procedures for sensitive container images |  |
| Medium | IVS-06 | CSA CCM v4.0 | Network Architecture | Ensure proper network security controls for ECR access | **Platform** - Configure VPC endpoints for ECR access and implement network access controls |  |
| **Critical** | AC-2 | NIST 800-53 Rev 5 | Account Management | Manage ECR user accounts and access permissions systematically | **Platform** - Implement automated account provisioning and deprovisioning for ECR access using IAM roles and policies |  |
| **Critical** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to ECR repositories | **IaC** - Configure repository policies and IAM policies to enforce least privilege access to specific ECR repositories and actions |  |
| **High** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted container images | **Platform** - Use HTTPS/TLS for all ECR API communications and Docker client interactions with ECR endpoints |  |
| **High** | SC-28 | NIST 800-53 Rev 5 | Protection of Information at Rest | Protect confidentiality and integrity of container images stored in ECR | **IaC** - Enable server-side encryption using AWS KMS customer-managed keys for ECR repositories |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Ensure ECR activities are logged for security monitoring | **Platform** - Enable CloudTrail logging for ECR API calls and configure CloudWatch for ECR events monitoring |  |
| Medium | SI-7 | NIST 800-53 Rev 5 | Software, Firmware, and Information Integrity | Verify integrity of container images stored in ECR | **User** - Implement image signing and verification using Docker Content Trust or similar mechanisms |  |
| **High** | ECR.1 | AWS Foundational Security Best Practices 1.0.0 | ECR private repositories should have image scanning configured | Enable vulnerability scanning for container images in ECR repositories | **IaC** - Configure scan on push for ECR repositories and implement automated scanning policies |  |
| Medium | ECR.2 | AWS Foundational Security Best Practices 1.0.0 | ECR private repositories should have tag immutability configured | Enable tag immutability to prevent image tags from being overwritten | **IaC** - Set imageTagMutability to IMMUTABLE when creating ECR repositories |  |
| Medium | ECR.3 | AWS Foundational Security Best Practices 1.0.0 | ECR repositories should have at least one lifecycle policy configured | Implement lifecycle policies to manage image retention and reduce storage costs | **IaC** - Configure lifecycle policies to automatically clean up old or unused container images based on age or count |  |
| **High** | ECR.1 | AWS Security Hub 2024 | ECR repositories should have image scanning configured | Automatically scan container images for vulnerabilities | **IaC** - Enable enhanced scanning with Inspector or basic scanning for ECR repositories |  |
| Medium | ECR.2 | AWS Security Hub 2024 | ECR repositories should have tag immutability enabled | Prevent accidental overwriting of container image tags | **IaC** - Configure tag immutability setting to IMMUTABLE for production ECR repositories |  |
| Medium | ECR.3 | AWS Security Hub 2024 | ECR repositories should have lifecycle policies | Automatically manage image lifecycle to control storage costs and maintain security | **IaC** - Create lifecycle policies to remove old images, untagged images, and limit repository size |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS ECR Cost Optimization Best Practices 2024 | Implement Lifecycle Policies | Automatically delete old and unused container images to reduce storage costs | **IaC** - Configure lifecycle policies to remove images older than specified days, limit number of images per repository, and delete untagged images |  |
| **High** | COST-02 | AWS ECR Cost Optimization Best Practices 2024 | Optimize Image Layers | Reduce image size and storage costs through efficient layering and multi-stage builds | **User** - Use multi-stage Docker builds, minimize layers, and remove unnecessary packages and files from container images |  |
| Medium | COST-03 | AWS ECR Cost Optimization Best Practices 2024 | Monitor Repository Usage | Track repository usage patterns to identify cost optimization opportunities | **Platform** - Use CloudWatch metrics and AWS Cost Explorer to monitor ECR storage usage, data transfer costs, and repository access patterns |  |
| Medium | COST-04 | AWS ECR Cost Optimization Best Practices 2024 | Implement Image Deduplication | Reduce storage costs by avoiding duplicate image layers across repositories | **User** - Use common base images across applications and implement consistent tagging strategies to maximize layer sharing |  |
| Medium | COST-05 | AWS ECR Cost Optimization Best Practices 2024 | Regular Repository Cleanup | Periodically review and clean up unused repositories and images | **User** - Establish processes to regularly audit ECR repositories, remove unused repositories, and clean up test/development images |  |
| Low | COST-06 | AWS ECR Cost Optimization Best Practices 2024 | Optimize Data Transfer | Minimize data transfer costs by using ECR in the same region as compute resources | **IaC** - Deploy ECR repositories in the same AWS region as ECS, EKS, or EC2 instances to avoid cross-region data transfer charges |  |

