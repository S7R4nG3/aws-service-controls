# AWS Amazon Certificate Manager
---


### CIS Controls v8.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Maintain Inventory of Certificate Assets | CIS-1.1 | Maintain an accurate inventory of all ACM certificates and their associated domains | **Platform** - Use AWS Config to track ACM certificate inventory and implement automated discovery of certificates across regions |  |
| **High** | Run Automated Certificate Vulnerability Scanning | CIS-3.1 | Monitor ACM certificates for vulnerabilities, weak encryption, and expiration risks | **Platform** - Use AWS Security Hub and third-party tools to scan certificate configurations and validate security settings |  |
| **High** | Establish Secure Certificate Configuration Process | CIS-4.1 | Define and enforce secure baseline configurations for ACM certificate requests and management | **IaC Template** - Use infrastructure as code templates with standardized certificate configurations including validation methods and key algorithms |  |
| Medium | Maintain Certificate Audit Logs | CIS-6.1 | Enable comprehensive logging for all ACM certificate lifecycle events and API operations | **Platform** - Enable CloudTrail logging for ACM API calls and configure CloudWatch Events for certificate lifecycle notifications |  |
| Medium | Activate Certificate Audit Logging | CIS-6.2 | Ensure audit logging captures all certificate management activities and access events | **Platform** - Configure CloudTrail with comprehensive ACM API logging and enable certificate transparency monitoring |  |
| Medium | Establish Data Recovery Processes | CIS-11.1 | Implement backup and recovery procedures for certificate management and private keys | **User** - Document certificate recovery procedures and maintain secure backups of certificate metadata and configurations |  |

### NIST Cybersecurity Framework v1.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Physical devices and systems are inventoried | NIST-ID.AM-1 | Maintain comprehensive inventory of all ACM certificates across all AWS regions and accounts | **Platform** - Use AWS Config Rules to monitor ACM certificate inventory and track certificate usage across services |  |
| Medium | Software platforms and applications are inventoried | NIST-ID.AM-2 | Track applications and services that utilize ACM certificates for SSL/TLS termination | **User** - Maintain documentation of certificate-to-service mappings and dependencies across the infrastructure |  |
| **High** | Identities and credentials are managed for authorized devices and users | NIST-PR.AC-1 | Implement proper IAM roles and policies for ACM certificate management and access | **IaC Template** - Define least-privilege IAM policies for ACM operations with resource-level permissions and conditions |  |
| **High** | Access permissions are managed, incorporating the principles of least privilege | NIST-PR.AC-4 | Apply least privilege access principles to ACM certificate operations and private key access | **User** - Create granular IAM policies limiting ACM actions to specific certificates and authorized personnel |  |
| **High** | Data-at-rest is protected | NIST-PR.DS-1 | Ensure ACM private keys are protected with appropriate encryption and access controls | **Platform** - Leverage ACM's built-in hardware security modules (HSMs) for private key protection and encryption |  |
| **High** | Data-in-transit is protected | NIST-PR.DS-2 | Ensure ACM certificates use strong encryption algorithms and are properly deployed for TLS termination | **IaC Template** - Configure ACM certificates with modern TLS versions and strong cipher suites in load balancers and services |  |
| Medium | Configuration change control processes are in place | NIST-PR.IP-3 | Implement change control processes for certificate lifecycle management and updates | **User** - Use infrastructure as code and approval workflows for certificate provisioning and renewal processes |  |
| Medium | The network is monitored to detect potential cybersecurity events | NIST-DE.CM-1 | Monitor certificate usage and detect potential certificate-related security events | **Platform** - Enable certificate transparency monitoring and configure alerts for certificate misuse or unauthorized issuance |  |
| Medium | Monitoring for unauthorized personnel, connections, devices, and software is performed | NIST-DE.CM-7 | Monitor for unauthorized certificate requests, modifications, or suspicious certificate activities | **Platform** - Use CloudTrail and CloudWatch to monitor ACM API calls and detect anomalous certificate management activities |  |
| Medium | Response plan is executed during or after an event | NIST-RS.RP-1 | Define incident response procedures for certificate compromise, expiration, or misuse events | **User** - Create runbooks for certificate incident response including revocation, reissuance, and service restoration procedures |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | ACM certificates should use DNS validation | AWS-ACM.1 | Use DNS validation method for ACM certificates to ensure automated renewal and reduce manual intervention | **IaC Template** - Configure ACM certificate requests with DNS validation method and automated Route 53 record creation |  |
| **High** | ACM certificates should be renewed before expiration | AWS-ACM.2 | Ensure ACM certificates are renewed automatically and monitor for renewal failures | **Platform** - Enable automatic renewal for ACM certificates and configure CloudWatch alarms for renewal failures |  |
| Medium | ACM certificates should not use wildcard domains unnecessarily | AWS-ACM.3 | Avoid using wildcard certificates when specific domain certificates provide better security | **User** - Review certificate usage and replace wildcard certificates with specific domain certificates where appropriate |  |
| Low | ACM certificates should be integrated with AWS services | AWS-ACM.4 | Use ACM certificates with AWS services rather than uploading external certificates | **IaC Template** - Configure AWS services to use ACM certificates directly rather than importing external certificates |  |
| Medium | ACM certificates should have appropriate subject alternative names | AWS-ACM.5 | Include all necessary domain names and subdomains in certificate subject alternative names | **IaC Template** - Configure ACM certificates with comprehensive SAN lists covering all required domains and subdomains |  |
| Low | ACM certificates should be monitored for transparency log entries | AWS-ACM.6 | Monitor certificate transparency logs for unauthorized certificate issuance | **Platform** - Implement certificate transparency monitoring using third-party tools or custom monitoring solutions |  |
| **High** | ACM private certificates should use appropriate key algorithms | AWS-ACM.7 | Use strong key algorithms and appropriate key sizes for ACM private certificates | **IaC Template** - Configure ACM Private CA with RSA 2048-bit or ECDSA P-256 key algorithms for certificate issuance |  |
| **High** | IAM policies should not allow full administrative privileges | AWS-IAM.1 | Avoid granting broad administrative permissions for ACM certificate management | **User** - Create specific IAM policies for ACM operations with resource-level permissions and conditions |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | AWS-CloudTrail.1 | Enable CloudTrail to audit ACM API calls across all regions | **Platform** - Configure multi-region CloudTrail with S3 bucket logging for ACM API activities and certificate lifecycle events |  |
| Medium | CloudWatch alarms should be configured for ACM certificate metrics | AWS-CloudWatch.1 | Set up monitoring and alerting for ACM certificate expiration and renewal failures | **IaC Template** - Configure CloudWatch alarms for certificate expiration warnings and renewal failure notifications |  |


