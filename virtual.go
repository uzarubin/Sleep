package Sleep

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Virtual struct {
	bools   map[string]bool
	ints    map[string]int
	strings map[string]string
	allElse map[string]interface{}
	ids     map[string]bson.ObjectId
	times   map[string]time.Time
}

func newVirtual() *Virtual {
	v := &Virtual{
		bools:   make(map[string]bool),
		ints:    make(map[string]int),
		strings: make(map[string]string),
		allElse: make(map[string]interface{}),
		ids:     make(map[string]bson.ObjectId),
		times:   make(map[string]time.Time)}

	return v
}

func (v *Virtual) Get(name string) (interface{}, bool) {
	val, ok := v.allElse[name]
	return val, ok
}

func (v *Virtual) Set(name string, val interface{}) {
	v.allElse[name] = val
}

func (v *Virtual) GetBool(name string) (bool, bool) {
	val, ok := v.bools[name]
	return val, ok
}

func (v *Virtual) SetBool(name string, val bool) {
	v.bools[name] = val
}

func (v *Virtual) GetInt(name string) (int, bool) {
	val, ok := v.ints[name]
	return val, ok
}

func (v *Virtual) SetInt(name string, val int) {
	v.ints[name] = val
}

func (v *Virtual) GetString(name string) (string, bool) {
	val, ok := v.strings[name]
	return val, ok
}

func (v *Virtual) SetString(name string, val string) {
	v.strings[name] = val
}

func (v *Virtual) GetObjectId(name string) (bson.ObjectId, bool) {
	val, ok := v.ids[name]
	return val, ok
}

func (v *Virtual) SetObjectId(name string, val bson.ObjectId) {
	v.ids[name] = val
}

func (v *Virtual) GetTime(name string) (time.Time, bool) {
	val, ok := v.times[name]
	return val, ok
}

func (v *Virtual) SetTime(name string, val time.Time) {
	v.times[name] = val
}