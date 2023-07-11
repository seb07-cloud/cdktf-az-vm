package shared

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Tags struct {
	Owner       string
	Service     string
	Environment string
}

func (t *Tags) Init(owner string, service string, environment string) {
	t.Owner = owner
	t.Service = service
	t.Environment = environment
}

func CreateTags(stack *cdktf.TerraformStack, tags Tags) cdktf.TerraformLocal {
	tagsMap := StructToMap(&tags)
	return cdktf.NewTerraformLocal(*stack, jsii.String("common_tags"), tagsMap)
}
