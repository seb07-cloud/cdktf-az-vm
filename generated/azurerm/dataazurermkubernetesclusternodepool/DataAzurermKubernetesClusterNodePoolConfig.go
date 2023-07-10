package dataazurermkubernetesclusternodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermKubernetesClusterNodePoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/data-sources/kubernetes_cluster_node_pool#kubernetes_cluster_name DataAzurermKubernetesClusterNodePool#kubernetes_cluster_name}.
	KubernetesClusterName *string `field:"required" json:"kubernetesClusterName" yaml:"kubernetesClusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/data-sources/kubernetes_cluster_node_pool#name DataAzurermKubernetesClusterNodePool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/data-sources/kubernetes_cluster_node_pool#resource_group_name DataAzurermKubernetesClusterNodePool#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/data-sources/kubernetes_cluster_node_pool#id DataAzurermKubernetesClusterNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/data-sources/kubernetes_cluster_node_pool#timeouts DataAzurermKubernetesClusterNodePool#timeouts}
	Timeouts *DataAzurermKubernetesClusterNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

