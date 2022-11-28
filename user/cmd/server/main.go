package main

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/tracer"
	"user-center-service/internal"
)

func main() {
	tracer.SetTraceId("", func() {
		gone.Serve(internal.MasterPriest)
	})
}
