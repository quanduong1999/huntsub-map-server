package ref

import (
	"reflect"
)

type Type struct {
	t        reflect.Type
	slice    reflect.Type
	ptr      reflect.Type
	slicePtr reflect.Type
}

func NewType(v interface{}) Type {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	slice := reflect.SliceOf(t)
	ptr := reflect.PtrTo(t)
	slicePtr := reflect.SliceOf(ptr)
	return Type{
		t:        t,
		slice:    slice,
		ptr:      ptr,
		slicePtr: slicePtr,
	}
}

func (t *Type) NewPtr() reflect.Value {
	return reflect.New(t.t)
}

func (t *Type) NewSliceOfPtr() reflect.Value {
	arr := reflect.MakeSlice(t.slicePtr, 0, 0)
	v := reflect.New(arr.Type())
	v.Elem().Set(arr)
	return v
}