## Operational Controls
---



## Cost Controls
---


### AWS Certificate Manager Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Use ACM public certificates instead of third-party certificates | COST-1 | Leverage free ACM public certificates instead of purchasing certificates from external CAs | **IaC Template** - Replace third-party certificates with ACM public certificates for AWS-integrated services |
| Medium | Optimize certificate coverage with appropriate SAN usage | COST-2 | Use subject alternative names efficiently to reduce the number of certificates needed | **User** - Consolidate multiple single-domain certificates into multi-domain certificates where security permits |
| **High** | Remove unused and expired certificates | COST-3 | Regularly audit and clean up unused ACM certificates to reduce management overhead | **User** - Implement regular certificate audits and automated cleanup of unused or expired certificates |
| **High** | Optimize ACM Private CA usage | COST-4 | Right-size ACM Private CA deployments and avoid unnecessary certificate authority instances | **User** - Evaluate Private CA requirements and consolidate certificate authorities where possible to reduce monthly costs |
| Medium | Use wildcard certificates judiciously | COST-5 | Balance security concerns with cost efficiency when deciding between wildcard and specific certificates | **User** - Evaluate the trade-offs between wildcard certificate convenience and security for cost-effective certificate strategy |
| Medium | Leverage automated certificate deployment | COST-6 | Use automated certificate deployment to reduce operational overhead and manual processes | **IaC Template** - Implement infrastructure as code for certificate provisioning and deployment automation |
| Low | Monitor certificate utilization across regions | COST-7 | Track certificate usage across AWS regions to identify optimization opportunities | **Platform** - Use AWS Config and custom scripts to monitor certificate deployment and utilization across regions |
| Medium | Implement proper tagging for certificate cost allocation | COST-8 | Tag ACM certificates and Private CAs for accurate cost tracking and allocation | **IaC Template** - Apply consistent tagging strategy including cost center, environment, and project tags to certificate resources |
| Low | Optimize certificate renewal processes | COST-9 | Ensure efficient certificate renewal processes to minimize operational overhead | **Platform** - Configure automatic renewal for ACM certificates and monitor renewal success rates |
| Medium | Evaluate certificate authority hierarchy efficiency | COST-10 | Design efficient certificate authority hierarchies for ACM Private CA to minimize costs | **User** - Review and optimize CA hierarchy design to balance security requirements with operational costs |
| Low | Use appropriate certificate validity periods | COST-11 | Configure certificate validity periods to balance security and operational efficiency | **IaC Template** - Set appropriate certificate validity periods for Private CA certificates based on security and operational requirements |
| Low | Monitor and optimize certificate API usage | COST-12 | Track ACM API usage patterns to identify potential cost optimization opportunities | **Platform** - Monitor ACM API call patterns and optimize certificate management workflows to reduce unnecessary API calls |


