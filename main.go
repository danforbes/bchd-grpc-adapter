package main

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/gcash/bchd/bchrpc/pb"
	"github.com/linkpoolio/bridges/bridge"
)

// Bchd is the bchd gRPC server
type Bchd struct {
	client pb.BchrpcClient
}

// Opts is the options for the bchd bridge
func (bchd *Bchd) Opts() *bridge.Opts {
	return &bridge.Opts{
		Name:   "bchd",
		Lambda: true,
	}
}

// Run is the main bchd gRPC adapter implementation
func (bchd *Bchd) Run(h *bridge.Helper) (interface{}, error) {
	switch h.GetParam("proc") {
	case "mempoolInfo":
		return bchd.client.GetMempoolInfo(context.Background(), &pb.GetMempoolInfoRequest{})
	case "mempool":
		fullTrx, err := strconv.ParseBool(h.GetParam("fullTrx"))
		if err != nil {
			fullTrx = true
		}

		return bchd.client.GetMempool(context.Background(), &pb.GetMempoolRequest{FullTransactions: fullTrx})
	case "blockchainInfo":
		return bchd.client.GetBlockchainInfo(context.Background(), &pb.GetBlockchainInfoRequest{})
	default:
		return nil, errors.New("unrecognized or unsupported bchd gRPC method")
	}
}

func getRPCURL() *url.URL {
	rawURL, ok := os.LookupEnv("BCHD_GRPC_URL")
	if !ok {
		rawURL = "bchd.greyh.at:8335"
	}

	url, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	return url
}

func getCredentials() credentials.TransportCredentials {
	certPath, ok := os.LookupEnv("BCHD_CERT_PATH")
	if !ok {
		return credentials.NewClientTLSFromCert(nil, os.Getenv("BCHD_SERVER_OVERRIDE"))
	}

	cert, err := credentials.NewClientTLSFromFile(certPath, os.Getenv("BCHD_SERVER_OVERRIDE"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return cert
}

func newBchd(cred credentials.TransportCredentials, url *url.URL) *Bchd {
	conn, err := grpc.Dial(url.String(), grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &Bchd{pb.NewBchrpcClient(conn)}
}

func main() {
	bridge.NewServer(newBchd(getCredentials(), getRPCURL())).Start(8080)
}
