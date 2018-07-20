package refl

import "reflect"

func NestedFields(obj interface{}) []reflect.StructField {

	var fields = make([]reflect.StructField, 0)

	ot := reflect.TypeOf(obj)
	ov := reflect.ValueOf(obj)

	// dereference
	if ot.Kind() == reflect.Ptr {
		ot = ot.Elem()
		ov = ov.Elem()
	}

	for i := 0; i < ot.NumField(); i++ {

		fv := ov.Field(i)
		ft := ot.Field(i)

		isTime := ft.Type.Name() == "Time" && ft.PkgPath == ""

		if fv.Kind() == reflect.Struct && !isTime {
			fields = append(fields, NestedFields(fv.Interface())...)
		} else {
			fields = append(fields, ft)
		}
	}

	return fields
}

func ComposedOf(item interface{}, parent interface{}) bool {

	it := reflect.TypeOf(item)
	if it.Kind() == reflect.Ptr {
		it = it.Elem()
	}

	pt := reflect.TypeOf(parent)
	if pt.Kind() == reflect.Ptr {
		pt = pt.Elem()
	}
	if pt.Kind() != reflect.Struct {
		panic("parent must be struct type")
	}

	// find field with parent's exact name
	f, ok := it.FieldByName(pt.Name())
	if !ok {
		return false
	}

	if !f.Anonymous {
		return false
	}

	if !f.Type.ConvertibleTo(pt) {
		return false
	}

	return true
}
