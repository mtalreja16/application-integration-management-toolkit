// Copyright 2023 Google LLC
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

package sfdc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/GoogleCloudPlatform/application-integration-management-toolkit/apiclient"
)

type instance struct {
	Name             string   `json:"name,omitempty"`
	DisplayName      string   `json:"displayName,omitempty"`
	Description      string   `json:"description,omitempty"`
	SfdcOrgId        string   `json:"sfdcOrgId,omitempty"`
	AuthConfigId     []string `json:"authConfigId,omitempty"`
	UpateTime        string   `json:"upateTime,omitempty"`
	CreateTime       string   `json:"createTime,omitempty"`
	DeleteTime       string   `json:"deleteTime,omitempty"`
	ServiceAuthority string   `json:"serviceAuthority,omitempty"`
}

type instances struct {
	SfdcInstances []instance `json:"sfdcInstances,omitempty"`
	NextPageToken string     `json:"nextPageToken,omitempty"`
}

type instanceExternal struct {
	DisplayName      string   `json:"displayName,omitempty"`
	Description      string   `json:"description,omitempty"`
	SfdcOrgId        string   `json:"sfdcOrgId,omitempty"`
	AuthConfigId     []string `json:"authConfigId,omitempty"`
	ServiceAuthority string   `json:"serviceAuthority,omitempty"`
}

// CreateInstance
func CreateInstance(name string, description string, sfdcOrgId string, serviceAuthority string, authConfig []string) (respBody []byte, err error) {
	u, _ := url.Parse(apiclient.GetBaseIntegrationURL())

	if len(authConfig) < 1 {
		return nil, fmt.Errorf("at least one authConfig must be sent")
	}

	instanceStr := []string{}
	instanceStr = append(instanceStr, "\"displayName\":\""+name+"\"")
	instanceStr = append(instanceStr, "\"description\":\""+description+"\"")
	instanceStr = append(instanceStr, "\"sfdcOrgId\":\""+sfdcOrgId+"\"")
	instanceStr = append(instanceStr, "\"serviceAuthority\":\""+serviceAuthority+"\"")

	authConfigsStr := "\"attributes\":[" + strings.Join(authConfig, ",") + "]"

	instanceStr = append(instanceStr, "\"authConfigId\":\""+authConfigsStr+"\"")

	payload := "{" + strings.Join(instanceStr, ",") + "}"
	u.Path = path.Join(u.Path, "sfdcInstances")
	respBody, err = apiclient.HttpClient(apiclient.GetPrintOutput(), u.String(), payload)

	return respBody, err
}

// GetInstance
func GetInstance(name string, minimal bool) (respBody []byte, err error) {
	u, _ := url.Parse(apiclient.GetBaseIntegrationURL())
	u.Path = path.Join(u.Path, "sfdcInstances", name)

	printSetting := apiclient.GetPrintOutput()
	if minimal {
		apiclient.SetPrintOutput(false)
	}
	respBody, err = apiclient.HttpClient(apiclient.GetPrintOutput(), u.String())
	if minimal {
		iversion := instance{}
		err := json.Unmarshal(respBody, &iversion)
		if err != nil {
			return nil, err
		}
		eversion := convertInternalInstanceToExternal(iversion)
		respBody, err = json.Marshal(eversion)
		if err != nil {
			return nil, err
		}
		if printSetting {
			apiclient.PrettyPrint(respBody)
		}
	}
	apiclient.SetPrintOutput(printSetting)
	return respBody, err
}

// ListInstances
func ListInstances() (respBody []byte, err error) {
	u, _ := url.Parse(apiclient.GetBaseIntegrationURL())
	u.Path = path.Join(u.Path, "sfdcInstances")
	respBody, err = apiclient.HttpClient(apiclient.GetPrintOutput(), u.String())
	return respBody, err
}

// FindInstance
func FindInstance(name string) (version string, respBody []byte, err error) {
	ilist := instances{}

	respBody, err = ListInstances()
	if err != nil {
		return "", nil, err
	}
	if err = json.Unmarshal(respBody, &ilist); err != nil {
		return "", nil, err
	}

	for _, i := range ilist.SfdcInstances {
		if i.DisplayName == name {
			version = i.Name[strings.LastIndex(i.Name, "/")+1:]
			respBody, err := json.Marshal(&i)
			return version, respBody, err
		}
	}
	return "", nil, fmt.Errorf("instance not found")
}

// convertInternalInstanceToExternal
func convertInternalInstanceToExternal(internalVersion instance) (externalVersion instanceExternal) {
	externalVersion = instanceExternal{}

	externalVersion.DisplayName = internalVersion.Name
	externalVersion.Description = internalVersion.Description
	externalVersion.ServiceAuthority = internalVersion.ServiceAuthority
	externalVersion.SfdcOrgId = internalVersion.SfdcOrgId
	externalVersion.AuthConfigId = internalVersion.AuthConfigId

	return externalVersion
}