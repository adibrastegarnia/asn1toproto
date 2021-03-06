// Copyright 2020-present Open Networking Foundation.
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

package protogen

import (
	"path"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var toCamelCase = func(value string) string {
	return strcase.ToCamel(value)
}

var toLowerCamelCase = func(value string) string {
	return strcase.ToLowerCamel(value)
}

var toLowerCase = func(value string) string {
	return strings.ToLower(value)
}

var toUpperCase = func(value string) string {
	return strings.ToUpper(value)
}

var upperFirst = func(value string) string {
	bytes := []byte(value)
	first := strings.ToUpper(string([]byte{bytes[0]}))
	return string(append([]byte(first), bytes[1:]...))
}

var quote = func(value string) string {
	return "\"" + value + "\""
}

var isLast = func(values interface{}, index int) bool {
	t := reflect.ValueOf(values)
	return index == t.Len()-1
}

var split = func(value, sep string) []string {
	return strings.Split(value, sep)
}

var trim = func(value string) string {
	return strings.Trim(value, " ")
}

var replaceHyphen = func(value string) string {
	return strings.Replace(value, "-", "_", -1)
}

var ternary = func(v1, v2 interface{}, b bool) interface{} {
	if b {
		return v1
	}
	return v2
}

// GetTemplate get an instance golang text template
func GetTemplate(filePath, file string) *template.Template {
	funcs := template.FuncMap{
		"toCamel":       toCamelCase,
		"toLowerCamel":  toLowerCamelCase,
		"lower":         toLowerCase,
		"upper":         toUpperCase,
		"upperFirst":    upperFirst,
		"quote":         quote,
		"isLast":        isLast,
		"split":         split,
		"trim":          trim,
		"ternary":       ternary,
		"replaceHyphen": replaceHyphen,
	}
	return template.Must(template.New(file).Funcs(funcs).ParseFiles(path.Join(filePath, file)))
}
