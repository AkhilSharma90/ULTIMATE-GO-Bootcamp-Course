package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	listenAddr   = ":9000"            // Server address
	serverSecret = "mysecretkey12345" // 16-byte AES key
)

// Encrypt data using AES
func encrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}

// Decrypt data using AES
func decrypt(data []byte, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plain := make([]byte, len(ciphertext))
	stream.XORKeyStream(plain, ciphertext)

	return plain, nil
}

// Server listens for client connections
func startServer() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
	defer listener.Close()

	fmt.Println("VPN server is listening on", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v\n", err)
			continue
		}

		go handleServerConnection(conn)
	}
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Client connected: %s\n", conn.RemoteAddr())

	for {
		var length uint32
		if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
			fmt.Println("Connection closed:", err)
			return
		}

		encryptedData := make([]byte, length)
		if _, err := io.ReadFull(conn, encryptedData); err != nil {
			fmt.Println("Failed to read data:", err)
			return
		}

		data, err := decrypt(encryptedData, serverSecret)
		if err != nil {
			fmt.Println("Failed to decrypt data:", err)
			return
		}

		fmt.Printf("Received: %s\n", string(data))
	}
}

// Client connects to the server and sends data
func startClient(serverAddr string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	defer conn.Close()

	fmt.Println("Connected to VPN server")

	for {
		fmt.Print("Enter message to send: ")
		var input string
		fmt.Scanln(&input)

		encryptedData, err := encrypt([]byte(input), serverSecret)
		if err != nil {
			fmt.Println("Failed to encrypt data:", err)
			continue
		}

		length := uint32(len(encryptedData))
		if err := binary.Write(conn, binary.BigEndian, length); err != nil {
			fmt.Println("Failed to write data length:", err)
			continue
		}

		if _, err := conn.Write(encryptedData); err != nil {
			fmt.Println("Failed to send data:", err)
			continue
		}

		fmt.Println("Message sent!")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vpn <server|client> [server address]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "server":
		startServer()
	case "client":
		if len(os.Args) < 3 {
			fmt.Println("Usage: vpn client <server address>")
			os.Exit(1)
		}
		startClient(os.Args[2])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

//to run this program - go run VPN.go server
//and in another terminal - go run VPN.go client 127.0.0.1:9000
