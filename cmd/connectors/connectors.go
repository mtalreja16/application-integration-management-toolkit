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

package connectors

import (
	"github.com/spf13/cobra"
)

// Cmd to manage preferences
var Cmd = &cobra.Command{
	Use:   "connectors",
	Short: "Manage connections for Integration Connectors",
	Long:  "Manage connections for Integration Connectors",
}

var region, project, name string

func init() {

	Cmd.PersistentFlags().StringVarP(&project, "proj", "p",
		"", "Integration GCP Project name")

	Cmd.PersistentFlags().StringVarP(&region, "reg", "r",
		"", "Integration region name")

	Cmd.AddCommand(CreateCmd)
	Cmd.AddCommand(DelCmd)
	Cmd.AddCommand(ListCmd)
	Cmd.AddCommand(GetCmd)
	Cmd.AddCommand(IamCmd)
	Cmd.AddCommand(NodeCountCmd)
	Cmd.AddCommand(ExportCmd)
	Cmd.AddCommand(ImportCmd)
	Cmd.AddCommand(PatchCmd)
}
