package main

import (
	"fmt"
	"testing"

	"github.com/narenarjun/go-micro-learning/sdk/client"
	"github.com/narenarjun/go-micro-learning/sdk/client/products"
)

func TestOurClient(t *testing.T){
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")

	c := client.NewHTTPClientWithConfig(nil,cfg)
	params := products.NewListProductParams()
	// params.
	prods, err := c.Products.ListProduct(params)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Printf("%#v",prods.GetPayload()[0])
}