package cmdImp

import (
	"github.com/openwhisk/openwhisk-wskdeploy/utils"
	"path"
	"github.com/openwhisk/openwhisk-wskdeploy/deployers"
	"log"
	"github.com/openwhisk/openwhisk-client-go/whisk"
	"regexp"
	"errors"
)

func Undeploy(params DeployParams) error {
	// TODO: Work your own magic here
	whisk.SetVerbose(params.Verbose)

	if params.ManifestPath == "" {
		if ok, _ := regexp.Match(ManifestFileNameYml, []byte(params.ManifestPath)); ok {
			params.ManifestPath = path.Join(params.ProjectPath, ManifestFileNameYml)
		} else {
			params.ManifestPath = path.Join(params.ProjectPath, ManifestFileNameYaml)
		}

	}

	if params.DeploymentPath == "" {
		if ok, _ := regexp.Match(DeploymentFileNameYml, []byte(params.ManifestPath)); ok {
			params.DeploymentPath = path.Join(params.ProjectPath, DeploymentFileNameYml)
		} else {
			params.DeploymentPath = path.Join(params.ProjectPath, DeploymentFileNameYaml)
		}

	}

	if utils.FileExists(params.ManifestPath) {

		var deployer = deployers.NewServiceDeployer()
		deployer.ProjectPath = params.ProjectPath
		deployer.ManifestPath = params.ManifestPath
		deployer.DeploymentPath = params.DeploymentPath

		deployer.IsInteractive = params.UseInteractive
		deployer.IsDefault = params.UseDefaults

		userHome := utils.GetHomeDirectory()
		propPath := path.Join(userHome, ".wskprops")

		whiskClient, clientConfig := deployers.NewWhiskClient(propPath, params.DeploymentPath, deployer.IsInteractive)
		deployer.Client = whiskClient
		deployer.ClientConfig = clientConfig

		verifiedPlan, err := deployer.ConstructUnDeploymentPlan()
		err = deployer.UnDeploy(verifiedPlan)
		if err != nil {
			utils.Check(err)
			return err
		} else {
			return nil
		}

	} else {
		log.Println("missing manifest.yaml file")
		return errors.New("missing manifest.yaml file")
	}
}