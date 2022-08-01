package bcapi

import (
	"app-hyperledger/pkg/models"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// const (
// 	mspID         = "Org1MSP"
// 	pathBC        = "/home/pluge/go/src/github.com/SiriusPluge/fabric-samples/"
// 	certPath      = pathBC + "test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem"
// 	keyPath       = pathBC + "test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/"
// 	tlsCertPath   = pathBC + "test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt"
// 	peerEndpoint  = "localhost:7051"
// 	gatewayPeer   = "peer0.org1.example.com"
// 	channelName   = "mychannel"
// 	chaincodeName = "private"
// )

func Init(cfg *models.ConfigBC) (*Instance, error) {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection, err := newGrpcConnection(cfg)
	if err != nil {
		return nil, ErrorHandling(err)
	}
	// defer clientConnection.Close()

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
	// defer gateway.Close()

	network := gateway.GetNetwork(cfg.ChannelName)
	contract := network.GetContract(cfg.ChaincodeName)

	return &Instance{
		Conn:     clientConnection,
		Gateway:  gateway,
		Signer:   sign,
		Contract: contract,
	}, nil
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection(cfg *models.ConfigBC) (*grpc.ClientConn, error) {
	certificate, err := loadCertificate(cfg.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cfg.GatewayPeer)

	connection, err := grpc.Dial(cfg.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection, nil
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(cfg *models.ConfigBC) *identity.X509Identity {
	certificate, err := loadCertificate(cfg.Cert)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(cfg.OrgID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(cfg *models.ConfigBC) identity.Sign {
	files, err := ioutil.ReadDir(cfg.PrivKey)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(cfg.PrivKey, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}
