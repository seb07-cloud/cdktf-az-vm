package datasharedatasetkustodatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataShareDatasetKustoDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/data_share_dataset_kusto_database#kusto_database_id DataShareDatasetKustoDatabase#kusto_database_id}.
	KustoDatabaseId *string `field:"required" json:"kustoDatabaseId" yaml:"kustoDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/data_share_dataset_kusto_database#name DataShareDatasetKustoDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/data_share_dataset_kusto_database#share_id DataShareDatasetKustoDatabase#share_id}.
	ShareId *string `field:"required" json:"shareId" yaml:"shareId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/data_share_dataset_kusto_database#id DataShareDatasetKustoDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/data_share_dataset_kusto_database#timeouts DataShareDatasetKustoDatabase#timeouts}
	Timeouts *DataShareDatasetKustoDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

