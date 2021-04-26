package test

import (
	"fmt"
	"testing"
	"time"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformWebServer(t *testing.T) {
	//setup params for the test to pass to CLI
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		
		//path where the terraform main.tf to be tested is located
		TerraformDir: "./modules/ec2",
		
		//input variables to pass
		Vars: map[string]interface{}{
			"servername" : "testwebserver",
			"region" : "eu-west-2",
			"size": "t2.micro"
		},
	})
	//run terraform init and apply with the 
	terraform.InitAndApply(t, terraformOptions)

	//terraform destroy when finished
	defer terraform.Destroy(t, terraformOptions)

	//get the publicIP (outputs.tf)
	publicIP := terraform.Output(t, terraformOptions, "public_ip")

	//create url with publicIP
	url := fmt.Sprintf("http://%s:8080", publicIP)
	
	//THE TEST
	// http get request to the url, check we get 200 back and he content is the string, run it for 30 times every 5 seconds
	http_helper.HttpGetWithRetry(t, url, nil, 200, "I made a Terraform Module!", 30, 5*time.Second)

}
