package main

import (
	"simplebattle/scenario"
)

func main() {

	// group of method and interface work
	scenario.GachaponRunner()

	// function and variable
	scenario.FunctionPointerRunner()
	scenario.VariableValueTypeRunner()
	scenario.VariablePointerTypeRunner()

	// Json
	scenario.SerializeObjectRunner()
	scenario.DeserializeObjectRunner()
	scenario.DeserializeObjectWithWrongJsonKeyRunner()
	scenario.DeserializeUnknowStruct()

}
