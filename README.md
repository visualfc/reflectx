# reflectx
Golang reflect package hack tools

[![Go1.14](https://github.com/goplus/reflectx/workflows/Go1.14/badge.svg)](https://github.com/goplus/reflectx/actions?query=workflow%3AGo1.14)
[![Go1.15](https://github.com/goplus/reflectx/workflows/Go1.15/badge.svg)](https://github.com/goplus/reflectx/actions?query=workflow%3AGo1.15)

struct unexported field can set
* reflectx.CanSet
* reflectx.Field
* reflectx.FieldByIndex
* reflectx.FieldByName
* reflectx.FieldByNameFunc

```
type Point struct {
    x int
    y int
}

x := &Point{10, 20}
v := reflect.ValueOf(x).Elem()
sf := v.Field(0)

fmt.Println(sf.CanSet()) // output: false
// sf.SetInt(102)        
// panic: reflect: reflect.Value.SetInt using value obtained using unexported field

sf = reflectx.CanSet(sf)
fmt.Println(sf.CanSet()) // output: true

sf.SetInt(102)           // x.x = 102
fmt.Println(x.x)         // output: 102

reflectx.Field(x,1).SetInt(100) // x.y = 100
```

embedded type more than one field
* reflectx.StructOf
```
type Buffer struct {
	*bytes.Buffer
	X int
	Y int
}

typ := reflect.TypeOf((*Buffer)(nil)).Elem()
var fs []reflect.StructField
for i := 0; i < typ.NumField(); i++ {
	fs = append(fs, typ.Field(i))
}

// reflect.StructOf(fs) 
// panic reflect: embedded type with methods not implemented if there is more than one field

reflectx.StructOf(fs)

```

different types of fields generated by name
* reflectx.NamedStructOf
```
fs := []reflect.StructField{
	reflect.StructField{Name: "X", Type: reflect.TypeOf(0)},
	reflect.StructField{Name: "Y", Type: reflect.TypeOf(0)},
}
t1 := reflectx.NamedStructOf("github.com/goplus/reflectx","Point", fs)
t2 := reflectx.NamedStructOf("github.com/goplus/reflectx","Point", fs)
t3 := reflectx.NamedStructOf("github.com/goplus/reflectx", "Point2", fs)

// t1 == t2
// t1 != t3
//
// t1.String() == reflectx.Point
// t1.Name() == Point
// t1.PkgPath() == github.com/goplus/reflectx
```
* reflectx.NamedTypeOf
```
typ := reflectx.NamedTypeOf("main","Int",reflect.TypeOf(int(0)))
// typ.Kind() == reflect.Int
// typ.String() == "main.Int"
// typ.Name() == "Int"
// typ.PkgPath() == "main"
```

* reflectx.IsNamed
* reflectx.ToNamed