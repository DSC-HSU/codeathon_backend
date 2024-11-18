// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package profile

import (
	"codeathon.runwayclub.dev/domain"
	"context"
	"errors"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "codeathon.runwayclub.dev/internal/profile/ProfileService",
		Iface: reflect.TypeOf((*ProfileService)(nil)).Elem(),
		Impl:  reflect.TypeOf(profileService{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return profileService_local_stub{impl: impl.(ProfileService), tracer: tracer, deleteMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "Delete", Remote: false, Generated: true}), getByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "GetById", Remote: false, Generated: true}), listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "List", Remote: false, Generated: true}), updateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "Update", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return profileService_client_stub{stub: stub, deleteMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "Delete", Remote: true, Generated: true}), getByIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "GetById", Remote: true, Generated: true}), listMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "List", Remote: true, Generated: true}), updateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/profile/ProfileService", Method: "Update", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return profileService_server_stub{impl: impl.(ProfileService), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return profileService_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[ProfileService] = (*profileService)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*profileService)(nil)

// Local stub implementations.

type profileService_local_stub struct {
	impl           ProfileService
	tracer         trace.Tracer
	deleteMetrics  *codegen.MethodMetrics
	getByIdMetrics *codegen.MethodMetrics
	listMetrics    *codegen.MethodMetrics
	updateMetrics  *codegen.MethodMetrics
}

// Check that profileService_local_stub implements the ProfileService interface.
var _ ProfileService = (*profileService_local_stub)(nil)

func (s profileService_local_stub) Delete(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	begin := s.deleteMetrics.Begin()
	defer func() { s.deleteMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "profile.ProfileService.Delete", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Delete(ctx, a0)
}

func (s profileService_local_stub) GetById(ctx context.Context, a0 string) (r0 *domain.Profile, err error) {
	// Update metrics.
	begin := s.getByIdMetrics.Begin()
	defer func() { s.getByIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "profile.ProfileService.GetById", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetById(ctx, a0)
}

func (s profileService_local_stub) List(ctx context.Context, a0 int, a1 *domain.ListOpts) (r0 *domain.ListResult[*domain.Profile], err error) {
	// Update metrics.
	begin := s.listMetrics.Begin()
	defer func() { s.listMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "profile.ProfileService.List", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.List(ctx, a0, a1)
}

func (s profileService_local_stub) Update(ctx context.Context, a0 *domain.Profile) (err error) {
	// Update metrics.
	begin := s.updateMetrics.Begin()
	defer func() { s.updateMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "profile.ProfileService.Update", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Update(ctx, a0)
}

// Client stub implementations.

type profileService_client_stub struct {
	stub           codegen.Stub
	deleteMetrics  *codegen.MethodMetrics
	getByIdMetrics *codegen.MethodMetrics
	listMetrics    *codegen.MethodMetrics
	updateMetrics  *codegen.MethodMetrics
}

// Check that profileService_client_stub implements the ProfileService interface.
var _ ProfileService = (*profileService_client_stub)(nil)

func (s profileService_client_stub) Delete(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.deleteMetrics.Begin()
	defer func() { s.deleteMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "profile.ProfileService.Delete", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

func (s profileService_client_stub) GetById(ctx context.Context, a0 string) (r0 *domain.Profile, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getByIdMetrics.Begin()
	defer func() { s.getByIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "profile.ProfileService.GetById", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += (4 + len(a0))
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.String(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 1, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_Profile_260fe93b(dec)
	err = dec.Error()
	return
}

func (s profileService_client_stub) List(ctx context.Context, a0 int, a1 *domain.ListOpts) (r0 *domain.ListResult[*domain.Profile], err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.listMetrics.Begin()
	defer func() { s.listMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "profile.ProfileService.List", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	enc.Int(a0)
	serviceweaver_enc_ptr_ListOpts_73a4ea72(enc, a1)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 2, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_ptr_ListResult_ptr_Profile_95df334e(dec)
	err = dec.Error()
	return
}

func (s profileService_client_stub) Update(ctx context.Context, a0 *domain.Profile) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.updateMetrics.Begin()
	defer func() { s.updateMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "profile.ProfileService.Update", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Encode arguments.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Profile_260fe93b(enc, a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 3, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.5 (codegen
version v0.24.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

// Server stub implementations.

type profileService_server_stub struct {
	impl    ProfileService
	addLoad func(key uint64, load float64)
}

// Check that profileService_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*profileService_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s profileService_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Delete":
		return s.delete
	case "GetById":
		return s.getById
	case "List":
		return s.list
	case "Update":
		return s.update
	default:
		return nil
	}
}

func (s profileService_server_stub) delete(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Delete(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s profileService_server_stub) getById(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 string
	a0 = dec.String()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetById(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Profile_260fe93b(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s profileService_server_stub) list(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()
	var a1 *domain.ListOpts
	a1 = serviceweaver_dec_ptr_ListOpts_73a4ea72(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.List(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_ListResult_ptr_Profile_95df334e(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s profileService_server_stub) update(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 *domain.Profile
	a0 = serviceweaver_dec_ptr_Profile_260fe93b(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Update(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type profileService_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that profileService_reflect_stub implements the ProfileService interface.
var _ ProfileService = (*profileService_reflect_stub)(nil)

func (s profileService_reflect_stub) Delete(ctx context.Context, a0 string) (err error) {
	err = s.caller("Delete", ctx, []any{a0}, []any{})
	return
}

func (s profileService_reflect_stub) GetById(ctx context.Context, a0 string) (r0 *domain.Profile, err error) {
	err = s.caller("GetById", ctx, []any{a0}, []any{&r0})
	return
}

func (s profileService_reflect_stub) List(ctx context.Context, a0 int, a1 *domain.ListOpts) (r0 *domain.ListResult[*domain.Profile], err error) {
	err = s.caller("List", ctx, []any{a0, a1}, []any{&r0})
	return
}

func (s profileService_reflect_stub) Update(ctx context.Context, a0 *domain.Profile) (err error) {
	err = s.caller("Update", ctx, []any{a0}, []any{})
	return
}

// Encoding/decoding implementations.

func serviceweaver_enc_ptr_Profile_260fe93b(enc *codegen.Encoder, arg *domain.Profile) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Profile_260fe93b(dec *codegen.Decoder) *domain.Profile {
	if !dec.Bool() {
		return nil
	}
	var res domain.Profile
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_ListOpts_73a4ea72(enc *codegen.Encoder, arg *domain.ListOpts) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_ListOpts_73a4ea72(dec *codegen.Decoder) *domain.ListOpts {
	if !dec.Bool() {
		return nil
	}
	var res domain.ListOpts
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_ptr_ListResult_ptr_Profile_95df334e(enc *codegen.Encoder, arg *domain.ListResult[*domain.Profile]) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.EncodeBinaryMarshaler(arg)
	}
}

func serviceweaver_dec_ptr_ListResult_ptr_Profile_95df334e(dec *codegen.Decoder) *domain.ListResult[*domain.Profile] {
	if !dec.Bool() {
		return nil
	}
	var res domain.ListResult[*domain.Profile]
	dec.DecodeBinaryUnmarshaler(&res)
	return &res
}
