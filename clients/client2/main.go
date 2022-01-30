package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"bitbucket.org/avd/go-ipc/fifo"
)

// Reads intergers from the pipe and calculate the Median value of the integers.
func main() {
	log.Println("Client2 is ready")

	for {
		conn, err := fifo.NewNamedPipe("pipe", os.O_CREATE|os.O_RDONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		buffer := make([]byte, 256)
		if n, err := conn.Read(buffer); err == nil {
			data := string(buffer[:n])

			numbers := make([]float64, 0)

			// 把資料轉成[]float64
			for _, v := range strings.Fields(data) {
				if s, err := strconv.ParseFloat(v, 64); err == nil {
					numbers = append(numbers, s)
				}
			}

			if len(numbers) > 0 {
				log.Println("Median is", Median(numbers))
			}
		} else {
			log.Println(err)
		}

		conn.Close()
	}
}

func Median(numbers []float64) float64 {
	// sort the numbers
	sort.Float64s(numbers)

	index := len(numbers) / 2

	if len(numbers)%2 == 0 {
		return (numbers[index-1] + numbers[index]) / 2
	}

	return numbers[index]
}
