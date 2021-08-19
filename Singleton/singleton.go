package Singleton

import "sync"

type SingleObject struct {
	name string
}

var (
	singleOnce   sync.Once
	singleObject *SingleObject
)

func initSignleInstance()  {
	if singleObject == nil {
		singleObject = &SingleObject{"singleinstance"}
	}
}

func GetInstance() *SingleObject {
	singleOnce.Do(initSignleInstance)
	return singleObject
}


func (s SingleObject) Show() string {
	return s.name
}