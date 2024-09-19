package main

import (
    "fmt"
    "os"
    "go-cliente-servidor/src/helpers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: ./programa <caminho_do_arquivo>")
		return
	}

	filePath := os.Args[1]
	hash, err := helpers.Sum(filePath)
	if err != nil {
		fmt.Printf("Erro ao calcular o hash: %v\n", err)
		return
	}

	fmt.Print("Hash do arquivo: ", hash)
	fmt.Print("\n")
}
