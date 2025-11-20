package utils

import "encoding/json"

func ConvertMaptoStruct[T any](data map[string]any) T {
	var result T

	dataByte, err := json.Marshal(data)
	if err != nil {
		return result
	}

	err = json.Unmarshal(dataByte, &result)
	if err != nil {
		return result
	}

	return result
}
