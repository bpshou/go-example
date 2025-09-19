package {{.ServicePackage}}

import "gin_app/app/types"

func {{camelcase .ServiceName}}Logic(req *types.{{camelcase .ServiceName}}Req) (types.{{camelcase .ServiceName}}Resp, error) {
	return types.{{camelcase .ServiceName}}Resp{}, nil
}
