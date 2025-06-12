package slogconfig

import (
	"context"
	"io"
	"log/slog"
	"testing"
)

type customHandler struct{}

func (h customHandler) Enabled(context.Context, slog.Level) bool {
	panic("not implemented")
}

func (h customHandler) Handle(context.Context, slog.Record) error {
	panic("not implemented")
}

func (h customHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("not implemented")
}

func (h customHandler) WithGroup(name string) slog.Handler {
	panic("not implemented")
}

func TestSlogConfig_NewHandler_Format(t *testing.T) {
	t.Run("text", func(t *testing.T) {
		c := SlogConfig{
			Format: "text",
		}
		h, err := c.NewHandler()
		if err != nil {
			t.Errorf("NewHandler() Errors: %v", err)
		}
		if _, ok := h.(*slog.TextHandler); !ok {
			t.Errorf("Expect handler type *slog.TextHandler. Got handler type: %T", h)
		}
	})
	t.Run("custom", func(t *testing.T) {
		RegisterLogFormat("customTest", func(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
			return customHandler{}
		})

		c := SlogConfig{
			Format: "customTest",
		}
		h, err := c.NewHandler()
		if err != nil {
			t.Errorf("NewHandler() Errors: %v", err)
		}
		if _, ok := h.(customHandler); !ok {
			t.Errorf("Expect handler type customHandler. Got handler type: %T", h)
		}
	})
	t.Run("json", func(t *testing.T) {
		c := SlogConfig{
			Format: "json",
		}
		h, err := c.NewHandler()
		if err != nil {
			t.Errorf("NewHandler() Errors: %v", err)
		}
		if _, ok := h.(*slog.JSONHandler); !ok {
			t.Errorf("Expect handler type *slog.JSONHandler. Got handler type: %T", h)
		}
	})
	t.Run("unknown", func(t *testing.T) {
		c := SlogConfig{
			Format: "114514",
		}
		_, err := c.NewHandler()
		if err == nil {
			t.Errorf("NewHandler() does not error.")
		}
		t.Logf("Got expected err: %v", err)
	})
}
