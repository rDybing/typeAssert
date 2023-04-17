package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type testDataT struct {
	NotBool interface{} `json:"notBool"`
	isBool  bool
}

func main() {
	jsonStr := (`[{"notBool": 1}, {"notBool": 0}, {"notBool": false}]`)
	var td []testDataT
	err := json.Unmarshal([]byte(jsonStr), &td)
	if err != nil {
		msg := fmt.Errorf("could not unmarshal testdata: %w", err)
		log.Panic(msg)
	}
	for i, v := range td {
		//fmt.Println(reflect.TypeOf(v.NotBool))
		switch v.NotBool.(type) {
		case float64:
			x := int(v.NotBool.(float64))
			v.isBool = x != 0
		case bool:
			v.isBool = v.NotBool.(bool)
		}
		fmt.Printf("index %02d is %v\n", i, v.isBool)
	}
}
