# AWS Route53
---


### Cloud Security Alliance Cloud Controls Matrix v5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | User Access Authorization | IAM-02 | Ensure only authorized users have access to Route53 resources with appropriate permissions | **IaC** - Implement IAM policies with least privilege access, define specific Route53 actions and resources in IAM policies |  |
| **Critical** | Privileged User Management | IAM-03 | Control and monitor privileged access to Route53 administrative functions | **Platform** - Use AWS IAM roles with time-based access and MFA requirements for Route53 administrative operations |  |
| **High** | Data Classification | DSI-01 | Classify DNS records and hosted zone data based on sensitivity and business impact | **User** - Implement data classification tags on hosted zones and establish handling procedures for sensitive DNS data |  |
| **High** | Secure Disposal | DSI-07 | Ensure secure deletion of DNS records and hosted zones when no longer needed | **IaC** - Implement automated lifecycle policies and secure deletion procedures for Route53 resources |  |
| Medium | Network Security | IVS-13 | Implement network-level security controls for DNS resolution | **IaC** - Configure Route53 Resolver with DNS filtering and private hosted zones for internal resources |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **Critical** | Account Management | AC-2 | Manage Route53 access accounts with proper provisioning and de-provisioning procedures | **Platform** - Implement AWS IAM account lifecycle management with automated provisioning and regular access reviews |  |
| **Critical** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to Route53 resources | **IaC** - Deploy IAM policies with explicit allow/deny rules and resource-based permissions for Route53 |  |
| **High** | Event Logging | AU-2 | Log Route53 API calls and DNS query activities for security monitoring | **Platform** - Enable CloudTrail for Route53 API logging and configure Route53 Resolver query logging |  |
| **High** | Audit Record Review | AU-6 | Review and analyze Route53 audit records for security incidents and anomalies | **Platform** - Implement automated log analysis using CloudWatch and AWS Security Hub for Route53 events |  |
| **High** | Secure Name/Address Resolution Service | SC-20 | Provide secure DNS resolution services with integrity and authenticity validation | **IaC** - Enable DNSSEC for hosted zones and implement DNS filtering through Route53 Resolver |  |
| Medium | System Backup | CP-9 | Backup Route53 configurations and DNS records for disaster recovery | **IaC** - Implement automated backup of Route53 configurations using AWS Config and infrastructure as code |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Route53 public hosted zones should log DNS queries | Route53.1 | Enable query logging for Route53 public hosted zones to monitor DNS resolution requests | **IaC** - Configure Route53 query logging to CloudWatch Logs for all public hosted zones |  |
| Medium | Route53 public hosted zones should not have unnecessary SRV records | Route53.2 | Review and remove unnecessary SRV records that may expose internal services | **User** - Regularly audit DNS records and remove unused or unnecessary SRV records from public zones |  |

### AWS Security Hub 2023
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | Route53 public hosted zones should log DNS queries | Route53.1 | DNS query logging should be enabled for Route53 public hosted zones | **IaC** - Enable DNS query logging using CloudFormation or Terraform templates for all public hosted zones |  |
| Low | Route53 public hosted zones should not have unnecessary SRV records | Route53.2 | Remove SRV records that are not required to reduce attack surface | **User** - Implement regular DNS record audits and establish approval processes for SRV record creation |  |


## Operational Controls
---



## Cost Controls
---


### AWS Route53 Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize Hosted Zone Count | COST-001 | Consolidate hosted zones where possible to reduce monthly hosting fees | **User** - Review and consolidate related domains into fewer hosted zones, use subdomains instead of separate zones where appropriate |
| **High** | Monitor DNS Query Volume | COST-002 | Track and optimize DNS query patterns to manage query-based charges | **Platform** - Implement CloudWatch monitoring for DNS query metrics and set up billing alerts for unexpected query volume increases |
| Medium | Optimize Health Check Configuration | COST-003 | Right-size health check frequency and endpoints to balance availability with cost | **IaC** - Configure health checks with appropriate intervals and disable unnecessary health checks for non-critical endpoints |
| Medium | Use Route53 Resolver Efficiently | COST-004 | Optimize Route53 Resolver rules and endpoints to minimize forwarding costs | **IaC** - Consolidate resolver rules and endpoints across VPCs, remove unused resolver configurations |
| Low | Implement TTL Optimization | COST-005 | Set appropriate TTL values to reduce DNS query frequency while maintaining performance | **User** - Configure longer TTL values for static records and shorter TTLs only for records requiring frequent updates |
| Low | Regular Resource Cleanup | COST-006 | Remove unused DNS records and hosted zones to eliminate unnecessary charges | **User** - Implement regular audits of Route53 resources and establish processes for removing unused zones and records |


