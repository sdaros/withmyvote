package db

import "sync"

type DB interface {
	//Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	//Range(f func(key, value interface{}) bool)
	//Store(key, value interface{})
}

type db struct {
	state *sync.Map
}

func Open() DB {
	return &db{state: &sync.Map{}}
}
func (d db) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return d.state.LoadOrStore(key, value)
}

func (d db) Load(key interface{}) (value interface{}, ok bool) {
	return d.state.Load(key)
}
