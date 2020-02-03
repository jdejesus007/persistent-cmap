package pcmap

import "sync"

var cacheLock sync.RWMutex

// PersistentSyncMap fixed map which is thread safe and persistent
type PersistentSyncMap struct {
	storage map[interface{}]interface{}
}

// New creates a new persistent concurrent map
func (p *PersistentSyncMap) New() {
	p.storage = make(map[interface{}]interface{}, 100)
}

// Set sets the key/value pair
func (p *PersistentSyncMap) Set(key interface{}, value interface{}) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	p.storage[key] = value
}

// Get gets the value with provided key
func (p *PersistentSyncMap) Get(key interface{}) interface{} {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	return p.storage[key]
}

// Delete deletes entry from map via key
func (p *PersistentSyncMap) Delete(key interface{}) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	delete(p.storage, key)
	return
}
