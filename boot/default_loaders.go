package boot

import (
	"fmt"

	"github.com/gookit/config/v2/dotnev"
	"github.com/gookit/event"
	"github.com/gookit/lako"
)

// EnvBootLoader for load .env file
func EnvBootLoader(dir string, envFiles ...string) lako.BootLoader {
	return lako.BootFunc(func(app *lako.Application) error {
		return dotnev.LoadExists(dir, envFiles...)
	})
}

// EnvBootLoader for load config files
func ConfigBootLoader(confFiles ...string) lako.BootLoader {
	return lako.BootFunc(func(app *lako.Application) error {
		app.MustFire(lako.OnBeforeConfig, event.M{
			"files":  confFiles,
			"config": app.Config,
		})

		fmt.Println("load config files:", confFiles)

		// load from files
		err := app.Config.LoadExists(confFiles...)

		// load from flags
		if err == nil {
			err =app.Config.LoadFlags([]string{"debug"})
		}

		app.MustFire(lako.OnAfterConfig, event.M{"config": app.Config})

		return err
	})
}
