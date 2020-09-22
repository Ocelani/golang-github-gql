package utils

// Response is the query response type (for http)
type Response struct {
	DataFrame DataFrame `json:"data"`
}

// DataFrame is the query result type from gql
type DataFrame struct {
	Search Search `json:"search"`
}

// Search is the main search result type from gql
type Search struct {
	Nodes    []Node   `json:"nodes"`
	PageInfo PageInfo `json:"pageInfo"`
}

// PageInfo type from gql
type PageInfo struct {
	EndCursor       string `json:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
}

// Node is the repository type
type Node struct {
	ID                 string          `json:"id"`
	Name               string          `json:"name"`
	Owner              Owner           `json:"owner"`
	ProgLanguage       PrimaryLanguage `json:"primaryLanguage"`
	URL                string          `json:"url"`
	StarsNumber        TotalCount      `json:"stargazers"`
	PullRequestsNumber TotalCount      `json:"pullRequests"`
	IssuesTotal        TotalCount      `json:"issuesTotal"`
	IssuesClosed       TotalCount      `json:"issuesClosed"`
	CreatedAt          string          `json:"createdAt"`
	UpdatedAt          string          `json:"updatedAt"`
}

// TotalCount is integer count from gql
type TotalCount struct {
	TotalCount int `json:"totalCount"`
}

// Owner is the username from github
type Owner struct {
	Login string `json:"login"`
}

// PrimaryLanguage is the main programming language from the repository
type PrimaryLanguage struct {
	Name string `json:"name"`
}
