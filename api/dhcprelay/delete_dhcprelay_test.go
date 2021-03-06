package dhcprelay

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteDhcpRelayAPI *DeleteDhcpRelayAPI

func setupDelete() {
	deleteDhcpRelayAPI = NewDelete("edge-50")
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteDhcpRelayAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-50/dhcp/config/relay", deleteDhcpRelayAPI.Endpoint())
}
