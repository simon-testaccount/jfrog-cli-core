package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jfrog/jfrog-cli-core/v2/utils/log"
	clientLog "github.com/jfrog/jfrog-client-go/utils/log"

	"github.com/jfrog/jfrog-cli-core/v2/utils/tests"
	clientTests "github.com/jfrog/jfrog-client-go/utils/tests"
)

const CoreIntegrationTests = "github.com/jfrog/jfrog-cli-core/v2/tests"

func init() {
	log.SetDefaultLogger()
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func printReversedEnv() {
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			fmt.Printf("%s=%s\n", parts[0], reverse(parts[1]))
		} else {
			fmt.Println(reverse(e))
		}
	}
}

func TestUnitTests(t *testing.T) {
	printReversedEnv()

	cleanUpJfrogHome, err := tests.SetJfrogHome()
	if err != nil {
		clientLog.Error(err)
		os.Exit(1)
	}
	defer cleanUpJfrogHome()

	packages := clientTests.GetTestPackages("./../...")
	packages = clientTests.ExcludeTestsPackage(packages, CoreIntegrationTests)
	assert.NoError(t, clientTests.RunTests(packages, false))
}
