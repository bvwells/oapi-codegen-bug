package main

import (
	"context"
	"fmt"

	"github.com/bwells/oapi-codegen-bug/petstore"
)

func main() {
	server := "http://petstore.swagger.io/api"
	client, err := petstore.NewClientWithResponses(server)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	validPets, err := client.ValidatePetsWithResponse(ctx, petstore.ValidatePetsJSONRequestBody{
		Names: []string{"cat", "dog"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(validPets)
}
