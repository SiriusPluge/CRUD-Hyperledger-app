package bcapi

import (
	"context"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	gwproto "github.com/hyperledger/fabric-protos-go/gateway"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
func ErrorHandling(err error) error {

	switch err := err.(type) {
	case *client.EndorseError:
		logrus.Errorf("endorse error with gRPC status %v: %s", status.Code(err), err)
	case *client.SubmitError:
		logrus.Errorf("submit error with gRPC status %v: %s", status.Code(err), err)
	case *client.CommitStatusError:
		if errors.Is(err, context.DeadlineExceeded) {
			logrus.Errorf("timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
		} else {
			logrus.Errorf("error obtaining commit status with gRPC status %v: %s", status.Code(err), err)
		}
	case *client.CommitError:
		logrus.Errorf("transaction %s failed to commit with status %d: %s", err.TransactionID, int32(err.Code), err)
	}
	/*
	 Any error that originates from a peer or orderer node external to the gateway will have its details
	 embedded within the gRPC status error. The following code shows how to extract that.
	*/
	statusErr := status.Convert(err)
	for _, detail := range statusErr.Details() {
		errDetail := detail.(*gwproto.ErrorDetail)
		logrus.Errorf("error from endpoint: %s, mspId: %s, message: %s", errDetail.Address, errDetail.MspId, errDetail.Message)
	}

	return err
}
