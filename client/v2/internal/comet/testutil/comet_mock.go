// Code generated by MockGen. DO NOT EDIT.
// Source: client/v2/internal/comet/comet.go
//
// Generated by this command:
//
//	mockgen -source=client/v2/internal/comet/comet.go -package testutil -destination client/v2/internal/comet/testutil/comet_mock.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	bytes "github.com/cometbft/cometbft/libs/bytes"
	client "github.com/cometbft/cometbft/rpc/client"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	types "github.com/cometbft/cometbft/types"
	gomock "go.uber.org/mock/gomock"
)

// MockCometRPC is a mock of CometRPC interface.
type MockCometRPC struct {
	ctrl     *gomock.Controller
	recorder *MockCometRPCMockRecorder
	isgomock struct{}
}

// MockCometRPCMockRecorder is the mock recorder for MockCometRPC.
type MockCometRPCMockRecorder struct {
	mock *MockCometRPC
}

// NewMockCometRPC creates a new mock instance.
func NewMockCometRPC(ctrl *gomock.Controller) *MockCometRPC {
	mock := &MockCometRPC{ctrl: ctrl}
	mock.recorder = &MockCometRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCometRPC) EXPECT() *MockCometRPCMockRecorder {
	return m.recorder
}

// ABCIInfo mocks base method.
func (m *MockCometRPC) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIInfo", ctx)
	ret0, _ := ret[0].(*coretypes.ResultABCIInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIInfo indicates an expected call of ABCIInfo.
func (mr *MockCometRPCMockRecorder) ABCIInfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIInfo", reflect.TypeOf((*MockCometRPC)(nil).ABCIInfo), ctx)
}

// ABCIQuery mocks base method.
func (m *MockCometRPC) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQuery", ctx, path, data)
	ret0, _ := ret[0].(*coretypes.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQuery indicates an expected call of ABCIQuery.
func (mr *MockCometRPCMockRecorder) ABCIQuery(ctx, path, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQuery", reflect.TypeOf((*MockCometRPC)(nil).ABCIQuery), ctx, path, data)
}

// ABCIQueryWithOptions mocks base method.
func (m *MockCometRPC) ABCIQueryWithOptions(ctx context.Context, path string, data bytes.HexBytes, opts client.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQueryWithOptions", ctx, path, data, opts)
	ret0, _ := ret[0].(*coretypes.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQueryWithOptions indicates an expected call of ABCIQueryWithOptions.
func (mr *MockCometRPCMockRecorder) ABCIQueryWithOptions(ctx, path, data, opts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQueryWithOptions", reflect.TypeOf((*MockCometRPC)(nil).ABCIQueryWithOptions), ctx, path, data, opts)
}

// Block mocks base method.
func (m *MockCometRPC) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", ctx, height)
	ret0, _ := ret[0].(*coretypes.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block.
func (mr *MockCometRPCMockRecorder) Block(ctx, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockCometRPC)(nil).Block), ctx, height)
}

// BlockByHash mocks base method.
func (m *MockCometRPC) BlockByHash(ctx context.Context, hash []byte) (*coretypes.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash)
	ret0, _ := ret[0].(*coretypes.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockCometRPCMockRecorder) BlockByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockCometRPC)(nil).BlockByHash), ctx, hash)
}

// BlockResults mocks base method.
func (m *MockCometRPC) BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", ctx, height)
	ret0, _ := ret[0].(*coretypes.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults.
func (mr *MockCometRPCMockRecorder) BlockResults(ctx, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockCometRPC)(nil).BlockResults), ctx, height)
}

// BlockSearch mocks base method.
func (m *MockCometRPC) BlockSearch(ctx context.Context, query string, page, perPage *int, orderBy string) (*coretypes.ResultBlockSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockSearch", ctx, query, page, perPage, orderBy)
	ret0, _ := ret[0].(*coretypes.ResultBlockSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockSearch indicates an expected call of BlockSearch.
func (mr *MockCometRPCMockRecorder) BlockSearch(ctx, query, page, perPage, orderBy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSearch", reflect.TypeOf((*MockCometRPC)(nil).BlockSearch), ctx, query, page, perPage, orderBy)
}

// BlockchainInfo mocks base method.
func (m *MockCometRPC) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockchainInfo", ctx, minHeight, maxHeight)
	ret0, _ := ret[0].(*coretypes.ResultBlockchainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockchainInfo indicates an expected call of BlockchainInfo.
func (mr *MockCometRPCMockRecorder) BlockchainInfo(ctx, minHeight, maxHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInfo", reflect.TypeOf((*MockCometRPC)(nil).BlockchainInfo), ctx, minHeight, maxHeight)
}

// BroadcastTxAsync mocks base method.
func (m *MockCometRPC) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxAsync", ctx, tx)
	ret0, _ := ret[0].(*coretypes.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxAsync indicates an expected call of BroadcastTxAsync.
func (mr *MockCometRPCMockRecorder) BroadcastTxAsync(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxAsync", reflect.TypeOf((*MockCometRPC)(nil).BroadcastTxAsync), ctx, tx)
}

// BroadcastTxCommit mocks base method.
func (m *MockCometRPC) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxCommit", ctx, tx)
	ret0, _ := ret[0].(*coretypes.ResultBroadcastTxCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxCommit indicates an expected call of BroadcastTxCommit.
func (mr *MockCometRPCMockRecorder) BroadcastTxCommit(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxCommit", reflect.TypeOf((*MockCometRPC)(nil).BroadcastTxCommit), ctx, tx)
}

// BroadcastTxSync mocks base method.
func (m *MockCometRPC) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxSync", ctx, tx)
	ret0, _ := ret[0].(*coretypes.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxSync indicates an expected call of BroadcastTxSync.
func (mr *MockCometRPCMockRecorder) BroadcastTxSync(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxSync", reflect.TypeOf((*MockCometRPC)(nil).BroadcastTxSync), ctx, tx)
}

// Commit mocks base method.
func (m *MockCometRPC) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, height)
	ret0, _ := ret[0].(*coretypes.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockCometRPCMockRecorder) Commit(ctx, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockCometRPC)(nil).Commit), ctx, height)
}

// Status mocks base method.
func (m *MockCometRPC) Status(arg0 context.Context) (*coretypes.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*coretypes.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockCometRPCMockRecorder) Status(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCometRPC)(nil).Status), arg0)
}

// Tx mocks base method.
func (m *MockCometRPC) Tx(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", ctx, hash, prove)
	ret0, _ := ret[0].(*coretypes.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockCometRPCMockRecorder) Tx(ctx, hash, prove any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockCometRPC)(nil).Tx), ctx, hash, prove)
}

// TxSearch mocks base method.
func (m *MockCometRPC) TxSearch(ctx context.Context, query string, prove bool, page, perPage *int, orderBy string) (*coretypes.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", ctx, query, prove, page, perPage, orderBy)
	ret0, _ := ret[0].(*coretypes.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch.
func (mr *MockCometRPCMockRecorder) TxSearch(ctx, query, prove, page, perPage, orderBy any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockCometRPC)(nil).TxSearch), ctx, query, prove, page, perPage, orderBy)
}

// Validators mocks base method.
func (m *MockCometRPC) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", ctx, height, page, perPage)
	ret0, _ := ret[0].(*coretypes.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators.
func (mr *MockCometRPCMockRecorder) Validators(ctx, height, page, perPage any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockCometRPC)(nil).Validators), ctx, height, page, perPage)
}
