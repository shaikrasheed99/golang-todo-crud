package helper

import "github.com/rs/zerolog/log"

func LogAndPanicError(err error) {
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
}
