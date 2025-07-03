# AWS API Gateway
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | IAM-01 | CSA Cloud Controls Matrix v4.0 | Entitlement | Provision and manage user account entitlements to enterprise assets and resources | **IaC** - Configure IAM policies, resource-based policies, and API Gateway authorizers (Lambda, Cognito, or IAM) to control access to APIs |  |
| **High** | DSI-05 | CSA Cloud Controls Matrix v4.0 | Data Transmission | Encrypt data in transit using strong cryptographic protocols and algorithms | **IaC** - Configure custom domain names with SSL/TLS certificates and enforce HTTPS-only communication through API Gateway settings |  |
| **High** | IVS-06 | CSA Cloud Controls Matrix v4.0 | Network Architecture | Design and implement network security architecture | **IaC** - Deploy private API Gateway endpoints within VPC, configure VPC endpoints, and implement WAF rules for protection |  |
| Medium | STA-04 | CSA Cloud Controls Matrix v4.0 | Denial of Service Protection | Implement protective measures against denial of service attacks | **IaC** - Configure throttling limits, usage plans, and integrate with AWS WAF for advanced threat protection |  |
| Medium | GRM-02 | CSA Cloud Controls Matrix v4.0 | Information System Regulatory Mapping | Map applicable regulatory requirements to information systems and associated controls | **User** - Implement API versioning strategies, deployment stages, and approval workflows for API changes |  |
| **High** | AC-3 | NIST SP 800-53 Rev 5 | Access Enforcement | Enforce approved authorizations for logical access to information and system resources | **IaC** - Implement API Gateway authorizers with fine-grained access control using IAM roles, Lambda authorizers, or Cognito User Pools |  |
| **High** | SC-8 | NIST SP 800-53 Rev 5 | Transmission Confidentiality and Integrity | Protect the confidentiality and integrity of transmitted information | **IaC** - Configure TLS 1.2+ encryption, certificate management, and mutual TLS authentication for sensitive APIs |  |
| **High** | AU-2 | NIST SP 800-53 Rev 5 | Event Logging | Identify the types of events that the system is capable of logging | **IaC** - Enable CloudTrail logging, API Gateway execution logs, and access logs with comprehensive event capture |  |
| Medium | SI-3 | NIST SP 800-53 Rev 5 | Malicious Code Protection | Implement malicious code protection mechanisms | **IaC** - Deploy AWS WAF with API Gateway to filter malicious requests and implement request validation |  |
| Medium | CP-2 | NIST SP 800-53 Rev 5 | Contingency Plan | Develop, document, and maintain a contingency plan for information systems | **User** - Implement multi-region deployments, backup strategies, and disaster recovery procedures for API Gateway configurations |  |
| **High** | APIGateway.2 | AWS Foundational Security Best Practices v1.0.0 | API Gateway REST API stages should be configured to use SSL certificates for backend authentication | Configure SSL certificates for backend authentication to ensure secure communication | **IaC** - Deploy SSL certificates via ACM and configure custom domain names with HTTPS endpoints |  |
| Medium | APIGateway.1 | AWS Foundational Security Best Practices v1.0.0 | API Gateway REST API execution logging should be enabled | Enable execution logging for REST APIs to capture detailed information about requests and responses | **IaC** - Configure CloudWatch Logs integration with API Gateway stages to enable INFO or ERROR level execution logging |  |
| Medium | APIGateway.4 | AWS Foundational Security Best Practices v1.0.0 | API Gateway should be associated with a WAF Web ACL | Associate API Gateway with AWS WAF to protect against common web exploits | **IaC** - Create WAF web ACL with appropriate rules and associate it with API Gateway stages |  |
| Medium | APIGateway.5 | AWS Foundational Security Best Practices v1.0.0 | API Gateway REST API cache data should be encrypted at rest | Enable encryption for API Gateway caching to protect cached data at rest | **IaC** - Configure cache encryption settings in API Gateway method responses and stage configurations |  |
| Low | APIGateway.3 | AWS Foundational Security Best Practices v1.0.0 | API Gateway REST API stages should have AWS X-Ray tracing enabled | Enable AWS X-Ray tracing to analyze and debug distributed API applications | **IaC** - Enable X-Ray tracing at the API Gateway stage level and configure sampling rules |  |
| **High** | APIGateway.10 | AWS Security Hub 2023.04.0 | API Gateway REST APIs should prohibit public read access | Ensure REST APIs do not allow unrestricted public read access | **IaC** - Configure proper authorization and resource policies to prevent anonymous public access to sensitive API operations |  |
| **High** | APIGateway.11 | AWS Security Hub 2023.04.0 | API Gateway REST APIs should prohibit public write access | Ensure REST APIs do not allow unrestricted public write access | **IaC** - Implement strong authorization controls for write operations including POST, PUT, DELETE methods |  |
| Medium | APIGateway.8 | AWS Security Hub 2023.04.0 | API Gateway routes should specify an authorization type | Ensure all API Gateway routes have appropriate authorization configured | **IaC** - Configure authorization type (IAM, Lambda, Cognito, or JWT) for all API Gateway routes and methods |  |
| Medium | APIGateway.9 | AWS Security Hub 2023.04.0 | Access logging should be configured for API Gateway V2 Stages | Enable access logging for HTTP API and WebSocket API stages | **IaC** - Configure access logging with CloudWatch Logs destination and appropriate log format for API Gateway v2 stages |  |

## Operational Controls
---



## Cost Controls
---

| Severity | Identifier | Framework | Title | Description | Implementation | Code |
| - | - | - | - | - | - | - |
| **High** | COST-01 | AWS API Gateway Cost Optimization Best Practices 2023 | Implement Caching Strategy | Enable API Gateway caching to reduce backend calls and improve performance while reducing costs | **IaC** - Configure method-level caching with appropriate TTL values and cache key parameters to minimize redundant backend requests |  |
| **High** | COST-02 | AWS API Gateway Cost Optimization Best Practices 2023 | Optimize API Gateway Type Selection | Choose the appropriate API Gateway type (HTTP API vs REST API) based on feature requirements | **User** - Use HTTP APIs for simple proxy integrations as they are up to 70% cheaper than REST APIs with lower latency |  |
| Medium | COST-03 | AWS API Gateway Cost Optimization Best Practices 2023 | Implement Usage Plans and API Keys | Control API usage and prevent unexpected costs through usage plans and throttling | **IaC** - Configure usage plans with quotas and throttling limits to control API consumption and associated backend costs |  |
| Medium | COST-04 | AWS API Gateway Cost Optimization Best Practices 2023 | Monitor and Optimize Request Routing | Optimize request routing to minimize data transfer and processing costs | **User** - Use regional endpoints instead of edge-optimized when clients are in the same region, and implement request/response transformations efficiently |  |
| Medium | COST-05 | AWS API Gateway Cost Optimization Best Practices 2023 | Implement Efficient Error Handling | Optimize error handling to reduce unnecessary backend invocations and associated costs | **IaC** - Configure gateway responses and request validation to catch errors early and avoid backend processing costs |  |
| Low | COST-06 | AWS API Gateway Cost Optimization Best Practices 2023 | Regular Cost Monitoring and Analysis | Implement regular monitoring of API Gateway costs and usage patterns | **Platform** - Set up CloudWatch dashboards and billing alerts to monitor API Gateway usage, costs, and identify optimization opportunities |  |

