// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "apps-management": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-apps-management/design
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-apps-management
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AppsController is the controller interface for the Apps actions.
type AppsController interface {
	goa.Muxer
	DeleteApp(*DeleteAppAppsContext) error
	Get(*GetAppsContext) error
	GetMyApps(*GetMyAppsAppsContext) error
	GetUserApps(*GetUserAppsAppsContext) error
	RegenerateClientSecret(*RegenerateClientSecretAppsContext) error
	RegisterApp(*RegisterAppAppsContext) error
	UpdateApp(*UpdateAppAppsContext) error
	VerifyApp(*VerifyAppAppsContext) error
}

// MountAppsController "mounts" a Apps resource controller on the given service.
func MountAppsController(service *goa.Service, ctrl AppsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteAppAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteApp(rctx)
	}
	service.Mux.Handle("DELETE", "/apps/:appId", ctrl.MuxHandler("deleteApp", h, nil))
	service.LogInfo("mount", "ctrl", "Apps", "action", "DeleteApp", "route", "DELETE /apps/:appId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/apps/:appId", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Apps", "action", "Get", "route", "GET /apps/:appId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetMyAppsAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetMyApps(rctx)
	}
	service.Mux.Handle("GET", "/apps/my", ctrl.MuxHandler("getMyApps", h, nil))
	service.LogInfo("mount", "ctrl", "Apps", "action", "GetMyApps", "route", "GET /apps/my")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserAppsAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.GetUserApps(rctx)
	}
	service.Mux.Handle("GET", "/apps/users/:userId/all", ctrl.MuxHandler("getUserApps", h, nil))
	service.LogInfo("mount", "ctrl", "Apps", "action", "GetUserApps", "route", "GET /apps/users/:userId/all")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegenerateClientSecretAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.RegenerateClientSecret(rctx)
	}
	service.Mux.Handle("PUT", "/apps/:appId/regenerate-secret", ctrl.MuxHandler("regenerateClientSecret", h, nil))
	service.LogInfo("mount", "ctrl", "Apps", "action", "RegenerateClientSecret", "route", "PUT /apps/:appId/regenerate-secret")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewRegisterAppAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AppPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.RegisterApp(rctx)
	}
	service.Mux.Handle("POST", "/apps", ctrl.MuxHandler("registerApp", h, unmarshalRegisterAppAppsPayload))
	service.LogInfo("mount", "ctrl", "Apps", "action", "RegisterApp", "route", "POST /apps")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAppAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AppPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateApp(rctx)
	}
	service.Mux.Handle("PUT", "/apps/:appId", ctrl.MuxHandler("updateApp", h, unmarshalUpdateAppAppsPayload))
	service.LogInfo("mount", "ctrl", "Apps", "action", "UpdateApp", "route", "PUT /apps/:appId")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewVerifyAppAppsContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AppCredentialsPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.VerifyApp(rctx)
	}
	service.Mux.Handle("POST", "/apps/verify", ctrl.MuxHandler("verifyApp", h, unmarshalVerifyAppAppsPayload))
	service.LogInfo("mount", "ctrl", "Apps", "action", "VerifyApp", "route", "POST /apps/verify")
}

// unmarshalRegisterAppAppsPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterAppAppsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &appPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAppAppsPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAppAppsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &appPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalVerifyAppAppsPayload unmarshals the request body into the context request data Payload field.
func unmarshalVerifyAppAppsPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &appCredentialsPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui/*filepath", "swagger-ui/dist")
	service.Mux.Handle("GET", "/swagger-ui/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist", "route", "GET /swagger-ui/*filepath")

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/swagger-ui/", "swagger-ui/dist/index.html")
	service.Mux.Handle("GET", "/swagger-ui/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger-ui/dist/index.html", "route", "GET /swagger-ui/")
}
