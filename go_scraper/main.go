package main

import (
	"log"
	"strings"
    "time"

	"github.com/gocolly/colly"
)

type JobOffer struct {
    Url interface{}
    Date interface{}
    TechStack []string
    Experience string
}

func main() {
    var techStack []string
    log.Println("Starting collector")
    c := colly.NewCollector()

    jobCollector := colly.NewCollector()


    c.OnHTML("div.MuiBox-root a", func(e *colly.HTMLElement) {
        inner := e.Attr("href")
        if inner != "" && strings.Contains(inner, "/offers/"){
            jobCollector.Visit(e.Request.AbsoluteURL(inner))
        }
    })

    jobCollector.OnHTML("div.MuiBox-root.css-jfr3nf", func(e *colly.HTMLElement) {
        e.ForEach("h4.MuiTypography-root.MuiTypography-subtitle2.css-7de92v", func(_ int, el *colly.HTMLElement) {
            // log.Println(el.Text)
            techStack = append(techStack, el.Text)
        })
    })
    
    jobCollector.OnRequest(func(r *colly.Request) {
        log.Println("visiting:", r.URL)
    })      

    jobCollector.OnScraped(func(r *colly.Response) {
        job := JobOffer{
            Url: r.Request.URL,
            Date: time.Now().Format("2006.01.02"), 
            TechStack: techStack,
        }
        log.Println(job)
        techStack = []string{}
    })
    
    c.Visit("https://justjoin.it/trojmiasto/python")
}
