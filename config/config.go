package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// Viewer is the user login session
var query struct {
	Viewer struct {
		Login     githubv4.String
		CreatedAt githubv4.DateTime
	}
}

// Init is responsible to connect with github gql api.
func Init() *githubv4.Client {
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
	}
	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

	return client
}
