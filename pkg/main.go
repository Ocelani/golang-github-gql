package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Node is the repository type
type Node struct {
	Repository struct {
		ID        string
		Name      string
		URL       string
		CreatedAt string
		UpdatedAt string
		Owner     struct {
			Login string
		}
		PrimaryLanguage struct {
			Name string
		}
		Stargazers struct {
			TotalCount int
		}
		IssuesTotal struct {
			TotalCount int
		} `graphql:"issuesTotal: issues"`
		IssuesClosed struct {
			TotalCount int
		} `graphql:"issuesClosed: issues(states: CLOSED)"`
		PullRequests struct {
			TotalCount int
		}
	} `graphql:" ... on Repository"`
}

func main() {
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	log.Println("⚡️Running...")
	run()
	log.Println("Done!✨")
}

func run() {
	// writeDataCsv(r)
	if err := runQuery(); err != nil {
		fmt.Printf("Error while running: %v", err)
	}
	return
}

func runQuery() (err error) {
	ctx := context.Background()
	// Auth token
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	// Use client...
	client := githubv4.NewClient(httpClient)
	{
		// q is the main query
		var q struct {
			Search struct {
				RepositoryCount githubv4.Int
				PageInfo        struct {
					EndCursor   githubv4.String
					HasNextPage githubv4.Boolean
				}
				Nodes []Node
			} `graphql:"search(query:$searchQuery, type:REPOSITORY, first:100, after:$afterCursor )"`
			RateLimit struct {
				Cost      githubv4.Int
				Limit     githubv4.Int
				Remaining githubv4.Int
				ResetAt   githubv4.DateTime
			}
		}
		// to set the query variables
		variables := map[string]interface{}{
			"searchQuery": githubv4.String("stars:>1000"),
			"afterCursor": (*githubv4.String)(nil),
		}
		// paginates the query
		var nodes []Node
		for {
			err := client.Query(ctx, &q, variables)
			if err != nil {
				fmt.Printf("Error on query: %v", err)
			}
			nodes = append(nodes, q.Search.Nodes...)
			if !q.Search.PageInfo.HasNextPage {
				break
			}
			variables["afterCursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
			fmt.Printf("\nCursor: %s", githubv4.String(q.Search.PageInfo.EndCursor))
			fmt.Println("\n.")
		}
		// finally
		writeJSON(nodes)
		writeCsv(nodes)
	}
	return
}

// writeJSON writes the JSON file
func writeJSON(v interface{}) {
	file, _ := os.OpenFile("./data.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	w := json.NewEncoder(file)
	w.SetIndent("", "\t")
	err := w.Encode(v)

	if err != nil {
		panic(err)
	}
	fmt.Println("json done!")
}

// writeCsv writes the file.csv
func writeCsv(v interface{}) {
	// Marshall JSON data from response variable
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal JSON data
	var n []Node
	err = json.Unmarshal(data, &n)
	if err != nil {
		fmt.Println(err)
	}

	csvdatafile, err := os.OpenFile("./data.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer csvdatafile.Close()

	writer := csv.NewWriter(csvdatafile)

	for i, node := range n {
		var record []string
		record = append(record, strconv.Itoa(i+1))
		record = append(record, string(node.Repository.ID))
		record = append(record, string(node.Repository.Name))
		record = append(record, string(node.Repository.URL))
		record = append(record, string(node.Repository.CreatedAt))
		record = append(record, string(node.Repository.UpdatedAt))
		record = append(record, string(node.Repository.Owner.Login))
		record = append(record, string(node.Repository.PrimaryLanguage.Name))
		record = append(record, strconv.Itoa(node.Repository.Stargazers.TotalCount))
		record = append(record, strconv.Itoa(node.Repository.IssuesTotal.TotalCount))
		record = append(record, strconv.Itoa(node.Repository.IssuesClosed.TotalCount))
		record = append(record, strconv.Itoa(node.Repository.PullRequests.TotalCount))
		writer.Write(record)
		i++
	}
	writer.Flush()
	fmt.Println("csv done!")
}
