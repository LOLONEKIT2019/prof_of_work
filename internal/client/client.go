package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/LOLONEKIT2019/prof_of_work/internal/config"
	"github.com/LOLONEKIT2019/prof_of_work/internal/contract"
	"github.com/LOLONEKIT2019/prof_of_work/internal/pow"
	"github.com/LOLONEKIT2019/prof_of_work/pgk/message"
)

type Client struct {
	conn net.Conn
}

func Start() error {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", config.Config.ServerHost, config.Config.ServerPort))
	if err != nil {
		return fmt.Errorf("initialize connection: %w", err)
	}
	defer conn.Close()

	client := Client{conn: conn}

	for {
		quote, err := client.getQuote(conn)
		if err != nil {
			return fmt.Errorf("get quote: %w", err)
		}

		fmt.Print("quote received: ", quote)
		time.Sleep(time.Second * 10)
	}
}

func (c *Client) getQuote(conn net.Conn) (string, error) {

	reader := bufio.NewReader(conn)
	writer := conn

	// 1 - request task
	err := sendMessage(writer, &message.Message{Type: contract.RequestTask})
	if err != nil {
		return "", fmt.Errorf("request task: %w", err)
	}
	fmt.Println("request task")

	// 2 - read task
	msg, err := readMessage(reader)
	if err != nil {
		return "", fmt.Errorf("read task: %w", err)
	}
	if msg.Type != contract.ResponseTask {
		return "", fmt.Errorf("received unknown read task message: %s", msg.String())
	}
	fmt.Println("read task")

	// 3 - do job
	job := pow.Pow{Data: msg.Body}
	hash, err := job.Hash(config.Config.MaxNonce, config.Config.ZerosCount)
	if err != nil {
		return "", fmt.Errorf("do job: %w", err)
	}
	fmt.Println("do job")

	// 4 - request quote
	err = sendMessage(writer, &message.Message{Type: contract.RequestQuote, Body: hash})
	if err != nil {
		return "", fmt.Errorf("request task: %w", err)
	}
	fmt.Println("request quote")

	// 5 - get result
	msg, err = readMessage(reader)
	if err != nil {
		return "", fmt.Errorf("get result: %w", err)
	}
	if msg.Type != contract.ResponseQuote {
		return "", fmt.Errorf("received unknown get result message: %s", msg.String())
	}

	// 6 - return result
	return msg.Body, nil
}

func readMessage(reader *bufio.Reader) (*message.Message, error) {
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("read message: %w", err)
	}
	msg, err := message.ParseMessage(data)
	if err != nil {
		return nil, fmt.Errorf("parse message: %w", err)
	}
	return msg, nil
}

func sendMessage(conn io.Writer, msg *message.Message) error {
	response := fmt.Sprintf("%s\n", msg.String())
	_, err := conn.Write([]byte(response))
	if err != nil {
		return fmt.Errorf("write message: %w", err)
	}
	return nil
}
