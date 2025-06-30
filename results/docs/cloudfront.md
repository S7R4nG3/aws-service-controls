# AWS CloudFront
---


### CSA CCM v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management | IAM-01 | Ensure proper identity and access management controls are implemented for CloudFront distributions | **IaC** - Configure IAM policies with least privilege access for CloudFront operations and implement origin access control (OAC) for S3 origins |  |
| **High** | Encryption Key Management | EKM-02 | Implement proper encryption key management for data in transit and at rest | **IaC** - Configure SSL/TLS certificates using AWS Certificate Manager and enforce HTTPS-only connections with security headers |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Protect sensitive data through proper classification and handling | **IaC** - Implement field-level encryption for sensitive data and configure appropriate cache behaviors based on data classification |  |
| Medium | Threat and Vulnerability Management | TVM-01 | Implement security monitoring and threat detection capabilities | **Platform** - Enable AWS WAF integration and configure security monitoring through CloudWatch and Security Hub |  |
| Medium | Audit Assurance and Compliance | LOG-01 | Ensure comprehensive logging and monitoring of CloudFront activities | **IaC** - Enable CloudFront access logs, real-time logs, and integrate with CloudTrail for API logging |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Configure Origin Access Control (OAC) and implement signed URLs/cookies for restricted content access |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information | **IaC** - Enforce TLS 1.2 or higher, implement HSTS headers, and configure secure cipher suites |  |
| **High** | Boundary Protection | SC-7 | Monitor and control communications at external boundaries and key internal boundaries | **Platform** - Integrate AWS WAF with CloudFront and configure geo-blocking and rate limiting as needed |  |
| Medium | Event Logging | AU-2 | Ensure system logs capture appropriate information for security monitoring | **IaC** - Configure CloudFront standard and real-time logs with appropriate log fields for security analysis |  |
| Medium | System Monitoring | SI-4 | Monitor system events and analyze for indicators of inappropriate or unusual activity | **Platform** - Set up CloudWatch alarms for CloudFront metrics and integrate with AWS Security Hub for centralized monitoring |  |
| Low | Contingency Plan | CP-2 | Develop and maintain contingency plans for system operations | **User** - Document CloudFront disaster recovery procedures and maintain multiple origin configurations where appropriate |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | CloudFront distributions should have a default root object configured | CloudFront.1 | CloudFront distributions should be configured with a default root object to prevent directory browsing | **IaC** - Configure default root object (e.g., index.html) in CloudFront distribution settings |  |
| **High** | CloudFront distributions should have origin access identity enabled | CloudFront.2 | CloudFront distributions should use Origin Access Control (OAC) to restrict direct access to S3 origins | **IaC** - Create and associate Origin Access Control with CloudFront distribution and update S3 bucket policy accordingly |  |
| **High** | CloudFront distributions should require encryption in transit | CloudFront.3 | CloudFront distributions should be configured to redirect HTTP requests to HTTPS | **IaC** - Set viewer protocol policy to 'redirect-to-https' or 'https-only' for all cache behaviors |  |
| Medium | CloudFront distributions should have origin failover configured | CloudFront.4 | CloudFront distributions should be configured with origin failover to improve availability | **IaC** - Configure origin groups with primary and secondary origins for automatic failover |  |
| Medium | CloudFront distributions should have logging enabled | CloudFront.5 | CloudFront distributions should have access logging enabled for security monitoring and analysis | **IaC** - Enable standard access logs and configure S3 bucket for log storage with appropriate retention policies |  |
| Medium | CloudFront distributions should have AWS WAF enabled | CloudFront.6 | CloudFront distributions should be associated with AWS WAF for additional security protection | **IaC** - Create and associate AWS WAF web ACL with CloudFront distribution, configure appropriate rules for protection |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | CloudFront distributions should use custom SSL/TLS certificates | CloudFront.7 | CloudFront distributions should use custom SSL certificates instead of default CloudFront certificates | **IaC** - Request or import SSL/TLS certificates in AWS Certificate Manager and associate with CloudFront distribution |  |
| **High** | CloudFront distributions should not use deprecated SSL protocols | CloudFront.8 | CloudFront distributions should be configured to use only secure SSL/TLS protocol versions | **IaC** - Configure minimum protocol version to TLSv1.2 or higher in the CloudFront distribution viewer certificate settings |  |
| Medium | CloudFront distributions should have field-level encryption enabled for sensitive data | CloudFront.9 | CloudFront distributions handling sensitive data should implement field-level encryption | **IaC** - Configure field-level encryption profiles and associate with cache behaviors that handle sensitive form data |  |
| Low | CloudFront distributions should have security headers configured | CloudFront.10 | CloudFront distributions should implement security headers for enhanced browser security | **IaC** - Use CloudFront Functions or Lambda@Edge to add security headers like HSTS, CSP, X-Frame-Options, and X-Content-Type-Options |  |


## Operational Controls
---



## Cost Controls
---


### AWS CloudFront Cost Optimization 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize CloudFront Price Class Selection | COST-01 | Select appropriate price class based on geographic distribution of users to minimize costs | **IaC** - Configure price class (Price Class 100, 200, or All) based on target audience geography to balance cost and performance |
| **High** | Implement Efficient Cache Behaviors | COST-02 | Configure cache behaviors to maximize cache hit ratios and reduce origin requests | **IaC** - Set appropriate TTL values, configure cache key normalization, and implement cache policies to increase cache efficiency |
| Medium | Optimize Origin Request Patterns | COST-03 | Minimize data transfer costs by optimizing requests between CloudFront and origins | **IaC** - Enable origin request policies to forward only necessary headers, cookies, and query strings to reduce origin load |
| Medium | Implement Content Compression | COST-04 | Enable automatic compression to reduce data transfer costs and improve performance | **IaC** - Enable automatic compression in cache behaviors for text-based content types to reduce bandwidth costs |
| Medium | Monitor and Analyze Usage Patterns | COST-05 | Regularly monitor CloudFront usage and costs to identify optimization opportunities | **Platform** - Set up CloudWatch dashboards and cost alerts to monitor data transfer, requests, and cache hit ratios |
| Low | Optimize Log Storage Costs | COST-06 | Implement cost-effective logging strategies to balance compliance and cost requirements | **IaC** - Configure S3 lifecycle policies for log storage, use appropriate storage classes, and consider log sampling for high-volume distributions |
| Low | Implement Reserved Capacity for Predictable Workloads | COST-07 | Consider CloudFront security savings bundle for predictable WAF usage patterns | **Platform** - Evaluate and purchase CloudFront security savings bundle if using AWS WAF consistently across distributions |


