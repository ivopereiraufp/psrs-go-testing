package beer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddCaseTesting(t *testing.T) {
	cart := NewCart()
	if len(cart.Cases) != 0 {
		t.Fatal("This is expected to be an empty cart")
	}
	superBockStout := FixtureBeer("Super Bock", "Stout", 10.0)
	cart.AddCase(FixtureCase(6, superBockStout, 12.0))
	if len(cart.Cases) != 1 {
		t.Fatal("We expect 1 case in the cart")
	}
}

func TestAddCaseAssert(t *testing.T) {
	cart := NewCart()
	assert.Equal(t, 0, len(cart.Cases))

	superBockStout := FixtureBeer("Super Bock", "Stout", 10.0)
	cart.AddCase(FixtureCase(6, superBockStout, 12.0))
	assert.Equal(t, 1, len(cart.Cases))
}
