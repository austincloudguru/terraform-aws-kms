package test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
  "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestExamplesTerraform(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())

	randId := strconv.Itoa(rand.Intn(100000))
	// keySuffix := []string{randId}

	terraformOpts := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Our Terraform code is in the /aws folder.
		TerraformDir: "../examples/terratest/",
		Vars: map[string]interface{}{
			"key_suffix": randId,
		},
		// BackendConfig: map[string]interface{}{
		// 	"path": "/go/src/terraform.state",
		// },
	})

	defer terraform.Destroy(t, terraformOpts)
	terraform.InitAndApply(t, terraformOpts)

	// Run `terraform output` to get the value of an output variable
	keyArn := terraform.Output(t, terraformOpts, "key_arn")
	// Verify we're getting back the outputs we expect
	assert.Contains(t, keyArn, "arn:aws:kms:us-west-2")

	// Run `terraform output` to get the value of an output variable
	aliasName := terraform.Output(t, terraformOpts, "alias_name")
	expectedAliasName := "alias/terratest-" + randId
	// Verify we're getting back the outputs we expect
	assert.Equal(t, expectedAliasName, aliasName)

}