# AWS Route53
---


### CIS Controls v8.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Maintain Inventory of Authorized DNS Resources | CIS-1.1 | Maintain an accurate inventory of all Route 53 hosted zones, domains, and DNS records | **Platform** - Use AWS Config to track Route 53 resources and implement automated discovery of hosted zones and records |  |
| **High** | Configure DNS Filtering Services | CIS-3.3 | Implement DNS filtering to block malicious domains and prevent DNS-based attacks | **IaC Template** - Configure Route 53 Resolver DNS Firewall with domain filtering rules and threat intelligence feeds |  |
| **High** | Establish Secure Configuration Process | CIS-4.1 | Define and enforce secure baseline configurations for Route 53 hosted zones and records | **IaC Template** - Use CloudFormation or Terraform templates with standardized Route 53 configurations including DNSSEC and security policies |  |
| Medium | Maintain Audit Logs | CIS-6.1 | Enable comprehensive logging for all Route 53 DNS queries and configuration changes | **Platform** - Enable Route 53 query logging to CloudWatch and configure CloudTrail for API activity logging |  |
| Medium | Activate Audit Logging | CIS-6.2 | Ensure audit logging captures all Route 53 administrative actions and DNS resolution activity | **Platform** - Configure CloudTrail with Route 53 API logging and enable query logging for all hosted zones |  |
| Medium | Ensure Network Infrastructure is Up-to-Date | CIS-12.1 | Keep Route 53 configurations current with security patches and feature updates | **User** - Regularly review and update Route 53 configurations, health checks, and security features |  |

### NIST Cybersecurity Framework v1.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Physical devices and systems are inventoried | NIST-ID.AM-1 | Maintain comprehensive inventory of all Route 53 hosted zones and DNS infrastructure | **Platform** - Use AWS Config Rules to monitor Route 53 resource inventory and compliance status |  |
| Medium | Organizational communication and data flows are mapped | NIST-ID.AM-3 | Document DNS resolution flows and dependencies for Route 53 hosted zones | **User** - Create network diagrams showing DNS resolution paths and document external dependencies |  |
| **High** | Identities and credentials are managed for authorized devices and users | NIST-PR.AC-1 | Implement proper IAM roles and policies for Route 53 access and management | **IaC Template** - Define least-privilege IAM policies for Route 53 operations with resource-level permissions |  |
| **High** | Access permissions are managed, incorporating the principles of least privilege | NIST-PR.AC-4 | Apply least privilege access to Route 53 hosted zones and DNS records | **User** - Create granular IAM policies limiting Route 53 actions to specific hosted zones and record types |  |
| **High** | Data-at-rest is protected | NIST-PR.DS-1 | Protect DNS configuration data and zone files with appropriate encryption | **Platform** - Enable DNSSEC for hosted zones and use AWS KMS for encrypting Route 53 query logs |  |
| **High** | Data-in-transit is protected | NIST-PR.DS-2 | Ensure DNS queries and responses are protected during transmission | **IaC Template** - Configure DNS over HTTPS (DoH) and DNS over TLS (DoT) where supported, implement DNSSEC validation |  |
| Medium | Configuration change control processes are in place | NIST-PR.IP-3 | Implement change control processes for Route 53 DNS record modifications | **User** - Use infrastructure as code and approval workflows for DNS record changes with version control |  |
| Medium | The network is monitored to detect potential cybersecurity events | NIST-DE.CM-1 | Monitor Route 53 DNS queries and responses for malicious activity | **Platform** - Enable Route 53 query logging and integrate with security monitoring tools for threat detection |  |
| Medium | Event data are aggregated and correlated from multiple sources | NIST-DE.AE-3 | Correlate Route 53 logs with other security events for comprehensive threat detection | **Platform** - Integrate Route 53 logs with AWS Security Hub and SIEM solutions for correlation analysis |  |
| Medium | Response plan is executed during or after an event | NIST-RS.RP-1 | Define incident response procedures for DNS-related security events | **User** - Create runbooks for DNS poisoning, DDoS, and other DNS-related security incidents |  |

