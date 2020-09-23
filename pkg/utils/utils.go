package utils

// Response is the query response type (for http)
type Response struct {
	DataFrame DataFrame
}

// DataFrame is the query result type from gql
type DataFrame struct {
	Search Search `json:"Search"`
}

// Search is the main search result type from gql
type Search struct {
	Nodes []Node `json:"Nodes"`
	// PageInfo PageInfo `json:"PageInfo"`
}

// PageInfo type from gql
type PageInfo struct {
	EndCursor       string `json:"EndCursor"`
	HasNextPage     bool   `json:"HasNextPage"`
	HasPreviousPage bool   `json:"HasPreviousPage"`
}

// Node is the gql node reference
type Node struct {
	Repository Repository `json:"Repository"`
}

// Repository is the repository type
type Repository struct {
	ID              string          `json:"ID"`
	Name            string          `json:"Name"`
	Owner           Owner           `json:"Owner"`
	PrimaryLanguage PrimaryLanguage `json:"PrimaryLanguage"`
	URL             string          `json:"Url"`
	Stargazers      TotalCount      `json:"Stargazers"`
	PullRequests    TotalCount      `json:"PullRequests"`
	IssuesTotal     TotalCount      `json:"IssuesTotal"`
	IssuesClosed    TotalCount      `json:"IssuesClosed"`
	CreatedAt       string          `json:"CreatedAt"`
	UpdatedAt       string          `json:"UpdatedAt"`
}

// TotalCount is integer count from gql
type TotalCount struct {
	TotalCount int `json:"TotalCount"`
}

// Owner is the username from github
type Owner struct {
	Login string `json:"Login"`
}

// PrimaryLanguage is the main programming language from the repository
type PrimaryLanguage struct {
	Name string `json:"Name"`
}
