package application_test

import (
	"hexagonal-architecture/src/application"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "The price must be greather than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	err := product.Disable()

	require.Equal(t, err.Error(), "The price must be zero in order to have the product disabled")
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = "Id"

	require.Equal(t, "Id", product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Name"

	require.Equal(t, "Name", product.Name)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.DISABLED

	require.Equal(t, application.DISABLED, product.Status)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	var num float64 = 10

	require.Equal(t, num, product.Price)
}
