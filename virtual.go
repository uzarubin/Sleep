package Sleep

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Virtual holds temporary/computed values related to the document.
//As of right now Virtual implements getters and setters
//for types most commonly used in web developement. It also implements a generic getter and setter
//for storing and retrieving any type as type interface{} and must be asserted to its proper type upon retrieval.
//
//These fields that are kept for the lifetime of the document in memory and are NOT persisted to the database.
type Virtual struct {
	bools   map[string]bool
	ints    map[string]int
	floats  map[string]float64
	strings map[string]string
	allElse map[string]interface{}
	ids     map[string]bson.ObjectId
	times   map[string]time.Time
}

func newVirtual() *Virtual {
	v := &Virtual{
		bools:   make(map[string]bool),
		ints:    make(map[string]int),
		floats:  make(map[string]float64),
		strings: make(map[string]string),
		allElse: make(map[string]interface{}),
		ids:     make(map[string]bson.ObjectId),
		times:   make(map[string]time.Time)}

	return v
}

// Get returns the stored value with the given name as type interface{}.
// It also returns a boolean value indicating whether a value was found.
//
// Get is a generic getter for any arbitrary type
func (v *Virtual) Get(name string) (interface{}, bool) {
	val, ok := v.allElse[name]
	return val, ok
}

// Set stores the value with the given name as type interface{}.
//
// Set is a generic setter for any arbitrary type
func (v *Virtual) Set(name string, val interface{}) {
	v.allElse[name] = val
}

// Get returns the stored boolean value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetBool(name string) (bool, bool) {
	val, ok := v.bools[name]
	return val, ok
}

// SetBool stores the boolean value with the given name.
func (v *Virtual) SetBool(name string, val bool) {
	v.bools[name] = val
}

// GetInt returns the stored int value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetInt(name string) (int, bool) {
	val, ok := v.ints[name]
	return val, ok
}

// SetInt stores the int value with the given name.
func (v *Virtual) SetInt(name string, val int) {
	v.ints[name] = val
}

// GetFloat returns the stored float64 value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetFloat(name string) (float64, bool) {
	val, ok := v.floats[name]
	return val, ok
}

// SetFloat stores the float64 value with the given name.
func (v *Virtual) SetFloat(name string, val float64) {
	v.floats[name] = val
}

// GetString returns the stored string value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetString(name string) (string, bool) {
	val, ok := v.strings[name]
	return val, ok
}

// SetString stores the string value with the given name.
func (v *Virtual) SetString(name string, val string) {
	v.strings[name] = val
}

// GetObjectId returns the stored bson.ObjectId value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetObjectId(name string) (bson.ObjectId, bool) {
	val, ok := v.ids[name]
	return val, ok
}

// SetObjectId stores the bson.ObjectId value with the given name.
func (v *Virtual) SetObjectId(name string, val bson.ObjectId) {
	v.ids[name] = val
}

// GetTime returns the stored time.Time value with the given name.
// It also returns a boolean value indicating whether a value was found.
func (v *Virtual) GetTime(name string) (time.Time, bool) {
	val, ok := v.times[name]
	return val, ok
}

// SetTime stores the time.Time value with the given name.
func (v *Virtual) SetTime(name string, val time.Time) {
	v.times[name] = val
}
