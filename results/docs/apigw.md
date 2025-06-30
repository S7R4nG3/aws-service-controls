# AWS API Gateway
---


### Cloud Security Alliance (CSA) Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management Policy and Procedures | IAM-01 | Establish and maintain policies and procedures for identity and access management for API Gateway resources | **IaC** - Implement IAM policies in CloudFormation/Terraform templates that define least privilege access to API Gateway resources, including resource-based policies and execution roles |  |
| **High** | Multi-Factor Authentication | IAM-05 | Require multi-factor authentication for administrative access to API Gateway console and APIs | **Platform** - Configure AWS IAM to require MFA for users accessing API Gateway management console and administrative APIs |  |
| **High** | Data Security and Information Lifecycle Management | DSI-01 | Implement data classification and protection mechanisms for data passing through API Gateway | **IaC** - Configure request/response data transformation and validation rules in API Gateway to sanitize and protect sensitive data |  |
| **High** | Encryption Key Management | EKM-01 | Implement proper encryption key management for API Gateway TLS certificates and custom domain names | **IaC** - Use AWS Certificate Manager or AWS KMS to manage TLS certificates and encryption keys for custom domains and client certificates |  |
| Medium | Network Security | IVS-04 | Implement network-level security controls for API Gateway endpoints | **IaC** - Configure VPC endpoints for private API access and implement resource policies to restrict access based on IP ranges or VPC endpoints |  |
| Medium | Security Event Logging | SEF-01 | Enable comprehensive logging of API Gateway security events and access patterns | **IaC** - Configure CloudTrail, CloudWatch Logs, and AWS Config to capture API Gateway management events, execution logs, and access logs |  |

### NIST 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to API Gateway resources and APIs | **IaC** - Implement API Gateway authorizers (Lambda, Cognito, or IAM) and resource policies to enforce granular access control |  |
| **High** | Identification and Authentication | IA-2 | Uniquely identify and authenticate organizational users accessing API Gateway | **User** - Configure appropriate authentication mechanisms such as API keys, AWS IAM, Amazon Cognito, or custom authorizers based on use case |  |
| **High** | Transmission Confidentiality and Integrity | SC-8 | Protect the confidentiality and integrity of transmitted information through API Gateway | **IaC** - Enforce HTTPS/TLS 1.2+ for all API Gateway endpoints and configure minimum TLS version and cipher suites |  |
| **High** | Event Logging | AU-2 | Ensure API Gateway generates audit records for security-relevant events | **IaC** - Enable CloudTrail for API Gateway management events and configure execution logging and access logging for runtime events |  |
| Medium | Boundary Protection | SC-7 | Monitor and control communications at the external boundary of API Gateway | **IaC** - Implement resource policies, WAF integration, and VPC endpoints to control network access to API Gateway |  |
| Medium | Information System Monitoring | SI-4 | Monitor API Gateway for attacks and indicators of potential attacks | **Platform** - Configure CloudWatch metrics, alarms, and AWS X-Ray tracing to monitor API performance, errors, and suspicious activity |  |

### AWS Foundational Security Best Practices 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | API Gateway REST API execution logging should be enabled | APIGateway.1 | Ensure execution logging is enabled for API Gateway REST APIs to capture detailed information about requests and responses | **IaC** - Configure stage-level execution logging in API Gateway deployment templates with appropriate log level (INFO or ERROR) |  |
| Medium | API Gateway REST API stages should be configured to use SSL certificates for backend authentication | APIGateway.2 | Configure SSL certificates for backend authentication to ensure secure communication between API Gateway and backend services | **IaC** - Configure client certificates in API Gateway stage settings and ensure backend services validate these certificates |  |
| Low | API Gateway REST API stages should have AWS X-Ray tracing enabled | APIGateway.3 | Enable X-Ray tracing for API Gateway stages to provide detailed tracing information for debugging and monitoring | **IaC** - Enable X-Ray tracing in API Gateway stage configuration and ensure proper IAM permissions for X-Ray service |  |
| Medium | API Gateway should be associated with an AWS WAF | APIGateway.4 | Associate API Gateway with AWS WAF to protect against common web exploits and attacks | **IaC** - Create and associate AWS WAF WebACL with API Gateway stage, implementing rules for common attack patterns and rate limiting |  |
| Medium | API Gateway REST API cache data should be encrypted at rest | APIGateway.5 | Enable encryption at rest for API Gateway cache data to protect cached responses | **IaC** - Configure cache encryption in API Gateway method settings when caching is enabled |  |

### AWS Security Hub 2023.1
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| Medium | API Gateway routes should specify an authorization type | APIGateway.8 | Ensure all API Gateway routes have an appropriate authorization type configured to prevent unauthorized access | **IaC** - Configure authorization type (AWS_IAM, COGNITO_USER_POOLS, CUSTOM, or JWT) for all API Gateway routes and methods |  |
| Medium | Access logging should be configured for API Gateway V2 Stages | APIGateway.9 | Configure access logging for API Gateway V2 stages to capture client access patterns and troubleshoot issues | **IaC** - Configure access logging settings in API Gateway V2 stage with CloudWatch Logs destination and appropriate log format |  |
| **High** | CloudTrail should be enabled and configured with at least one multi-Region trail | CloudTrail.1 | Enable CloudTrail to capture API Gateway management events across all regions | **Platform** - Configure CloudTrail with multi-region trail to capture API Gateway management API calls and configuration changes |  |
| Medium | AWS Config should be enabled | Config.1 | Enable AWS Config to track configuration changes to API Gateway resources | **Platform** - Configure AWS Config to monitor API Gateway resource configurations and detect configuration drift |  |


## Operational Controls
---



## Cost Controls
---


### AWS API Gateway Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Implement Response Caching | COST-01 | Enable response caching to reduce backend calls and API Gateway request charges for frequently accessed data | **IaC** - Configure caching at the stage or method level with appropriate TTL values and cache key parameters to optimize cache hit ratios |
| **High** | Optimize API Gateway Type Selection | COST-02 | Choose the appropriate API Gateway type (REST vs HTTP vs WebSocket) based on feature requirements and cost considerations | **User** - Evaluate using HTTP APIs instead of REST APIs when advanced features are not needed, as HTTP APIs cost up to 70% less |
| Medium | Implement Request Throttling and Rate Limiting | COST-03 | Configure throttling and rate limiting to control API usage and prevent unexpected cost spikes | **IaC** - Set up usage plans with throttling limits and quotas to control API consumption and implement burst limits |
| Medium | Optimize Data Transfer | COST-04 | Minimize data transfer costs by implementing response compression and optimizing payload sizes | **IaC** - Enable content encoding (gzip) and implement request/response transformation to reduce payload sizes |
| Medium | Monitor and Alert on Usage Patterns | COST-05 | Implement monitoring and alerting for API usage to identify cost optimization opportunities and usage anomalies | **Platform** - Configure CloudWatch metrics and billing alerts to monitor API Gateway usage, costs, and set up notifications for unusual patterns |
| Low | Implement Regional Optimization | COST-06 | Deploy API Gateway in regions closest to users and integrate with CloudFront for global optimization | **IaC** - Use regional API endpoints and configure CloudFront distribution for caching and global edge optimization |
| Low | Optimize Custom Domain and Certificate Usage | COST-07 | Consolidate APIs under fewer custom domains to reduce certificate and domain management costs | **User** - Use path-based routing to serve multiple APIs from a single custom domain rather than creating separate domains for each API |


