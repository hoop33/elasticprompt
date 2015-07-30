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
		// CPU
		processors := "processors"
		if nodeOS.AvailableProcessors == 1 {
			processors = "processor"
		}
		buffer.WriteString(fmt.Sprintf("%s %s -- %d %s, %dMHz, %dkb cache\n%d sockets; %d cores\n",
			nodeOS.CPU.Vendor,
			nodeOS.CPU.Model,
			nodeOS.AvailableProcessors,
			processors,
			nodeOS.CPU.MHz,
			nodeOS.CPU.CacheSizeInBytes,
			nodeOS.CPU.TotalSockets,
			nodeOS.CPU.TotalCores))

		// Memory/Swap
		buffer.WriteString(fmt.Sprintf("Memory/Swap: %dkb/%dkb\n", nodeOS.Mem.TotalInBytes/1024, nodeOS.Swap.TotalInBytes/1024))
	}
	return buffer.String()
}
