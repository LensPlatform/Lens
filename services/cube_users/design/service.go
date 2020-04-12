package design

import (
	model "github.com/LensPlatform/Lens/services/cube_users/pkg/models/proto"

	. "github.com/goadesign/goa/dsl"
)

var service = Service("users-microservice", func() {
	Description("The users microservice exposes endpoints useful in accessing various schema types")
	Error("unauthorized", String, "Credentials are invalid")
	// The "timeout" error is also defined at the service level.
	Error("timeout", ErrorResult, "operation timed out, retry later.", func() {
		// Timeout indicates an error due to a timeout.
		Timeout()
		// Temporary indicates that the request may be retried.
		Temporary()
	})
	HTTP(func() {
		Path("/users-microservice")
		Response("unauthorized", StatusUnauthorized)
	})

	GRPC(func() {
		Response("unauthorized", CodeUnauthenticated)
	})

	Method("signin", func() {
		Description("Creates a valid JWT")

		// The signin endpoint is secured via basic auth
		Security(BasicAuth)

		Payload(func() {
			Description("Credentials used to authenticate to retrieve JWT token")
			UsernameField(1, "username", String, "Username used to perform signin", func() {
				Example("username")
			})
			PasswordField(2, "password", String, "Password used to perform signin", func() {
				Example("password")
			})
			Field(1, "email", String, "Email used to perform sigin", func() {
				Example("userame@gmail.com")
			})
			Required("username", "password")
		})

		Result(Credentials)

		HTTP(func() {
			POST("/signin")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("CreateUser", func() {
		Payload("CreateUser", func() {
			Attribute("user", model.User{}, "User to be created", func() {
				Meta("rpc:tag", "1")
			})
			Required("user")
		})

		Result(Int)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			// Requests to the service consist of HTTP POST requests.
			// The payload fields are encoded as path parameters.
			POST("/user/user-account?payload={user}")
			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
			Response("timeout", StatusGatewayTimeout)
		})

		// GRPC describes the gRPC transport mapping.
		GRPC(func() {
			// Responses use a "OK" gRPC code.
			// The result is encoded in the response message.
			Response(CodeOK)
			Response("timeout", CodeDeadlineExceeded)
		})
	})

	Method("CreateProfile", func() {
		Payload("CreateProfile", func() {
			Attribute("profile", model.Profile{}, "Profile", func() {
				Meta("rpc:tag", "1")
			})
			Attribute("user_id", Int64, "user id token which the profile is tied to", func() {
				Meta("rpc:tag", "2")
			})
			Required("profile", "user_id")
		})

		Result(Int)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			POST("/user/user-profile?payload1={profile}&payload2={user_id}")
			Response(StatusOK)
			Response("timeout", StatusGatewayTimeout)
		})

		// GRPC describes the gRPC transport mapping.
		GRPC(func() {
			Response(CodeOK)
			Response("timeout", CodeDeadlineExceeded)
		})
	})

	Method("CreateUserSubscription", func() {
		Payload("CreateUserSubscription", func() {
			Attribute("subscription", model.Subscriptions{}, "User Subscription", func() {
				Meta("rpc:tag", "1")
			})
			Attribute("user_id", Int, "user id to which the subscription is to be created for", func() {
				Meta("rpc:tag", "2")
			})
			Required("subscription", "user_id")
		})

		Result(Int)

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests.
			// The payload fields are encoded as path parameters.
			POST("/user/subscription?payload1={subscription}&&payload2={user_id}")
			// Responses use a "200 OK" HTTP status.
			// The result is encoded in the response body.
			Response(StatusOK)
			Response("timeout", StatusGatewayTimeout)
		})

		// GRPC describes the gRPC transport mapping.
		GRPC(func() {
			Response(CodeOK)
			Response("timeout", CodeDeadlineExceeded)
		})
	})

	Method("GetUser", func() {
		Payload("GetUser", func() {
			Attribute("user_id", Int64, "User id", func() {
				Meta("rpc:tag", "1")
			})
			Required("user_id")
		})

		Result(model.UserORM{})

		// HTTP describes the HTTP transport mapping.
		HTTP(func() {
			GET("/user/{user_id}")
			Response(StatusOK)
			Response("timeout", StatusGatewayTimeout)
		})

		// GRPC describes the gRPC transport mapping.
		GRPC(func() {
			Response(CodeOK)
			Response("timeout", CodeDeadlineExceeded)
		})
	})

	// Serve the file with relative path ../../gen/http/openapi.json for
	// requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json")
})

var swaggerService = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})
	Files("/swagger.json", "gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})