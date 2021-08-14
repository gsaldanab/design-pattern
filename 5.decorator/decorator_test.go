package decorator_test

import (
	"testing"

	decorator "github.com/gsaldanab/design-pattern/5.decorator"
)

func TestDecorator(t *testing.T) {
	gc := decorator.NewInitDecorator(decorator.NewFinDecorator(&decorator.GopherCreator{}))
	gc.Create()
}
