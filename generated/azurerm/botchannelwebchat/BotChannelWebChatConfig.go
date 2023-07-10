package botchannelwebchat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BotChannelWebChatConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#bot_name BotChannelWebChat#bot_name}.
	BotName *string `field:"required" json:"botName" yaml:"botName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#location BotChannelWebChat#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#resource_group_name BotChannelWebChat#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#site_names BotChannelWebChat#site_names}.
	SiteNames *[]*string `field:"required" json:"siteNames" yaml:"siteNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#id BotChannelWebChat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/bot_channel_web_chat#timeouts BotChannelWebChat#timeouts}
	Timeouts *BotChannelWebChatTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

