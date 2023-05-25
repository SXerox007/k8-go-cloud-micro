package logger

import (
	"context"
	"fmt"
	"testing"

	meta "github.com/grpc-ecosystem/go-grpc-middleware/v2/metadata"
	"github.com/sirupsen/logrus/hooks/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestFromContext(t *testing.T) {
	md := metadata.Pairs(correlationIDKey, "1111-2222-3333")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	logger := FromContext(ctx)

	cid := logger.Data[correlationIDKey].(string)
	if cid != "1111-2222-3333" {
		t.Errorf("correlation id in logger from context is not expected '1111-2222-3333': %s", cid)
	}
}

func TestWithContext(t *testing.T) {
	md := metadata.Pairs(correlationIDKey, "1111-2222-3333")
	ctx := metadata.NewIncomingContext(context.Background(), md)
	logger := WithContext(ctx, NewLogger())
	cid := logger.Data[correlationIDKey].(string)
	if cid != "1111-2222-3333" {
		t.Errorf("correlation id in logger from context is not expected '1111-2222-3333': %s", cid)
	}
}

func TestCorrelationIDFromContext(t *testing.T) {
	t.Run("WithCorrelationIDInContextMetadata", func(t *testing.T) {
		md := metadata.Pairs(correlationIDKey, "1111-2222-3333")
		ctx := metadata.NewIncomingContext(context.Background(), md)
		cid := CorrelationIDFromContext(ctx)
		if cid != "1111-2222-3333" {
			t.Errorf("correlation id in context is not expected '1111-2222-3333': %s", cid)
		}
	})

	t.Run("WithoutCorrelationIDInContextMetadata", func(t *testing.T) {
		md := metadata.Pairs("not-correlation-key", "1234")
		ctx := metadata.NewIncomingContext(context.Background(), md)

		cid := CorrelationIDFromContext(ctx)
		if cid != "" {
			t.Errorf("correlation id in context is not empty: %s", cid)
		}
	})

	t.Run("WithoutMetadataInContext", func(t *testing.T) {
		ctx := ContextWithCorrelationID(context.Background())
		cid := CorrelationIDFromContext(ctx)
		if cid == "" {
			t.Errorf("correlation id in context is empty")
		}
	})
}

func TestContextWithCorrelationID(t *testing.T) {
	t.Run("WithCorrelationIDInContextMetadata", func(t *testing.T) {
		md := metadata.Pairs(correlationIDKey, "1111-2222-3333")
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx = ContextWithCorrelationID(ctx)
		cid := CorrelationIDFromContext(ctx)
		if cid != "1111-2222-3333" {
			t.Errorf("correlation id in context is not expected '1111-2222-3333': %s", cid)
		}
	})

	t.Run("WithoutCorrelationIDInContextMetadata", func(t *testing.T) {
		md := metadata.Pairs("not-correlation-key", "1234")
		ctx := metadata.NewIncomingContext(context.Background(), md)
		ctx = ContextWithCorrelationID(ctx)
		cid := CorrelationIDFromContext(ctx)
		if cid == "" {
			t.Errorf("correlation id in context empty")
		}
	})

	t.Run("WithoutMetadataInContext", func(t *testing.T) {
		ctx := ContextWithCorrelationID(context.Background())
		cid := CorrelationIDFromContext(ctx)
		if cid == "" {
			t.Errorf("correlation id in context empty")
		}
	})
}

func TestUnaryClientInterceptor(t *testing.T) {
	inv := UnaryClientInterceptor()
	logr, hook := test.NewNullLogger()
	ctx := WithLogger(context.Background(), logr)

	t.Run("ValidMetadataLogForward", func(t *testing.T) {
		k := "not-correlation-key"
		lCtx := WithLogger(ctx, logr.WithField(k, "1234"))
		v := func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, opts ...grpc.CallOption,
		) error {
			md := meta.ExtractOutgoing(ctx)
			if md.Get(logLead+k) != "1234" {
				return fmt.Errorf("not forwarding the key %s, expecting forward it", k)
			}
			return nil
		}

		if err := inv(lCtx, "", nil, nil, nil, v); err != nil {
			t.Error(err)
		}

		if t.Failed() || testing.Verbose() {
			for _, e := range hook.Entries {
				s, _ := e.String()
				t.Log(s)
			}
		}
	})

	t.Run("InvalidMetadataLogForward", func(t *testing.T) {
		k := "not-correlation<key"
		lCtx := WithLogger(ctx, logr.WithField(k, "1234"))
		v := func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, opts ...grpc.CallOption,
		) error {
			md := meta.ExtractOutgoing(ctx)
			if md.Get(logLead+k) != "" {
				return fmt.Errorf("forwarding the key %s-%s, expecting not", k, md.Get(k))
			}
			return nil
		}

		if err := inv(lCtx, "", nil, nil, nil, v); err != nil {
			t.Error(err)
		}

		if t.Failed() || testing.Verbose() {
			for _, e := range hook.Entries {
				s, _ := e.String()
				t.Log(s)
			}
		}
	})

	t.Run("SpaceMetadataLogForward", func(t *testing.T) {
		k := "not-correlation-key"
		ks := "not correlation key"
		lCtx := WithLogger(ctx, logr.WithField(ks, "1234"))
		v := func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, opts ...grpc.CallOption,
		) error {
			md := meta.ExtractOutgoing(ctx)
			if md.Get(logLead+k) == "" {
				return fmt.Errorf("forwarding the key %s-%+v, expecting not", k, md.Get(k))
			}
			return nil
		}

		if err := inv(lCtx, "", nil, nil, nil, v); err != nil {
			t.Error(err)
		}

		if t.Failed() || testing.Verbose() {
			for _, e := range hook.Entries {
				s, _ := e.String()
				t.Log(s)
			}
		}
	})

	t.Run("BinaryMetadataLogForward", func(t *testing.T) {
		k := "not-correlation-key"
		lCtx := WithLogger(ctx, logr.WithField(k, "\n"))
		v := func(ctx context.Context, method string, req, reply interface{},
			cc *grpc.ClientConn, opts ...grpc.CallOption,
		) error {
			md := meta.ExtractOutgoing(ctx)
			if md.Get(logLead+k+"-bin") == "" {
				return fmt.Errorf("not forwarding the key %s-%+v, expecting forward it", k, md.Get(k))
			}

			if md.Get(logLead+k) != "" {
				return fmt.Errorf("forwarding the key %s-%+v, expecting not", k, md.Get(k))
			}
			return nil
		}

		if err := inv(lCtx, "", nil, nil, nil, v); err != nil {
			t.Error(err)
		}

		if t.Failed() || testing.Verbose() {
			for _, e := range hook.Entries {
				s, _ := e.String()
				t.Log(s)
			}
		}
	})
}
