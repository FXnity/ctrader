package ctrader

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/fxnity/ctrader/openapi"
)

type mockClient struct {
	mock.Mock
	clientTransport
	t     *testing.T
	count atomic.Int64
}

func (m *mockClient) send(payload []byte) error {
	var msg openapi.ProtoMessage
	require.NoError(m.t, proto.Unmarshal(payload, &msg))
	require.Equal(m.t, msg.GetPayloadType(), uint32(openapi.ProtoPayloadType_HEARTBEAT_EVENT))
	m.count.Add(1)
	return nil
}

func TestClientKeepAlive(t *testing.T) {
	t.Parallel()
	mc := mockClient{t: t}
	c := Client{transport: &mc}
	c.keepalive()
	time.Sleep(21 * time.Second)
	require.Equal(t, int64(2), mc.count.Load())
	c.stopSignal.Store(true)
	time.Sleep(11 * time.Second)
	require.Equal(t, int64(2), mc.count.Load())
}
