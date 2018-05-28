// ipc_test
package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"OK", "ECHO: " + method + " ~ " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcSever(&EchoServer{})

	clent1 := NewIpcClient(server)
	clent2 := NewIpcClient(server)

	resp1, _ := clent1.Call("foo", "From client1")
	resp2, _ := clent2.Call("foo", "From client2")

	if resp1.Body != "ECHO: foo ~ From client1" ||
		resp2.Body != "ECHO: foo ~ From client2" {
		t.Error("IpcClient.Call failed. resp1 :", resp1, "resp2: ", resp2)

	}
	clent1.Close()
	clent2.Close()
}
