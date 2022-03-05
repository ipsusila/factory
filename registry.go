package factory

import (
	"sort"
	"sync"
)

var (
	factoriesMu sync.RWMutex
	factories   = make(map[string]*Factory)
)

// Register makes a factory available by the provided name.
// If register is called twice with the same name or if factory is nil, it panics.
func register(name string, factory *Factory) {
	factoriesMu.Lock()
	defer factoriesMu.Unlock()
	if factory == nil {
		panic("factory: Register factory is nil")
	}
	if _, dup := factories[name]; dup {
		panic("factory: Register called twice for factory " + name)
	}
	factories[name] = factory
}

// Factories returns a sorted list of the names of the registered factories.
func Factories() []*Factory {
	factoriesMu.RLock()
	defer factoriesMu.RUnlock()
	list := make([]*Factory, 0, len(factories))
	for _, factory := range factories {
		list = append(list, factory)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].info.Name < list[j].info.Name
	})
	return list
}

// Get return factory with given name.
// If Factory does not exists, it will return nil
func Get(name string) *Factory {
	factoriesMu.RLock()
	defer factoriesMu.RUnlock()

	for key, factory := range factories {
		if key == name {
			return factory
		}
	}
	return nil
}
