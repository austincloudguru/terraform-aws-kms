# Terraform AWS KMS Module
[![Terratest](https://github.com/austincloudguru/terraform-aws-kms/workflows/Terratest/badge.svg)](https://github.com/austincloudguru/terraform-aws-kms/actions?query=workflow%3ATerratest) 
![Latest Version](https://img.shields.io/github/v/tag/austincloudguru/terraform-aws-kms?sort=semver&label=Latest%20Version)
[![License](https://img.shields.io/github/license/austincloudguru/terraform-aws-kms)](https://github.com/austincloudguru/terraform-aws-kms/blob/master/LICENSE)

Terraform module that creates a KMS Key on AWS to use with any AWS Resource.  It will also create the KMS alias

## Usage
```hcl
module "kms" {
  source                  = "AustinCloudGuru/alb/aws/kms"
  # You should pin the module to a specific version
  # version                 = "x.x.x"
  key_alias               = "aws_config_bucket"
  description             = "KMS Key for the AWS Config Bucket"
  deletion_window_in_days = 7
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.8.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_kms_alias.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_alias) | resource |
| [aws_kms_key.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_customer_master_key_spec"></a> [customer\_master\_key\_spec](#input\_customer\_master\_key\_spec) | Specifies whether the key contains a symmetric key or an asymmetric key pair and the encryption algorithms or signing algorithms that the key supports. Valid values: `SYMMETRIC_DEFAULT`, `RSA_2048`, `RSA_3072`, `RSA_4096`, `ECC_NIST_P256`, `ECC_NIST_P384`, `ECC_NIST_P521`, or `ECC_SECG_P256K1`. | `string` | `"SYMMETRIC_DEFAULT"` | no |
| <a name="input_deletion_window_in_days"></a> [deletion\_window\_in\_days](#input\_deletion\_window\_in\_days) | The waiting period, specified in number of days. After the waiting period ends, AWS KMS deletes the KMS key. If you specify a value, it must be between 7 and 30, inclusive. If you do not specify a value, it defaults to 30. | `number` | `30` | no |
| <a name="input_description"></a> [description](#input\_description) | The description of the key as viewed in AWS console. | `string` | n/a | yes |
| <a name="input_enable_key_rotation"></a> [enable\_key\_rotation](#input\_enable\_key\_rotation) | Specifies whether key rotation is enabled. Defaults to false. | `bool` | `true` | no |
| <a name="input_key_alias"></a> [key\_alias](#input\_key\_alias) | The display name of the alias. | `string` | n/a | yes |
| <a name="input_key_usage"></a> [key\_usage](#input\_key\_usage) | Specifies the intended use of the key. Valid values: ENCRYPT\_DECRYPT or SIGN\_VERIFY. Defaults to ENCRYPT\_DECRYPT. | `string` | `"ENCRYPT_DECRYPT"` | no |
| <a name="input_multi_region"></a> [multi\_region](#input\_multi\_region) | Indicates whether the KMS key is a multi-Region (true) or regional (false) key. Defaults to false. | `bool` | `false` | no |
| <a name="input_policy"></a> [policy](#input\_policy) | A valid policy JSON document. Although this is a key policy, not an IAM policy, an aws\_iam\_policy\_document, in the form that designates a principal, can be used. | `string` | `null` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | A map of tags to add to all resources | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_alias_arn"></a> [alias\_arn](#output\_alias\_arn) | Alias ARN |
| <a name="output_alias_name"></a> [alias\_name](#output\_alias\_name) | Alias name |
| <a name="output_key_arn"></a> [key\_arn](#output\_key\_arn) | Key ARN |
| <a name="output_key_id"></a> [key\_id](#output\_key\_id) | Key ID |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
