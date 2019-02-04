package testutil

import (
	"testing"

	"github.com/elsevier-core-engineering/replicator/replicator/structs"
	"github.com/hashicorp/consul/testutil"
)

// MakeClientWithConfig will be written by Eric.
func MakeClientWithConfig(t *testing.T) (*structs.Config, *testutil.TestServer) {

	srv1 := testutil.NewTestServerConfig(t, func(c *testutil.TestServerConfig) {
		c.Bind = "127.0.0.1"
	})

	config := &structs.Config{
		ConsulKeyRoot: "replicator/config",
		Notification:  &structs.Notification{},
		Consul:        srv1.HTTPAddr,
	}

	return config, srv1
}
