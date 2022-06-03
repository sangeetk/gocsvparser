package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Device struct {
	RetailBranding string `csv:"Retail Branding"`
	MarketingName  string `csv:"Marketing Name"`
	Device         string `csv:"Device"`
	Model          string `csv:"Model"`
}

func main() {

	// Open the CSV file
	file, err := os.Open("supported_devices.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Read one line at a time
	device := Device{}
	for scanner.Scan() {
		line := scanner.Text()

		if err := UnmarshalDevice(line, &device); err != nil {
			log.Println("line:", line)
			log.Fatal("Unmarshal err:", err)
		}

		j, err := json.Marshal(device)
		if err != nil {
			log.Fatal("json.Marshal err:", err)
		}

		log.Println("device:", string(j))
	}

}

func UnmarshalDevice(line string, device *Device) error {
	return nil
}
