package operations

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/spaceuptech/space-cloud/space-cli/cmd/model"
	"github.com/spaceuptech/space-cloud/space-cli/cmd/utils"
)

// Update upgrades existing space cloud cluster
func Update(setValuesFlag, valuesYamlFile, chartLocation, version string) error {
	_ = utils.CreateDirIfNotExist(utils.GetSpaceCloudDirectory())

	charList, err := utils.HelmList(model.HelmSpaceCloudNamespace)
	if err != nil {
		return err
	}
	if len(charList) < 1 {
		utils.LogInfo("Space cloud cluster not found, setup a new cluster using the setup command")
		return nil
	}

	clusterID := charList[0].Name
	isOk := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Space cloud cluster with id (%s) will be upgraded, Do you want to continue", clusterID),
	}
	if err := survey.AskOne(prompt, &isOk); err != nil {
		return err
	}
	if !isOk {
		return nil
	}

	valuesFileObj, err := utils.ExtractValuesObj(setValuesFlag, valuesYamlFile)
	if err != nil {
		return err
	}

	// set clusterId of existing cluster
	charInfo, err := utils.HelmGet(clusterID)
	if err != nil {
		return err
	}

	valuesFileObj["clusterId"] = charInfo.Config["clusterId"]

	_, err = utils.HelmUpgrade(clusterID, chartLocation, utils.GetHelmChartDownloadURL(model.HelmSpaceCloudChartDownloadURL, version), "", valuesFileObj)
	if err != nil {
		return err
	}

	fmt.Println()
	utils.LogInfo(fmt.Sprintf("Space Cloud (cluster id: \"%s\") has been successfully upgraded! 👍", charList[0].Name))
	return nil
}
