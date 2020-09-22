package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

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
		// Node is the repository type
		type Node struct {
			Repository struct {
				ID        githubv4.ID
				Name      githubv4.String
				URL       githubv4.URI
				CreatedAt githubv4.DateTime
				UpdatedAt githubv4.DateTime
				Owner     struct {
					Login githubv4.String
				}
				PrimaryLanguage struct {
					Name githubv4.String
				}
				Stargazers struct {
					TotalCount githubv4.Int
				}
				IssuesTotal struct {
					TotalCount githubv4.Int
				} `graphql:"issuesTotal: issues"`
				IssuesClosed struct {
					TotalCount githubv4.Int
				} `graphql:"issuesClosed: issues(states: CLOSED)"`
				PullRequests struct {
					TotalCount githubv4.Int
				}
			} `graphql:" ... on Repository"`
		}
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

		variables := map[string]interface{}{
			"searchQuery": githubv4.String("stars:>10000"),
			"afterCursor": (*githubv4.String)(nil),
		}

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
		printJSON(q)
	}
	return
}

// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
func printJSON(v interface{}) {
	file, _ := os.OpenFile("out/second.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	w := json.NewEncoder(file)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}
