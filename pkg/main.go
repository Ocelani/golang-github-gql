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
		Name           string
		CreatedAt      string
		UpdatedAt      string
		StargazerCount int
		ForkCount      int
		Owner          struct {
			Login string
		}
		PrimaryLanguage struct {
			Name string
		}
		Releases struct {
			TotalCount int
		} `graphql:"releases(orderBy: {field:CREATED_AT, direction:DESC}, first:100)"`
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
	c := make(chan bool, 2)
	c <- runQuery("Python stars:>1000", "python")
	c <- runQuery("Java stars:>1000", "java")
	for q := range c {
		if q != true {
			fmt.Println("Task not finished on channel: ", c)
		}
	}
	return
}

func runQuery(search string, file string) bool {
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
			} `graphql:"search(query:$what, type:REPOSITORY, first:100)"`
			RateLimit struct {
				Cost      githubv4.Int
				Limit     githubv4.Int
				Remaining githubv4.Int
				ResetAt   githubv4.DateTime
			}
		}
		// to set the query variables
		variables := map[string]interface{}{
			"what": githubv4.String(search),
			// "afterCursor": (*githubv4.String)(nil),
		}
		// paginates the query
		var nodes []Node
		// for {
		err := client.Query(ctx, &q, variables)
		if err != nil {
			fmt.Printf("Error on query: %v", err)
		}
		nodes = append(nodes, q.Search.Nodes...)
		// if !q.Search.PageInfo.HasNextPage {
		// 	break
		// }
		// variables["afterCursor"] = githubv4.NewString(q.Search.PageInfo.EndCursor)
		fmt.Printf("\nCursor: %s", githubv4.String(q.Search.PageInfo.EndCursor))
		fmt.Println("\n.")
		// }
		// finally
		writeJSON(nodes, file)
		writeCsv(nodes, file)
		return true
	}
}

// writeJSON writes the JSON file
func writeJSON(v interface{}, file string) {
	f, err := os.OpenFile("./"+file+".json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := json.NewEncoder(f)
	w.SetIndent("", "\t")
	err = w.Encode(v)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("json done!")
	}
}

// writeCsv writes the file.csv
func writeCsv(v interface{}, file string) {
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

	csvfile, err := os.OpenFile("./"+file+".csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)

	for i, node := range n {
		var record []string
		record = append(record, strconv.Itoa(i+1))
		record = append(record, string(node.Repository.Name))
		record = append(record, string(node.Repository.CreatedAt))
		record = append(record, string(node.Repository.UpdatedAt))
		record = append(record, string(node.Repository.StargazerCount))
		record = append(record, string(node.Repository.ForkCount))
		record = append(record, string(node.Repository.Owner.Login))
		record = append(record, string(node.Repository.PrimaryLanguage.Name))
		record = append(record, string(node.Repository.Releases.TotalCount))
		writer.Write(record)
		i++
	}
	writer.Flush()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("csv done!")
	}
}
