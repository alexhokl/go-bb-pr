package models

import (
	"fmt"
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

// Commit struct
type Commit struct {
	Branch Branch     `json:"branch"`
	Commit CommitInfo `json:"commit"`
}

// CommitInfo struct
type CommitInfo struct {
	Hash string `json:"hash"`
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
	Content   CommentContent `json:"content"`
	CreatedOn time.Time      `json:"created_on"`
	UpdatedOn time.Time      `json:"updated_on"`
	User      User           `json:"user"`
}

// CommentContent struct
type CommentContent struct {
	Raw string `json:"raw"`
}

// PullRequestCreateRequest struct
type PullRequestCreateRequest struct {
	Destination Commit     `json:"destination"`
	Source      Commit     `json:"source"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Reviewers   []Reviewer `json:"reviewers"`
}

// Repository struct
type Repository struct {
	Org  string
	Name string
}

// UserCredential struct
type UserCredential struct {
	Username string
	Password string
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

// ToShortDescription retursn a short description of a pull request
func (pr PullRequestDetail) ToShortDescription(isIncludeCreatedOn bool) string {
	approveStr := ""
	for _, reviewer := range pr.Participants {
		if reviewer.Approved {
			approveStr = fmt.Sprintf("%s\n\tApproved by %s", approveStr, reviewer.User)
		}
	}

	if isIncludeCreatedOn {
		return fmt.Sprintf("%d Updated:%s Created by:%s at %s\n\t%s -> %s\n\t%s\n%s\n\n",
			pr.ID,
			formatLocalTime(pr.UpdatedOn),
			pr.Author.DisplayName,
			formatLocalTime(pr.CreatedOn),
			pr.Source.Branch.Name,
			pr.Destination.Branch.Name,
			pr.Title,
			approveStr)
	}
	return fmt.Sprintf("%d Updated:%s Created by:%s\n\t%s -> %s\n\t%s\n%s\n\n",
		pr.ID,
		formatLocalTime(pr.UpdatedOn),
		pr.Author.DisplayName,
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title,
		approveStr)
}

// ToOneLiner retursn a short description of a pull request
func (pr PullRequestDetail) ToOneLiner() string {
	return fmt.Sprintf("%d Updated:%s %s %s\n",
		pr.ID,
		formatLocalTime(pr.UpdatedOn),
		pr.Author.DisplayName,
		pr.Title)
}

// ToString returns the description of a pull request
func (pr PullRequestDetail) ToString() string {
	return fmt.Sprintf("%d Updated:%s Created by:%s at %s\n\t%s -> %s\n\t%s\n%s\n",
		pr.ID,
		formatLocalTime(pr.UpdatedOn),
		pr.Author.DisplayName,
		formatLocalTime(pr.CreatedOn),
		pr.Source.Branch.Name,
		pr.Destination.Branch.Name,
		pr.Title,
		pr.Description)
}

// ToString returns the description of a comment
func (c Comment) ToString() string {
	return fmt.Sprintf(
		"Comment by %s (%s): %s",
		c.User.DisplayName,
		formatLocalTime(c.UpdatedOn),
		c.Content.Raw)
}

// ToString returns the description of an update
func (c Update) ToString() string {
	return fmt.Sprintf(
		"Commit made by %s (%s): %s",
		c.Author.DisplayName,
		formatLocalTime(c.Date),
		c.Source.Commit.Hash)
}

func formatLocalTime(t time.Time) string {
	loc, _ := time.LoadLocation("Local")
	lt := t.In(loc)
	return lt.Format("2006-01-02 15:04")
}
