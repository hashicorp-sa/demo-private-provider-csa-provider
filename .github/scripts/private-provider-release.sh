#!/bin/bash

providerName="" # The name of the provider as it appears in the private registry
organizationName="" # The Terraform Enterprise or Cloud organization name
terraformUrl="" # The Terraforn Enterprise or Cloud URL
version="" # The version of the provider. Must be SemVer in the format 0.0.0 or v0.0.0
terraformToken="" # The Terraform Enterprise or Cloud API token

while getopts o:u:t:v:p: flag
do
    case "${flag}" in
        o) organizationName=${OPTARG};;
        u) terraformUrl=${OPTARG};;
        t) terraformToken=${OPTARG};;
        v) version=${OPTARG};;
        p) providerName=${OPTARG};;
    esac
done

if [[ -z $terraformToken ]]; then
    terraformToken=`cat ./temptoken.txt`
fi

echo "Publishing provider $providerName"

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

echo "Checking if GPG key exists..."

gpgKey=$(curl -s \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request GET \
  "https://$terraformUrl/api/registry/private/v2/gpg-keys/$organizationName/$gpgKeyId")

if [[ $gpgKey == "{\"errors\":[\"Not Found\"]}" ]]; then
  echo "GPG key does not exist, creating..."
  curl -s \
    --header "Authorization: Bearer $terraformToken" \
    --header "Content-Type: application/vnd.api+json" \
    --request POST \
    --data @gpg_payload.json \
    https://$terraformUrl/api/registry/private/v2/gpg-keys
fi


echo "Creating provider $providerName"

cat >provider_payload.json <<-EOF 
{
  "data": {
    "type": "registry-providers",
    "attributes": {
      "name": "$providerName",
      "namespace": "$organizationName",
      "registry-name": "private"
    }
  }
}
EOF

providerOutput=$(curl -s \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @provider_payload.json \
  https://$terraformUrl/api/v2/organizations/$organizationName/registry-providers)


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

providerVersion=$(curl -s \
  --header "Authorization: Bearer $terraformToken" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @provider_version_payload.json \
  https://$terraformUrl/api/v2/organizations/$organizationName/registry-providers/private/$organizationName/$providerName/versions)

platformsJson=$(cat dist/artifacts.json)

shasumsUploadUrl=$(echo $providerVersion | jq -r '.data.links."shasums-upload"')
shasumsSigUploadUrl=$(echo $providerVersion | jq -r '.data.links."shasums-sig-upload"')

platforms=$(echo $platformsJson | jq -r '.[] | select(.type == "Checksum") | @base64')
for row in $platforms; do
    _jq() {
     echo ${row} | base64 --decode | jq -r ${1}
    }
   shasumsFile=$(_jq '.path')
done

platforms=$(echo $platformsJson | jq -r '.[] | select(.type == "Signature") | @base64')
for row in $platforms; do
    _jq() {
     echo ${row} | base64 --decode | jq -r ${1}
    }
   shasumsSigFile=$(_jq '.path')
done

echo "Uploading shasums file..."
curl -s \
  --header "Content-Type: application/octet-stream" \
  --request PUT \
  --data-binary @$shasumsFile \
  $shasumsUploadUrl

echo "Uploading shasums signature file..."
curl -s \
  --header "Content-Type: application/octet-stream" \
  --request PUT \
  --data-binary @$shasumsSigFile \
  $shasumsSigUploadUrl


platforms=$(echo $platformsJson | jq -r '.[] | select(.type == "Archive") | @base64')

echo "Uploading binaries for each platform..."

for row in $platforms; do
    _jq() {
     echo ${row} | base64 --decode | jq -r ${1}
    }
   goos=$(_jq '.goos')
   goarch=$(_jq '.goarch')
   filename=$(_jq '.name')
   filepath=$(_jq '.path')
   shasum=$(grep "$filename" $shasumsFile | cut -d " " -f1)

   echo "Creating provider platform $filename for $goos $goarch"

   cat >provider_platform_payload.json <<-EOF 
    {
        "data": {
            "type": "registry-provider-version-platforms",
            "attributes": {
                "os": "$goos",
                "arch": "$goarch",
                "shasum": "$shasum",
                "filename": "$filename"
            }
        }
    }
EOF

    echo "Uploading $filename..."

    platformUpload=$(curl -s \
    --header "Authorization: Bearer $terraformToken" \
    --header "Content-Type: application/vnd.api+json" \
    --request POST \
    --data @provider_platform_payload.json \
    https://$terraformUrl/api/v2/organizations/$organizationName/registry-providers/private/$organizationName/$providerName/versions/$version/platforms)

    platformUploadUrl=$(echo $platformUpload | jq -r '.data.links."provider-binary-upload"')

    curl -s \
        --header "Content-Type: application/octet-stream" \
        --request PUT \
        --data-binary @$filepath \
        $platformUploadUrl

done

echo "Completed publishing provider $providerName"