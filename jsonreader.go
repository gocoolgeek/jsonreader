package jsonreader

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var configurations map[string]interface{} = make(map[string]interface{})

func visit(path string, f os.FileInfo, err error) error {
	var jsonMap map[string]interface{}
	if !(f.IsDir()) {
		contentBytes, err := ioutil.ReadFile(path)
		if hasError(err) {
			return err
		}
		if isJSON(contentBytes) {
			err = json.Unmarshal(contentBytes, &jsonMap)
			if hasError(err) {
				return err
			}
			configurations = merge(configurations, jsonMap, 0)
		}
	}
	return nil
}

// Get the entire configuration(s) as map
func GetAll() map[string]interface{} {
	return configurations
}

// Get a map of values for the specified key
func GetMap(key string) map[string]interface{} {
	valueMap := configurations[key]
	if valueMap == nil {
		return make(map[string]interface{})
	}
	return TransformInterfaceToMap(valueMap)
}

// Get the value for dot seperated key path or just key
// Refer to test for examples
func GetValue(dotSeperatedPath string) string {
	var value interface{} = ""
	var mapValues map[string]interface{} = configurations
	splits := strings.Split(dotSeperatedPath, ".")
	for _, split := range splits {
		tempMapValues := mapValues[split]
		if tempMapValues == nil {
			value = ""
			break
		}
		if reflect.TypeOf(tempMapValues).Kind() == reflect.String {
			value = tempMapValues
		} else {
			mapValues = TransformInterfaceToMap(tempMapValues)
		}
	}
	return value.(string)
}

// Load a specific json file or folder which has set of json files
// Only valid json files will be loaded
func Load(root string) error {
	var err error
	err = filepath.Walk(root, visit)
	if hasError(err) {
		return err
	}
	return nil
}

func TransformInterfaceToMap(generic interface{}) map[string]interface{} {
	return generic.(map[string]interface{})
}

func merge(dst, src map[string]interface{}, depth int) map[string]interface{} {
	for key, srcVal := range src {
		if dstVal, ok := dst[key]; ok {
			srcMap, srcMapOk := mapify(srcVal)
			dstMap, dstMapOk := mapify(dstVal)
			if srcMapOk && dstMapOk {
				srcVal = merge(dstMap, srcMap, depth+1)
			}
		}
		dst[key] = srcVal
	}
	return dst
}

func mapify(i interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, k := range value.MapKeys() {
			m[k.String()] = value.MapIndex(k).Interface()
		}
		return m, true
	}
	return map[string]interface{}{}, false
}

func hasError(e error) bool {
	return (e != nil)
}

func isJSONString(s []byte) bool {
	var js string
	return json.Unmarshal(s, &js) == nil
}

func isJSON(s []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(s, &js) == nil
}
