package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"
	"unicode"
)

type Device struct {
	RetailBranding string `json:"retail_branding",csv:"Retail Branding"`
	MarketingName  string `json:",marketing_name",csv:"Marketing Name"`
	Device         string `json:"device",csv:"Device"`
	Model          string `json:"model",csv:"Model"`
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
	n := 0

	for scanner.Scan() {
		line := scanner.Text()

		if err := UnmarshalDevice(line, &device); err != nil {
			// log.Println("line:", line)
			// log.Fatal("Unmarshal err:", err)
			continue
		}
		n++
		// log.Printf("device[%d]: %+v\n", n, device)

		j, err := json.Marshal(device)
		if err != nil {
			// log.Println("json.Marshal err:", err)
		} else {
			// log.Println("line:", line)
			log.Println("json:", string(j))
		}
	}

}

func UnmarshalDevice(input string, d *Device) error {

	vals := strings.Split(input, ",")

	if len(vals) == 4 {
		d.RetailBranding = Normalize(vals[0])
		d.MarketingName = Normalize(vals[1])
		d.Device = Normalize(vals[2])
		d.Model = Normalize(vals[3])
		return nil
	}

	// line := strings.ReplaceAll(input, `,"`, "|")
	// line = strings.ReplaceAll(line, `",`, "|")
	// line = strings.ReplaceAll(line, `""`, `"`)

	return errors.New("parsing error")
}

func Normalize(s string) string {
	v := strings.ReplaceAll(s, `""`, `"`)
	// remove starting and ending quotes
	return Sanitize(v)
}

func Sanitize(s string) string {
	sanitizedString := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, s)
	return sanitizedString
}
