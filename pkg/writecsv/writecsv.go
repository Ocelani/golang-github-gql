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
	// read data from file
	jsondatafromfile, err := ioutil.ReadFile("./out/data.json")
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal JSON data
	var jsondata utils.DataFrame
	err = json.Unmarshal([]byte(jsondatafromfile), &jsondata)
	if err != nil {
		fmt.Println(err)
	}

	csvdatafile, err := os.Create("./datatwo.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)

	for _, repo := range jsondata.Search.Nodes {
		var record []string
		record = append(record, repo.Repository.ID)
		record = append(record, repo.Repository.Name)
		record = append(record, repo.Repository.URL)
		record = append(record, repo.Repository.CreatedAt)
		record = append(record, repo.Repository.UpdatedAt)
		record = append(record, repo.Repository.Owner.Login)
		record = append(record, repo.Repository.PrimaryLanguage.Name)
		record = append(record, strconv.Itoa(repo.Repository.Stargazers.TotalCount))
		record = append(record, strconv.Itoa(repo.Repository.IssuesTotal.TotalCount))
		record = append(record, strconv.Itoa(repo.Repository.IssuesClosed.TotalCount))
		record = append(record, strconv.Itoa(repo.Repository.PullRequests.TotalCount))
		writer.Write(record)
	}

	// remember to flush!
	writer.Flush()
}
