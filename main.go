package main

import (
	"context"
	"fmt"

	"github.com/fabruun/go-customers/api"
)

func main() {
	app := api.New()

	err := app.Start(context.Background())
	if err != nil {
		fmt.Println("failed to start app", err)
	}
}
