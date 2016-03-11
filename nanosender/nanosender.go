package nanosender

import (
	"errors"
	"sync"
	"time"

	// "github.com/deepglint/glog"
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pair"
	"github.com/gdamore/mangos/protocol/req"
	"github.com/gdamore/mangos/transport/ipc"
	// "github.com/gdamore/mangos/transport/tcp"
	// nano "github.com/op/go-nanomsg"
)

var (
	ERROR_SOCKET = errors.New("can not create socket connection")
	ERROR_LISTEN = errors.New("can not create ipc listener")
)

type NanoRequest struct {
	SessionId string
	Cmd       string
	Params    map[string]interface{}
}

type NanoResponse struct {
	SessionId string
	Status    int
	Body      interface{}
}

type NanoSender struct {
	Mutex  *sync.Mutex
	Socket mangos.Socket
}

func NewNanoSender(url string) (*NanoSender, error) {

	var sock mangos.Socket
	var err error
	if sock, err = req.NewSocket(); err != nil {
		return nil, ERROR_SOCKET
	}
	sock.SetOption(mangos.OptionRetryTime, 0)
	sock.AddTransport(ipc.NewTransport())
	// sock.AddTransport(tcp.NewTransport())

	if err = sock.Dial(url); err != nil {
		return nil, ERROR_LISTEN
	}

	mutex := new(sync.Mutex)

	return &NanoSender{mutex, sock}, nil
}

func NewNanoPairSocket(url string) (*NanoSender, error) {

	var sock mangos.Socket
	var err error
	if sock, err = pair.NewSocket(); err != nil {
		return nil, ERROR_SOCKET
	}
	sock.AddTransport(ipc.NewTransport())
	// sock.AddTransport(tcp.NewTransport())

	if err = sock.Dial(url); err != nil {
		return nil, ERROR_LISTEN
	}

	mutex := new(sync.Mutex)

	return &NanoSender{mutex, sock}, nil
}

func (this *NanoSender) Send(data []byte) (err error) {
	this.Mutex.Lock()
	// glog.Infoln(string(data))
	err = this.Socket.Send(data)

	return
}

func (this *NanoSender) Recv() (body []byte, err error) {
	body, err = this.Socket.Recv()
	this.Mutex.Unlock()
	// glog.Infoln(string(body))
	return
}

func (this *NanoSender) RecvTimeout(o int) (body []byte, err error) {
	timeout := time.Duration(o) * time.Second

	this.Socket.SetOption(mangos.OptionRecvDeadline, timeout)

	body, err = this.Socket.Recv()
	this.Mutex.Unlock()
	// glog.Infoln(string(body))
	return
}

func (this *NanoSender) Close() error {
	return this.Socket.Close()
}

/*
type NanoSender struct {
	SockFd nano.ReqSocket
}

func NewNanoSender(url string, resendInteval, sendTimout, recvTimeout int) *NanoSender {
	sock, err := nano.NewReqSocket()
	if err != nil {
		glog.Errorln("Create new nano socket failed: ", err)
		return nil
	}

	if _, err = sock.Connect(url); err != nil {
		glog.Errorf("Connect to nano socket %s failed: ", url, err)
		return nil
	}

	sock.SetResendInterval(time.Duration(resendInteval) * time.Second)
	sock.SetSendTimeout(time.Duration(sendTimout) * time.Second)
	sock.SetRecvTimeout(time.Duration(recvTimeout) * time.Second)

	return &NanoSender{
		SockFd: *sock}
}

func (this *NanoSender) Send(data []byte, block int) (int, error) {
	return this.SockFd.Send(data, block)
}

func (this *NanoSender) Recv(block int) ([]byte, error) {
	return this.SockFd.Recv(block)
}
*/
