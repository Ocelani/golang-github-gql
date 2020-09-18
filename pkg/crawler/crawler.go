// package main

// import (
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// type pageInfo struct {
// 	StatusCode int
// 	Links      map[string]int
// }

// type RepoLink struct {
// 	Name string
// 	Link string
// }

// var w http.ResponseWriter

// func scrape(jobs chan string, backs chan<- string) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			println("Recovered for", interface2string(r))
// 			jobs <- interface2string(r)
// 			go scrape(jobs, backs)
// 		}
// 	}()
// 	for p := range jobs {
// 		var doc *goquery.Document
// 		var e error
// 		result := "\n#### " + "\n"

// 		if doc, e = goquery.NewDocument("https://github.com/search?p=" + p + "&q=stars%3A%3E100000&type=Repositories"); e != nil {
// 			println("Error:", e.Error())
// 		}

// 		doc.Find("a.v-align-middle").Each(func(i int, s *goquery.Selection) {
// 			title := s.Text()
// 			url := "https://github.com/" + title
// 			var stars = "0"
// 			var forks = "0"
// 			s.Find("a.muted-link.mr-3").Each(func(i int, contentSelection *goquery.Selection) {
// 				if temp, ok := contentSelection.Find("svg").Attr("aria-label"); ok {
// 					switch temp {
// 					case "star":
// 						stars = contentSelection.Text()
// 					case "fork":
// 						forks = contentSelection.Text()
// 					}
// 				}
// 			})
// 			result = result + "* [" + title + " (" + strings.TrimSpace(stars) + "s/" + strings.TrimSpace(forks) + "f)](" + url + ") : " + "\n"
// 		})
// 		backs <- result
// 	}
// }

// //interface to string
// func interface2string(inter interface{}) string {
// 	var tempStr string
// 	switch inter.(type) {
// 	case string:
// 		tempStr = inter.(string)
// 		break
// 	case float64:
// 		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
// 		break
// 	case int64:
// 		tempStr = strconv.FormatInt(inter.(int64), 10)
// 		break
// 	case int:
// 		tempStr = strconv.Itoa(inter.(int))
// 		break
// 	}
// 	return tempStr
// }

// func main() {
// 	scrape()
// }
