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
	doc := `Intel Itanium -- 2 processors, 4800MHz, 512kb cache
2 sockets; 4 cores
Memory/Swap: 1048576kb/2097152kb
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
	doc := `  -- 1 processor, 0MHz, 0kb cache
0 sockets; 0 cores
Memory/Swap: 0kb/0kb
`
	nodeOS := &elastic.NodesInfoNodeOS{
		RefreshInterval:     "5s",
		AvailableProcessors: 1,
	}
	assert.Equal(t, doc, NewOSWrapper(nodeOS).String())
}
