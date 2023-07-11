package shared

import (
	resourcegroup "cdk.tf/go/stack/generated/azurerm/resourcegroup"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// a struct for resource group config
type ResourceGroupConfig struct {
	Name     string
	Location string
	Tags     *cdktf.TerraformLocal
}

// a function to create a resource group and return a pointer to it.
func (rgrp *ResourceGroupConfig) Init(name string, location string, tags *cdktf.TerraformLocal) {
	rgrp.Name = name
	rgrp.Location = location
	rgrp.Tags = tags
}

// a function which takes a resource group config and creates a resourcegroup.ResourceGroup
func CreateAzResourceGroup(stack cdktf.TerraformStack, rgrp ResourceGroupConfig) resourcegroup.ResourceGroup {
	return resourcegroup.NewResourceGroup(stack, jsii.String("resourcegroup"),
		&resourcegroup.ResourceGroupConfig{
			Name:     jsii.String(rgrp.Name),
			Location: jsii.String(rgrp.Location),
		},
	)
}
