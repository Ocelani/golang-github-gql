package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// PageInfo gql
type PageInfo struct {
	EndCursor   string
	HasNextPage bool
}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var after string
	query(after)
	resp := runQuery()
	writeResult(resp)
	nextPage, cursor := paginate()

	pageInfo := PageInfo{
		EndCursor:   cursor,
		HasNextPage: nextPage,
	}
}

func runQuery() io.ReadCloser {
	// Prepares the json stub to send
	query, err := json.Marshal(query)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(bytes.NewBuffer(query))

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
	return resp.Body
}

func writeResult(r io.ReadCloser) []byte {
	var buffer bytes.Buffer
	// Collects the data
	data, _ := ioutil.ReadAll(r)
	// Writes into a json file
	json.Indent(&buffer, data, "", "\t")
	err := ioutil.WriteFile("out/repos.json", buffer.Bytes(), 0777)
	if err != nil {
		log.Println("Error trying to write file.\n[ERRO] -", err)
	}
	log.Println("Page done!✨")
	return data
}

func paginate() (bool, string) {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}
	var dat map[string]interface{}
	if err := json.Unmarshal(content, &dat); err != nil {
		panic(err)
	}
	p := dat["hasNextPage"].(bool)
	c := dat["endCursor"].(string)
	return p, c
}

// Query repares the json stub to send
func query(after string) map[string]string {
	query := map[string]string{
		"query": fmt.Sprintf(`
  {
    search(query: "stars:>100", type: REPOSITORY, first: 10, %s) {
      repositoryCount
      pageInfo {
        hasNextPage
        endCursor
      }
      nodes {
        ... on Repository {
          id
          name
          owner{login}
          primaryLanguage{name}
          url
          stargazers{totalCount}
          issues{totalCount}
          pullRequests{totalCount}
          createdAt
          updatedAt
        }
      }
    }
  }
  `, after),
	}
	return query
}

/*
total_pages = 1

result = run_query(json, headers)

nodes = result['data']['search']['nodes']
next_page  = result["data"]["search"]["pageInfo"]["hasNextPage"]

#paginating
while (next_page and total_pages < 3):
    // total_pages += 1

    // cursor = result["data"]["search"]["pageInfo"]["endCursor"]

    next_query = query.replace("{AFTER}", ", after: \"%s\"" % cursor)

    json["query"] = next_query

    result = run_query(json, headers)

    nodes += result['data']['search']['nodes']

    next_page  = result["data"]["search"]["pageInfo"]["hasNextPage"]
*/
//--------------------------------------
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
