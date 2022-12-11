package entity_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/raphaellorencini/golang/internal/order/entity"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID:"123"}
	assert.Error(t, order.IsValid(), "invalid price")
}