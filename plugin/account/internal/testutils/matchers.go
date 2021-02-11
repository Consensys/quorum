package testutils

import (
	"fmt"

	protobuf "github.com/golang/protobuf/proto"
	"github.com/jpmorganchase/quorum-account-plugin-sdk-go/proto"
)

type StatusRequestMatcher struct {
	R *proto.StatusRequest
}

func (m StatusRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.StatusRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m StatusRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type OpenRequestMatcher struct {
	R *proto.OpenRequest
}

func (m OpenRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.OpenRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m OpenRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type CloseRequestMatcher struct {
	R *proto.CloseRequest
}

func (m CloseRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.CloseRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m CloseRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type AccountsRequestMatcher struct {
	R *proto.AccountsRequest
}

func (m AccountsRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.AccountsRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m AccountsRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type ContainsRequestMatcher struct {
	R *proto.ContainsRequest
}

func (m ContainsRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.ContainsRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m ContainsRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type SignRequestMatcher struct {
	R *proto.SignRequest
}

func (m SignRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.SignRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m SignRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type UnlockAndSignRequestMatcher struct {
	R *proto.UnlockAndSignRequest
}

func (m UnlockAndSignRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.UnlockAndSignRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m UnlockAndSignRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type TimedUnlockRequestMatcher struct {
	R *proto.TimedUnlockRequest
}

func (m TimedUnlockRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.TimedUnlockRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m TimedUnlockRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type LockRequestMatcher struct {
	R *proto.LockRequest
}

func (m LockRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.LockRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m LockRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type NewAccountRequestMatcher struct {
	R *proto.NewAccountRequest
}

func (m NewAccountRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.NewAccountRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m NewAccountRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}

type ImportRawKeyRequestMatcher struct {
	R *proto.ImportRawKeyRequest
}

func (m ImportRawKeyRequestMatcher) Matches(x interface{}) bool {
	r, ok := x.(*proto.ImportRawKeyRequest)
	if !ok {
		return false
	}
	return protobuf.Equal(m.R, r)
}

func (m ImportRawKeyRequestMatcher) String() string {
	return fmt.Sprintf("is %v", m.R)
}
