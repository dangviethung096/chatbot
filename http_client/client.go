package http_client

import (
	"sync"

	"github.com/dangviethung096/core"
)

// Create a new pool of http client
var httpClientPool = sync.Pool{
	New: func() interface{} {
		return core.NewClient()
	},
}

func Init(ctx core.Context) core.HttpClientBuilder {
	builder := httpClientPool.Get().(core.HttpClientBuilder)
	builder.SetContext(ctx)
	return builder
}
