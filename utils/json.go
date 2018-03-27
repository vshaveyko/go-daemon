package utils

import (
	"encoding/json"
	"log"
)

func Jsonify(tar interface{}) string {
	json, err := json.Marshal(tar)

	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}
