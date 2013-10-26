package ipc

import (
    "testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
    resp := Response{method, params}
    return &resp
}

func (server *EchoServer) Name() string {
    return "EchoServer"
}

func TestIpc(t *testing.T) {
    server := NewIpcServer(&EchoServer{})

    client1 := NewIpcClient(server)
    client2 := NewIpcClient(server)

    resp1, _ := client1.Call("from client1", "get")
    resp2, _ := client2.Call("from client2", "get")

    if resp1.Body != "ECHO:from client1" || resp2.Body != "ECHO:from client2" {
        t.Error("Ipclient.call failed, resp1:", resp1, "resp2:", resp2)
    }

    client1.Close()
    client2.Close()
}
