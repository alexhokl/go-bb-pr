package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/alexhokl/go-bb-pr/models"
)

// APIClient interface
type APIClient interface {
	ListRequests(cred *models.UserCredential, repo *models.Repository) ([]models.PullRequestInfo, error)
	GetRequest(cred *models.UserCredential, repo *models.Repository, id int) (*models.PullRequestDetail, error)
	ApproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	UnapproveRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	DeclineRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	MergeRequest(cred *models.UserCredential, repo *models.Repository, id int) error
	ActivityRequest(cred *models.UserCredential, repo *models.Repository, id int) ([]models.PullRequestActivity, error)
	CreateRequest(cred *models.UserCredential, repo *models.Repository, req *models.PullRequestCreateRequest) error
}

// Client struct
type Client struct {
	client *http.Client
}

// NewClient creates a new Client instance
func NewClient() *Client {
	client := &http.Client{}

	return &Client{
		client: client,
	}
}

func getBasePath(repo *models.Repository) string {
	return fmt.Sprintf(
		"https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests",
		repo.Org,
		repo.Name)
}

func newRequest(cred *models.UserCredential, verb string, path string) *http.Request {
	req, _ := http.NewRequest(verb, path, nil)
	req.SetBasicAuth(cred.Username, cred.Password)
	return req
}

func newPostRequest(cred *models.UserCredential, path string, data interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest("POST", path, buf)
	req.SetBasicAuth(cred.Username, cred.Password)
	return req, err
}

func dumpResponse(resp *http.Response) error {
	_, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func getErrorResponseMessage(resp *http.Response) string {
	return fmt.Sprintf(
		"Failed response (status code: %d): %s", resp.StatusCode, resp.Status)
}
