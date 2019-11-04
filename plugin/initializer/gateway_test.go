package initializer

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/plugin/proto"

	"github.com/ethereum/go-ethereum/plugin/proto/mock_proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPluginGateway_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &proto.PluginInitialization_Request{
		HostIdentity:     "arbitraryName",
		RawConfiguration: []byte("arbitrary config"),
	}

	mockClient := mock_proto.NewMockPluginInitializerClient(ctrl)
	mockClient.
		EXPECT().
		Init(gomock.Any(), gomock.Eq(req)).
		Return(&proto.PluginInitialization_Response{}, nil)

	testObject := &PluginGateway{client: mockClient}

	err := testObject.Init(context.Background(), req.HostIdentity, req.RawConfiguration)

	assert.NoError(t, err)
}
