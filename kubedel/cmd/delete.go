/*
Copyright © 2020 Abhishek Singh Saini <abhi.taker20@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

func FindObjects(kind, namespace string) {

	var (
		instruct string
		out []string
	)
	
	ObjectNameList := []string{}


	if namespace == "" {
		instruct = "kubectl get " +  kind + " -o=name"
	} else {
		instruct = "kubectl get " +  kind + " -n " + namespace + " -o=name"
	}

	cmd := exec.Command("sh", "-c", instruct)
	res, _ := cmd.Output()
	out = strings.Fields(string(res))

	for ind:=0; ind<len(out); ind+=1 {
		name := strings.Split(out[ind], "/")[1]
		ObjectNameList = append(ObjectNameList, name)
	}

	DeleteObjects(kind, namespace, ObjectNameList)
}

func DeleteObjects(kind, namespace string, ObjectNameList []string) {

	var instruct string

	for _, objectName := range ObjectNameList {
		if namespace == "" {
			instruct = "kubectl delete " +  kind + " " + objectName
		} else {
			instruct = "kubectl delete " +  kind + " " + objectName + " -n " + namespace 
		}

		cmd := exec.Command("sh", "-c", instruct)
		res, _ := cmd.Output()
		fmt.Printf("%s", res)
	}
}


// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Command to delete Kubernetes Object",
	Long: `This Command helps you to delete Kubernetes Objects`,

	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		
		for ind:=0; ind<len(args); ind+=1 {
			FindObjects(args[ind], namespace)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("namespace", "n", "", "Namespace where the object to be removed resides")
}
