package random

import (
	"testing"

	"github.com/panovateam/go-micro/selector"
)

func TestRandom(t *testing.T) {
	selector.Tests(t, NewSelector())
}
