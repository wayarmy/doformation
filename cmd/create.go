// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"

	"github.com/doformation/doformation/manifests"
	"github.com/doformation/doformation/resources"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources on DigitalOcean",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		create()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().StringP("template", "t", "", "Template of stack")
	createCmd.PersistentFlags().StringP("name", "n", "", "Stack name")

	viper.BindPFlag("template", createCmd.PersistentFlags().Lookup("template"))
	viper.BindPFlag("name", createCmd.PersistentFlags().Lookup("name"))
}

type rParser struct {
	rName     string
	rResource interface{}
}

func create() {
	// stackName := viper.GetString("name")
	template := viper.GetString("template")
	tParser, err := manifests.Parser(template)
	if err != nil {
		panic(err)
	}

	var arr []rParser
	for k, v := range tParser.Resources {
		switch v.Type {
		case "DO::Droplet::Server":
			rYAML, err := yaml.Marshal(v.Properties)
			if err != nil {
				panic(err)
			}

			var d resources.DoDroplet
			err = yaml.Unmarshal(rYAML, &d)
			if err != nil {
				panic(err)
			}

			arr = append(arr, rParser{
				rName:     k,
				rResource: d,
			})
		default:
			panic(errors.New("Type " + v.Type + " is not supported yet"))
		}
	}

	fmt.Println(arr[0].rResource)
}
