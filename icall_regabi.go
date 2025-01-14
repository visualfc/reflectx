//go:build ((go1.17 && goexperiment.regabireflect) || (go1.19 && goexperiment.regabiargs) || (go1.18 && amd64) || (go1.19 && arm64) || (go1.19 && ppc64) || (go1.19 && ppc64le) || (go1.20 && riscv64)) && (!js || (js && wasm))
// +build go1.17,goexperiment.regabireflect go1.19,goexperiment.regabiargs go1.18,amd64 go1.19,arm64 go1.19,ppc64 go1.19,ppc64le go1.20,riscv64
// +build !js js,wasm

package reflectx

import (
	"reflect"
	"unsafe"
)

const capacity = 256

type methodUsed struct {
	fun reflect.Value
	ptr unsafe.Pointer
}

type provider struct {
	used map[int]*methodUsed
}

//go:linkname callReflect reflect.callReflect
func callReflect(ctxt unsafe.Pointer, frame unsafe.Pointer, retValid *bool, r unsafe.Pointer)

//go:linkname moveMakeFuncArgPtrs reflect.moveMakeFuncArgPtrs
func moveMakeFuncArgPtrs(ctx unsafe.Pointer, r unsafe.Pointer)

func i_x(c unsafe.Pointer, frame unsafe.Pointer, retValid *bool, r unsafe.Pointer, index int) {
	ptr := mp.used[index].ptr
	moveMakeFuncArgPtrs(ptr, r)
	callReflect(ptr, frame, retValid, r)
}

func spillArgs()
func unspillArgs()

