package config

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewRuntimeContext returns context & cancel func listening to :
// - os.Interrupt
// - syscall.SIGTERM
// - syscall.SIGINT.
func NewRuntimeContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	return
}
