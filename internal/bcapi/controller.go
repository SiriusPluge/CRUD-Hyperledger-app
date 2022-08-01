package bcapi

import (
	"app-hyperledger/pkg/models"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
)

type Instance struct {
	Conn     *grpc.ClientConn
	Gateway  *client.Gateway
	Signer   identity.Sign
	Contract *client.Contract
}

// Возможность создавать новое подключение к BC с новыми входными конфигурациями
func CreateInstance(cfg *models.ConfigBC) (*Instance, error) {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection, err := newGrpcConnection(cfg)
	if err != nil {
		return nil, ErrorHandling(err)
	}

	id := newIdentity(cfg)
	sign := newSign(cfg)

	// Create a Gateway connection for a specific client identity
	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}

	network := gateway.GetNetwork(cfg.ChannelName)
	contract := network.GetContract(cfg.ChaincodeName)

	return &Instance{
		Conn:     clientConnection,
		Gateway:  gateway,
		Signer:   sign,
		Contract: contract,
	}, nil
}
