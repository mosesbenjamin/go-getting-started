package main

import (
	"fmt"

	"github.com/mosesbenjamin/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tracy",
		LastName:  "Mark",
	}
	fmt.Println((u))
}
