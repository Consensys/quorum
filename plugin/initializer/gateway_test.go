package initializer

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/plugin/gen/proto_common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPluginGateway_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &proto_common.PluginInitialization_Request{
		HostIdentity:     "arbitraryName",
		RawConfiguration: []byte("arbitrary config"),
	}

	mockClient := proto_common.NewMockPluginInitializerClient(ctrl)
	mockClient.
		EXPECT().
		Init(gomock.Any(), gomock.Eq(req)).
		Return(&proto_common.PluginInitialization_Response{}, nil)

	testObject := &PluginGateway{client: mockClient}

	err := testObject.Init(context.Background(), req.HostIdentity, req.RawConfiguration)

	assert.NoError(t, err)
}
