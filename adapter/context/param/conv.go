package param

import (
	"reflect"

	beecontext "github.com/xiaobo88michael/beego/adapter/context"
	"github.com/xiaobo88michael/beego/server/web/context"
	"github.com/xiaobo88michael/beego/server/web/context/param"
)

// ConvertParams converts http method params to values that will be passed to the method controller as arguments
func ConvertParams(methodParams []*MethodParam, methodType reflect.Type, ctx *beecontext.Context) (result []reflect.Value) {
	nps := make([]*param.MethodParam, 0, len(methodParams))
	for _, mp := range methodParams {
		nps = append(nps, (*param.MethodParam)(mp))
	}
	return param.ConvertParams(nps, methodType, (*context.Context)(ctx))
}
