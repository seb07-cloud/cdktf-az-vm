package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfigApplicationStack struct {
	// docker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#docker LinuxFunctionAppSlot#docker}
	Docker interface{} `field:"optional" json:"docker" yaml:"docker"`
	// The version of .Net. Possible values are `3.1` and `6.0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#dotnet_version LinuxFunctionAppSlot#dotnet_version}
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// The version of Java to use. Possible values are `8`, and `11`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#java_version LinuxFunctionAppSlot#java_version}
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// The version of Node to use. Possible values include `12`, and `14`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#node_version LinuxFunctionAppSlot#node_version}
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// The version of PowerShell Core to use. Possibles values are `7`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#powershell_core_version LinuxFunctionAppSlot#powershell_core_version}
	PowershellCoreVersion *string `field:"optional" json:"powershellCoreVersion" yaml:"powershellCoreVersion"`
	// The version of Python to use. Possible values include `3.9`, `3.8`, and `3.7`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#python_version LinuxFunctionAppSlot#python_version}
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#use_custom_runtime LinuxFunctionAppSlot#use_custom_runtime}.
	UseCustomRuntime interface{} `field:"optional" json:"useCustomRuntime" yaml:"useCustomRuntime"`
	// Should the DotNet process use an isolated runtime. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.0.2/docs/resources/linux_function_app_slot#use_dotnet_isolated_runtime LinuxFunctionAppSlot#use_dotnet_isolated_runtime}
	UseDotnetIsolatedRuntime interface{} `field:"optional" json:"useDotnetIsolatedRuntime" yaml:"useDotnetIsolatedRuntime"`
}

