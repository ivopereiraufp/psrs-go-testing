package beer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestSubTotal(t *testing.T) {
	cart := NewCart()
	assert.Equal(t, 0, len(cart.Cases))

	superBockStout := FixtureBeer("Super Bock", "Stout", 10.0)
	cart.AddCase(FixtureCase(6, superBockStout, 12.0))

	sagresPreta := FixtureBeer("Sagres", "Preta", 8.0)
	cart.AddCase(FixtureCase(12, sagresPreta, 25.0))

	assert.Equal(t, 37.0, cart.Subtotal())
}

func TestCartSubtotal(t *testing.T) {
	tests := []struct {
		name string
		cart *Cart
		want float64
	}{
		{
			name: "Testing Empty Cart",
			cart: &Cart{},
			want: 0,
		},
		{
			name: "Superbock and Sagres",
			cart: &Cart{Cases: []*Case{
				FixtureCase(6, FixtureBeer("Super Bock", "Stout", 10.0), 12.0),
				FixtureCase(12, FixtureBeer("Sagres", "Preta", 8.0), 25.0),
			}},
			want: 37.0,
		},
		{
			name: "Price Negative?",
			cart: &Cart{Cases: []*Case{
				FixtureCase(6, FixtureBeer("Super Bock", "Stout", 10.0), -12.0),
				FixtureCase(12, FixtureBeer("Sagres", "Preta", 8.0), 25.0),
			}},
			want: 13.0,
		},
		{
			name: "Testing FixtureCart",
			cart: FixtureCart(),
			want: 14.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.cart.Subtotal())
		})
	}
}

func TestSubscriptionTimer(t *testing.T) {
	contxt := context.Background()

	cart1 := &Cart{Cases: []*Case{
		FixtureCase(6, FixtureBeer("Super Bock", "Stout", 10.0), -12.0),
		FixtureCase(12, FixtureBeer("Sagres", "Preta", 8.0), 25.0),
	}}

	subscription := &Subscription{
		cart:        cart1,
		interval:    time.Duration(1) * time.Second,
		messageChan: make(chan interface{}),
	}

	go subscription.startSubscriptionTimer(contxt)
	msg := <-subscription.messageChan
	order, ok := msg.(*Cart)
	if !ok {
		t.Fatal("message invalid")
	}
	assert.Equal(t, cart1, order)

}

func TestStartOrderHandler(t *testing.T) {
	handler := &OrderHandler{
		messageChan: make(chan interface{}),
	}

	go handler.startOrderHandler(context.Background())
	assert.Equal(t, 0, len(handler.ProcessedOrders))

	handler.messageChan <- FixtureCart()
	handler.messageChan <- &Cart{Cases: []*Case{
		FixtureCase(6, FixtureBeer("Super Bock", "Stout", 10.0), -12.0),
		FixtureCase(12, FixtureBeer("Sagres", "Preta", 8.0), 25.0),
	}}
	handler.messageChan <- FixtureCase(6, FixtureBeer("Super Bock", "Stout", 10.0), -12.0)

	assert.Equal(t, 2, len(handler.ProcessedOrders))

}
