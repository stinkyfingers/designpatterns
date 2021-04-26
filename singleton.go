package designpatterns

import "sync"

/*
 The Singleton pattern assures that only a single instance of an object is created.
*/

type SingletonObject struct {
}

var instance *SingletonObject

var once sync.Once

var mutex sync.Mutex

// NotThreadSafeSingleton demonstrates a bad example. A race could occur
// causing instance to be created more than once by separate callers.
func NotThreadSafeSingleton() *SingletonObject {
	if instance == nil {
		instance = &SingletonObject{}
	}
	return instance
}

// ThreadSafeSingleton demonstrates a threadsafe instantiation using the sync pkg
func ThreadSafeSingleton() *SingletonObject {
	once.Do(func() {
		instance = &SingletonObject{}
	})
	return instance
}

// ThreadSafeMutexSingleton demonstrates a threadsafe instantiation using a mutex
func ThreadSafeMutexSingleton() *SingletonObject {
	mutex.Lock()
	defer mutex.Unlock()
	if instance == nil {
		instance = &SingletonObject{}
	}
	return instance
}
