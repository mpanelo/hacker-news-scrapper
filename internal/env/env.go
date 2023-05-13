package env

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

func ReadString(key string) string {
	return os.Getenv(key)
}

func ReadInt(key string) int {
	i, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatal().Err(err).Msgf("environment variable %s is not an integer", key)
	}
	return i
}
