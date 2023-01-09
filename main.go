package main

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
)

var lengte, getallen, symbolen int

func init() {
	fmt.Println("Type hoe lang u uw wachtwoord wil maken.")
	fmt.Scan(&lengte)
	fmt.Println("Type hoeveel getallen u wilt genereren in uw wachtwoord.")
	fmt.Scan(&getallen)
	fmt.Println("Type hoeveel symbolen u wilt genereren in uw wachtwoord")
	fmt.Scan(&symbolen)
}

func main() {
	wachtwoord()
}

func wachtwoord() {
	res, err := password.Generate(lengte, getallen, symbolen, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)

}
