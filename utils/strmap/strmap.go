package strmap

import (
	"log"

	"github.com/spf13/cast"
)

type StrMap map[string]interface{}

func (s StrMap) Get(key string) StrMap {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToStringMapE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return StrMap(value)
}

func (s StrMap) GetString(key string) string {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToStringE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}

func (s StrMap) GetStringSlice(key string) []string {
	if s == nil || s[key] == nil {
		log.Panicf("no key: %s", key)
	}
	value, err := cast.ToStringSliceE(s[key])
	if err != nil {
		log.Panic(err)
	}
	return value
}
