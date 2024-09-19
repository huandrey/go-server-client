package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"strings"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "search" {
		fmt.Println("Uso correto: ./programa search hash_de_arquivo")
		return
	}

	hash := os.Args[2]
	servers := []struct {
		ip   string
		port string
	}{
		// {"150.165.42.133", "9000"},
		{"150.165.42.135", "9000"},
		// {"150.165.42.134", "9000"},
		// {"150.165.42.136", "9000"}, 
	}

	for _, server := range servers {
		response, err := checkHashOnServer(server.ip, server.port, hash)
		response = strings.TrimSpace(response)
		
		if err != nil {
			fmt.Printf("Erro ao conectar ao servidor %s:%s: %v\n", server.ip, server.port, err)
			continue
		}
	
		if response == "found" {
			err := receiveFile(server.ip, server.port, hash)
			if err != nil {
				fmt.Printf("Erro ao receber o arquivo: %v\n", err)
			}
			fmt.Printf("%s:%s\n", server.ip, server.port)
		} else {
			fmt.Printf("Hash não encontrado no servidor %s:%s\n", server.ip, server.port)
		}
	}
}

// Função para conectar ao servidor e enviar o hash
func checkHashOnServer(ip, port, hash string) (string, error) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\n", hash)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}

	return message, nil
}

// Função para receber o arquivo do servidor
func receiveFile(ip, port, hash string) error {
	fmt.Println("receive file function")

	conn, err := net.Dial("tcp", ip+":"+port)

	fmt.Println("conexao")
	fmt.Println(conn)

	if err != nil {
		return err
	}
	defer conn.Close()

	file, err := os.Create(hash + ".download")
	if err != nil {
		return err
	}

	fmt.Println("file")
	fmt.Println(file)
	defer file.Close()

	// Recebe o arquivo do servidor
	_, err = io.Copy(file, conn)
	if err != nil {
		return err
	}
	fmt.Println("response")
	// fmt.Println(response)
	
	fmt.Println("Arquivo salvo como:", hash+".download")
	return nil
}
