/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"github.com/apache/incubator-openwhisk-wskdeploy/utils"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

// TODO() i18n
var versionCmd = &cobra.Command{
	Use:   "version",
	SuggestFor: []string {"edition", "release"},
	Short: "Print the version number of openwhisk-wskdeploy",
	Long:  `Print the version number of openwhisk-wskdeploy`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("openwhisk-wskdeploy version is %s--%s\n", utils.Flags.CliBuild, utils.Flags.CliVersion)
	},
}
