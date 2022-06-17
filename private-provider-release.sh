#!/bin/bash

providerName="terraform_provider_csa"
organizationName="provider-demo"
terraformUrl="tfe.hashicorpdemo.net"
version="v1.0.7"
terraformToken=""

while getopts o:u:t:v: flag
do
    case "${flag}" in
        o) organizationName=${OPTARG};;
        u) terraformUrl=${OPTARG};;
        t) terraformToken=${OPTARG};;
        v) version=${OPTARG};;
    esac
done

if [[ -z $terraformToken ]]; then
    terraformToken=`cat ./temptoken.txt`
fi

echo "Here we go!"
ls dist

gpg_public_key=$(awk '{printf "%s\\n", $0}' gpg_public_key.txt)

cat >gpg_payload.json <<-EOF 
{ 
  "data": {
    "type": "gpg-keys",
    "attributes": {
      "namespace": "$organizationName",
      "ascii-armor": "$gpg_public_key"
    }  
  }
}
EOF

gpgKeyId=$(gpg --show-keys gpg_public_key.txt | sed -n 2p | xargs | tail -c 17)

gpgKey=$(curl \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request GET \
  "https://$terraformUrl/api/registry/private/v2/gpg-keys/$organizationName/$gpgKeyId")

if [[ $gpgKey == "{\"errors\":[\"Not Found\"]}" ]]; then
  curl \
    --header "Authorization: Bearer $terraformToken" \
    --header "Content-Type: application/vnd.api+json" \
    --request POST \
    --data @gpg_payload.json \
    https://$terraformUrl/api/registry/private/v2/gpg-keys
fi


providerShortName=$(echo $providerName | cut -d'-' -f3)

echo "Creating provider $providerShortName"

cat >provider_payload.json <<-EOF 
{
  "data": {
    "type": "registry-providers",
    "attributes": {
      "name": "$providerShortName",
      "namespace": "$organizationName",
      "registry-name": "private"
    }
  }
}
EOF

curl \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @provider_payload.json \
  https://$terraformUrl/api/v2/organizations/$organizationName/registry-providers


version=$(echo $version | tr -d 'v')
echo "Creating provider version $version"

cat >provider_version_payload.json <<-EOF 
{
  "data": {
    "type": "registry-provider-versions",
    "attributes": {
      "version": "$version",
      "key-id": "$gpgKeyId",
      "protocols": ["5.0"]
    }
  }
}
EOF

providerVersion=$(curl \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @provider_version_payload.json \
  https://$terraformUrl/api/v2/organizations/$organizationName/registry-providers/private/$organizationName/$providerShortName/versions)

shasumsUploadUrl=$(echo $providerVersion | jq -r '.data.links."shasums-upload"')
shasumsSigUploadUrl=$(echo $providerVersion | jq -r '.data.links."shasums-sig-upload"')
shasumsFile="dist/${providerName}_${version}_SHA256SUMS"
shasumsSigFile="dist/${providerName}_${version}_SHA256SUMS.sig"

curl \
  --header "Content-Type: application/octet-stream" \
  --request PUT \
  --data-binary @$shasumsFile \
  $shasumsUploadUrl

curl \
  --header "Content-Type: application/octet-stream" \
  --request PUT \
  --data-binary @$shasumsSigFile \
  $shasumsSigUploadUrl

cat >provider_platform_payload.json <<-EOF 
{
  "data": {
    "type": "registry-provider-version-platforms",
    "attributes": {
      "os": "linux",
      "arch": "amd64",
      "shasum": "8f69533bc8afc227b40d15116358f91505bb638ce5919712fbb38a2dec1bba38",
      "filename": "terraform-provider-aws_3.1.1_linux_amd64.zip"
    }
  }
}
EOF

curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @payload.json \
  https://$terraformUrl/api/v2/organizations/hashicorp/registry-providers/private/$organizationName/aws/versions/$version/platforms