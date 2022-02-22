package tcp

import (
	"bufio"
	"fmt"
	"net"

	"gorm.io/gorm"
)

// TCP Server
func Start(ip string, port string, sql *gorm.DB) error {
	// Listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return err
	}

	// Close
	defer listener.Close()

	// Accept new connections
	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}

		go processData(connection, sql)
	}
}

// Process data
func processData(connection net.Conn, sql *gorm.DB) error {
	reader := bufio.NewReader(connection)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			connection.Close()
			return err
		}

		// Handle data
		connection.Write([]byte(Handler(sql, string(data))))
	}
}
