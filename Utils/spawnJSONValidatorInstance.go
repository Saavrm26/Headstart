package Utils

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func SpawnJSONValidator(jsonPath string, nInstances uint) ([]*jsonschema.Schema, error) {
	if nInstances < 1 {
		nInstances = 1
	}

	sch, err := jsonschema.Compile(jsonPath)
	if err != nil {
		return nil, err
	}

	schemaList := make([]*jsonschema.Schema, nInstances)
	schemaList[0] = sch
  
  for i := uint(1); i < nInstances; i++ {
    var cpy *jsonschema.Schema
    *cpy = *sch
    schemaList = append(schemaList, cpy)
  }

  return schemaList, nil
}
