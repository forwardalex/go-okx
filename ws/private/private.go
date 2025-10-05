package private

import (
	"github.com/forwardalex/go-okx/common"
	"github.com/forwardalex/go-okx/ws"
	"github.com/gorilla/websocket"
)

type Private struct {
	Auth common.Auth
	C    *ws.Client
}

// new Private
func NewPrivate(auth common.Auth) *Private {
	private := &Private{
		Auth: auth,
		C:    ws.DefaultClientPrivate,
	}
	if auth.Simulated {
		private.C = ws.DefaultClientPrivateSimulated
	}
	return private
}

// subscribe
func (p *Private) Subscribe(args interface{}, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(subscribe, p.Login)
}

// loging private
func (p *Private) Login(conn *websocket.Conn) error {
	args := ws.NewArgsLoginFromAuth(p.Auth)
	login := ws.NewOperateLogin(args)
	return p.C.MessageOperate(conn, login)
}
