package col

import "reflect"

// FindOne searches inside a slice (of structs) and returns the first match which
// satisifes the given value for the given field. If nothing is found then
// it send nil
func FindOne(sliceOfStruct interface{}, fldName string, fldValue interface{}) interface{} {
	switch reflect.TypeOf(sliceOfStruct).Kind() {
	case reflect.Slice:
		v := reflect.ValueOf(sliceOfStruct)
		for i := 0; i < v.Len(); i++ {
			fld := v.Index(i).FieldByName(fldName)
			if fld.IsValid() && fld.Interface() == fldValue {
				return v.Index(i).Interface()
			}
		}
	default:
		panic("slice expected")
	}

	return nil
}
