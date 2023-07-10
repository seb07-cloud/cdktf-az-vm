package kubernetescluster


type KubernetesClusterLinuxProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/kubernetes_cluster#admin_username KubernetesCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// ssh_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/kubernetes_cluster#ssh_key KubernetesCluster#ssh_key}
	SshKey *KubernetesClusterLinuxProfileSshKey `field:"required" json:"sshKey" yaml:"sshKey"`
}

