package client

import "testing"

func TestConsul_NewConsulClient(t *testing.T) {

	addr := "http://consul.tiorap.systems"
	token := "afb3bc3a-6acd-11e7-b70c-784f43a63381"

	_, err := NewConsulClient(addr, token)

	if err != nil {
		t.Fatalf("error creating Consul client %s", err)
	}
}
