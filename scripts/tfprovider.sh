#!/bin/bash

# Read the JSON from a file
json=$(cat ../cdktf.json)

providers='["hashicorp/azurerm@~> 3.6.4"]'

# Update the 'terraformProviders' field
updated_json=$(echo "$json" | jq --argjson providers "$providers" '.terraformProviders = $providers')

echo "$updated_json"
