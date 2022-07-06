package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/Pencroff/JsonStruct/experiment"
)

func JsonStructFactory() djs.JStructOps {
	return &djs.JsonStruct{}
}

func JsonStructConvertibleFactory() djs.JStructConvertibleOps {
	return &djs.JsonStruct{}
}

func JsonStructValueFactory() djs.JStructOps {
	return &experiment.JsonStructValue{}
}

func JsonStructPointerFactory() djs.JStructOps {
	return &experiment.JsonStructPtr{}
}
