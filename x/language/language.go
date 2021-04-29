package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Translation(language, key string) string {
	fileName := fmt.Sprintf("./translation/%s.json", language)
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return fmt.Sprintf("%v", result[key])
}
