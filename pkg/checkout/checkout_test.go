package checkout

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
)

type checkoutCtxKey struct{}

func iCheckout(arg1 int, arg2 string) error {
	return godog.ErrPending
}

func thePriceOfAIsC(ctx context.Context, arg1 string, arg2 int) (context.Context, error) {
	return context.WithValue(ctx, checkoutCtxKey{}, arg2), nil
}

func theTotalPriceShouldBeC(arg1 int) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I checkout (\d+) "([^"]*)"$`, iCheckout)
	ctx.Step(`^the price of a "([^"]*)" is (\d+)c$`, thePriceOfAIsC)
	ctx.Step(`^the total price should be (\d+)c$`, theTotalPriceShouldBeC)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
