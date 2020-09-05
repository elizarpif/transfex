package transfex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
)

func createStructNameFromFile(filename string) string {
	filename = strings.Title(filename)

	_, f := path.Split(filename)

	if ind := strings.Index(f, "."); ind > 0 {
		return f[:ind]
	}
	return f
}

func GenerateCode(filename, pack string) error {
	mapa, err := ReadJson(filename)
	if err != nil {
		return err
	}

	name := createStructNameFromFile(filename)
	data, err := GenerateStructs(name, mapa)
	if err != nil {
		return err
	}

	structs := []string{}
	for k, v := range data{
		structs = append(structs, fmt.Sprintf("type %s struct {\n%s\n}\n", k, v))
	}

	dat := "package " + pack + "\n" + strings.Join(structs,"\n")

	newFile := strings.ToLower(path.Join(pack, name) + ".go")

	_, err = os.Stat(pack)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(pack, 0755)
		if errDir != nil {
			return nil
		}
	}

	file, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println(dat)
	_, err = file.Write([]byte(dat))
	return err
}

func ReadJson(name string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var unmarshalled map[string]interface{}

	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func toCamelCase(old string) string {
	splitted := strings.Split(old, "_")

	news := []string{}
	for _, w := range splitted {
		n := strings.Title(w)
		news = append(news, n)
	}

	return strings.Join(news, "")
}

func GenerateStructs(name string, data map[string]interface{}) (map[string]string, error) {
	namesOfStructs := []string{}

	ret := map[string]string{}

	for key, value := range data {
		str := toCamelCase(key)

		t := reflect.ValueOf(value)
		switch t.Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32:
			{
				str += fmt.Sprintf(" int `json:\"%s\"`", key)
			}
		case reflect.Float64:
			{
				str += fmt.Sprintf(" float64 `json:\"%s\"`", key)
			}
		case reflect.Float32:
			{
				str += fmt.Sprintf(" float32 `json:\"%s\"`", key)
			}
		case reflect.String:
			{
				str += fmt.Sprintf(" string `json:\"%s\"`", key)
			}
		case reflect.Bool:
			{
				str += fmt.Sprintf(" bool `json:\"%s\"`", key)
			}
		case reflect.Slice, reflect.Array:
			{

				tt := t.Index(0).Elem()
				switch tt.Kind() {
				case reflect.Int, reflect.Int64, reflect.Int32:
					{
						str += fmt.Sprintf(" []int `json:\"%s\"`", key)
					}
				case reflect.Float64:
					{
						str += fmt.Sprintf(" []float64 `json:\"%s\"`", key)
					}
				case reflect.Float32:
					{
						str += fmt.Sprintf(" []float32 `json:\"%s\"`", key)
					}
				case reflect.String:
					{
						str += fmt.Sprintf(" []string `json:\"%s\"`", key)
					}
				case reflect.Bool:
					{
						str += fmt.Sprintf(" []bool `json:\"%s\"`", key)
					}
				default:
					{
						tt := t.Index(0).Interface()
						if is, ok := tt.(map[string]interface{}); ok {
							newStruct, err := GenerateStructs(str, is)
							if err != nil {
								return nil, err
							}
							for k, v := range newStruct{
								ret[k] = v
							}

							str += fmt.Sprintf(" *%s `json:\"%s\"`", str, key)
						}
					}
				}

			}
		case reflect.Map:
			{
				if is, ok := value.(map[string]interface{}); ok {
					newStruct, err := GenerateStructs(str, is)
					if err != nil {
						return nil, err
					}
					for k, v := range newStruct{
						ret[k] = v
					}

					str += fmt.Sprintf(" *%s `json:\"%s\"`", str, key)
				}
			}

		}
		namesOfStructs = append(namesOfStructs, str)
	}

	ret[strings.Title(name)]=strings.Join(namesOfStructs, "\n")
	return ret, nil
}
