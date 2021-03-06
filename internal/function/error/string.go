package error

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnString{})
}

type fnString struct {
}

// Name returns the name of the function
func (fnString) Name() string {
	return "error.string"
}

// Sig returns the function signature
func (fnString) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeAny}, false
}

// Eval executes the function
func (fnString) Eval(params ...interface{}) (interface{}, error) {
	return params[0].(error).Error(), nil
}
