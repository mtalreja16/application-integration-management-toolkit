// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package integrations

import (
	"fmt"

	"internal/apiclient"

	"internal/client/integrations"

	"github.com/spf13/cobra"
)

// ListExecCmd to list executions of an integration version
var ListExecCmd = &cobra.Command{
	Use:   "list",
	Short: "List all executions of an integration version",
	Long:  "List all executions of an integration version",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if err = apiclient.SetRegion(region); err != nil {
			return err
		}
		if allVersions && pageSize != -1 {
			return fmt.Errorf("allVersions and pageSize cannot be combined")
		}
		if allVersions && pageToken != "" {
			return fmt.Errorf("allVersions and pageToken cannot be combined")
		}
		return apiclient.SetProjectID(project)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		_, err = integrations.ListExecutions(name, pageSize, pageToken, filter, orderBy)
		return

	},
}

func init() {
	ListExecCmd.Flags().StringVarP(&name, "name", "n",
		"", "Integration flow name")
	ListExecCmd.Flags().IntVarP(&pageSize, "pageSize", "",
		-1, "The maximum number of versions to return")
	ListExecCmd.Flags().StringVarP(&pageToken, "pageToken", "",
		"", "A page token, received from a previous call")
	ListExecCmd.Flags().StringVarP(&filter, "filter", "",
		"", "Filter results")
	ListExecCmd.Flags().StringVarP(&orderBy, "orderBy", "",
		"", "The results would be returned in order")

	_ = ListExecCmd.MarkFlagRequired("name")
}
