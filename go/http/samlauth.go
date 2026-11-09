package http

import (
	"net/http"

	saml "github.com/RobotsAndPencils/go-saml"
)

// samlIdPCertPath is the identity provider certificate that signs SAML requests.
const samlIdPCertPath = "/etc/orchestrator/saml-idp.crt"

// samlRequestValid reports whether req carries a SAML AuthnRequest signed by the identity provider.
func samlRequestValid(req *http.Request) bool {
	authnRequest, err := saml.ParseEncodedRequest(req.FormValue("SAMLRequest"))
	if err != nil {
		return false
	}
	return authnRequest.Validate(samlIdPCertPath) == nil
}
