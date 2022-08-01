package models

func (c *UserCert) MspId() string {
	return "Org" + c.OrgID + "MSP"
}
