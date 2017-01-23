package client

import (
	"errors"
	"fmt"

	"github.com/alexhokl/go-bb-pr/models"
)

// GetRequest makes an API call to retrieve a pull request
func (client *Client) GetRequest(cred *models.UserCredential, repo *models.Repository, id int) (*models.PullRequestDetail, error) {
	path := fmt.Sprintf("%s/%d", getBasePath(repo), id)
	req := newRequest(cred, "GET", path)
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		msg := getErrorResponseMessage(resp)
		return nil, errors.New(msg)
	}
	return parse(resp)
}
