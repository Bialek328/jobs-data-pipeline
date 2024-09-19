package main

import (
    "log"
	"github.com/gocolly/colly"
)

func main() {
    log.Println("Starting collector")
    c := colly.NewCollector()

    c.OnHTML("div.MuiBox-root", func(e *colly.HTMLElement) {
        // inner := e.ChildText(".MuiTypography-root")
        inner := e.Text
        log.Println("content:", inner)
    })

    c.Visit("https://justjoin.it/trojmiasto")
}
