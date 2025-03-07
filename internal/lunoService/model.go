package lunoService

import (
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"os"
	"os/user"
)

type LunoKeys = struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func ProvideLunoAPIKeyID() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: "LunoAPIKeyID",
			Target: func(data *LunoKeys) string {

				return data.Key
			},
		})
}

func ProvideLunoAPIKeySecret() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: "LunoAPIKeySecret",
			Target: func(data *LunoKeys) string {
				return data.Secret
			},
		})
}

func ProvideLunoKeys(fromFile bool, useThis *LunoKeys) fx.Option {
	if fromFile {
		return fx.Provide(
			fx.Annotated{
				Target: func() (*LunoKeys, error) {
					data := &LunoKeys{}
					current, err := user.Current()
					if err != nil {
						return nil, err
					}
					f, err := os.Open(fmt.Sprintf("%v/.luno/keys.json", current.HomeDir))
					if err != nil {
						return nil, err
					}

					defer func() {
						_ = f.Close()
					}()
					decoder := json.NewDecoder(f)
					err = decoder.Decode(data)
					if err != nil {
						return nil, err
					}
					return data, nil
				},
			},
		)
	}
	if useThis == nil {
		return fx.Error(fmt.Errorf("provide LunoKeys == nil"))
	}
	return fx.Provide(
		fx.Annotated{
			Target: func() *LunoKeys {
				return useThis
			},
		},
	)
}
