## Description

This repository demonstrates building, releases and publishing a custom Terraform provider to the private registry in Terraform Enterprise and Cloud.

It consists of:

- A provider written in golang based on the https://github.com/hashicorp/terraform-provider-scaffolding template.
- A GitHub Action to build and release the provider using [goreleaser](https://goreleaser.com/) and the [goreleaser GitHub Action](https://github.com/goreleaser/goreleaser-action).
- A [bash script](.github/scripts/private-provider-release.sh) example of calling the Terraform Enterprise / Cloud API to create and publish the provider to the private registry.

## Getting Started & Documentation

The GitHub Action [release.yml](release.yml) requires the following secrets to be set at the repository or environment scope:

- `GPG_PRIVATE_KEY`: The GPG private key. See below for how to generate this.
- `PASSPHRASE`: The GPG private key passphrase. See [below](generating-gpg-keys) for how to generate this.
- `TF_URL`: The url for Terraform Enterprise or Cloud.
- `TF_TOKEN`: The API token for Terraform Enterprise or Cloud.
- `TF_ORG`: The organization for the private registry Terraform Enterprise or Cloud. 

Follow these steps to setup and run the demo:

1. Setup and instance of Terraform Enterprise or Cloud and create an organization.
1. Fork and clone this repository.
1. Generate a public and private key GPG following [these instructions](generating-gpg-keys).
1. Save the public key in [gpg_public_key.txt](gpg_public_key.txt).
1. Add the private key and passphrase to the `GPG_PRIVATE_KEY` and `PASSPHRASE` secrets in GitHub.
1. Generate an API token for Terraform Enterprise or Cloud.
1. Add the Terraform Enterprise or Cloud URL, API token and organization to the `TF_URL`, `TF_TOKEN` and `TF_ORG` secrets in GitHub.
1. Commit the [gpg_public_key.txt](gpg_public_key.txt) and push to origin.
1. Add a new tag and push it to origin. e.g. `git tag v1.0.1` and `git push --tags`.
1. The new tag will trigger the [release.yml](release.yml) GitHub Action to build and release the provider.
1. Watch the GitHub Action complete and then navigate to the private registry in Terraform Enterprise or Cloud to see your published provider.
1. Head over to the [demo](https://github.com/HashiCorp-CSA/demo-private-provider-csa-provider-demo) repository for the next steps.

More details about the steps to release to the private registry can be found [here](https://www.terraform.io/cloud-docs/registry/publish-providers#publishing-a-provider-and-creating-a-version).

### Generating GPG Keys
In order to run the [release.yml](.github/workflows/release.yml) workflow, you'll need to set the GPG private key (`GPG_PRIVATE_KEY`) and passphrase (`PASSPHRASE`). Follow (these steps)[https://learn.hashicorp.com/tutorials/terraform/provider-release-publish?in=terraform/providers#generate-gpg-signing-key
] to do that.

## What does it do?

The [release.yml](.github/workflows/release.yml) GitHub Action will build the provider and publish it to the private registry in Terraform Enterprise or Cloud. It completes the following steps:

1. Checkout the code.
1. Pulls the git history down, so it can see the tags.
1. Installs `go` based on the version required by the provider.
1. Imports the GPG key using the private key and passphrase.
1. Runs `goreleaser` to build, package and release the provider for multiple platforms specified in the [.goreleaser.yml](.goreleaser.yml) file.
    1. The [.goreleaser.yml](.goreleaser.yml) file contains `project_name: terraform-provider-demo` which determines the prefix of the binary name it generates. Terraform expects the binary to be prefixed with `terraform-provider-` followed by the provider name. If you don't follow this pattern, then `terraform init` will fail.
    1. goreleaser will publish the binaries to the `./dist` folder as well as releasing them. The `./dist` folder is used by the [private-provider-release.sh](.github/scripts/private-provider-release.sh) step.
1. Gets the latest version tag into an environment variable.
1. Runs the [private-provider-release.sh](.github/scripts/private-provider-release.sh) script to publish the provider to the private registry.

The [private-provider-release.sh](.github/scripts/private-provider-release.sh) script runs the following steps:

1. It gets the follow input variables:
    1. `u`: The url for Terraform Enterprise or Cloud.
    1. `o`: The organization for Terraform Enterprise or Cloud.
    1. `t`: The API token for Terraform Enterprise or Cloud.
    1. `v`: The version tag (e.g. v1.0.1).
    1. `p`: Name of the provider in the private registry (e.g. demo).
1. Gets the public GPG key and extrapolates the name of the key.
1. Checks if the GPG key already exists in Terraform Cloud or Enterprise.
1. If the GPG key does not exist in Terraform Cloud or Enterprise, it creates it.
1. Creates the provider in the private registry if it does not already exist.
1. Creates the provider version in the private registry, which returns the urls to upload the SHA sum and signature files to.
1. Interrogates the `./dist/artifacts.json` file to get the paths to the SHA sum and signature files.
1. Uploads the SHA sum file.
1. Uploads the signature file.
1. Interrogates the `./dist/artifacts.json` file to get the path of the `.tar.gz`, `os` and `architecture` of all the provider binaries generated by goreleaser.
1. Iterates over each binary configuration performing the following steps:
    1. Creates a platform in the private registry, which returns an upload url for the binary `.tar.gz`.
    2. Uploads the binary `.tar.gz` file to the platform.
1. The provider is now available in the private registry.

## Contributing
Please create a fork and raise a Pull Request.

## Disclaimer
“By using the software in this repository (the “Software”), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an “as-is” basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.”