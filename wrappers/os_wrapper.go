package wrappers

import (
	"bytes"
	"fmt"

	"github.com/olivere/elastic"
)

type OSWrapper struct {
	NodeOS *elastic.NodesInfoNodeOS
}

func NewOSWrapper(nodeOS *elastic.NodesInfoNodeOS) *OSWrapper {
	return &OSWrapper{
		NodeOS: nodeOS,
	}
}

func (wrapper *OSWrapper) String() string {
	var buffer bytes.Buffer
	if nodeOS := wrapper.NodeOS; nodeOS != nil {
		// Memory
		buffer.WriteString(fmt.Sprintf("Memory: %s/%d\n", nodeOS.Mem.Total, nodeOS.Mem.TotalInBytes))

		// Swap
		buffer.WriteString(fmt.Sprintf("Swap: %s/%d\n", nodeOS.Swap.Total, nodeOS.Swap.TotalInBytes))

		// Procs and Refresh
		processors := "processors"
		if nodeOS.AvailableProcessors == 1 {
			processors = "processor"
		}
		buffer.WriteString(fmt.Sprintf("%d %s; %s refresh interval\n", nodeOS.AvailableProcessors, processors, nodeOS.RefreshInterval))

		// CPU
		buffer.WriteString(fmt.Sprintf("%s %s %dMHz %db cache\n%d sockets; %d cores; %d cores/socket\n",
			nodeOS.CPU.Vendor,
			nodeOS.CPU.Model,
			nodeOS.CPU.MHz,
			nodeOS.CPU.CacheSizeInBytes,
			nodeOS.CPU.TotalSockets,
			nodeOS.CPU.TotalCores,
			nodeOS.CPU.CoresPerSocket))
	}
	return buffer.String()
}
