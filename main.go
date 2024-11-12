package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// Porta e endereço onde o servidor irá escutar
const (
	host = "0.0.0.0" // Escuta em todas as interfaces de rede
	port = "9000"    // Porta que você pode configurar conforme necessário
)

func main() {
	// Inicia o servidor TCP na porta definida
	listener, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Servidor TCP escutando em %s:%s\n", host, port)

	// Aceita conexões continuamente
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão: %v", err)
			continue
		}
		fmt.Printf("Nova conexão de %s\n", conn.RemoteAddr().String())

		// Manipula a conexão em uma nova goroutine
		go handleConnection(conn)
	}
}

// handleConnection trata a conexão de um cliente específico
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Abre o arquivo de log (ou cria caso não exista) para registrar as mensagens recebidas
	logFile, err := os.OpenFile("mensagens.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Erro ao abrir arquivo de log: %v", err)
		return
	}
	defer logFile.Close()

	writer := bufio.NewWriter(logFile)

	// Lê a mensagem do cliente
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// Exibe a mensagem recebida no console
		fmt.Printf("[%s] Mensagem recebida de %s: %s\n", timestamp, conn.RemoteAddr().String(), msg)

		// Grava a mensagem no arquivo de log com o timestamp
		logEntry := fmt.Sprintf("[%s] %s: %s\n", timestamp, conn.RemoteAddr().String(), msg)
		if _, err := writer.WriteString(logEntry); err != nil {
			log.Printf("Erro ao escrever no arquivo de log: %v", err)
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Erro ao ler do cliente %s: %v", conn.RemoteAddr().String(), err)
	}
}
