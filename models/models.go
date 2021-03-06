package models

import (
	"time"
)

// PullRequest interface
type PullRequest interface {
	ToString() string
}

// PullRequestList struct
type PullRequestList struct {
	PageLen int               `json:"pagelen"`
	Page    int               `json:"page"`
	Size    int               `json:"size"`
	Items   []PullRequestInfo `json:"values"`
}

// PullRequestInfo struct
type PullRequestInfo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	Author      User      `json:"author"`
	Destination Commit    `json:"destination"`
	Source      Commit    `json:"source"`
	Description string    `json:"description"`
	Links       Links     `json:"links"`
}

// PullRequestDetail struct
type PullRequestDetail struct {
	PullRequestInfo
	Participants []Reviewer `json:"participants"`
}

// Links struct
type Links struct {
	Html Link `json:"html"`
}

// Link struct
type Link struct {
	Href string `json:"href"`
}

// Reviewer struct
type Reviewer struct {
	User     User `json:"user"`
	Approved bool `json:"approved"`
}

// User struct
type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

// CommitBranch struct
type CommitBranch struct {
	Branch Branch `json:"branch"`
}

// Commit struct
type Commit struct {
	CommitBranch
	Commit CommitInfo `json:"commit"`
}

// CommitInfo struct
type CommitInfo struct {
	Hash    string     `json:"hash"`
	Summary RawContent `json:"summary"`
}

// Branch struct
type Branch struct {
	Name string `json:"name"`
}

// PullRequestActivity struct
type PullRequestActivity struct {
	Update  Update  `json:"update,omitempty"`
	Comment Comment `json:"comment,omitempty"`
}

// Update struct
type Update struct {
	Date   time.Time `json:"date"`
	Source Commit    `json:"source"`
	Author User      `json:"author"`
}

// Comment struct
type Comment struct {
	Content   RawContent `json:"content"`
	CreatedOn time.Time  `json:"created_on"`
	UpdatedOn time.Time  `json:"updated_on"`
	User      User       `json:"user"`
}

// RawContent struct
type RawContent struct {
	Raw string `json:"raw"`
}

// PullRequestCreateRequest struct
type PullRequestCreateRequest struct {
	Destination CommitBranch `json:"destination"`
	Source      CommitBranch `json:"source"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Reviewers   []Reviewer   `json:"reviewers"`
}

// CommentRequest contains fields for making a create comment request
type CommentRequest struct {
	Content RawContent `json:"content"`
}

// IssueRequest contains fields for making a request to JIRA API issue endpoints
type IssueRequest struct {
	Update UpdateIssue `json:"update"`
}

// UpdateIssue contains fields to update a JIRA issue
type UpdateIssue struct {
	Labels []LabelReqeuest `json:"labels"`
}

// LabelReqeuest contains fields to update labels of a JIRA issue
type LabelReqeuest struct {
	Add string `json:"add"`
}

// Repository struct
type Repository struct {
	Org  string
	Name string
}

// UserCredential struct
type UserCredential struct {
	AccessToken      string
	RefreshToken     string
	JiraEmailAddress string
	JiraAPIKey       string
}

// HasJiraCredentials return true if JIRA credentials has been configured
func (cred UserCredential) HasJiraCredentials() bool {
	return cred.JiraEmailAddress != "" && cred.JiraAPIKey != ""
}

// IsApproved checks the pull request has been approved by
// user with the specified username
func (pr PullRequestDetail) IsApproved(username string) bool {
	for _, reviewer := range pr.Participants {
		if reviewer.Approved && reviewer.User.Username == username {
			return true
		}
	}
	return false
}
