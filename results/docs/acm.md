# AWS Certificate Manager
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | IAM-01 | CSA Cloud Controls Matrix v4.0.10 | Entitlement | Implement proper identity and access management controls for certificate management operations | **IaC** - Configure IAM policies with least privilege access for ACM operations, implement role-based access control with specific permissions for certificate lifecycle management |  |
| **High** | EKM-02 | CSA Cloud Controls Matrix v4.0.10 | Key Generation | Ensure proper encryption key management for certificate private keys | **Platform** - AWS ACM automatically manages private key security using AWS KMS integration, ensure KMS key policies are properly configured |  |
| Medium | AIS-01 | CSA Cloud Controls Matrix v4.0.10 | Web Application Firewall | Secure API access to certificate management functions | **IaC** - Implement API Gateway with proper authentication, use VPC endpoints for ACM API calls, enable API logging and monitoring |  |
| Medium | IPY-01 | CSA Cloud Controls Matrix v4.0.10 | Policy | Ensure certificate portability and integration capabilities | **User** - Document certificate dependencies, maintain certificate inventory, plan for certificate migration scenarios |  |
| **High** | SC-12 | NIST 800-53 Rev 5 | Cryptographic Key Establishment and Management | Establish and manage cryptographic keys for required cryptography employed within the system | **Platform** - Leverage ACM's automated key generation and management, integrate with AWS KMS for additional key management controls |  |
| **High** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to certificate management functions | **IaC** - Implement IAM policies with condition keys, use resource-based policies for certificate access control, enable MFA for sensitive operations |  |
| **High** | AU-12 | NIST 800-53 Rev 5 | Audit Record Generation | Generate audit records for certificate lifecycle events | **IaC** - Enable CloudTrail logging for ACM API calls, configure CloudWatch Events for certificate state changes, implement centralized logging |  |
| Medium | SC-13 | NIST 800-53 Rev 5 | Cryptographic Protection | Implement cryptographic mechanisms to prevent unauthorized disclosure of information | **Platform** - Use ACM certificates with strong encryption algorithms, ensure TLS 1.2+ implementation, configure cipher suites appropriately |  |
| Medium | CM-8 | NIST 800-53 Rev 5 | System Component Inventory | Maintain an inventory of certificates and their associated systems | **User** - Implement certificate tracking using AWS Config, maintain certificate lifecycle documentation, use tagging for certificate categorization |  |
| Medium | ACM.1 | AWS Foundational Security Standard 1.0.0 | ACM certificates should have transparency logging enabled | Ensure Certificate Transparency logging is enabled for public certificates | **Platform** - ACM automatically enables Certificate Transparency logging for public certificates, verify logging is active for compliance |  |
| Medium | ACM.1 | AWS Security Hub 2023 | ACM certificates should have transparency logging enabled | ACM certificates should have transparency logging enabled to detect mistakenly or maliciously issued certificates | **Platform** - Enable Certificate Transparency logging through ACM console or API, monitor CT logs for unauthorized certificates |  |
| Medium | ACM.2 | AWS Security Hub 2023 | RSA certificates managed by ACM should use a key length of at least 2048 bits | Ensure RSA certificates use sufficient key length for cryptographic strength | **User** - Request certificates with RSA-2048 or higher key lengths, validate existing certificates meet minimum requirements |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS ACM Cost Optimization Best Practices 2023 | Use ACM Public Certificates for Internet-Facing Applications | Utilize free ACM public certificates instead of purchasing third-party certificates for AWS services | **User** - Replace third-party certificates with ACM public certificates for ELB, CloudFront, and API Gateway to eliminate certificate procurement costs |  |
| **High** | COST-02 | AWS ACM Cost Optimization Best Practices 2023 | Implement Automated Certificate Renewal | Leverage ACM's automatic renewal to reduce operational overhead and prevent certificate expiration incidents | **Platform** - Enable automatic renewal for ACM certificates to avoid manual renewal processes and potential downtime costs |  |
| Medium | COST-03 | AWS ACM Cost Optimization Best Practices 2023 | Optimize Certificate Usage Across Services | Use single wildcard certificates for multiple subdomains to reduce certificate management overhead | **User** - Request wildcard certificates (*.domain.com) for multiple subdomain applications to consolidate certificate management |  |
| Medium | COST-04 | AWS ACM Cost Optimization Best Practices 2023 | Monitor and Clean Up Unused Certificates | Regularly audit and remove unused certificates to maintain clean certificate inventory | **User** - Implement periodic reviews of certificate usage, remove certificates not associated with active AWS resources |  |
| Low | COST-05 | AWS ACM Cost Optimization Best Practices 2023 | Use Private Certificates Judiciously | Evaluate the cost-benefit of ACM Private CA versus self-managed CA solutions for internal certificates | **User** - Assess Private CA monthly fees against certificate volume and management overhead, consider alternatives for low-volume scenarios |  |

