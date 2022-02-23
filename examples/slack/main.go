package main

import (
	"fmt"

	"github.com/ncostamagna/streetflow/slack"
)

func main() {
	slack, err := slack.NewSlackBuilder("birthday", "My token").Build()

	if err != nil {
		fmt.Println("Has been an error")
	}

	response := slack.SendMessage("<@U01CDEPA3T9> hoy es el cumple de %s, saludalo en su dia: https://api.whatsapp.com/send?phone=541130100415&text=aassa")

	fmt.Println(response)
}
