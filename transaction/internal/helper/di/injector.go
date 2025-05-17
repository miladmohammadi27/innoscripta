package di

import (
	"log"
	"os"

	"github.com/samber/do"
)

// isDev checks whether application runs in development mode
// before actual config reader was injected
func isDev() bool {
	env := os.Getenv("TRANSACTION_LOGGER_ENV")
	if env == "development" || env == "" {
		return true
	}
	return false
}

// NewInjector create and configure a new samber do injector
func NewInjector() *do.Injector {
	return do.NewWithOpts(&do.InjectorOpts{
		Logf: func(format string, args ...any) {
			// use default logger because we don't have our own logger injected yet
			if isDev() {
				log.Printf(format, args...)
			}
		},
		HookAfterRegistration: func(_ *do.Injector, serviceName string) {
			if isDev() {
				log.Printf("service %s is registered", serviceName)
			}
		},
		HookAfterShutdown: func(_ *do.Injector, serviceName string) {
			if isDev() {
				log.Printf("service %s is unregistered", serviceName)
			}
		},
	})
}
