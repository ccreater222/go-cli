package common

import (
	"reflect"
)
type Command interface {
	Execute()
}


type cmd struct {
	name string
	t interface{}
	tag string
}

var cmds []cmd

func init(){
	cmds = make([]cmd,0)
}

func RegisterCmd(name string,t interface{},tag string){
	cmds = append(cmds,cmd{
		name: name,
		t:t,
		tag:tag,
	})
}

func GetStruct() interface{}{
	var sfs []reflect.StructField
	for _,v := range cmds {
		sf := reflect.StructField{
			Name: v.name,
			Type: reflect.PtrTo(reflect.TypeOf(v.t)),
			Tag: reflect.StructTag(v.tag),
		}
		sfs = append(sfs, sf)
	}
	st := reflect.StructOf(sfs)
	so := reflect.New(st)

	return so.Interface()
}
