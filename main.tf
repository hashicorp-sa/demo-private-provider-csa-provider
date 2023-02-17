provider "github" {
  # token = var.token # or `GITHUB_TOKEN`
}

data "github_actions_public_key" "my_public_key" {
  repository = "demo-private-provider-csa-provider"
}


# https://developer.hashicorp.com/terraform/tutorials/providers/provider-release-publish?in=terraform%2Fproviders#add-github-secrets-for-github-action
resource "github_actions_secret" "gpg_private_key" {
  repository       = data.github_actions_public_key.my_public_key.repository
  secret_name      = "GPG_PRIVATE_KEY"
  plaintext_value  =file(var.gpg_private_key)
}

resource "github_actions_secret" "passphrase" {
  repository       = data.github_actions_public_key.my_public_key.repository
  secret_name      = "PASSPHRASE"
  plaintext_value  = var.passphrase
}

resource "github_actions_secret" "tf_url" {
  repository       = data.github_actions_public_key.my_public_key.repository
  secret_name      = "TF_URL"
  plaintext_value  = var.tf_url
}

resource "github_actions_secret" "tf_token" {
  repository       = data.github_actions_public_key.my_public_key.repository
  secret_name      = "TF_TOKEN"
  plaintext_value  = var.tf_token
}

resource "github_actions_secret" "tf_org" {
  repository       = data.github_actions_public_key.my_public_key.repository
  secret_name      = "TF_ORG"
  plaintext_value  = var.tf_org
}
