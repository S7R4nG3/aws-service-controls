# AWS Certificate Manager (ACM)
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management Policy and Procedures | IAM-01 | Establish formal identity and access management policies for ACM certificate lifecycle management | **IaC** - Implement IAM policies through CloudFormation/Terraform templates defining least-privilege access to ACM operations including certificate request, validation, renewal, and deletion |  |
| **High** | Data Security and Information Lifecycle Management | DSI-07 | Ensure proper handling and protection of certificate private keys and sensitive certificate data | **Platform** - ACM automatically manages private key protection using AWS KMS encryption and HSMs, with no user access to private keys |  |
| **High** | Encryption and Key Management | EKM-01 | Implement proper encryption key management for certificate private keys | **Platform** - ACM uses AWS managed HSMs and KMS for private key generation, storage, and protection with FIPS 140-2 Level 3 compliance |  |
| Medium | Audit Logging / Audit Trails | LOG-01 | Enable comprehensive logging of certificate management activities | **IaC** - Configure CloudTrail to log all ACM API calls and integrate with CloudWatch for monitoring certificate operations and lifecycle events |  |
| Medium | Threat and Vulnerability Management | TVM-01 | Monitor for certificate vulnerabilities and ensure timely renewal | **User** - Set up CloudWatch alarms for certificate expiration and implement automated renewal processes where possible |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Account Management | AC-2 | Manage accounts with access to certificate management functions | **IaC** - Implement IAM roles and policies with time-bound access and regular access reviews for ACM administrative functions |  |
| **High** | Least Privilege | AC-6 | Enforce least privilege access to ACM certificate operations | **IaC** - Create granular IAM policies limiting ACM permissions to specific actions and resources based on job functions |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect certificate data during transmission | **Platform** - ACM API calls are encrypted in transit using TLS 1.2+ and certificate deployment to AWS services maintains encryption |  |
| Medium | Event Logging | AU-2 | Log certificate management events for security monitoring | **IaC** - Enable CloudTrail logging for ACM events and configure log retention policies in CloudWatch Logs |  |
| Medium | System Backup | CP-9 | Ensure certificate configuration and metadata backup | **User** - Implement automated backup of certificate metadata and configuration using AWS Config and maintain certificate documentation |  |
| Low | System Monitoring | SI-4 | Monitor certificate usage and health | **IaC** - Configure CloudWatch metrics and alarms for certificate status monitoring and expiration notifications |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Imported and ACM-issued certificates should be renewed after a specified time period | ACM.1 | Ensure certificates are renewed before expiration to maintain service availability | **User** - Set up automated renewal for ACM-managed certificates and establish renewal procedures for imported certificates with CloudWatch alarms |  |
| Medium | RSA certificates managed by ACM should use a key length of at least 2048 bits | ACM.2 | Ensure adequate cryptographic strength for certificate keys | **User** - Configure certificate requests to use RSA-2048 or higher key lengths, or use ECDSA P-256 certificates for better security |  |

### AWS Security Hub Custom Controls 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Certificate Validation Method Security | ACM-CUSTOM-1 | Use secure validation methods for certificate domain validation | **User** - Prefer DNS validation over email validation for domain ownership verification and ensure DNS records are properly secured |  |
| Medium | Cross-Region Certificate Management | ACM-CUSTOM-2 | Implement consistent certificate management across regions | **IaC** - Deploy certificates in multiple regions as needed and maintain consistent naming and tagging strategies across regions |  |
| Low | Certificate Transparency Monitoring | ACM-CUSTOM-3 | Monitor Certificate Transparency logs for unauthorized certificates | **User** - Implement monitoring tools to track certificate issuance in Certificate Transparency logs for your domains |  |


## Operational Controls
---



## Cost Controls
---


### AWS ACM Cost Optimization Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Use ACM Public Certificates Instead of Third-Party | COST-01 | Utilize free ACM public certificates instead of purchasing third-party certificates | **User** - Replace third-party SSL certificates with free ACM public certificates for AWS-integrated services to eliminate certificate purchase costs |
| Medium | Optimize Multi-Domain Certificates | COST-02 | Use Subject Alternative Names (SAN) to consolidate multiple domains under single certificates | **User** - Request certificates with multiple domain names in SAN field instead of individual certificates to reduce management overhead |
| Medium | Eliminate Unused Certificates | COST-03 | Regular cleanup of unused or expired imported certificates | **User** - Implement automated processes to identify and remove unused imported certificates to avoid unnecessary storage costs |
| Medium | Leverage Wildcard Certificates | COST-04 | Use wildcard certificates for multiple subdomains of the same domain | **User** - Request wildcard certificates (*.example.com) for applications with multiple subdomains to reduce certificate management complexity |
| Low | Optimize Regional Certificate Placement | COST-05 | Deploy certificates only in required regions to minimize cross-region data transfer costs | **IaC** - Provision ACM certificates only in regions where they are actively used by services to optimize network costs |
| Low | Automate Certificate Lifecycle Management | COST-06 | Implement automation to reduce operational overhead and manual certificate management costs | **IaC** - Use Infrastructure as Code and AWS Lambda functions to automate certificate provisioning, renewal notifications, and lifecycle management |


