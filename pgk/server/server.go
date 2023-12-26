package server

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/LOLONEKIT2019/prof_of_work/pgk/message"
)

type Handler func(ctx context.Context, userInfo string, messageType message.Type, body string) (*message.Message, error)

const (
	connectionTimeout = time.Second * 5
)

func RunServer(port int, handler Handler) error {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("server started on port", port)

	for {

		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("cannot accept connection: %w", err)
		}

		go handleConnection(conn, handler)
	}
}

func handleConnection(conn net.Conn, handler Handler) {

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	reader := bufio.NewReader(conn)

	for {
		userInfo := conn.RemoteAddr().String()
		fmt.Println("received request from", userInfo)

		request, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read request:", err)
			return
		}

		message, err := message.ParseMessage(request)
		if err != nil {
			fmt.Println("parse message:", err)
			return
		}

		msg, err := handler(ctx, userInfo, message.Type, message.Body)
		if err != nil {
			fmt.Println("handle message:", err)
			return
		}

		response := fmt.Sprintf("%s\n", msg.String())
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("write message:", err)
			return
		}
	}
}
