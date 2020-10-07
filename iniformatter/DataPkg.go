package iniformatter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type DataPkg struct {
	pkgName string
	attr map[string]string
}

func NewDataPkg(name string)*DataPkg {
	m := make(map[string]string)
	return &DataPkg{pkgName:name, attr:m}
}

func(dp *DataPkg)SetAttr(name string, val interface{}) {
	valStr := fmt.Sprintf("%v", val)
	dp.attr[name] = valStr
}

func(dp *DataPkg)GetAttrStr(attrName string)(string, error) {
	v, ok := dp.attr[attrName]
	if ok {
		return v, nil
	} else {
		return "", errors.New("this "+ attrName +" is invalid")
	}
}

func(dp *DataPkg)GetAttrVal(typ string ,attrName string)(v reflect.Value, err error) {
	switch typ {
	case "int":
		var i int
		i, err = strconv.Atoi(dp.attr[attrName])
		if err != nil {
			break
		}
		v = newValue(i)
		v.SetInt(int64(i))
	case "string":
		var s string = dp.attr[attrName]
		v = newValue(s)
		v.SetString(s)
	default:
		v = reflect.ValueOf(nil)
		err = errors.New("GetAttrVal:don't used this type as the attr type")
	}

	return
}

func newValue(data interface{}) reflect.Value {
	t := reflect.TypeOf(data)
	vp := reflect.New(t)
	return vp.Elem()
}

// sample 只是一个结构体类型样例，并不包括具体的值，用于构建结构体映射Value类型
func (dp *DataPkg) GetAttrValByStruct(sample interface{})(v reflect.Value, err error) {
	err = nil
	typ := reflect.TypeOf(sample)
	v = reflect.ValueOf(nil)
	name := typ.Name()
	if typ.Kind() != reflect.Struct {
		err = errors.New("GetAttrValByStruct:type of sample mush be struct")
		return
	}

	matchErr := errors.New("GetAttrValByStruct:this type struct '" + name + "' didn't match the struct '" + dp.pkgName + "'")
	val := reflect.New(typ).Elem()
	length := val.NumField()
	if length != len(dp.attr) {
		err = matchErr
		return
	}

	for i := 0;i < length;i++ {
		fieldInfo := typ.Field(i)
		key := fieldInfo.Tag.Get("ini")
		_, ok := dp.attr[key]
		if ok {
			var attrVal reflect.Value
			attrVal, err = dp.GetAttrVal(fieldInfo.Type.Name(), key)
			if err != nil {
				return
			}
			val.Field(i).Set(attrVal)
		} else {
			err = matchErr
			return
		}
	}

	v = val
	return
}