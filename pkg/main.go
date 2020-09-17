package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// DataFrame is the query result type
type DataFrame struct {
	Repos []*Node
}

// Node is the repository type
type Node struct {
	RepoName           string `json:"name"`
	StarsNumber        int64  `json:"stargazers"`
	CreatedAt          string `json:"createdAt"`
	PullRequestsNumber int64  `json:"pullRequests"`
	// ProgLanguage       string    `json:"primaryLanguage{name}"`
}

func run() {
	// Prepares the json stub to send
	gql := map[string]string{
		"query": `
  {
    search(query: "stars:>100", type: REPOSITORY, first: 10) {
      repositoryCount
      pageInfo {
        hasNextPage
        endCursor
      }
      nodes {
        ... on Repository {
          name
          owner{login}
          stargazers{totalCount}
          url
          primaryLanguage{name}
          issues{totalCount}
          issues(states: CLOSED){totalCount}
          createdAt
          updatedAt
          issues{totalCount}
          issues(states: CLOSED){totalCount}
          pullRequests(states: MERGED, first: 10) {
            totalCount
            pageInfo {
              hasNextPage
              endCursor
            }
          }
        }
      }
    }
  }
  `,
	}

	query, err := json.Marshal(gql)
	if err != nil {
		log.Println(err)
	}

	// Bearer token auth header
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")

	// Creates the request
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(query))
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}

	// Receives resp from the req
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	// Closes the response when it has ended
	defer resp.Body.Close()

	// Collects the data
	data, _ := ioutil.ReadAll(resp.Body)

	// Indent
	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")

	// Writes into a json file
	err = ioutil.WriteFile("out/repos.json", out.Bytes(), 0777)
	if err != nil {
		log.Println("Error trying to write file.\n[ERRO] -", err)
	}

}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	run()
}

/*
func writeCsv() {
VVVVVVVVVVVVVVVV
err = json.Unmarshal([]byte(data), &node)
reading data from JSON File
read, err := ioutil.ReadFile("../out/repos.json")
if err != nil {
	fmt.Println(err)
}
var node []Node
// Unmarshal JSON data
err = json.Unmarshal([]byte(read), &node)
if err != nil {
	fmt.Println(err)
}
// Create a csv file
csvFile, err := os.Create("./data.csv")
if err != nil {
	fmt.Println(err)
}
defer csvFile.Close()
// Write Unmarshaled json data to CSV file
writer := csv.NewWriter(csvFile)
for _, obj := range node {
	var row []string
	row = append(row, obj.RepoName)
	row = append(row, strconv.FormatInt(obj.StarsNumber, 10))
	row = append(row, obj.CreatedAt)
	row = append(row, strconv.FormatInt(obj.PullRequestsNumber, 10))
	writer.Write(row)
}
writer.Flush()
}
*/
