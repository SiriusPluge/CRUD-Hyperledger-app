package models

type ReqBody struct {
	ConfigBC    `json:"config_bc"`
	Services    `json:"services,omitempty"`
	InputType   string `json:"input_type,omitempty"`
	InputPhone  string `json:"input_phone,omitempty"`
	InputStatus string `json:"input_status,omitempty"`
}

// Конфигурация для BC
type ConfigBC struct {
	// MspID            string `json:"mspID"`
	// CertPath         string `json:"certPath"`
	// KeyPath          string `json:"keyPath"`
	TlsCertPath   string `json:"tls_cert_path"`
	PeerEndpoint  string `json:"peer_endpoint"`
	GatewayPeer   string `json:"gateway_peer"`
	ChannelName   string `json:"channel_name"`
	ChaincodeName string `json:"chaincode_name"`
	UserCert      `json:"user_cert"`
}

// Пользовательский сертификат для BC API
type UserCert struct {
	OrgID    string `json:"org_id"`
	ClientID string `json:"client_id"`
	Cert     string `json:"cert"`
	PrivKey  string `json:"priv"`
}

// Услуга
type Services struct {
	TypeServices string `json:"type_services"`       // Тип услуги
	Comment      string `json:"comment,omitempty"`   // Комментарий
	Email        string `json:"email"`               // Рабочий E-mail адрес
	FirstName    string `json:"first_name"`          // Фамилия
	LastName     string `json:"last_name,omitempty"` // Имя
	Phone        string `json:"phone"`               // Телефон
	Address      string `json:"address"`             // Адрес
	Status       int    `json:"status"`              // Статус
}
