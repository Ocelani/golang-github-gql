package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Ocelani/medexp/pkg/utils"
)

func main() {
	// var df utils.Query
	// data, _ := ioutil.ReadFile("out/data.json")
	// // Unmarshal JSON data
	// json.Unmarshal([]byte(data), &df)
	// // Create a csv file
	// file, err := os.OpenFile("out/data.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // Write Unmarshaled json data to CSV file
	// w := csv.NewWriter(file)
	// // Set csv headers
	// defer file.Close()
	// read data from file
	jsonDataFromFile, err := ioutil.ReadFile("out/data.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal JSON data
	var df utils.DataFrame
	err = json.Unmarshal([]byte(jsonDataFromFile), &df)
	if err != nil {
		fmt.Println(err)
	}

	csvFile, err := os.Create("out/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	// Appends the data to the csv
	for _, repo := range df.Search.Nodes {
		var row []string
		row = append(row, repo.ID, repo.Owner.Login, repo.Name, repo.URL, repo.ProgLanguage.Name, strconv.Itoa(repo.StarsNumber.TotalCount), strconv.Itoa(repo.PullRequestsNumber.TotalCount), strconv.Itoa(repo.IssuesTotal.TotalCount), strconv.Itoa(repo.IssuesClosed.TotalCount), repo.CreatedAt, repo.UpdatedAt)
		fmt.Println()
		w.Write(row)
	}
	w.Flush()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data written at -> out/data.csv")
	return
}
