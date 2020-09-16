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

type Animal struct {
	Name  string
	Order string
}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	run()
}

func run() {
	jsonData := map[string]string{
		"query": `
	  {
	    search(query: "stars:>100", type: REPOSITORY, first: 50) {
	      repositoryCount
	      pageInfo {
	        hasNextPage
	        endCursor
	      }
	      nodes {
	        ... on Repository {
	          name
	          owner{login}
	          url
	          stargazers{totalCount}
            issues{totalCount}
            createdAt
	          primaryLanguage {
	            name
	          }
	        }
	      }
	    }
	  }
	  `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	// Bearer token auth header
	var bearer = "Bearer " + os.Getenv("GITHUB_TOKEN")
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	// Receives resp from the req
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile("out/repos.json", data, 0777)

}

/*
query = """
    query{
      search(query: "stars:>100", type: REPOSITORY, first: 50 %s) {
        repositoryCount
        pageInfo {
          hasNextPage
          endCursor
        }
        nodes {
          ... on Repository {
            name
            owner{login}
            url
            stargazers{totalCount}
            issues{totalCount}
            primaryLanguage {
              name
            }
          }
        }
      }

    }
"""
*/
