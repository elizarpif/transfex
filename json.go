package transfex

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJson(name string) (map[string]interface{}, error){
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