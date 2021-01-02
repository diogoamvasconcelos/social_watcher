package lib

import (
	"encoding/json"
	"fmt"
)

func JSONStringify(obj interface{}) (string, error) {
	out, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(string(out)), nil
}
