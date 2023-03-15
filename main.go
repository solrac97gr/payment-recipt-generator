package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron"
	"github.com/solrac97gr/payment-recipt-generator/factories"
)

func main() {
	configApplication := factories.NewConfigFileApplication("config.json")
	config, err := configApplication.GetConfig()
	if err != nil {
		panic(err)
	}

	_PDFGenerator := factories.NewPDFGeneratorApp("./out/", config)

	c := cron.New()

	// Schedule the job to run every 15th at 1:00 AM
	c.AddFunc("0 0 1 15 * *", func() {
		fmt.Println("Running job on 15th of the month at 1:00 AM")
		err = _PDFGenerator.GenerateInvoicePDF()
		if err != nil {
			panic(err)
		}
	})

	// Schedule the job to run at the end of the month at 1:00 AM
	c.AddFunc("0 0 1 * * *", func() {
		// Check if it's the last day of the month
		t := time.Now()
		lastDay := t.AddDate(0, 1, 0).Add(-time.Hour * 24)
		if t.Day() == lastDay.Day() {
			fmt.Println("Running job at the end of the month at 1:00 AM")
			err = _PDFGenerator.GenerateInvoicePDF()
			if err != nil {
				panic(err)
			}
		}
	})

	c.Start()

	select {}
}
