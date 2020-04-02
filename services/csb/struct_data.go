package csb

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in csb response
type Data struct {
	Id                   int64                    `json:"Id" xml:"Id"`
	UpdateCount          int                      `json:"UpdateCount" xml:"UpdateCount"`
	CurrentPage          int                      `json:"CurrentPage" xml:"CurrentPage"`
	Total                int                      `json:"Total" xml:"Total"`
	PageNumber           int                      `json:"PageNumber" xml:"PageNumber"`
	Exist                bool                     `json:"Exist" xml:"Exist"`
	ProjectNameList      []string                 `json:"ProjectNameList" xml:"ProjectNameList"`
	Credentials          Credentials              `json:"Credentials" xml:"Credentials"`
	Order                Order                    `json:"Order" xml:"Order"`
	Service              Service                  `json:"Service" xml:"Service"`
	Instance             Instance                 `json:"Instance" xml:"Instance"`
	RegionList           []Region                 `json:"RegionList" xml:"RegionList"`
	ProjectList          []Project                `json:"ProjectList" xml:"ProjectList"`
	LinkRuleList         []LinkRuleListItem       `json:"LinkRuleList" xml:"LinkRuleList"`
	CredentialList       []Credential             `json:"CredentialList" xml:"CredentialList"`
	OrderList            []OrderInFindOrderedList `json:"OrderList" xml:"OrderList"`
	ItemList             []Item                   `json:"ItemList" xml:"ItemList"`
	ServiceList          []Service                `json:"ServiceList" xml:"ServiceList"`
	InstanceNodeList     []InstanceNode           `json:"InstanceNodeList" xml:"InstanceNodeList"`
	MonitorStatisticData []ServiceStatisticData   `json:"MonitorStatisticData" xml:"MonitorStatisticData"`
}
