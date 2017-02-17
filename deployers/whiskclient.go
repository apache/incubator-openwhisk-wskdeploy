package deployers

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/openwhisk/openwhisk-client-go/whisk"
	"github.com/openwhisk/openwhisk-wskdeploy/parsers"
	"github.com/openwhisk/openwhisk-wskdeploy/utils"
)

func NewWhiskClient(proppath string, deploymentPath string, isInteractive bool) (*whisk.Client, *whisk.Config) {
	var clientConfig *whisk.Config

	configs, err := utils.LoadConfiguration(proppath)
	utils.Check(err)

	credential := configs[2]
	namespace := configs[0]
	//we need to get Apihost from property file which currently not defined in sample deployment file.

	u := configs[1]
	var baseURL *url.URL

	if u == "" && isInteractive == true {
		host, err := promptForValue("\nPlease provide the hostname for OpenWhisk [openwhisk.ng.bluemix.net]: ")
		utils.Check(err)
		if host == "" {
			host = "openwhisk.ng.bluemix.net"
		}

		fmt.Println("Host set to " + host)

		baseURL, err = utils.GetURLBase(host)
		utils.Check(err)

	} else if u == "" {
		// handle some error
	} else {
		baseURL, err = utils.GetURLBase(configs[1])
		utils.Check(err)
	}

	if deploymentPath != "" {
		mm := parsers.NewYAMLParser()
		deployment := mm.ParseDeployment(deploymentPath)
		// We get the first package from the sample deployment file.
		credentialDep := deployment.Application.Credential
		namespaceDep := deployment.Application.Namespace
		baseUrlDep := deployment.Application.BaseUrl

		if credentialDep != "" {
			credential = credentialDep
		}

		if namespaceDep != "" {
			namespace = namespaceDep
		}

		if baseUrlDep != "" {
			u, err := url.Parse(baseUrlDep)
			utils.Check(err)

			baseURL = u
		}

	}

	if credential == "" && isInteractive == true {
		cred, err := promptForValue("\nPlease provide an authentication token: ")
		utils.Check(err)
		credential = cred

		fmt.Println("Authentication token set.")
	}

	if namespace == "" && isInteractive == true {
		ns, err := promptForValue("\nPlease provide a namespace [default]: ")
		utils.Check(err)

		if ns == "" {
			ns = "_"
		}

		namespace = ns
		fmt.Println("Namespace set to '" + namespace + "'")
	}

	clientConfig = &whisk.Config{
		AuthToken: credential, //Authtoken
		Namespace: namespace,  //Namespace
		BaseURL:   baseURL,
		Version:   "v1",
		Insecure:  true, // true if you want to ignore certificate signing

	}

	// Setup network client
	client, err := whisk.NewClient(http.DefaultClient, clientConfig)
	utils.Check(err)
	return client, clientConfig

}

func promptForValue(msg string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	return text, nil

}
