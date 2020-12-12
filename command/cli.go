package command

import (
	"github.com/alexhokl/go-bb-pr/client"
	"github.com/alexhokl/go-bb-pr/models"
	"github.com/spf13/cobra"
)

// Cli interface
type Cli interface {
	Client() client.APIClient
	UserCredential() *models.UserCredential
	Repo() *models.Repository
}

// ManagerCli struct
type ManagerCli struct {
	client     client.APIClient
	credential *models.UserCredential
	repo       *models.Repository
}

type idOption struct {
	id int
}

// NewManagerCli creates a new manager cli instance
func NewManagerCli(cred *models.UserCredential, repo *models.Repository) *ManagerCli {
	cli := ManagerCli{
		client:     client.NewClient(),
		credential: cred,
		repo:       repo,
	}
	return &cli
}

// Client returns API client
func (cli *ManagerCli) Client() client.APIClient {
	return cli.client
}

// UserCredential returns user credential
func (cli *ManagerCli) UserCredential() *models.UserCredential {
	return cli.credential
}

// Repo returns information on Repository
func (cli *ManagerCli) Repo() *models.Repository {
	return cli.repo
}

// ShowHelp shows the command help
func (cli *ManagerCli) ShowHelp(cmd *cobra.Command, args []string) error {
	cmd.HelpFunc()(cmd, args)
	return nil
}
