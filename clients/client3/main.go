package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/avd/go-ipc/mmf"
	"bitbucket.org/avd/go-ipc/shm"
)

// Reads intergers from the shared memory, and then calculating the Mode value of the integers.
func main() {
	log.Println("Client3 is ready")

	object, err := shm.NewWindowsNativeMemoryObject("shm", os.O_CREATE|os.O_RDWR, 1024)
	if err != nil {
		log.Fatal(err)
	}

	for {
		region, err := mmf.NewMemoryRegion(object, mmf.MEM_READWRITE, 0, 1024)
		if err != nil {
			log.Fatal(err)
		}

		numbers := make([]float64, 0)

		// 把資料轉成[]float64
		for _, v := range strings.Fields(string(region.Data())) {
			if s, err := strconv.ParseFloat(v, 64); err == nil {
				numbers = append(numbers, s)
			}
		}

		region.Close()

		if len(numbers) > 0 {
			log.Println("Median is", Mode(numbers))
			object.Close()
			object, err = shm.NewWindowsNativeMemoryObject("shm", os.O_CREATE|os.O_RDWR, 1024)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Mode Mode
func Mode(numbers []float64) []float64 {
	m := map[float64]int{}
	var max int

	for _, num := range numbers {
		m[num]++
		if m[num] > max {
			max = m[num]
		}
	}

	var freq []float64
	for num, v := range m {
		if max == v {
			freq = append(freq, num)
		}
	}

	return freq
}
