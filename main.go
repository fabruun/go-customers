package main

import (
	"context"
	"fmt"

	application "github.com/fabruun/go-customers/application/app"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app", err)
	}
}
