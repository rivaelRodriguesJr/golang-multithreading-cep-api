package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	buscacep "github.com/rivaelRodriguesJr/golang-multithreading-cep-api/buscaCep"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: program <cep>")
		fmt.Println("Example: program 13025-320 or 13025320")
		os.Exit(1)
	}

	cep := os.Args[1]

	validCEP := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	if !validCEP.MatchString(cep) {
		fmt.Println("Error: Invalid CEP format")
		fmt.Println("CEP must have 8 digits (e.g., 13025-320 or 13025320)")
		os.Exit(1)
	}

	fmt.Printf("Searching for CEP: %s\n\n", cep)

	chBrasilApi := make(chan buscacep.BrasilApi)
	chViaCep := make(chan buscacep.ViaCep)
	go func() {
		res, _ := buscacep.BuscaBrasilApi(cep)
		chBrasilApi <- *res
	}()
	go func() {
		res, _ := buscacep.BuscaViaCep(cep)
		chViaCep <- *res
	}()

	select {
	case msg := <-chBrasilApi:
		{
			fmt.Println("Data from: Brasil API")
			fmt.Println(msg)
		}
	case msg := <-chViaCep:
		{
			fmt.Println("Data from: Via CEP")
			fmt.Println(msg)
		}
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
