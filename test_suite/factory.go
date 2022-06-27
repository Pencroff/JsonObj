package test_suite

import (
	djs "github.com/Pencroff/JsonStruct"
	"github.com/Pencroff/JsonStruct/experiment"
)

func JsonStructValueFactory() djs.JsonStructOps {
	return &experiment.JsonStructValue{}
}

func JsonStructPointerFactory() djs.JsonStructOps {
	return &experiment.JsonStructPtr{}
}
