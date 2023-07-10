package datafactorydatasetpostgresql


type DataFactoryDatasetPostgresqlSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/data_factory_dataset_postgresql#name DataFactoryDatasetPostgresql#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/data_factory_dataset_postgresql#description DataFactoryDatasetPostgresql#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/data_factory_dataset_postgresql#type DataFactoryDatasetPostgresql#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

