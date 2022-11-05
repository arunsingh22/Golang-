package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type data struct {
	before int64
	after  int64
	RTT    int64
}

var records = make([]data, 0, 500)

func main() {
	//dumping into a file.
	// curr := time.Now().Unix()

	//writes to CSV file
	//writes to CSV file
	fileName := "TIMESTAMPS:" + time.Now().Format("01-02-2006_15:04:05") + ".csv"
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	tk := time.NewTicker(1 * time.Second)

	i := 0
	for range tk.C {
		//fetchig data from the channel and writing to the csv file.
		i++

		if i >= 5 {
			tk.Stop()
			break
		}
		fmt.Println("Writing into CSV..")
		for i := 0; i < 1; i++ {
			row := []string{"message1", time.Now().Format("01-02-2006_15:04:05"), time.Now().Format("01-02-2006_15:04:05"), strconv.Itoa(23525)}
			w.Write(row)
		}
	}

	// for i := 0; i < 10; i++ {
	// 	records = append(records, data{curr, curr, curr})
	// }

	// for _, record := range records {
	//     row := []string{record.ID, strconv.Itoa(record.Age)}
	//     data = append(data, row)
	// }
	// w.WriteAll(data)

}
