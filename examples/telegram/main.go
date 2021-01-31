package main

import (
	"fmt"

	"github.com/ncostamagna/streetflow/telegram"
)

func main() {
	transport := telegram.NewClient("1536608370:AAErsMmopurv4JhVp1ondOuld8GRUJxohOY", telegram.HTTP)
	err := telegram.NewTelegramBuilder(transport).Message("Hola Nahuel,\nhoy es el cumpleaños del Random, recorda saludarlo en su dia\n\nhttps://wa.me/5491151579872?text=Feliz%20cumple%20Random!%0AEspero%20que%20lo%20pases%20de%20lo%20mejor!%0ATe%20mando%20un%20abrazo%20y%20muchos%20exitos!").Send()

	if err != nil {
		fmt.Println("Has been an error:")
		fmt.Println(err)
		return
	}

	fmt.Printf("The message has been sent")
}
