package design

import (
	. "github.com/goadesign/goa/dsl"
)

// API describes the global properties of the API server.
var _ = API("Users-Microservice", server)

// Service describes a service
var _ = service