package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=10"

func main() {
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()
	// fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobswithGoRoutines(jobs)
	// writeJobs(jobs)
	// fmt.Println(jobs)
}

// write csg
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjobs?jk=" + job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func writeJobswithGoRoutines(jobs []extractedJob) {
	wCsv, err := ccsv.NewCsvWriter("sample.csv")
	checkErr(err)

	defer wCsv.Close()

	wCsv.Write([]string{"Link", "Title", "Location", "Salary", "Summary"})
	done := make(chan bool)
	for _, job := range jobs {
		go func(job extractedJob) {
			wCsv.Write([]string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.location})
			done
		}(job)
	}
	for i := 0; i < len(jobs); i++ {
	}
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&satar=" + strconv.Itoa(page*10)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		//jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location}
	//fmt.Println(id, title, location)
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination-list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
