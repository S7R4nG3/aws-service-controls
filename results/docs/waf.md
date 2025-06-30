# AWS Web Application Firewall (WAF)
---


### CSA Cloud Controls Matrix v4.0.10
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management Policy and Procedures | IAM-01 | Establish and maintain identity and access management policies for WAF resources | **IaC** - Implement IAM policies and roles with least privilege access to WAF resources using CloudFormation or Terraform templates |  |
| **High** | User Access Provisioning | IAM-02 | Ensure proper user access provisioning and deprovisioning for WAF management | **Platform** - Configure IAM users, groups, and roles with appropriate permissions for WAF administration through AWS IAM console |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Protect data in transit and at rest for WAF configurations and logs | **IaC** - Enable encryption for WAF logs stored in S3, CloudWatch, and Kinesis Data Firehose destinations |  |
| Medium | Secure Transmission | STA-01 | Ensure secure transmission of WAF configurations and logs | **Platform** - Use TLS 1.2 or higher for all API communications and enable HTTPS for management interfaces |  |
| Medium | Information System Documentation | IVS-01 | Maintain documentation of WAF rules, configurations, and security policies | **User** - Document WAF rule sets, IP allow/deny lists, and security group configurations in version-controlled repositories |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Account Management | AC-2 | Manage accounts with access to WAF resources and configurations | **Platform** - Implement account lifecycle management for WAF administrators using AWS IAM with regular access reviews |  |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to WAF resources | **IaC** - Deploy resource-based policies and IAM policies that enforce least privilege access to WAF APIs and console |  |
| **High** | System Monitoring | SI-4 | Monitor WAF for security events and potential threats | **IaC** - Configure CloudWatch metrics, alarms, and logging for WAF to detect and alert on security events |  |
| **High** | Event Logging | AU-2 | Ensure comprehensive logging of WAF events and rule matches | **IaC** - Enable WAF logging to CloudWatch Logs, S3, or Kinesis Data Firehose with appropriate retention policies |  |
| Medium | Boundary Protection | SC-7 | Implement boundary protection through WAF rules and policies | **IaC** - Deploy managed rule groups and custom rules to protect against OWASP Top 10 and other web application threats |  |
| Medium | System Component Inventory | CM-8 | Maintain inventory of WAF components and configurations | **User** - Use AWS Config to track WAF resources and maintain inventory of Web ACLs, rule groups, and IP sets |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | AWS WAF Classic Global Web ACL logging should be enabled | WAF.1 | Enable logging for AWS WAF Classic Global Web ACLs to monitor and analyze web requests | **IaC** - Configure logging destination for WAF Classic Web ACLs using CloudFormation or Terraform |  |
| Medium | AWS WAF Classic Regional Web ACL logging should be enabled | WAF.2 | Enable logging for AWS WAF Classic Regional Web ACLs | **IaC** - Enable regional WAF logging through infrastructure templates with appropriate log destinations |  |
| Medium | AWS WAF Classic Global Web ACL should have at least one rule or rule group | WAF.3 | Ensure WAF Classic Global Web ACLs contain protective rules | **IaC** - Deploy Web ACLs with managed rule groups or custom rules for application protection |  |
| Medium | AWS WAF Classic Regional Web ACL should have at least one rule or rule group | WAF.4 | Ensure WAF Classic Regional Web ACLs contain protective rules | **IaC** - Implement regional Web ACLs with appropriate rule sets for regional application protection |  |
| Medium | AWS WAF Global Web ACL logging should be enabled | WAF.6 | Enable logging for AWS WAF v2 Global Web ACLs | **IaC** - Configure logging for WAF v2 Global Web ACLs with CloudWatch Logs, S3, or Kinesis Data Firehose |  |
| Medium | AWS WAF Regional Web ACL logging should be enabled | WAF.7 | Enable logging for AWS WAF v2 Regional Web ACLs | **IaC** - Enable comprehensive logging for regional WAF v2 Web ACLs through infrastructure automation |  |
| Medium | AWS WAF Global Web ACL should have at least one rule or rule group | WAF.8 | Ensure WAF v2 Global Web ACLs contain protective rules | **IaC** - Deploy managed rule groups such as Core Rule Set, Known Bad Inputs, and application-specific rules |  |
| Medium | AWS WAF Regional Web ACL should have at least one rule or rule group | WAF.10 | Ensure WAF v2 Regional Web ACLs contain protective rules | **IaC** - Implement regional Web ACLs with managed and custom rule groups for comprehensive protection |  |

### AWS Security Hub 2023.1.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | WAF Classic logging should be enabled | WAF.1 | Enable comprehensive logging for WAF Classic to support security monitoring and incident response | **IaC** - Configure WAF Classic logging destinations and ensure log retention meets compliance requirements |  |
| Medium | WAF v2 logging should be enabled | WAF.2 | Enable logging for WAF v2 Web ACLs to capture detailed request information | **IaC** - Deploy WAF v2 logging configuration with appropriate filtering and sampling rates |  |
| Medium | WAF Web ACLs should have protective rules | WAF.3 | Ensure Web ACLs implement adequate protection through rules and rule groups | **IaC** - Implement a combination of AWS managed rules and custom rules tailored to application requirements |  |


## Operational Controls
---



## Cost Controls
---


### AWS WAF Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Optimize Web ACL Rule Evaluation | COST-WAF-01 | Minimize Web Request Units (WRU) consumption by optimizing rule order and complexity | **IaC** - Order rules by likelihood of match and use efficient rule patterns to reduce WRU consumption |
| **High** | Right-size Managed Rule Groups | COST-WAF-02 | Select only necessary managed rule groups to avoid unnecessary WRU charges | **User** - Regularly review and disable unused managed rule groups based on application requirements and threat landscape |
| Medium | Implement Intelligent Request Sampling | COST-WAF-03 | Configure appropriate logging sampling rates to reduce storage costs while maintaining visibility | **IaC** - Configure sampling rates based on traffic volume and compliance requirements, typically 1-10% for high-volume applications |
| Medium | Optimize Log Storage and Retention | COST-WAF-04 | Implement lifecycle policies for WAF logs to manage storage costs effectively | **IaC** - Configure S3 lifecycle policies to transition logs to IA and Glacier storage classes, and set appropriate retention periods |
| Medium | Monitor and Alert on WAF Costs | COST-WAF-05 | Implement cost monitoring and alerting for WAF resource consumption | **Platform** - Set up CloudWatch billing alarms and AWS Cost Explorer reports to track WAF WRU consumption and associated costs |
| Low | Regularly Review Rule Effectiveness | COST-WAF-06 | Conduct periodic reviews of rule effectiveness to remove unused or ineffective rules | **User** - Use WAF metrics and logs to identify rules with low match rates and evaluate their continued necessity |
| Low | Consolidate Web ACLs Where Appropriate | COST-WAF-07 | Consolidate multiple Web ACLs with similar rule sets to reduce management overhead and costs | **IaC** - Evaluate opportunities to merge Web ACLs serving similar applications or environments while maintaining security boundaries |


