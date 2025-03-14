package sdk

import (
	"log"
	"net"
)

// GetHostFreePort finds an available port on the system by letting the OS assign an available port.
// It returns the port number and an error if any occurred during the process.
// The function creates a TCP listener on port 0, which tells the OS to find an available port.
// Once the port is assigned, the listener is closed and the port number is returned.
func GetHostFreePort() (int, error) {
	log.Println("Finding an available port on the system")
	listener, err := net.Listen("tcp", ":0") // Let OS assign an available port
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}
