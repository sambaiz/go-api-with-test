package conf

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Conf struct {
	Db string `envconfig:"db"` // e.g. export MBOARD_DB="root:@tcp(localhost:3306)/mboard"
}

func Parse() (*Conf, error) {
	conf := &Conf{}
	err := envconfig.Process("mboard", conf)
	if err != nil {
		panic(err)
	}
	if conf.Db == ""{
		panic(errors.New("MBOARD_DB is required."))
	}
	return conf, nil
}