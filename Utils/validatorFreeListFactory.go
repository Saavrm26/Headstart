package Utils

import "github.com/santhosh-tekuri/jsonschema/v5"

func ValidatorFreeListFactory(filePath string, nInstance uint,
	                            freeList chan<- *jsonschema.Schema) func() error {
	return func() error {
		list, err := SpawnJSONValidator(filePath, nInstance)
		if err != nil {
			return err
		}

		for _, v := range list {
			select {
			case freeList <- v:
			}
		}

		return nil
	}
}
