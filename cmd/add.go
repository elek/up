package cmd

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/cobra"
	"storj.io/storj-up/cmd/files/templates"
	"storj.io/storj-up/pkg/common"
)

func AddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add <selector>",
		Short: "add more services to the docker compose file. " + selectorHelp,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			composeProject, err := common.LoadComposeFromFile(ComposeFile)
			if err != nil {
				return err
			}
			templateProject, err := common.LoadComposeFromBytes(templates.ComposeTemplate)
			if err != nil {
				return err
			}
			updatedComposeProject, err := AddToCompose(composeProject, templateProject, args)
			if err != nil {
				return err
			}
			return common.WriteComposeFile(updatedComposeProject)
		},
	}
}

func init() {
	rootCmd.AddCommand(AddCmd())
}

func AddToCompose(compose *types.Project, template *types.Project, services []string) (*types.Project, error) {
	for _, service := range common.ResolveServices(services) {
		if !common.ContainsService(compose.Services, service) {
			newService, err := template.GetService(service)
			if err != nil {
				return nil, err
			}
			compose.Services = append(compose.Services, newService)
		}
	}
	return compose, nil
}
