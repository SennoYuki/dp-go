package dp

import "sync"

// Singleton 饿汉单例
type Singleton struct {
	val interface{}
}

func NewSingleton(f func() interface{}) *Singleton {
	return &Singleton{val: f()}
}

func (p *Singleton) Get() interface{} {
	return p.val
}

type LazySingleton struct {
	f    func() interface{}
	val  interface{}
	done bool
}

func NewLazySingleton(f func() interface{}) *LazySingleton {
	return &LazySingleton{f: f}
}

func (p *LazySingleton) Get() interface{} {
	if p.done {
		return p.val
	}
	p.val = p.f()
	p.done = true
	return p.val
}

// LazySingletonC 支持并发的懒汉单例
type LazySingletonC struct {
	once sync.Once
	f    func() interface{}
	val  interface{}
}

func NewLazySingletonC(f func() interface{}) *LazySingletonC {
	return &LazySingletonC{f: f}
}

func (p *LazySingletonC) Get() interface{} {
	p.once.Do(func() {
		p.val = p.f
	})
	return p.val
}
