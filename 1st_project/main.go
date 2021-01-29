package main

import (
	"fmt"

	"github.com/sophieLiu/1st_go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Yue",
		LastName:  "Liu",
	}
	fmt.Println(u)
}
