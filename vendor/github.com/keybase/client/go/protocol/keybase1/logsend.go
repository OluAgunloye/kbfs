// Auto-generated by avdl-compiler v1.3.16 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/logsend.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type PrepareLogsendArg struct {
}

func (o PrepareLogsendArg) DeepCopy() PrepareLogsendArg {
	return PrepareLogsendArg{}
}

type LogsendInterface interface {
	PrepareLogsend(context.Context) error
}

func LogsendProtocol(i LogsendInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.logsend",
		Methods: map[string]rpc.ServeHandlerDescription{
			"prepareLogsend": {
				MakeArg: func() interface{} {
					ret := make([]PrepareLogsendArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.PrepareLogsend(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LogsendClient struct {
	Cli rpc.GenericClient
}

func (c LogsendClient) PrepareLogsend(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.logsend.prepareLogsend", []interface{}{PrepareLogsendArg{}}, nil)
	return
}
