package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vinihss/telecine/messages"
	"google.golang.org/grpc"
)

func main() {
	// Conecta ao servidor gRPC
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	// Envia uma mensagem ao servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.MessageRequest{
		Content:  "Olá, servidor gRPC!",
		ClientId: "Cliente_1",
	}

	res, err := client.SendMessage(ctx, req)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}
	log.Printf("Resposta do servidor: %s", res.Status)
}
