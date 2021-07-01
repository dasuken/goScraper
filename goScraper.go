package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
)

func main() {

	fName := "data.csv"

	file, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
		)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement){
		writer.Write([]string {
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	for i := 0; i < 10; i ++ {
		fmt.Printf("Scraping Page: %d\n", i)

		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete \n")
	log.Println(c)

}