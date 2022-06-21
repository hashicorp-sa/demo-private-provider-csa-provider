## Description

This repository demonstrates building, releases and publishing a custom Terraform provider to the private registry in Terraform Enterprise and Cloud.

It consists of:

- A provider written in golang based on the https://github.com/hashicorp/terraform-provider-scaffolding template.
- A GitHub Action to build and release the provider using [goreleaser](https://goreleaser.com/) and the [goreleaser GitHub Action](https://github.com/goreleaser/goreleaser-action).
- A [bash script](.github/scripts/private-provider-release.sh) example of calling the Terraform Enterprise / Cloud API to create and publish the provider to the private registry.

## Getting Started & Documentation

The GitHub Action [release.yml](release.yml) requires the following secrets to be set at the repository or environment scope:

- `GPG_PRIVATE_KEY`: The GPG private key. See below for how to generate this.
- `PASSPHRASE`: The GPG private key passphrase. See below for how to generate this.
- `TF_URL`: The url for Terraform Enterprise or Cloud.
- `TF_TOKEN`: The API token for Terraform Enterprise or Cloud.
- `TF_ORG`: The organization for the private registry Terraform Enterprise or Cloud. 

In order to run the [release.yml](.github/workflows/release.yml) workflow, you'll need to set the GPG private key (`GPG_PRIVATE_KEY`) and passphrase (`PASSPHRASE`). Follow (these steps)[https://learn.hashicorp.com/tutorials/terraform/provider-release-publish?in=terraform/providers#generate-gpg-signing-key
] to do that.

More details about the steps to release to the private registry can be found [here](https://www.terraform.io/cloud-docs/registry/publish-providers#publishing-a-provider-and-creating-a-version).


## Contributing
Please create a fork and raise a Pull Request.

## Disclaimer
“By using the software in this repository (the “Software”), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an “as-is” basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.”