package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func respError(c *gin.Context, err error) {
	logrus.Error(err)
	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func respOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

const (
	certAdmin = `-----BEGIN CERTIFICATE-----
	MIICrzCCAlWgAwIBAgIUd9iAPN5Rhhrue3fLjH0izUicDa4wCgYIKoZIzj0EAwIw
	cDELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMQ8wDQYDVQQH
	EwZEdXJoYW0xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh
	Lm9yZzEuZXhhbXBsZS5jb20wHhcNMjIwNzMxMjM0MDAwWhcNMjMwNzMxMjM0NTAw
	WjBgMQswCQYDVQQGEwJVUzEXMBUGA1UECBMOTm9ydGggQ2Fyb2xpbmExFDASBgNV
	BAoTC0h5cGVybGVkZ2VyMQ4wDAYDVQQLEwVhZG1pbjESMBAGA1UEAxMJb3JnMWFk
	bWluMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEanBniOQPrMO6GkwdVldSS/25
	pN+B2wikeCLHm/ziRyBPV8x3Nzcg5ysUopTZSDPkDEtdkZgcVTEP73Dsm7SOAaOB
	3DCB2TAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4EFgQUU/P9
	azgVxM05qHqngu9g1/yAFZUwHwYDVR0jBBgwFoAUFG6gfed8CZ+3e04aZIjKGUTD
	D54wHAYDVR0RBBUwE4IRYmVlbGluZXJvdXRlci5uZXQwWwYIKgMEBQYHCAEET3si
	YXR0cnMiOnsiaGYuQWZmaWxpYXRpb24iOiIiLCJoZi5FbnJvbGxtZW50SUQiOiJv
	cmcxYWRtaW4iLCJoZi5UeXBlIjoiYWRtaW4ifX0wCgYIKoZIzj0EAwIDSAAwRQIh
	AIe4Xrc0o5e08aSwkJIKn2BgqoEmkBpelN31aHqdnUOdAiBegPW4yxp6xKnW/qI/
	hevbDMZy6YWDpMxFEqwK3g9I/A==
	-----END CERTIFICATE-----`
	privKey = `-----BEGIN PRIVATE KEY-----
	MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgVb23HWtFege3CPkr
	c3RYDsOuuf3qPs5DNl3Ol8VdTUehRANCAARqcGeI5A+sw7oaTB1WV1JL/bmk34Hb
	CKR4Iseb/OJHIE9XzHc3NyDnKxSilNlIM+QMS12RmBxVMQ/vcOybtI4B
	-----END PRIVATE KEY-----`
)
