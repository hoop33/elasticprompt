package wrappers

import (
	"testing"

	"github.com/olivere/elastic"
	"github.com/stretchr/testify/assert"
)

func TestNilOSReturnsEmpty(t *testing.T) {
	assert.Equal(t, "", NewOSWrapper(nil).String())
}

func TestOSInformationIsPresent(t *testing.T) {
	doc := `Memory: 1gb/1073741824
Swap: 2gb/2147483648
2 processors; 5s refresh interval
Intel Itanium 4800MHz 512b cache
2 sockets; 4 cores; 2 cores/socket
`

	nodeOS := &elastic.NodesInfoNodeOS{
		RefreshInterval:     "5s",
		AvailableProcessors: 2,
		Mem: struct {
			Total        string `json:"total"`
			TotalInBytes int    `json:"total_in_bytes"`
		}{
			Total:        "1gb",
			TotalInBytes: 1073741824,
		},
		Swap: struct {
			Total        string `json:"total"`
			TotalInBytes int    `json:"total_in_bytes"`
		}{
			Total:        "2gb",
			TotalInBytes: 2147483648,
		},
		CPU: struct {
			Vendor           string `json:"vendor"`
			Model            string `json:"model"`
			MHz              int    `json:"mhz"`
			TotalCores       int    `json:"total_cores"`
			TotalSockets     int    `json:"total_sockets"`
			CoresPerSocket   int    `json:"cores_per_socket"`
			CacheSizeInBytes int    `json:"cache_size_in_bytes"`
		}{
			Vendor:           "Intel",
			Model:            "Itanium",
			MHz:              4800,
			TotalCores:       4,
			TotalSockets:     2,
			CoresPerSocket:   2,
			CacheSizeInBytes: 512,
		}}
	assert.Equal(t, doc, NewOSWrapper(nodeOS).String())
}

func TestSingularFormsDropS(t *testing.T) {
	doc := `Memory: /0
Swap: /0
1 processor; 5s refresh interval
  0MHz 0b cache
0 sockets; 0 cores; 0 cores/socket
`
	nodeOS := &elastic.NodesInfoNodeOS{
		RefreshInterval:     "5s",
		AvailableProcessors: 1,
	}
	assert.Equal(t, doc, NewOSWrapper(nodeOS).String())
}
