package utils

import "fmt"

// DataFrame is the query result type
type DataFrame struct {
	Search Search `json:"{data{search{}}}"`
}

// Search ...
type Search struct {
	Repos []Repo `json:"nodes[{}]"`
}

// Repo is the repository type
type Repo struct {
	ID                 int    `json:"id"`
	RepoName           string `json:"name"`
	StarsNumber        int64  `json:"stargazers"`
	CreatedAt          string `json:"createdAt"`
	PullRequestsNumber int64  `json:"pullRequests"`
	// ProgLanguage       string    `json:"primaryLanguage{name}"`
}

// Query repares the json stub to send
func Query(after string) map[string]string {
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
