package letsencrypt

import (
	"github.com/spf13/cobra"

	"github.com/spaceuptech/space-cloud/space-cli/cmd/utils"
)

// GenerateSubCommands is the list of commands the ingress module exposes
func GenerateSubCommands() []*cobra.Command {

	var generateletsencrypt = &cobra.Command{
		Use:     "letsencrypt [path to config file]",
		RunE:    actionGenerateLetsEncryptDomain,
		Aliases: []string{"letsencrypts"},
		Example: "space-cli generate letsencrypt config.yaml --project myproject --log-level info",
	}
	return []*cobra.Command{generateletsencrypt}
}

// GetSubCommands is the list of commands the ingress module exposes
func GetSubCommands() []*cobra.Command {

	var getletsencrypt = &cobra.Command{
		Use:  "letsencrypt",
		RunE: actionGetLetsEncrypt,
	}

	return []*cobra.Command{getletsencrypt}
}

func actionGetLetsEncrypt(cmd *cobra.Command, args []string) error {
	// Get the project and url parameters
	project, check := utils.GetProjectID()
	if !check {
		return utils.LogError("Project not specified in flag", nil)
	}
	commandName := "letsencrypt"

	params := map[string]string{}
	obj, err := GetLetsEncryptDomain(project, commandName, params)
	if err != nil {
		return err
	}

	if err := utils.PrintYaml(obj); err != nil {
		return err
	}
	return nil
}

func actionGenerateLetsEncryptDomain(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return utils.LogError("incorrect number of arguments. Use -h to check usage instructions", nil)
	}
	dbruleConfigFile := args[0]
	dbrule, err := generateLetsEncryptDomain()
	if err != nil {
		return err
	}

	return utils.AppendConfigToDisk(dbrule, dbruleConfigFile)
}

// DeleteSubCommands is the list of commands the ingress module exposes
func DeleteSubCommands() []*cobra.Command {

	var getletsencrypt = &cobra.Command{
		Use:     "letsencrypt",
		RunE:    actionDeleteLetsEncrypt,
		Example: "space-cli delete letsencrypt --project myproject",
	}

	return []*cobra.Command{getletsencrypt}
}

func actionDeleteLetsEncrypt(cmd *cobra.Command, args []string) error {
	// Get the project
	project, check := utils.GetProjectID()
	if !check {
		return utils.LogError("Project not specified in flag", nil)
	}

	return deleteLetsencryptDomains(project)
}
