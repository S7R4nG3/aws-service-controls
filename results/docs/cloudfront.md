# AWS CloudFront
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **Critical** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Entitlement | Establish comprehensive identity and access management policies for CloudFront distribution access | **IaC** - Implement IAM policies and roles through CloudFormation/Terraform templates to control CloudFront distribution management access |  |
| **Critical** | DSI-01 | CSA Cloud Controls Matrix v4.0 | Classification | Implement data classification and protection measures for content distributed through CloudFront | **Platform** - Configure SSL/TLS certificates, enable HTTPS-only viewer protocol policy, and implement Origin Access Control (OAC) |  |
| **High** | EKM-01 | CSA Cloud Controls Matrix v4.0 | Entitlement | Ensure proper encryption of data in transit and at rest for CloudFront distributions | **IaC** - Configure field-level encryption, use AWS Certificate Manager for SSL certificates, and enforce minimum TLS version 1.2 |  |
| **High** | IVS-01 | CSA Cloud Controls Matrix v4.0 | Network Architecture | Secure CloudFront infrastructure configurations and edge locations | **Platform** - Configure security headers, implement WAF integration, and enable geo-restriction capabilities |  |
| Medium | AIS-01 | CSA Cloud Controls Matrix v4.0 | Anti-Virus / Malicious Software | Secure application interfaces and APIs served through CloudFront | **User** - Implement custom headers validation, configure cache behaviors for API endpoints, and set appropriate TTL values |  |
| **Critical** | AC-3 | NIST 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to CloudFront distributions and cached content | **IaC** - Implement signed URLs/cookies, configure Origin Access Control, and use IAM policies for distribution management |  |
| **Critical** | SC-8 | NIST 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information through CloudFront | **Platform** - Enable HTTPS viewer protocol policy, configure HSTS headers, and implement field-level encryption for sensitive data |  |
| **High** | AU-2 | NIST 800-53 Rev 5 | Event Logging | Ensure CloudFront generates audit records for security-relevant events | **IaC** - Enable CloudFront access logging, configure real-time logs, and integrate with CloudTrail for API logging |  |
| **High** | SI-4 | NIST 800-53 Rev 5 | System Monitoring | Monitor CloudFront distributions for indicators of compromise and unauthorized activities | **Platform** - Configure CloudWatch alarms, enable AWS WAF logging, and implement GuardDuty for threat detection |  |
| Medium | SC-7 | NIST 800-53 Rev 5 | Boundary Protection | Monitor and control communications at external boundaries through CloudFront edge locations | **User** - Configure geo-restriction, implement custom error pages, and use response headers policies |  |
| **Critical** | CloudFront.1 | AWS Foundational Security Best Practices 1.0.0 | CloudFront distributions should have a default root object configured | Prevent directory listing and information disclosure by configuring a default root object | **IaC** - Set default root object property in CloudFront distribution configuration (e.g., index.html) |  |
| **Critical** | CloudFront.2 | AWS Foundational Security Best Practices 1.0.0 | CloudFront distributions should have origin access control enabled | Restrict direct access to S3 origins by using Origin Access Control (OAC) | **IaC** - Configure Origin Access Control for S3 origins and update S3 bucket policies to allow only CloudFront access |  |
| **High** | CloudFront.3 | AWS Foundational Security Best Practices 1.0.0 | CloudFront distributions should require encryption in transit | Ensure all viewer communications use HTTPS to protect data in transit | **Platform** - Set viewer protocol policy to 'Redirect HTTP to HTTPS' or 'HTTPS Only' for all cache behaviors |  |
| **High** | CloudFront.4 | AWS Foundational Security Best Practices 1.0.0 | CloudFront distributions should have logging configured | Enable access logging to track requests and detect potential security issues | **IaC** - Configure access logging with appropriate S3 bucket and prefix for log storage and analysis |  |
| Medium | CloudFront.5 | AWS Foundational Security Best Practices 1.0.0 | CloudFront distributions should not use deprecated SSL protocols | Use modern TLS versions to ensure strong encryption standards | **User** - Configure minimum protocol version to TLSv1.2 or higher in viewer certificate settings |  |
| **Critical** | CloudFront.10 | AWS Security Hub 2023.4 | CloudFront distributions should not use deprecated SSL protocols | Ensure CloudFront distributions use secure SSL/TLS protocol versions | **Platform** - Configure viewer certificate to use security policy with TLSv1.2 minimum or TLS-1-2-2021 security policy |  |
| **Critical** | CloudFront.12 | AWS Security Hub 2023.4 | CloudFront distributions should not point to non-existent S3 origins | Prevent potential subdomain takeover attacks by ensuring S3 origins exist and are properly configured | **IaC** - Validate S3 bucket existence and ownership before configuring as CloudFront origin, implement resource dependency management |  |
| **High** | CloudFront.9 | AWS Security Hub 2023.4 | CloudFront distributions should have AWS WAF enabled | Protect against common web attacks by integrating AWS WAF with CloudFront distributions | **IaC** - Create and associate AWS WAF Web ACL with CloudFront distribution, configure appropriate rules for threat protection |  |
| **High** | CloudFront.13 | AWS Security Hub 2023.4 | CloudFront distributions should use SNI to serve HTTPS requests | Use Server Name Indication (SNI) for cost-effective SSL certificate management | **Platform** - Configure SNI-only SSL certificate support instead of dedicated IP SSL certificates unless specifically required |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS CloudFront Cost Optimization Best Practices 2024 | Optimize Cache Behaviors and TTL Settings | Configure appropriate TTL values to maximize cache hit ratios and reduce origin requests | **User** - Set longer TTL values for static content, shorter TTL for dynamic content, and use cache-control headers effectively |  |
| **High** | COST-02 | AWS CloudFront Cost Optimization Best Practices 2024 | Select Appropriate Price Class | Choose the most cost-effective price class based on geographic distribution requirements | **IaC** - Configure price class to 'Use Only US, Canada and Europe' or 'Use Only US and Europe' if global distribution is not required |  |
| Medium | COST-03 | AWS CloudFront Cost Optimization Best Practices 2024 | Implement Compression | Enable compression to reduce data transfer costs and improve performance | **Platform** - Enable automatic compression in CloudFront cache behaviors for compressible file types |  |
| Medium | COST-04 | AWS CloudFront Cost Optimization Best Practices 2024 | Optimize Origin Request Patterns | Minimize origin requests through effective caching strategies and origin shielding | **User** - Enable origin shield for frequently requested content and configure appropriate cache key parameters |  |
| Medium | COST-05 | AWS CloudFront Cost Optimization Best Practices 2024 | Monitor and Analyze Usage Patterns | Regularly review CloudFront usage reports to identify cost optimization opportunities | **Platform** - Set up CloudWatch dashboards and billing alerts to monitor data transfer costs and cache performance metrics |  |
| Low | COST-06 | AWS CloudFront Cost Optimization Best Practices 2024 | Use Reserved Capacity for Predictable Workloads | Consider CloudFront Reserved Capacity for predictable, high-volume usage patterns | **Platform** - Analyze historical usage patterns and purchase reserved capacity for consistent data transfer volumes |  |
| Low | COST-07 | AWS CloudFront Cost Optimization Best Practices 2024 | Optimize Request Methods and Headers | Configure cache behaviors to avoid unnecessary cache variations based on headers, cookies, and query strings | **User** - Whitelist only necessary headers, cookies, and query string parameters for caching to improve cache hit ratios |  |

