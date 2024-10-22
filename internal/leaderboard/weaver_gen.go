// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package leaderboard

import (
	"codeathon.runwayclub.dev/domain"
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "codeathon.runwayclub.dev/internal/leaderboard/LeaderboardService",
		Iface: reflect.TypeOf((*LeaderboardService)(nil)).Elem(),
		Impl:  reflect.TypeOf(leaderboardService{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return leaderboardService_local_stub{impl: impl.(LeaderboardService), tracer: tracer, getByCIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/leaderboard/LeaderboardService", Method: "GetByCId", Remote: false, Generated: true}), recalculateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/leaderboard/LeaderboardService", Method: "Recalculate", Remote: false, Generated: true})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return leaderboardService_client_stub{stub: stub, getByCIdMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/leaderboard/LeaderboardService", Method: "GetByCId", Remote: true, Generated: true}), recalculateMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "codeathon.runwayclub.dev/internal/leaderboard/LeaderboardService", Method: "Recalculate", Remote: true, Generated: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return leaderboardService_server_stub{impl: impl.(LeaderboardService), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(string, context.Context, []any, []any) error) any {
			return leaderboardService_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[LeaderboardService] = (*leaderboardService)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*leaderboardService)(nil)

// Local stub implementations.

type leaderboardService_local_stub struct {
	impl               LeaderboardService
	tracer             trace.Tracer
	getByCIdMetrics    *codegen.MethodMetrics
	recalculateMetrics *codegen.MethodMetrics
}

// Check that leaderboardService_local_stub implements the LeaderboardService interface.
var _ LeaderboardService = (*leaderboardService_local_stub)(nil)

func (s leaderboardService_local_stub) GetByCId(ctx context.Context, a0 string, a1 *domain.ListOpts) (r0 *Leaderboard, err error) {
	// Update metrics.
	begin := s.getByCIdMetrics.Begin()
	defer func() { s.getByCIdMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "leaderboard.LeaderboardService.GetByCId", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetByCId(ctx, a0, a1)
}

func (s leaderboardService_local_stub) Recalculate(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	begin := s.recalculateMetrics.Begin()
	defer func() { s.recalculateMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "leaderboard.LeaderboardService.Recalculate", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Recalculate(ctx, a0)
}

// Client stub implementations.

type leaderboardService_client_stub struct {
	stub               codegen.Stub
	getByCIdMetrics    *codegen.MethodMetrics
	recalculateMetrics *codegen.MethodMetrics
}

// Check that leaderboardService_client_stub implements the LeaderboardService interface.
var _ LeaderboardService = (*leaderboardService_client_stub)(nil)

func (s leaderboardService_client_stub) GetByCId(ctx context.Context, a0 string, a1 *domain.ListOpts) (r0 *Leaderboard, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getByCIdMetrics.Begin()
	defer func() { s.getByCIdMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "leaderboard.LeaderboardService.GetByCId", trace.WithSpanKind(trace.SpanKindClient))
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
	enc.String(a0)
	serviceweaver_enc_ptr_ListOpts_73a4ea72(enc, a1)
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
	r0 = serviceweaver_dec_ptr_Leaderboard_de222c7e(dec)
	err = dec.Error()
	return
}

func (s leaderboardService_client_stub) Recalculate(ctx context.Context, a0 string) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.recalculateMetrics.Begin()
	defer func() { s.recalculateMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "leaderboard.LeaderboardService.Recalculate", trace.WithSpanKind(trace.SpanKindClient))
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
	err = dec.Error()
	return
}

// Note that "weaver generate" will always generate the error message below.
// Everything is okay. The error message is only relevant if you see it when
// you run "go build" or "go run".
var _ codegen.LatestVersion = codegen.Version[[0][24]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.24.6 (codegen
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

type leaderboardService_server_stub struct {
	impl    LeaderboardService
	addLoad func(key uint64, load float64)
}

// Check that leaderboardService_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*leaderboardService_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s leaderboardService_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetByCId":
		return s.getByCId
	case "Recalculate":
		return s.recalculate
	default:
		return nil
	}
}

func (s leaderboardService_server_stub) getByCId(ctx context.Context, args []byte) (res []byte, err error) {
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
	var a1 *domain.ListOpts
	a1 = serviceweaver_dec_ptr_ListOpts_73a4ea72(dec)

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetByCId(ctx, a0, a1)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_ptr_Leaderboard_de222c7e(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

func (s leaderboardService_server_stub) recalculate(ctx context.Context, args []byte) (res []byte, err error) {
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
	appErr := s.impl.Recalculate(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type leaderboardService_reflect_stub struct {
	caller func(string, context.Context, []any, []any) error
}

// Check that leaderboardService_reflect_stub implements the LeaderboardService interface.
var _ LeaderboardService = (*leaderboardService_reflect_stub)(nil)

func (s leaderboardService_reflect_stub) GetByCId(ctx context.Context, a0 string, a1 *domain.ListOpts) (r0 *Leaderboard, err error) {
	err = s.caller("GetByCId", ctx, []any{a0, a1}, []any{&r0})
	return
}

func (s leaderboardService_reflect_stub) Recalculate(ctx context.Context, a0 string) (err error) {
	err = s.caller("Recalculate", ctx, []any{a0}, []any{})
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Leaderboard)(nil)

type __is_Leaderboard[T ~struct {
	weaver.AutoMarshal
	TotalPage int                  "json:\"total_page\""
	Data      []*domain.Submission "json:\"data\""
}] struct{}

var _ __is_Leaderboard[Leaderboard]

func (x *Leaderboard) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Leaderboard.WeaverMarshal: nil receiver"))
	}
	enc.Int(x.TotalPage)
	serviceweaver_enc_slice_ptr_Submission_6530ab64(enc, x.Data)
}

func (x *Leaderboard) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Leaderboard.WeaverUnmarshal: nil receiver"))
	}
	x.TotalPage = dec.Int()
	x.Data = serviceweaver_dec_slice_ptr_Submission_6530ab64(dec)
}

func serviceweaver_enc_ptr_Submission_54a2faef(enc *codegen.Encoder, arg *domain.Submission) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Submission_54a2faef(dec *codegen.Decoder) *domain.Submission {
	if !dec.Bool() {
		return nil
	}
	var res domain.Submission
	(&res).WeaverUnmarshal(dec)
	return &res
}

func serviceweaver_enc_slice_ptr_Submission_6530ab64(enc *codegen.Encoder, arg []*domain.Submission) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		serviceweaver_enc_ptr_Submission_54a2faef(enc, arg[i])
	}
}

func serviceweaver_dec_slice_ptr_Submission_6530ab64(dec *codegen.Decoder) []*domain.Submission {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]*domain.Submission, n)
	for i := 0; i < n; i++ {
		res[i] = serviceweaver_dec_ptr_Submission_54a2faef(dec)
	}
	return res
}

// Encoding/decoding implementations.

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

func serviceweaver_enc_ptr_Leaderboard_de222c7e(enc *codegen.Encoder, arg *Leaderboard) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		(*arg).WeaverMarshal(enc)
	}
}

func serviceweaver_dec_ptr_Leaderboard_de222c7e(dec *codegen.Decoder) *Leaderboard {
	if !dec.Bool() {
		return nil
	}
	var res Leaderboard
	(&res).WeaverUnmarshal(dec)
	return &res
}