### AWS Foundational Security Best Practices v1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Route 53 public hosted zones should log DNS queries | AWS-Route53.1 | Enable query logging for Route 53 hosted zones to monitor DNS resolution activity | **IaC Template** - Configure query logging for all public hosted zones with CloudWatch Logs as destination |  |
| **High** | Route 53 public hosted zones should have DNSSEC enabled | AWS-Route53.2 | Enable DNSSEC signing for public hosted zones to prevent DNS spoofing attacks | **IaC Template** - Enable DNSSEC signing for hosted zones and configure key-signing key (KSK) management |  |
| Medium | Route 53 health checks should be configured for critical resources | AWS-Route53.3 | Implement health checks for critical endpoints to ensure DNS failover functionality | **IaC Template** - Configure Route 53 health checks with appropriate failure thresholds and notification settings |  |
| **High** | Route 53 Resolver should use DNS Firewall | AWS-Route53.4 | Configure Route 53 Resolver DNS Firewall to block malicious domains | **IaC Template** - Implement DNS Firewall rules with domain filtering lists and custom block/allow rules |  |
| **High** | Route 53 hosted zones should not be publicly writable | AWS-Route53.5 | Restrict write access to hosted zones to authorized users and services only | **User** - Review and restrict IAM policies to prevent unauthorized modifications to DNS records |  |
| Low | Route 53 domain registration should have privacy protection enabled | AWS-Route53.6 | Enable domain privacy protection to prevent exposure of registrant information | **IaC Template** - Configure domain registration with privacy protection enabled in Route 53 domain settings |  |
| Medium | Route 53 hosted zones should have MX records configured securely | AWS-Route53.7 | Ensure MX records point to secure mail servers and include appropriate SPF/DKIM records | **IaC Template** - Configure MX records with proper priorities and implement SPF, DKIM, and DMARC records |  |
| **High** | IAM policies should not allow full administrative privileges | AWS-IAM.1 | Avoid granting broad administrative permissions for Route 53 operations | **User** - Create specific IAM policies for Route 53 with resource-level permissions and conditions |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | AWS-CloudTrail.1 | Enable CloudTrail to audit Route 53 API calls across all regions | **Platform** - Configure multi-region CloudTrail with S3 bucket logging for Route 53 API activities |  |
| Medium | CloudWatch alarms should be configured for Route 53 metrics | AWS-CloudWatch.1 | Set up monitoring and alerting for Route 53 health checks and query metrics | **IaC Template** - Configure CloudWatch alarms for health check failures, query volumes, and resolver errors |  |


## Operational Controls
---



## Cost Controls
---


### AWS Route 53 Cost Optimization Best Practices 2024
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize hosted zone usage | COST-1 | Consolidate domains into fewer hosted zones where appropriate to reduce monthly charges | **User** - Review domain architecture and consolidate subdomains within single hosted zones when possible |
| **High** | Optimize query volume costs | COST-2 | Implement DNS caching and optimize TTL values to reduce billable query volumes | **IaC Template** - Configure appropriate TTL values for DNS records and implement client-side DNS caching strategies |
| Medium | Use geolocation and latency-based routing efficiently | COST-3 | Optimize routing policies to balance performance with cost for complex routing scenarios | **IaC Template** - Review and optimize routing policies to use simpler routing types where complex routing is unnecessary |
| Medium | Optimize health check frequency and locations | COST-4 | Configure health checks with appropriate intervals and geographic distribution to minimize costs | **IaC Template** - Set health check intervals based on actual availability requirements and limit checking locations |
| **High** | Remove unused DNS records and hosted zones | COST-5 | Regularly audit and clean up unused DNS resources to avoid unnecessary charges | **User** - Implement regular audits of DNS records and hosted zones, removing unused resources |
| Medium | Optimize DNS Firewall rule usage | COST-6 | Use DNS Firewall rules efficiently to minimize processing costs while maintaining security | **IaC Template** - Optimize DNS Firewall rule ordering and use efficient pattern matching to reduce processing overhead |
| Medium | Monitor and optimize query logging costs | COST-7 | Balance security monitoring needs with query logging costs by selective logging | **Platform** - Enable query logging selectively for critical zones and implement log retention policies |
| Low | Use Private DNS efficiently | COST-8 | Optimize private hosted zone usage for VPC DNS resolution to avoid unnecessary charges | **IaC Template** - Share private hosted zones across VPCs where appropriate and avoid duplicate private zones |
| Medium | Implement proper tagging for cost allocation | COST-9 | Tag Route 53 resources for accurate cost tracking and chargeback allocation | **IaC Template** - Apply consistent tagging strategy including cost center, environment, and project tags to hosted zones |
| Low | Monitor domain registration renewal costs | COST-10 | Track domain registration and renewal costs, considering alternative registrars when appropriate | **User** - Set up monitoring for domain renewal dates and compare costs with alternative registration options |
| Medium | Optimize resolver endpoint usage | COST-11 | Right-size Route 53 Resolver endpoints and minimize unnecessary endpoint deployments | **IaC Template** - Deploy resolver endpoints only where required and share endpoints across multiple VPCs when possible |
| Low | Use aliases instead of CNAME records where possible | COST-12 | Use Route 53 alias records for AWS resources to avoid additional query charges | **IaC Template** - Configure alias records for AWS resources like ELBs and CloudFront distributions instead of CNAME records |


