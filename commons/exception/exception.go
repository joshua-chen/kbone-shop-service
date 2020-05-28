/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 21:47:23
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:59:53
 */
package exception

import (
	"errors"
	"sync"

)

type Exception struct {
}

var (
	instance *Exception
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Exception {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Exception{}
		}
	}
	return instance
}

func (e *Exception) Fatal(err error) {
	panic(err)
}
func (e *Exception) FatalS(err string) {
	panic(errors.New(err))
}
func (e *Exception) Error(err error) {

}
func (e *Exception) ErrorS(err string) {

}
