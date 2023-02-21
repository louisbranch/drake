package mock

import (
	"testing"

	"github.com/louisbranch/drake/web"
)

func TestMockSatisfiesInterfaces(t *testing.T) {
	var _ web.Template = &Template{}
}
