package di

import (
	"errors"

	"transaction/internal/helper/envy"

	"github.com/samber/do"
)

var ErrGetCfg = errors.New("failed to get config from DI")

func GetConfigFromDI[T any](i *do.Injector, cfg *T) error {
	configurator := do.MustInvoke[envy.Envy](i)
	return configurator.Parse(cfg)
}
