package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type JobOffer struct {
    Url interface{} `json:"url"`
    Date interface{} `json:"date"`
    TechStack []string `json:"techStack"`
    Info []string `json:"info"`
}
func timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

func main() {
    defer timer("main")()
    var techStack []string
    var info []string

    log.Println("Starting collector")
    c := colly.NewCollector()

    jobCollector := colly.NewCollector(colly.Async(true))


    c.OnHTML("div.MuiBox-root a", func(e *colly.HTMLElement) {
        inner := e.Attr("href")
        if inner != "" && strings.Contains(inner, "/offers/"){
            jobCollector.Visit(e.Request.AbsoluteURL(inner))
            jobCollector.Wait()
        }
    })

    jobCollector.OnHTML("div.MuiBox-root.css-jfr3nf", func(e *colly.HTMLElement) {
        e.ForEach("h4.MuiTypography-root.MuiTypography-subtitle2.css-7de92v", func(_ int, el *colly.HTMLElement) {
            // log.Println(el.Text)
            techStack = append(techStack, el.Text)
        })
    })
    
    jobCollector.OnHTML("div.MuiBox-root.css-r1n8l8", func(en *colly.HTMLElement) {
        en.ForEach("ul li", func(_ int, ele *colly.HTMLElement) {
            info = append(info, ele.Text)
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
            Info: info,
        }
        log.Println(job)
        techStack = []string{}
        info = []string{}
    })
    
    c.Visit("https://justjoin.it/trojmiasto/python")
}
