package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

type listOptions struct {
	isQuiet                     bool
	isOneLiner                  bool
	isIncldeCreationTime        bool
	isHideAuthoredByCurrentUser bool
}

// NewListCommand returns definition of command list
func NewListCommand(cli *ManagerCli) *cobra.Command {
	opts := listOptions{}

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List pull requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				cli.ShowHelp(cmd, args)
				return nil
			}
			return runList(cli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVarP(&opts.isQuiet, "quiet", "q", false, "List IDs only")
	flags.BoolVar(&opts.isOneLiner, "oneline", false, "List in oneliners")
	flags.BoolVar(
		&opts.isIncldeCreationTime, "created-time", false, "Include created time")
	flags.BoolVarP(
		&opts.isHideAuthoredByCurrentUser,
		"hide-current", "c", false, "Hide pull requests created by current user")

	return cmd
}

func runList(cli *ManagerCli, opts listOptions) error {
	client := cli.Client()
	cred := cli.UserCredential()
	repo := cli.Repo()

	prList, err := client.ListRequests(cred, repo)
	if err != nil {
		return err
	}

	if len(prList) == 0 {
		fmt.Println("There are no open pull requests.")
		return nil
	}

	for _, pr := range prList {
		if opts.isHideAuthoredByCurrentUser {
			if pr.Author.Username == cred.Username {
				continue
			}
		}
		prInfo, _ := client.GetRequest(cred, repo, pr.ID)
		printFunc := getPrint(prInfo, cred)
		if opts.isQuiet {
			printFunc("%d", prInfo.ID)
		} else if opts.isOneLiner {
			printFunc(prInfo.ToOneLiner())
		} else {
			printFunc(prInfo.ToShortDescription(opts.isIncldeCreationTime))
		}
	}
	return nil
}
