package applicationgateway


type ApplicationGatewayProbeMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/application_gateway#status_code ApplicationGateway#status_code}.
	StatusCode *[]*string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.64.0/docs/resources/application_gateway#body ApplicationGateway#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
}

