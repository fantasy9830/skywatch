package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// Reads intergers from the socket and calculate the Mean value of the integers.
func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Client1 is ready")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleClient(conn)
	}

}

// Mean Mean
func Mean(numbers []float64) float64 {
	var total float64
	for _, v := range numbers {
		total += v
	}

	return total / float64(len(numbers))
}

func handleClient(conn net.Conn) {
	buffer := make([]byte, 256)
	defer conn.Close()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}

		data := string(buffer[:n])

		numbers := make([]float64, 0)

		// 把資料轉成[]float64
		for _, v := range strings.Fields(data) {
			if s, err := strconv.ParseFloat(v, 64); err == nil {
				numbers = append(numbers, s)
			}
		}

		if len(numbers) > 0 {
			log.Println("Mean is", Mean(numbers))
		}
	}
}
