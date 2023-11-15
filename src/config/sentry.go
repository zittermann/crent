package config

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
)

func SentryInit() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DNS"),
		EnableTracing: true,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}
}
