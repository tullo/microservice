package ${SERVICE}

// file is autogenerated, do not modify here, see
// generator and template: templates/client_client.go.tpl

import (
	"net/http"

	"${MODULE}/rpc/${SERVICE}"
)

// New creates a ${SERVICE_CAMEL} RPC client.
func New() ${SERVICE}.${SERVICE_CAMEL}Service {
	return NewCustom("http://${SERVICE}.service:3000", &http.Client{})
}

// NewCustom creates a ${SERVICE_CAMEL} RPC client with custom Address/HTTP Client.
func NewCustom(addr string, client ${SERVICE}.HTTPClient) ${SERVICE}.${SERVICE_CAMEL}Service {
	return ${SERVICE}.New${SERVICE_CAMEL}ServiceJSONClient(addr, client)
}