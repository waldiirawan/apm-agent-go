// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package apmechov4 // import "github.com/waldiirawan/apm-agent-go/module/apmechov4/v2"

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/labstack/echo/v4"

	"github.com/waldiirawan/apm-agent-go/module/apmhttp/v2"
	"github.com/waldiirawan/apm-agent-go/v2"
)

// Middleware returns a new Echo middleware handler for tracing
// requests and reporting errors.
//
// This middleware will recover and report panics, so it can
// be used instead of echo/middleware.Recover.
//
// By default, the middleware will use apm.DefaultTracer().
// Use WithTracer to specify an alternative tracer.
func Middleware(o ...Option) echo.MiddlewareFunc {
	opts := options{
		tracer: apm.DefaultTracer(),
	}
	for _, o := range o {
		o(&opts)
	}
	if opts.requestIgnorer == nil {
		opts.requestIgnorer = apmhttp.NewDynamicServerRequestIgnorer(opts.tracer)
	}
	if opts.requestName == nil {
		opts.requestName = func(c echo.Context) string {
			return c.Request().Method + " " + c.Path()
		}
	}
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		m := &middleware{
			tracer:         opts.tracer,
			handler:        h,
			requestIgnorer: opts.requestIgnorer,
			requestName:    opts.requestName,
		}
		return m.handle
	}
}

// RequestNameFunc should return span name for the given echo context
type RequestNameFunc func(c echo.Context) string

type middleware struct {
	handler        echo.HandlerFunc
	tracer         *apm.Tracer
	requestIgnorer apmhttp.RequestIgnorerFunc
	requestName    RequestNameFunc
}

func (m *middleware) handle(c echo.Context) error {
	req := c.Request()
	if !m.tracer.Recording() || m.requestIgnorer(req) {
		return m.handler(c)
	}
	name := m.requestName(c)
	tx, body, req := apmhttp.StartTransactionWithBody(m.tracer, name, req)
	defer tx.End()
	c.SetRequest(req)

	resp := c.Response()
	var handlerErr error
	defer func() {
		if v := recover(); v != nil {
			err, ok := v.(error)
			if !ok {
				err = errors.New(fmt.Sprint(v))
			}
			c.Error(err)

			e := m.tracer.Recovered(v)
			e.SetTransaction(tx)
			setContext(&e.Context, req, resp, body)
			e.Send()
		}
		if handlerErr != nil {
			e := m.tracer.NewError(handlerErr)
			setContext(&e.Context, req, resp, body)
			e.SetTransaction(tx)
			e.Handled = true
			e.Send()
		}
		tx.Result = apmhttp.StatusCodeResult(resp.Status)
		if tx.Sampled() {
			setContext(&tx.Context, req, resp, body)
		}
		body.Discard()
	}()

	handlerErr = m.handler(c)
	if handlerErr != nil {
		resp.Status = http.StatusInternalServerError
		if handlerErr, ok := handlerErr.(*echo.HTTPError); ok {
			resp.Status = handlerErr.Code
			reqPath := req.URL.RawPath
			if reqPath == "" {
				reqPath = req.URL.Path
			}
			if c.Path() == reqPath {
				// When c.Path() matches the request path exactly,
				// that means either there's no matching route, or
				// there's an exactly matching route.
				//
				// When ErrNotFound or ErrMethodNotAllowed are
				// returned, it's probably because there's no
				// matching route, as opposed to the handler
				// returning them. We can confirm this by looking
				// for exact-matching routes.
				var unknownRoute bool
				switch handlerErr {
				case echo.ErrNotFound:
					unknownRoute = isNotFoundHandler(c.Handler())
				case echo.ErrMethodNotAllowed:
					unknownRoute = isMethodNotAllowedHandler(c.Handler())
				}
				if unknownRoute {
					tx.Name = apmhttp.UnknownRouteRequestName(req)
				}
			}
		}
	} else if !resp.Committed {
		resp.WriteHeader(http.StatusOK)
	}
	return handlerErr
}

func setContext(ctx *apm.Context, req *http.Request, resp *echo.Response, body *apm.BodyCapturer) {
	ctx.SetFramework("echo", echo.Version)
	ctx.SetHTTPRequest(req)
	ctx.SetHTTPRequestBody(body)
	ctx.SetHTTPStatusCode(resp.Status)
	ctx.SetHTTPResponseHeaders(resp.Header())
}

type options struct {
	tracer         *apm.Tracer
	requestIgnorer apmhttp.RequestIgnorerFunc
	requestName    RequestNameFunc
}

// Option sets options for tracing.
type Option func(*options)

// WithTracer returns an Option which sets t as the tracer
// to use for tracing server requests.
func WithTracer(t *apm.Tracer) Option {
	if t == nil {
		panic("t == nil")
	}
	return func(o *options) {
		o.tracer = t
	}
}

// WithRequestName returns an Option which sets r as the function
// to use to obtain the transaction name for the given echo request.
func WithRequestName(r RequestNameFunc) Option {
	if r == nil {
		panic("r == nil")
	}

	return func(o *options) {
		o.requestName = r
	}
}

// WithRequestIgnorer returns a Option which sets r as the
// function to use to determine whether or not a request should
// be ignored. If r is nil, all requests will be reported.
func WithRequestIgnorer(r apmhttp.RequestIgnorerFunc) Option {
	if r == nil {
		r = apmhttp.IgnoreNone
	}
	return func(o *options) {
		o.requestIgnorer = r
	}
}

func isNotFoundHandler(h echo.HandlerFunc) bool {
	return isHandler(h, notFoundHandlerIdentity, &echo.NotFoundHandler)
}

func isMethodNotAllowedHandler(h echo.HandlerFunc) bool {
	return isHandler(h, methodNotAllowedHandlerIdentity, &echo.MethodNotAllowedHandler)
}

func isHandler(h echo.HandlerFunc, ident handlerFuncIdentity, handlerVar *func(echo.Context) error) bool {
	rv := reflect.ValueOf(h)
	ptr := rv.Pointer()
	if ptr == ident.rv.Pointer() {
		return true
	}
	// A sufficiently smart compiler could perform whole program optimisation
	// to determine that echo.NotFoundHandler and/or echo.MethodNotAllowedHandler
	// are only written to once to a defined function, enabling callers to inline
	// the assigned function. In this case, the function PC will not match.
	name := runtime.FuncForPC(ptr).Name()
	if name == ident.name {
		return true
	}
	// The global variables could have been reassigned since we read
	// their values during package init.
	ident = getHandlerFuncIdentity(*handlerVar)
	return ptr == ident.rv.Pointer() || name == ident.name
}

var (
	notFoundHandlerIdentity         = getHandlerFuncIdentity(echo.NotFoundHandler)
	methodNotAllowedHandlerIdentity = getHandlerFuncIdentity(echo.MethodNotAllowedHandler)
)

type handlerFuncIdentity struct {
	rv   reflect.Value
	name string
}

func getHandlerFuncIdentity(h func(echo.Context) error) handlerFuncIdentity {
	rv := reflect.ValueOf(h)
	return handlerFuncIdentity{
		rv:   rv,
		name: runtime.FuncForPC(rv.Pointer()).Name(),
	}
}
