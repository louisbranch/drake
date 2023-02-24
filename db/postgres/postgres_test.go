package postgres

import (
	"testing"

	"github.com/louisbranch/drake"
)

func TestDBInterface(t *testing.T) {
	var _ drake.Database = &DB{}

}
