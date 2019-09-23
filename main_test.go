package main

import (
	"testing"

	"github.com/gcash/bchd/bchrpc/pb"

	"github.com/linkpoolio/bridges/bridge"
	"github.com/stretchr/testify/assert"
)

func TestBchd_Opts(t *testing.T) {
	bchd := Bchd{}
	opts := bchd.Opts()

	assert.Equal(t, opts.Name, "bchd")
	assert.True(t, opts.Lambda)
}

func TestBchd_BlockchainInfo(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc": "blockchainInfo",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetBlockchainInfoResponse{}, obj)
}

func TestBchd_MempoolInfo(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc": "mempoolInfo",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetMempoolInfoResponse{}, obj)
}

func TestBchd_MempoolNoParam(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc": "mempool",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetMempoolResponse{}, obj)
	resp := obj.(*pb.GetMempoolResponse)
	assert.IsType(t, &pb.GetMempoolResponse_TransactionData_Transaction{}, resp.TransactionData[0].TxidsOrTxs)
}

func TestBchd_MempoolFullTrx(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc":    "mempool",
		"fullTrx": "true",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetMempoolResponse{}, obj)
	resp := obj.(*pb.GetMempoolResponse)
	assert.IsType(t, &pb.GetMempoolResponse_TransactionData_Transaction{}, resp.TransactionData[0].TxidsOrTxs)
}

func TestBchd_MempoolTrxHash(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc":    "mempool",
		"fullTrx": "false",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetMempoolResponse{}, obj)
	resp := obj.(*pb.GetMempoolResponse)
	assert.IsType(t, &pb.GetMempoolResponse_TransactionData_TransactionHash{}, resp.TransactionData[0].TxidsOrTxs)
}

func TestBchd_MempoolBadTrx(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc":    "mempool",
		"fullTrx": "not a boolean value",
	}

	query, _ := bridge.ParseInterface(data)
	obj, err := bchd.Run(bridge.NewHelper(query))
	assert.Nil(t, err)
	assert.IsType(t, &pb.GetMempoolResponse{}, obj)
	resp := obj.(*pb.GetMempoolResponse)
	assert.IsType(t, &pb.GetMempoolResponse_TransactionData_Transaction{}, resp.TransactionData[0].TxidsOrTxs)
}

func TestBchd_BadProc(t *testing.T) {
	bchd := newBchd(getCredentials(), getRPCURL())
	data := map[string]interface{}{
		"proc": "not a proc",
	}

	query, _ := bridge.ParseInterface(data)
	_, err := bchd.Run(bridge.NewHelper(query))
	assert.NotNil(t, err)
	assert.Equal(t, "unrecognized or unsupported bchd gRPC method", err.Error())
}
