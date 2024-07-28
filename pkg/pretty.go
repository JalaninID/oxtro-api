package pkg

import "encoding/json"

func Prt(i any) {
	json, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(json))
}
