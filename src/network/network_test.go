package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"regexp"
	"net/http"
	"net/http/httptest"
)

func TestGetLocalIPAddress(t *testing.T) {
	IPAddress := GetLocalIPAddress()
	IPAddressRegexpPattern := "([0-9]*\\.){3}[0-9]*"
	re := regexp.MustCompile(IPAddressRegexpPattern)
	assert.Equal(t, true, re.MatchString(IPAddress))
}

func TestAddProtocolAndPortToIP(t *testing.T) {
	assert.Equal(t, "http://10.0.0.126:1234", AddProtocolAndPortToIP("10.0.0.126", 1234))
}

func TestSetMasterIP(t *testing.T) {

}

func TestSendSlaveURLToMaster(t *testing.T) {
	var numberOfMessagesSent = 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		numberOfMessagesSent++
	})
	testServer := httptest.NewServer(handler)

	sendSlaveURLToMaster("testSlaveName", "http://localhost:8080", testServer.URL)
	assert.Equal(t, 1, numberOfMessagesSent)
}

func TestSendSlaveURLToMaster_DEFAULT_SLAVE_NAME(t *testing.T) {
	var numberOfMessagesSent = 0
	handler := http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		numberOfMessagesSent++
	})
	testServer := httptest.NewServer(handler)

	sendSlaveURLToMaster("DEFAULT_SLAVE_NAME", "http://localhost:8080", testServer.URL)
	assert.Equal(t, 1, numberOfMessagesSent)
}
