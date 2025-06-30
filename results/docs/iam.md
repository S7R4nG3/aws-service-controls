# AWS Identity and Access Management (IAM)
---


### CSA Cloud Controls Matrix v5.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Identity and Access Management Policy | IAM-01 | Establish comprehensive identity and access management policies with role-based access controls | **IaC** - Define IAM policies, roles, and groups using CloudFormation or Terraform templates with least privilege principles |  |
| **High** | Multi-Factor Authentication | IAM-02 | Enforce multi-factor authentication for all privileged accounts and sensitive operations | **User** - Configure MFA devices for root and IAM users through AWS Console or CLI, enforce via IAM policies |  |
| **High** | Privileged Account Management | IAM-03 | Implement strict controls and monitoring for privileged administrative accounts | **Platform** - Use AWS Organizations SCPs and IAM Access Analyzer to monitor and restrict privileged access |  |
| Medium | Access Reviews and Certification | IAM-04 | Conduct regular access reviews and certifications to ensure appropriate permissions | **Platform** - Implement automated access reviews using AWS Access Analyzer and IAM credential reports |  |
| Medium | Identity Federation | IAM-05 | Implement federated identity management for centralized authentication | **IaC** - Configure SAML 2.0 or OIDC identity providers and IAM roles for federated access |  |

### NIST SP 800-53 Rev 5
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | Account Management | AC-2 | Implement comprehensive account management procedures including creation, modification, and deletion | **Platform** - Use AWS Organizations and CloudTrail to track account lifecycle events and implement automated account provisioning |  |
| **High** | Access Enforcement | AC-3 | Enforce approved authorizations for logical access to information and system resources | **IaC** - Implement IAM policies with explicit deny statements and service control policies to enforce access boundaries |  |
| **High** | Least Privilege | AC-6 | Employ the principle of least privilege for specific security functions and privileged accounts | **IaC** - Create granular IAM policies with minimal required permissions and use IAM Access Analyzer for policy validation |  |
| **High** | Identification and Authentication | IA-2 | Uniquely identify and authenticate organizational users and associate with accounts | **User** - Implement strong password policies and MFA requirements through IAM password policy and MFA enforcement |  |
| Medium | Event Logging | AU-2 | Identify the types of events that the system is capable of logging | **Platform** - Enable CloudTrail logging for all IAM API calls and integrate with CloudWatch for monitoring |  |
| Medium | Session Authenticity | SC-23 | Protect the authenticity of communications sessions | **Platform** - Use AWS STS temporary credentials and implement session duration limits in IAM roles |  |

### AWS Foundational Security Best Practices Standard 1.0.0
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | IAM policies should not allow full '*' administrative privileges | IAM.1 | Avoid granting full administrative privileges using wildcard permissions | **IaC** - Review and remediate IAM policies to remove wildcard permissions and implement specific resource-based permissions |  |
| **High** | IAM users should not have IAM policies attached | IAM.2 | Avoid attaching policies directly to IAM users, use groups or roles instead | **IaC** - Migrate user-attached policies to IAM groups and implement role-based access patterns |  |
| **High** | IAM users' access keys should be rotated every 90 days or less | IAM.3 | Regularly rotate IAM user access keys to reduce security risks | **User** - Implement automated access key rotation using AWS Secrets Manager or custom Lambda functions |  |
| **Critical** | IAM root access key should not exist | IAM.4 | Remove and do not create access keys for the root account | **User** - Delete existing root access keys and implement IAM users with appropriate permissions for administrative tasks |  |
| **High** | MFA should be enabled for all IAM users that have a console password | IAM.5 | Enable multi-factor authentication for console access to enhance security | **User** - Configure virtual or hardware MFA devices for all console users and enforce through IAM policies |  |
| **Critical** | Hardware MFA should be enabled for the root user | IAM.6 | Enable hardware-based MFA for the root account for maximum security | **User** - Configure hardware MFA device for root account through AWS Console account settings |  |
| Medium | Password policies for IAM users should have strong configurations | IAM.7 | Implement strong password policies including length, complexity, and rotation requirements | **IaC** - Configure IAM account password policy with minimum length, complexity requirements, and rotation policies |  |
| Medium | Unused IAM user credentials should be removed | IAM.8 | Identify and remove unused IAM user credentials to reduce attack surface | **Platform** - Use IAM credential reports and Access Analyzer to identify and remove unused credentials automatically |  |

### AWS Security Hub 2023.04.04
| Severity | Title | Identifier | Description | Implementation | Code |
| - | - | - | - | - | - |
| **High** | IAM customer managed policies should not allow decryption actions on all KMS keys | IAM.21 | Restrict IAM policies from allowing broad KMS decryption permissions | **IaC** - Review and update IAM policies to specify explicit KMS key ARNs instead of wildcard permissions for KMS actions |  |
| Medium | IAM user credentials unused for 45 days should be removed | IAM.22 | Remove IAM user credentials that have not been used for 45 days or more | **Platform** - Implement automated credential cleanup using Lambda functions triggered by CloudWatch Events based on credential reports |  |
| Medium | Security contact information should be provided for an AWS account | Account.1 | Ensure security contact information is configured for incident response | **User** - Configure security contact information in AWS account settings through the AWS Console |  |
| Medium | AWS accounts should be part of an AWS Organizations organization | Account.2 | Use AWS Organizations for centralized account management and governance | **Platform** - Create or join an AWS Organizations structure and implement service control policies for governance |  |


## Operational Controls
---



## Cost Controls
---


### AWS IAM Cost Optimization Best Practices 2023
| Severity | Title | Identifier | Description | Implementation |
| - | - | - | - | - |
| **High** | Eliminate Unused IAM Users and Roles | COST-IAM-01 | Regularly identify and remove unused IAM users, roles, and policies to reduce management overhead | **Platform** - Use IAM Access Analyzer and credential reports to identify unused identities and implement automated cleanup processes |
| Medium | Optimize Cross-Account Role Usage | COST-IAM-02 | Use cross-account IAM roles instead of creating duplicate users across multiple accounts | **IaC** - Implement cross-account trust relationships and assume role patterns to reduce user proliferation |
| Medium | Leverage AWS SSO for User Management | COST-IAM-03 | Use AWS IAM Identity Center (SSO) to reduce IAM user management overhead and costs | **Platform** - Migrate from individual IAM users to AWS SSO for centralized identity management and reduced operational costs |
| Medium | Implement Policy Consolidation | COST-IAM-04 | Consolidate similar IAM policies to reduce management complexity and overhead | **IaC** - Review and merge similar IAM policies using IAM Access Analyzer policy validation and optimization recommendations |
| Low | Use Service-Linked Roles | COST-IAM-05 | Utilize AWS service-linked roles instead of custom roles where possible to reduce management overhead | **Platform** - Replace custom service roles with AWS-managed service-linked roles when supported by the service |
| Medium | Implement Automated Access Reviews | COST-IAM-06 | Automate access reviews to reduce manual effort and operational costs | **Platform** - Use AWS Config rules and Lambda functions to automate access certification and policy compliance checks |


