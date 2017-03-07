package parsers

import (
	"log"

	"github.com/openwhisk/openwhisk-wskdeploy/utils"
	"gopkg.in/yaml.v2"
)

func (dm *YAMLParser) UnmarshalDeployment(input []byte, deploy *DeploymentYAML) error {
	err := yaml.Unmarshal(input, deploy)
	if err != nil {
		log.Printf("error happened during unmarshal :%v", err)
		return err
	}
	return nil
}

func (dm *YAMLParser) MarshalDeployment(deployment *DeploymentYAML) (output []byte, err error) {
	data, err := yaml.Marshal(deployment)
	if err != nil {
		log.Printf("err happened during marshal :%v", err)
		return nil, err
	}
	return data, nil
}

func (dm *YAMLParser) ParseDeployment(dply string) *DeploymentYAML {
	dplyyaml := DeploymentYAML{}
	content, err := new(utils.ContentReader).LocalReader.ReadLocal(dply)
	utils.Check(err)
	err = dm.UnmarshalDeployment(content, &dplyyaml)
	utils.Check(err)
	dplyyaml.Filepath = dply
	return &dplyyaml
}

//********************Application functions*************************//
//This is for parse the deployment yaml file.
func (app *Application) GetPackageList() []Package {
	var s1 []Package = make([]Package, 0)
	for _, pkg := range app.Packages {
		pkg.Packagename = pkg.Packagename
		s1 = append(s1, pkg)
	}
	return s1
}
