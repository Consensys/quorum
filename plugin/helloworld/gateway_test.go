package helloworld

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/mock_proto"
	"github.com/jpmorganchase/quorum-hello-world-plugin-sdk-go/proto"
	"github.com/stretchr/testify/assert"
)

func TestPluginPingPong_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	req := &proto.PluginHelloWorld_Request{
		Msg: "arbitrary msg",
	}
	mockClient := mock_proto.NewMockPluginGreetingClient(ctrl)
	mockClient.
		EXPECT().
		Greeting(gomock.Any(), gomock.Eq(req)).
		Return(&proto.PluginHelloWorld_Response{
			Msg: "arbitrary response",
		}, nil)
	testObject := &PluginGateway{client: mockClient}

	resp, err := testObject.Greeting(context.Background(), "arbitrary msg")

	assert.NoError(t, err)
	assert.Equal(t, "arbitrary response", resp)
}
