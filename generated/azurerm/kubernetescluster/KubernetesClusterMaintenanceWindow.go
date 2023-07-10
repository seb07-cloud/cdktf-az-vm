package kubernetescluster


type KubernetesClusterMaintenanceWindow struct {
	// allowed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/kubernetes_cluster#allowed KubernetesCluster#allowed}
	Allowed interface{} `field:"optional" json:"allowed" yaml:"allowed"`
	// not_allowed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/kubernetes_cluster#not_allowed KubernetesCluster#not_allowed}
	NotAllowed interface{} `field:"optional" json:"notAllowed" yaml:"notAllowed"`
}

