package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Bialek328/jobs-data-pipeline/db"
	"github.com/gocolly/colly"
    "github.com/lib/pq"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

type JobOffer struct {
	Url       string `json:"url"`
	Date      interface{} `json:"date"`
	TechStack []string    `json:"techStack"`
	MiscInfo  []string    `json:"info"`
	Salary    string      `json:"salary"`
}

func InsertJobListingIntoDB(db *sql.DB, job JobOffer) error {
	insertQuery := `INSERT INTO job_listings (url, tech_stack, misc_info, salary)
                    VALUES ($1, $2, $3, $4);`
	_, err := db.Exec(insertQuery, string(job.Url), pq.Array(job.TechStack), pq.Array(job.MiscInfo), job.Salary)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dbInstance, err := db.InitDB()
    defer dbInstance.Close()
	if err != nil {
		log.Fatal(err)
        return 
	}
    defer timer("main")()
	var techStack []string
	var info []string
    var salary string

	log.Println("Starting collector")
	c := colly.NewCollector()

	jobCollector := colly.NewCollector(colly.Async(true))

	c.OnHTML("div.MuiBox-root.css-ggjav7 a", func(e *colly.HTMLElement) {
		inner := e.Attr("href")
		if inner != "" && strings.Contains(inner, "/job-offer/") {
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

	jobCollector.OnHTML("div.MuiBox-root.css-1km0bek span.css-1pavfqb", func(e *colly.HTMLElement) {
		str := []string{}
		e.ForEach("span", func(_ int, elele *colly.HTMLElement) {
			str = append(str, e.Text)
		})
		salary = str[0]
		log.Println(salary)
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
			Url:       r.Request.URL.String(),
			Date:      time.Now().Format("2006.01.02"),
			TechStack: techStack,
            Salary: salary,
			MiscInfo:  info,
		}
        log.Println(job)
		err := InsertJobListingIntoDB(dbInstance, job)
		if err != nil {
			log.Fatal(err)
            return
		}
		techStack = []string{}
		info = []string{}
        salary = ""
	})

	c.Visit("https://justjoin.it/job-offers/all-locations/python?remote=yes")

}
