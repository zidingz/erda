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

package strutil

import "regexp"

func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, s string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(s), -1) {
		var groups []string
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, s[v[i]:v[i+1]])
		}

		result += s[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + s[lastIndex:]
}

// docker repository 正则：^[a-z0-9]+(?:(?:(?:[._]|__|[-]*)[a-z0-9]+)+)?$
// ingress 域名不允许有 . ，spring cloud 不允许有 _
var reg = regexp.MustCompile(`^[a-z0-9]+(?:(?:(?:[-]*)[a-z0-9]+)+)?$`)

// IsValidPrjOrAppName 是否是一个合法的项目或应用名
// 需要满足docker repository的规则，和ingress域名的规则
func IsValidPrjOrAppName(repo string) bool {
	return reg.MatchString(repo)
}

func InSlice(item string, dst []string) bool {
	for _, v := range dst {
		if item == v {
			return true
		}
	}
	return false
}
