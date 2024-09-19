package main

import (
  "bufio"
  "fmt"
  "io"
  "net"
  "os"
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
    err := checkHashOnServerAndReceiveFile(server.ip, server.port, hash)
    if err != nil {
      fmt.Printf("Erro ao processar o servidor %s:%s: %v\n", server.ip, server.port, err)
    }
  }
}

// Função para conectar ao servidor, enviar o hash e receber o arquivo se o hash for encontrado
func checkHashOnServerAndReceiveFile(ip, port, hash string) error {
  conn, err := net.Dial("tcp", ip+":"+port)
  if err != nil {
    return fmt.Errorf("erro ao conectar ao servidor: %v", err)
  }
  defer conn.Close()

  // Enviar o hash
  fmt.Fprintf(conn, "%s\n", hash)

  // Ler a resposta do servidor
  message, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    return fmt.Errorf("erro ao ler a resposta do servidor: %v", err)
  }

  message = strings.TrimSpace(message)
  if message == "found" {
    fmt.Println("Arquivo encontrado, recebendo...")

    // Receber o arquivo na mesma conexão
    err := receiveFile(conn, hash)
    if err != nil {
      return fmt.Errorf("erro ao receber o arquivo: %v", err)
    }

    fmt.Printf("Arquivo recebido com sucesso do servidor %s:%s\n", ip, port)
  } else {
    fmt.Printf("Hash não encontrado no servidor %s:%s\n", ip, port)
  }

  return nil
}

// Função para receber o arquivo do servidor
<<<<<<< HEAD
func receiveFile(conn net.Conn, hash string) error {
  // Criar o arquivo local para salvar os dados recebidos
  file, err := os.Create(hash + ".txt")
  if err != nil {
    return fmt.Errorf("erro ao criar arquivo: %v", err)
  }
  defer file.Close()

  // Copiar os dados recebidos para o arquivo
  bytesCopied, err := io.Copy(file, conn)
  if err != nil {
    return fmt.Errorf("erro ao copiar dados: %v", err)
  }

  fmt.Printf("Arquivo recebido (%d bytes) e salvo como: %s\n", bytesCopied, hash+".txt")
  return nil
}
=======
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
>>>>>>> d01be5e82ee96aa5b03b84a05314bcdd09c14bc2
