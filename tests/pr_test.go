package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terraform-ibm-modules/ibmcloud-terratest-wrapper/testprojects"
)

const resourceGroup = "geretain-test-resources"

func TestProjectsFullTest(t *testing.T) {

	options := testprojects.TestProjectOptionsDefault(&testprojects.TestProjectsOptions{
		Testing:              t,
		Prefix:               "stack-template", // Add this prefix to the project name to make it easier to identify in the IBM Cloud console
		DeployTimeoutMinutes: 30,               // Set the timeout for the stack deployment to a reasonable value for your stack
		CatalogProductName:   "stack-template", // Set the Product name from the ibm_catalog.json
		CatalogFlavorName:    "stack-template", // Set the Flavor name from the ibm_catalog.json, for any configurations set here the default values will be applied to the stack
	})

	// Test inputs override all other input values from the stack definiton and the ibm_catalog.json
	options.StackMemberInputs = map[string]map[string]interface{}{ // Set the inputs for the stack members
		"1a-primary-da": {
			"prefix": fmt.Sprintf("p%s", options.Prefix),
		},
		"secondary-da": {
			"prefix": fmt.Sprintf("s%s", options.Prefix),
		},
	}
	options.StackInputs = map[string]interface{}{ // Set the inputs for the stack
		"resource_group_name": resourceGroup,
		"ibmcloud_api_key":    os.Getenv("TF_VAR_ibmcloud_api_key"),
	}

	err := options.RunProjectsTest()
	if assert.NoError(t, err) {
		t.Log("TestProjectsFullTest Passed")
	} else {
		t.Error("TestProjectsFullTest Failed")
	}
}
