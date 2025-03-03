// Copyright (c) 2021 Terminus, Inc.
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

package fileFormModal

import (
	protocol "github.com/erda-project/erda/modules/openapi/component-protocol"
)

type ComponentFileFormModal struct {
	CtxBdl     protocol.ContextBundle
	Props      Props                  `json:"props"`
	State      State                  `json:"state"`
	Operations map[string]interface{} `json:"operations"`
}

type State struct {
	FormVisible bool     `json:"formVisible"`
	SceneSetKey int      `json:"sceneSetKey"`
	Visible     bool     `json:"visible"`
	FormData    FormData `json:"formData"`
	ActionType  string   `json:"actionType"`
	SceneId     uint64   `json:"sceneId"`
}

type FormData struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	// SetID       uint64 `json:"setId"`
	ScenesSet *uint64 `json:"scenesSet"`
}

type Props struct {
	Title  string  `json:"title"`
	Fields []Entry `json:"fields"`
}

type Entry struct {
	Key            string         `json:"key"`
	Label          string         `json:"label"`
	Required       bool           `json:"required"`
	Component      string         `json:"component"`
	Rules          []Rule         `json:"rules"`
	ComponentProps ComponentProps `json:"componentProps"`
}

type Rule struct {
	Pattern string `json:"pattern"`
	Msg     string `json:"msg"`
}
type ComponentProps struct {
	MaxLength   int           `json:"maxLength"`
	Placeholder string        `json:"placeholder"`
	Options     []interface{} `json:"options"`
}

type Operation struct {
	Key    string `json:"key"`
	Reload bool   `json:"reload"`
}
