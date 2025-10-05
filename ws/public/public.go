package public

import (
	"github.com/forwardalex/go-okx/ws"
)

type Public struct {
	C *ws.Client
}

func NewPublic(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientPublic,
	}
	if simulated {
		public.C = ws.DefaultClientPublicSimulated
	}
	return public
}

func NewKlinePublic(simulated bool) *Public {
	public := &Public{
		C: ws.DefaultClientKlinePublic,
	}
	if simulated {
		public.C = ws.DefaultClientKlineSimulated
	}
	return public
}

// subscribe
func (p *Public) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}

// subscribes
func (p *Public) Subscribes(argss []interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribes(argss, handler, handlerError)
	return p.C.Operate(subscribe, nil)
}

// unbscribe
func (p *Public) UnSubscribe(args interface{}) error {
	unsubscribe := ws.NewOperationUnSubscribe(args)
	return p.C.Operate(unsubscribe, nil)
}
