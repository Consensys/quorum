package qlight

import (
	"context"
	"testing"

	"github.com/ConsenSys/quorum-qlight-token-manager-plugin-sdk-go/mock_proto"
	"github.com/ConsenSys/quorum-qlight-token-manager-plugin-sdk-go/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPluginPingPong_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	req := &proto.TokenRefresh_Request{
		CurrentToken: "arbitrary msg",
		Psi:          "psi",
	}
	mockClient := mock_proto.NewMockPluginQLightTokenRefresherClient(ctrl)
	mockClient.
		EXPECT().
		TokenRefresh(gomock.Any(), gomock.Eq(req)).
		Return(&proto.TokenRefresh_Response{
			Token: "new token",
		}, nil)
	testObject := &PluginGateway{client: mockClient}

	resp, err := testObject.TokenRefresh(context.Background(), "arbitrary msg", "psi")

	assert.NoError(t, err)
	assert.Equal(t, "new token", resp)
}