func (p *provider) Insert(info *MethodInfo) (unsafe.Pointer, int) {
	var index = -1
	for i := 0; i < capacity; i++ {
		if _, ok := p.used[i]; !ok {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, -1
	}
	var fn reflect.Value
	if (!info.Pointer && !info.OnePtr) || info.Indirect {
		ftyp := info.Func.Type()
		numIn := ftyp.NumIn()
		numOut := ftyp.NumOut()
		in := make([]reflect.Type, numIn, numIn)
		out := make([]reflect.Type, numOut, numOut)
		in[0] = reflect.PtrTo(info.Type)
		for i := 1; i < numIn; i++ {
			in[i] = ftyp.In(i)
		}
		for i := 0; i < numOut; i++ {
			out[i] = ftyp.Out(i)
		}
		ftyp = reflect.FuncOf(in, out, info.Variadic)
		if info.Variadic {
			fn = reflect.MakeFunc(ftyp, func(args []reflect.Value) []reflect.Value {
				args[0] = args[0].Elem()
				return info.Func.CallSlice(args)
			})
		} else {
			fn = reflect.MakeFunc(ftyp, func(args []reflect.Value) []reflect.Value {
				args[0] = args[0].Elem()
				return info.Func.Call(args)
			})
		}
	} else {
		fn = info.Func
	}
	p.used[index] = &methodUsed{
		fun: fn,
		ptr: (*struct{ typ, ptr unsafe.Pointer })(unsafe.Pointer(&fn)).ptr,
	}
	icall := icall_fn[index]
	return unsafe.Pointer(reflect.ValueOf(icall).Pointer()), index
}

func (p *provider) Remove(indexs []int) {
	for _, n := range indexs {
		delete(p.used, n)
	}
}

func (p *provider) Available() int {
	return capacity - len(p.used)
}

func (p *provider) Used() int {
	return len(p.used)
}

func (p *provider) Cap() int {
	return capacity
}

func (p *provider) Clear() {
	p.used = make(map[int]*methodUsed)
}

var (
	mp = &provider{
		used: make(map[int]*methodUsed),
	}
)

func init() {
	AddMethodProvider(mp)
}

func f0()
func f1()
func f2()
func f3()
func f4()
func f5()
func f6()
func f7()
func f8()
func f9()
func f10()
func f11()
func f12()
func f13()
func f14()
func f15()
func f16()
func f17()
func f18()
func f19()
func f20()
func f21()
func f22()
func f23()
func f24()
func f25()
func f26()
func f27()
func f28()
func f29()
func f30()
func f31()
func f32()
func f33()
func f34()
func f35()
func f36()
func f37()
func f38()
func f39()
func f40()
func f41()
func f42()
func f43()
func f44()
func f45()
func f46()
func f47()
func f48()
func f49()
func f50()
func f51()
func f52()
func f53()
func f54()
func f55()
func f56()
func f57()
func f58()
func f59()
func f60()
func f61()
func f62()
func f63()
func f64()
func f65()
func f66()
func f67()
func f68()
func f69()
func f70()
func f71()
func f72()
func f73()
func f74()
func f75()
func f76()
func f77()
func f78()
func f79()
func f80()
func f81()
func f82()
func f83()
func f84()
func f85()
func f86()
func f87()
func f88()
func f89()
func f90()
func f91()
func f92()
func f93()
func f94()
func f95()
func f96()
func f97()
func f98()
func f99()
func f100()
func f101()
func f102()
func f103()
func f104()
func f105()
func f106()
func f107()
func f108()
func f109()
func f110()
func f111()
func f112()
func f113()
func f114()
func f115()
func f116()
func f117()
func f118()
func f119()
func f120()
func f121()
func f122()
func f123()
func f124()
func f125()
func f126()
func f127()
func f128()
func f129()
func f130()
func f131()
func f132()
func f133()
func f134()
func f135()
func f136()
func f137()
func f138()
func f139()
func f140()
func f141()
func f142()
func f143()
func f144()
func f145()
func f146()
func f147()
func f148()
func f149()
func f150()
func f151()
func f152()
func f153()
func f154()
func f155()
func f156()
func f157()
func f158()
func f159()
func f160()
func f161()
func f162()
func f163()
func f164()
func f165()
func f166()
func f167()
func f168()
func f169()
func f170()
func f171()
func f172()
func f173()
func f174()
func f175()
func f176()
func f177()
func f178()
func f179()
func f180()
func f181()
func f182()
func f183()
func f184()
func f185()
func f186()
func f187()
func f188()
func f189()
func f190()
func f191()
func f192()
func f193()
func f194()
func f195()
func f196()
func f197()
func f198()
func f199()
func f200()
func f201()
func f202()
func f203()
func f204()
func f205()
func f206()
func f207()
func f208()
func f209()
func f210()
func f211()
func f212()
func f213()
func f214()
func f215()
func f216()
func f217()
func f218()
func f219()
func f220()
func f221()
func f222()
func f223()
func f224()
func f225()
func f226()
func f227()
func f228()
func f229()
func f230()
func f231()
func f232()
func f233()
func f234()
func f235()
func f236()
func f237()
func f238()
func f239()
func f240()
func f241()
func f242()
func f243()
func f244()
func f245()
func f246()
func f247()
func f248()
func f249()
func f250()
func f251()
func f252()
func f253()
func f254()
func f255()

var (
	icall_fn = []func(){f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27, f28, f29, f30, f31, f32, f33, f34, f35, f36, f37, f38, f39, f40, f41, f42, f43, f44, f45, f46, f47, f48, f49, f50, f51, f52, f53, f54, f55, f56, f57, f58, f59, f60, f61, f62, f63, f64, f65, f66, f67, f68, f69, f70, f71, f72, f73, f74, f75, f76, f77, f78, f79, f80, f81, f82, f83, f84, f85, f86, f87, f88, f89, f90, f91, f92, f93, f94, f95, f96, f97, f98, f99, f100, f101, f102, f103, f104, f105, f106, f107, f108, f109, f110, f111, f112, f113, f114, f115, f116, f117, f118, f119, f120, f121, f122, f123, f124, f125, f126, f127, f128, f129, f130, f131, f132, f133, f134, f135, f136, f137, f138, f139, f140, f141, f142, f143, f144, f145, f146, f147, f148, f149, f150, f151, f152, f153, f154, f155, f156, f157, f158, f159, f160, f161, f162, f163, f164, f165, f166, f167, f168, f169, f170, f171, f172, f173, f174, f175, f176, f177, f178, f179, f180, f181, f182, f183, f184, f185, f186, f187, f188, f189, f190, f191, f192, f193, f194, f195, f196, f197, f198, f199, f200, f201, f202, f203, f204, f205, f206, f207, f208, f209, f210, f211, f212, f213, f214, f215, f216, f217, f218, f219, f220, f221, f222, f223, f224, f225, f226, f227, f228, f229, f230, f231, f232, f233, f234, f235, f236, f237, f238, f239, f240, f241, f242, f243, f244, f245, f246, f247, f248, f249, f250, f251, f252, f253, f254, f255}
)
