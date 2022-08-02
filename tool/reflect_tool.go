package tool

import "reflect"

func CallMethod(obj interface{}, method string, params ...interface{}) interface{} {
	valueParams := make([]reflect.Value, len(params))
	for i, p := range params {
		valueParams[i] = reflect.ValueOf(p)
	}
	ptr := reflect.ValueOf(obj)
	methodValue := ptr.MethodByName(method)
	res := methodValue.Call(valueParams)
	if len(res) == 0 {
		return nil
	}
	return res[0].Interface()
}
