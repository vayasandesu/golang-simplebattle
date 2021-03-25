package scenario

import (
	"encoding/json"
	"fmt"
	"simplebattle/jsonexample"
)

func SerializeObjectRunner() {
	obj := jsonexample.JsonStruct{
		Name: "Serializer",
		Age:  1,
	}
	text := jsonexample.Serialize(obj)
	fmt.Println("Serialize object runner")
	fmt.Println(text)
	fmt.Println("=======================")
}

func DeserializeObjectRunner() {
	text := `{"myName" : "benten", "age" : 88 }`
	obj := jsonexample.Deserialize(text)

	fmt.Println("Deserialize object runner")
	fmt.Println(obj)
	fmt.Println("=======================")
}

func DeserializeObjectWithWrongJsonKeyRunner() {
	text := `{"Name" : "benten", "Age" : 88 }`
	obj := jsonexample.Deserialize(text)

	fmt.Println("Deserialize object runner")
	fmt.Println(obj)
	fmt.Println("Name property not deserialize because JSON name is MyName")
	fmt.Println("=======================")
}

func DeserializeUnknowStruct() {
	var result map[string]interface{}
	jsonText := `{
		"name": "poring", 
		"level": 99, 
		"stat" : {
			"hp" : 66,
			"atk": 10
		}
	}`

	json.Unmarshal([]byte(jsonText), &result)

	fmt.Println("Deserialize unknow struct object runner")
	fmt.Println(result)
	fmt.Println("=======================================")
}
