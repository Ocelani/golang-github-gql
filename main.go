package main

import (
	"context"
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

	err := run()
	if err != nil {
		log.Println(err)
	}
}

var query struct {
	Viewer struct {
		Login     githubv4.String
		CreatedAt githubv4.DateTime
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	{
		err := client.Query(context.Background(), &query, nil)
		if err != nil {
			log.Fatal("Error on query!")
		}
		fmt.Println("    Login:", query.Viewer.Login)
		fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

		/* 		type githubV4Actor struct {
		   			Login     githubv4.String
		   			AvatarURL githubv4.URI `graphql:"avatarUrl(size:72)"`
		   			URL       githubv4.URI
		   		}

		   		var q struct {
		   			Repository struct {
		   				DatabaseID githubv4.Int
		   				URL        githubv4.URI

		   				Issue struct {
		   					Author         githubV4Actor
		   					PublishedAt    githubv4.DateTime
		   					LastEditedAt   *githubv4.DateTime
		   					Editor         *githubV4Actor
		   					Body           githubv4.String
		   					ReactionGroups []struct {
		   						Content githubv4.ReactionContent
		   						Users   struct {
		   							Nodes []struct {
		   								Login githubv4.String
		   							}

		   							TotalCount githubv4.Int
		   						} `graphql:"users(first:10)"`
		   						ViewerHasReacted githubv4.Boolean
		   					}
		   					ViewerCanUpdate githubv4.Boolean

		   					Comments struct {
		   						Nodes []struct {
		   							Body   githubv4.String
		   							Author struct {
		   								Login githubv4.String
		   							}
		   							Editor struct {
		   								Login githubv4.String
		   							}
		   						}
		   						PageInfo struct {
		   							EndCursor   githubv4.String
		   							HasNextPage githubv4.Boolean
		   						}
		   					} `graphql:"comments(first:$commentsFirst,after:$commentsAfter)"`
		   				} `graphql:"issue(number:$issueNumber)"`
		   			} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
		   			Viewer struct {
		   				Login      githubv4.String
		   				CreatedAt  githubv4.DateTime
		   				ID         githubv4.ID
		   				DatabaseID githubv4.Int
		   			}
		   			RateLimit struct {
		   				Cost      githubv4.Int
		   				Limit     githubv4.Int
		   				Remaining githubv4.Int
		   				ResetAt   githubv4.DateTime
		   			}
		   		}
		   		variables := map[string]interface{}{
		   			"repositoryOwner": githubv4.String("shurcooL"),
		   			"repositoryName":  githubv4.String("githubv4"),
		   			"issueNumber":     githubv4.Int(1),
		   		}
		   		err := client.Query(context.Background(), &q, variables)
		   		if err != nil {
		   			return err
		   		}
		   		printJSON(q)
		   		//goon.Dump(out)
		   		//fmt.Println(github.Stringify(out))
		   	} */

	}

	// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
	// func printJSON(v interface{}) {
	// 	w := json.NewEncoder(os.Stdout)
	// 	w.SetIndent("", "\t")
	// 	err := w.Encode(v)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	return nil
}
