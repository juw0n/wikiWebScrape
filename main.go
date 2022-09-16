package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// getting the page url
	url := "https://www.flipkart.com/exercise-fitness/fitness-accessories/pr?sid=qoc%2Cacb&p%5B%5D=facets.fulfilled_by%255B%255D%3DPlus%2B%2528FAssured%2529&hpid=wXxDXjUDi0JocvmBiNvc9Kp7_Hsxr70nj65vMAAFKlc%3D&ctx=eyJjYXJkQ29udGV4dCI6eyJhdHRyaWJ1dGVzIjp7InZhbHVlQ2FsbG91dCI6eyJtdWx0aVZhbHVlZEF0dHJpYnV0ZSI6eyJrZXkiOiJ2YWx1ZUNhbGxvdXQiLCJpbmZlcmVuY2VUeXBlIjoiVkFMVUVfQ0FMTE9VVCIsInZhbHVlcyI6WyJGcm9tIOKCuTEzOSJdLCJ2YWx1ZVR5cGUiOiJNVUxUSV9WQUxVRUQifX0sImhlcm9QaWQiOnsic2luZ2xlVmFsdWVBdHRyaWJ1dGUiOnsia2V5IjoiaGVyb1BpZCIsImluZmVyZW5jZVR5cGUiOiJQSUQiLCJ2YWx1ZSI6IlJUQkcyWUFRUzlIRTZNRkciLCJ2YWx1ZVR5cGUiOiJTSU5HTEVfVkFMVUVEIn19LCJ0aXRsZSI6eyJtdWx0aVZhbHVlZEF0dHJpYnV0ZSI6eyJrZXkiOiJ0aXRsZSIsImluZmVyZW5jZVR5cGUiOiJUSVRMRSIsInZhbHVlcyI6WyJHeW0gRXNzZW50aWFscyJdLCJ2YWx1ZVR5cGUiOiJNVUxUSV9WQUxVRUQifX19fX0%3D&fm=neo%2Fmerchandising&iid=M_37a4c345-e213-4c4b-abe1-a85b59705056_6.I6MRHUZZIZNC&ppt=None&ppn=None&ssid=n30z80ixb40000001663368218654&otracker=hp_omu_Beauty%252C%2BFood%252C%2BToys%2B%2526%2Bmore_6_6.dealCard.OMU_I6MRHUZZIZNC_5&otracker1=hp_omu_PINNED_neo%2Fmerchandising_Beauty%252C%2BFood%252C%2BToys%2B%2526%2Bmore_NA_dealCard_cc_6_NA_view-all_5&cid=I6MRHUZZIZNC"

	resp := getHTML(url)
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	getLinks(doc)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func getHTML(url string) *http.Response {
	resp, err := http.Get(url)
	checkErr(err)

	if resp.StatusCode > 400 {
		fmt.Println("status code", resp.StatusCode)
	}

	return resp
}

func getLinks(doc *goquery.Document) {
	doc.Find("div.mw-body-content").Each(func(i int, item *goquery.Selection) {
		a := item.Find("")
	})
}
