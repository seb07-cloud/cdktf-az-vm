# Infra-Logic-as-Code

This project is based on CDKTF and creates resources in Azure.

![Infra-Logic-as-Code](https://github.com/seb07-cloud/infra-logic-as-code/blob/main/images/cdktf.png)

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository
2. Install the necessary dependencies
3. Build the project
4. Deploy the stack to Azure

### Prerequisites

- Terraform
- CDK for Terraform
- Azure CLI
- Go
- Node.js

### Installing

To install the necessary dependencies, run the following command:

cdktf init --template=go --local
cdktf provider add azurerm@~> 3.6.4

### Run the tfprovider script to update the provider config

./scripts/tfprovider.sh

### Fetch Providers

cdktf get

### Deploying

To deploy the stack to Azure, run the following command:

cdktf synth

cdktf deploy

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
