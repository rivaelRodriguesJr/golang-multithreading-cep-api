# Golang Multithreading CEP API

A simple Go program that fetches Brazilian address information (CEP) from two different APIs **at the same time** and returns whichever responds first.

## What it does

You give it a CEP (Brazilian postal code), and it queries:

- [ViaCEP](http://viacep.com.br) — `http://viacep.com.br/ws/{cep}/json/`
- [BrasilAPI](https://brasilapi.com.br) — `https://brasilapi.com.br/api/cep/v1/{cep}`

Both requests run in parallel (goroutines). The first one to respond wins, and the result is printed. If neither responds within **1 second**, you get a timeout message.

This is a classic pattern to learn about **goroutines** and **channels** in Go.

## Project structure
```
.
├── main.go              # program entry point
├── buscaCep/
│   └── buscaCep.go      # API clients + structs + String() methods
├── go.mod
└── go.sum
```

## Requirements
[Go](https://go.dev/dl/) 1.26 or newer


## How to run

Clone the repo and run:

```bash
go run main.go 13025-320
```

Or with the CEP without hyphen:
```bash
go run main.go 13025320
```

## Invalid input

```bash
go run main.go abc
# Error: Invalid CEP format
# CEP must have 8 digits (e.g., 13025-320 or 13025320)

go run main.go
# Usage: program <cep>
# Example: program 13025-320 or 13025320
``` 

## Example output

```
Searching for CEP: 13025-320

Via CEP
CEP:          13025-320
Logradouro:   Rua Exemplo
Complemento:  
Unidade:      
Bairro:       Centro
Localidade:   Campinas
UF:           SP
Estado:       São Paulo
Região:       Sudeste
IBGE:         3509502
GIA:          
DDD:          19
SIAFI:        6633
```

> Since the two APIs run in parallel, sometimes you'll see the Brasil API result instead — whichever answers first.

