resource "aws_kms_key" "this" {
  description              = var.description
  key_usage                = var.key_usage
  customer_master_key_spec = var.customer_master_key_spec
  deletion_window_in_days  = var.deletion_window_in_days
  enable_key_rotation      = var.enable_key_rotation
  multi_region             = var.multi_region
  policy                   = var.policy
  tags = merge(
    {
      "name"      = var.key_alias
      "terraform" = "true"
    },
    var.tags
  )
  lifecycle {
    ignore_changes = [tags, ]
  }
}

resource "aws_kms_alias" "this" {
  name          = join("/", ["alias", var.key_alias])
  target_key_id = aws_kms_key.this.key_id
}
