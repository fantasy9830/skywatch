package main

import (
	"bufio"
	"log"
	"net"
	"os"

	"bitbucket.org/avd/go-ipc/fifo"
	"bitbucket.org/avd/go-ipc/mmf"
	"bitbucket.org/avd/go-ipc/shm"
)

func main() {
	log.Println("Server is ready. You can type intergers and then click [ENTER].  Clients will show the mean, median, and mode of the input values.")
	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		Client1(data)
		Client2(data)
		Client3(data)
	}
}

// Client1 Client1
func Client1(data string) error {
	network := "tcp"
	tcpAddr, err := net.ResolveTCPAddr(network, ":8888")
	if err != nil {
		return err
	}

	conn, err := net.DialTCP(network, nil, tcpAddr)
	if err != nil {
		return err
	}
	conn.Write([]byte(data))
	conn.Close()

	return nil
}

// Client2 Client2
func Client2(data string) error {
	obj, err := fifo.New("pipe", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer obj.Close()

	obj.Write([]byte(data))

	return nil
}

// Client3 Client3
func Client3(data string) error {
	object, err := shm.NewWindowsNativeMemoryObject("shm", os.O_CREATE|os.O_RDWR, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer object.Close()

	region, err := mmf.NewMemoryRegion(object, mmf.MEM_READWRITE, 0, 1024)
	if err != nil {
		log.Fatal(err)
	}
	defer region.Close()

	copy(region.Data(), []byte(data))

	return nil
}
