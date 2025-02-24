package lambda

import (
	"context"
	"testing"
)

func TestDeployer_Deploy(t *testing.T) {
	ctx := context.Background()
	deployer := NewDeployer(ctx, "./testdata")

	err := deployer.Deploy("test-function", "test.zip")
	if err != nil {
		t.Errorf("Deploy failed: %v", err)
	}
}
