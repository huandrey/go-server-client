package main

import (
	"fmt"
	"go-cliente-servidor/src/helpers"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var hash string
	fmt.Fscanf(conn, "%s\n", &hash)

	directory := "./tmp"
	found, calculatedHash, err := helpers.FindHash(hash, directory)

	if calculatedHash >=0 {
		fmt.Println("O arquivo de hash", calculatedHash, "foi encontrado.")
	}
	if err != nil {
		fmt.Fprintf(conn, "error: %v\n", err)
		return
	}

	if found {
		fmt.Fprintf(conn, "found\n")
	} else {
		fmt.Fprintf(conn, "not found\n")
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso correto: ./servidor <porta>")
		os.Exit(1)
	}
	port := os.Args[1]

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Printf("Servidor rodando na porta %s...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleConnection(conn)
	}
}
