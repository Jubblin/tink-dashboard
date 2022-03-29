// Copyright 2022 Tinker codeowners.
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

package types

// TemplateList is the view model for template list page
type TemplateList struct {
	Base
	Templates []Template
}

// Template represents a workflow template
type Template struct {
	ID          string
	Name        string
	Data        string
	LastUpdated string
}

// UpdateTemplate represents an update request for a template
type UpdateTemplate struct {
	ID   string
	Name string
	Data string
}
