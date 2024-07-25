package syntaxBase

import "testing"

func TestOption(t *testing.T) {
	server := NewServer()

	if server.Name != "server" {
		t.Errorf("server name is not server, got %s", server.Name)
	}

	if server.Addr != "0.0.0.0" {
		t.Errorf("server addr is not 0.0.0.0, got %s", server.Addr)
	}

	if server.Port != 8080 {
		t.Errorf("server port is not 8080, got %d", server.Port)
	}

	server = NewServer(
		WithAddr("127.0.0.1"),
		WithName("server2"),
		WithPort(8081),
	)

	if server.Name != "server2" {
		t.Errorf("server name is not server2, got %s", server.Name)
	}

	if server.Addr != "127.0.0.1" {
		t.Errorf("server addr is not 127.0.0.1, got %s", server.Addr)
	}

	if server.Port != 8081 {
		t.Errorf("server port is not 8081, got %d", server.Port)
	}
}
