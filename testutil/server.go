package testutil

import (
	"testing"

	"github.com/elsevier-core-engineering/replicator/replicator/structs"
	"github.com/hashicorp/consul/testutil"
)

// MakeClientWithConfig will be written by Eric.
func MakeClientWithConfig(t *testing.T) (*structs.Config, *testutil.TestServer) {

	srv1 := testutil.NewTestServer(t)

	config := &structs.Config{
		ConsulKeyRoot: "replicator/config",
		Notification:  &structs.Notification{},
		Consul:        srv1.HTTPAddr,
	}

	return config, srv1
}
