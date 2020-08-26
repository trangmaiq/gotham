package sqlcon

import "runtime"

func maxParallelism() int {
	var (
		maxProcs = runtime.GOMAXPROCS(0)
		numCPU   = runtime.NumCPU()
	)

	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
