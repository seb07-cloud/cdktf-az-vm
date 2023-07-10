package containerregistrytask


type ContainerRegistryTaskAgentSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/container_registry_task#cpu ContainerRegistryTask#cpu}.
	Cpu *float64 `field:"required" json:"cpu" yaml:"cpu"`
}

