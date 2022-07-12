//go:build (go1.17 && goexperiment.regabireflect) || (go1.19 && goexperiment.regabiargs) || (go1.18 && amd64) || (go1.19 && arm64) || (go1.19 && ppc64) || (go1.19 && ppc64le)
// +build go1.17,goexperiment.regabireflect go1.19,goexperiment.regabiargs go1.18,amd64 go1.19,arm64 go1.19,ppc64 go1.19,ppc64le

package icall

import (
	"reflect"
	"unsafe"

	"github.com/goplus/reflectx"
)

const capacity = 1024

type provider struct {
}

//go:linkname callReflect reflect.callReflect
func callReflect(ctxt unsafe.Pointer, frame unsafe.Pointer, retValid *bool, r unsafe.Pointer)

//go:linkname moveMakeFuncArgPtrs reflect.moveMakeFuncArgPtrs
func moveMakeFuncArgPtrs(ctx unsafe.Pointer, r unsafe.Pointer)

var infos []*reflectx.MethodInfo
var funcs []reflect.Value
var fnptr []unsafe.Pointer

func i_x(c unsafe.Pointer, frame unsafe.Pointer, retValid *bool, r unsafe.Pointer, index int) {
	moveMakeFuncArgPtrs(fnptr[index], r)
	callReflect(fnptr[index], frame, retValid, r)
}

func spillArgs()
func unspillArgs()

func (p *provider) Push(info *reflectx.MethodInfo) (ifn unsafe.Pointer) {
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
	funcs = append(funcs, fn)
	fnptr = append(fnptr, (*struct{ typ, ptr unsafe.Pointer })(unsafe.Pointer(&fn)).ptr)
	icall := icall_fn[len(infos)]
	infos = append(infos, info)
	return unsafe.Pointer(reflect.ValueOf(icall).Pointer())
}

func (p *provider) Len() int {
	return len(infos)
}

func (p *provider) Cap() int {
	return capacity
}

func (p *provider) Clear() {
	infos = nil
	funcs = nil
	fnptr = nil
}

var (
	mp provider
)

func init() {
	reflectx.AddMethodProvider(&mp)
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
func f256()
func f257()
func f258()
func f259()
func f260()
func f261()
func f262()
func f263()
func f264()
func f265()
func f266()
func f267()
func f268()
func f269()
func f270()
func f271()
func f272()
func f273()
func f274()
func f275()
func f276()
func f277()
func f278()
func f279()
func f280()
func f281()
func f282()
func f283()
func f284()
func f285()
func f286()
func f287()
func f288()
func f289()
func f290()
func f291()
func f292()
func f293()
func f294()
func f295()
func f296()
func f297()
func f298()
func f299()
func f300()
func f301()
func f302()
func f303()
func f304()
func f305()
func f306()
func f307()
func f308()
func f309()
func f310()
func f311()
func f312()
func f313()
func f314()
func f315()
func f316()
func f317()
func f318()
func f319()
func f320()
func f321()
func f322()
func f323()
func f324()
func f325()
func f326()
func f327()
func f328()
func f329()
func f330()
func f331()
func f332()
func f333()
func f334()
func f335()
func f336()
func f337()
func f338()
func f339()
func f340()
func f341()
func f342()
func f343()
func f344()
func f345()
func f346()
func f347()
func f348()
func f349()
func f350()
func f351()
func f352()
func f353()
func f354()
func f355()
func f356()
func f357()
func f358()
func f359()
func f360()
func f361()
func f362()
func f363()
func f364()
func f365()
func f366()
func f367()
func f368()
func f369()
func f370()
func f371()
func f372()
func f373()
func f374()
func f375()
func f376()
func f377()
func f378()
func f379()
func f380()
func f381()
func f382()
func f383()
func f384()
func f385()
func f386()
func f387()
func f388()
func f389()
func f390()
func f391()
func f392()
func f393()
func f394()
func f395()
func f396()
func f397()
func f398()
func f399()
func f400()
func f401()
func f402()
func f403()
func f404()
func f405()
func f406()
func f407()
func f408()
func f409()
func f410()
func f411()
func f412()
func f413()
func f414()
func f415()
func f416()
func f417()
func f418()
func f419()
func f420()
func f421()
func f422()
func f423()
func f424()
func f425()
func f426()
func f427()
func f428()
func f429()
func f430()
func f431()
func f432()
func f433()
func f434()
func f435()
func f436()
func f437()
func f438()
func f439()
func f440()
func f441()
func f442()
func f443()
func f444()
func f445()
func f446()
func f447()
func f448()
func f449()
func f450()
func f451()
func f452()
func f453()
func f454()
func f455()
func f456()
func f457()
func f458()
func f459()
func f460()
func f461()
func f462()
func f463()
func f464()
func f465()
func f466()
func f467()
func f468()
func f469()
func f470()
func f471()
func f472()
func f473()
func f474()
func f475()
func f476()
func f477()
func f478()
func f479()
func f480()
func f481()
func f482()
func f483()
func f484()
func f485()
func f486()
func f487()
func f488()
func f489()
func f490()
func f491()
func f492()
func f493()
func f494()
func f495()
func f496()
func f497()
func f498()
func f499()
func f500()
func f501()
func f502()
func f503()
func f504()
func f505()
func f506()
func f507()
func f508()
func f509()
func f510()
func f511()
func f512()
func f513()
func f514()
func f515()
func f516()
func f517()
func f518()
func f519()
func f520()
func f521()
func f522()
func f523()
func f524()
func f525()
func f526()
func f527()
func f528()
func f529()
func f530()
func f531()
func f532()
func f533()
func f534()
func f535()
func f536()
func f537()
func f538()
func f539()
func f540()
func f541()
func f542()
func f543()
func f544()
func f545()
func f546()
func f547()
func f548()
func f549()
func f550()
func f551()
func f552()
func f553()
func f554()
func f555()
func f556()
func f557()
func f558()
func f559()
func f560()
func f561()
func f562()
func f563()
func f564()
func f565()
func f566()
func f567()
func f568()
func f569()
func f570()
func f571()
func f572()
func f573()
func f574()
func f575()
func f576()
func f577()
func f578()
func f579()
func f580()
func f581()
func f582()
func f583()
func f584()
func f585()
func f586()
func f587()
func f588()
func f589()
func f590()
func f591()
func f592()
func f593()
func f594()
func f595()
func f596()
func f597()
func f598()
func f599()
func f600()
func f601()
func f602()
func f603()
func f604()
func f605()
func f606()
func f607()
func f608()
func f609()
func f610()
func f611()
func f612()
func f613()
func f614()
func f615()
func f616()
func f617()
func f618()
func f619()
func f620()
func f621()
func f622()
func f623()
func f624()
func f625()
func f626()
func f627()
func f628()
func f629()
func f630()
func f631()
func f632()
func f633()
func f634()
func f635()
func f636()
func f637()
func f638()
func f639()
func f640()
func f641()
func f642()
func f643()
func f644()
func f645()
func f646()
func f647()
func f648()
func f649()
func f650()
func f651()
func f652()
func f653()
func f654()
func f655()
func f656()
func f657()
func f658()
func f659()
func f660()
func f661()
func f662()
func f663()
func f664()
func f665()
func f666()
func f667()
func f668()
func f669()
func f670()
func f671()
func f672()
func f673()
func f674()
func f675()
func f676()
func f677()
func f678()
func f679()
func f680()
func f681()
func f682()
func f683()
func f684()
func f685()
func f686()
func f687()
func f688()
func f689()
func f690()
func f691()
func f692()
func f693()
func f694()
func f695()
func f696()
func f697()
func f698()
func f699()
func f700()
func f701()
func f702()
func f703()
func f704()
func f705()
func f706()
func f707()
func f708()
func f709()
func f710()
func f711()
func f712()
func f713()
func f714()
func f715()
func f716()
func f717()
func f718()
func f719()
func f720()
func f721()
func f722()
func f723()
func f724()
func f725()
func f726()
func f727()
func f728()
func f729()
func f730()
func f731()
func f732()
func f733()
func f734()
func f735()
func f736()
func f737()
func f738()
func f739()
func f740()
func f741()
func f742()
func f743()
func f744()
func f745()
func f746()
func f747()
func f748()
func f749()
func f750()
func f751()
func f752()
func f753()
func f754()
func f755()
func f756()
func f757()
func f758()
func f759()
func f760()
func f761()
func f762()
func f763()
func f764()
func f765()
func f766()
func f767()
func f768()
func f769()
func f770()
func f771()
func f772()
func f773()
func f774()
func f775()
func f776()
func f777()
func f778()
func f779()
func f780()
func f781()
func f782()
func f783()
func f784()
func f785()
func f786()
func f787()
func f788()
func f789()
func f790()
func f791()
func f792()
func f793()
func f794()
func f795()
func f796()
func f797()
func f798()
func f799()
func f800()
func f801()
func f802()
func f803()
func f804()
func f805()
func f806()
func f807()
func f808()
func f809()
func f810()
func f811()
func f812()
func f813()
func f814()
func f815()
func f816()
func f817()
func f818()
func f819()
func f820()
func f821()
func f822()
func f823()
func f824()
func f825()
func f826()
func f827()
func f828()
func f829()
func f830()
func f831()
func f832()
func f833()
func f834()
func f835()
func f836()
func f837()
func f838()
func f839()
func f840()
func f841()
func f842()
func f843()
func f844()
func f845()
func f846()
func f847()
func f848()
func f849()
func f850()
func f851()
func f852()
func f853()
func f854()
func f855()
func f856()
func f857()
func f858()
func f859()
func f860()
func f861()
func f862()
func f863()
func f864()
func f865()
func f866()
func f867()
func f868()
func f869()
func f870()
func f871()
func f872()
func f873()
func f874()
func f875()
func f876()
func f877()
func f878()
func f879()
func f880()
func f881()
func f882()
func f883()
func f884()
func f885()
func f886()
func f887()
func f888()
func f889()
func f890()
func f891()
func f892()
func f893()
func f894()
func f895()
func f896()
func f897()
func f898()
func f899()
func f900()
func f901()
func f902()
func f903()
func f904()
func f905()
func f906()
func f907()
func f908()
func f909()
func f910()
func f911()
func f912()
func f913()
func f914()
func f915()
func f916()
func f917()
func f918()
func f919()
func f920()
func f921()
func f922()
func f923()
func f924()
func f925()
func f926()
func f927()
func f928()
func f929()
func f930()
func f931()
func f932()
func f933()
func f934()
func f935()
func f936()
func f937()
func f938()
func f939()
func f940()
func f941()
func f942()
func f943()
func f944()
func f945()
func f946()
func f947()
func f948()
func f949()
func f950()
func f951()
func f952()
func f953()
func f954()
func f955()
func f956()
func f957()
func f958()
func f959()
func f960()
func f961()
func f962()
func f963()
func f964()
func f965()
func f966()
func f967()
func f968()
func f969()
func f970()
func f971()
func f972()
func f973()
func f974()
func f975()
func f976()
func f977()
func f978()
func f979()
func f980()
func f981()
func f982()
func f983()
func f984()
func f985()
func f986()
func f987()
func f988()
func f989()
func f990()
func f991()
func f992()
func f993()
func f994()
func f995()
func f996()
func f997()
func f998()
func f999()
func f1000()
func f1001()
func f1002()
func f1003()
func f1004()
func f1005()
func f1006()
func f1007()
func f1008()
func f1009()
func f1010()
func f1011()
func f1012()
func f1013()
func f1014()
func f1015()
func f1016()
func f1017()
func f1018()
func f1019()
func f1020()
func f1021()
func f1022()
func f1023()

var (
	icall_fn = []func(){f0,f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13,f14,f15,f16,f17,f18,f19,f20,f21,f22,f23,f24,f25,f26,f27,f28,f29,f30,f31,f32,f33,f34,f35,f36,f37,f38,f39,f40,f41,f42,f43,f44,f45,f46,f47,f48,f49,f50,f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65,f66,f67,f68,f69,f70,f71,f72,f73,f74,f75,f76,f77,f78,f79,f80,f81,f82,f83,f84,f85,f86,f87,f88,f89,f90,f91,f92,f93,f94,f95,f96,f97,f98,f99,f100,f101,f102,f103,f104,f105,f106,f107,f108,f109,f110,f111,f112,f113,f114,f115,f116,f117,f118,f119,f120,f121,f122,f123,f124,f125,f126,f127,f128,f129,f130,f131,f132,f133,f134,f135,f136,f137,f138,f139,f140,f141,f142,f143,f144,f145,f146,f147,f148,f149,f150,f151,f152,f153,f154,f155,f156,f157,f158,f159,f160,f161,f162,f163,f164,f165,f166,f167,f168,f169,f170,f171,f172,f173,f174,f175,f176,f177,f178,f179,f180,f181,f182,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f193,f194,f195,f196,f197,f198,f199,f200,f201,f202,f203,f204,f205,f206,f207,f208,f209,f210,f211,f212,f213,f214,f215,f216,f217,f218,f219,f220,f221,f222,f223,f224,f225,f226,f227,f228,f229,f230,f231,f232,f233,f234,f235,f236,f237,f238,f239,f240,f241,f242,f243,f244,f245,f246,f247,f248,f249,f250,f251,f252,f253,f254,f255,f256,f257,f258,f259,f260,f261,f262,f263,f264,f265,f266,f267,f268,f269,f270,f271,f272,f273,f274,f275,f276,f277,f278,f279,f280,f281,f282,f283,f284,f285,f286,f287,f288,f289,f290,f291,f292,f293,f294,f295,f296,f297,f298,f299,f300,f301,f302,f303,f304,f305,f306,f307,f308,f309,f310,f311,f312,f313,f314,f315,f316,f317,f318,f319,f320,f321,f322,f323,f324,f325,f326,f327,f328,f329,f330,f331,f332,f333,f334,f335,f336,f337,f338,f339,f340,f341,f342,f343,f344,f345,f346,f347,f348,f349,f350,f351,f352,f353,f354,f355,f356,f357,f358,f359,f360,f361,f362,f363,f364,f365,f366,f367,f368,f369,f370,f371,f372,f373,f374,f375,f376,f377,f378,f379,f380,f381,f382,f383,f384,f385,f386,f387,f388,f389,f390,f391,f392,f393,f394,f395,f396,f397,f398,f399,f400,f401,f402,f403,f404,f405,f406,f407,f408,f409,f410,f411,f412,f413,f414,f415,f416,f417,f418,f419,f420,f421,f422,f423,f424,f425,f426,f427,f428,f429,f430,f431,f432,f433,f434,f435,f436,f437,f438,f439,f440,f441,f442,f443,f444,f445,f446,f447,f448,f449,f450,f451,f452,f453,f454,f455,f456,f457,f458,f459,f460,f461,f462,f463,f464,f465,f466,f467,f468,f469,f470,f471,f472,f473,f474,f475,f476,f477,f478,f479,f480,f481,f482,f483,f484,f485,f486,f487,f488,f489,f490,f491,f492,f493,f494,f495,f496,f497,f498,f499,f500,f501,f502,f503,f504,f505,f506,f507,f508,f509,f510,f511,f512,f513,f514,f515,f516,f517,f518,f519,f520,f521,f522,f523,f524,f525,f526,f527,f528,f529,f530,f531,f532,f533,f534,f535,f536,f537,f538,f539,f540,f541,f542,f543,f544,f545,f546,f547,f548,f549,f550,f551,f552,f553,f554,f555,f556,f557,f558,f559,f560,f561,f562,f563,f564,f565,f566,f567,f568,f569,f570,f571,f572,f573,f574,f575,f576,f577,f578,f579,f580,f581,f582,f583,f584,f585,f586,f587,f588,f589,f590,f591,f592,f593,f594,f595,f596,f597,f598,f599,f600,f601,f602,f603,f604,f605,f606,f607,f608,f609,f610,f611,f612,f613,f614,f615,f616,f617,f618,f619,f620,f621,f622,f623,f624,f625,f626,f627,f628,f629,f630,f631,f632,f633,f634,f635,f636,f637,f638,f639,f640,f641,f642,f643,f644,f645,f646,f647,f648,f649,f650,f651,f652,f653,f654,f655,f656,f657,f658,f659,f660,f661,f662,f663,f664,f665,f666,f667,f668,f669,f670,f671,f672,f673,f674,f675,f676,f677,f678,f679,f680,f681,f682,f683,f684,f685,f686,f687,f688,f689,f690,f691,f692,f693,f694,f695,f696,f697,f698,f699,f700,f701,f702,f703,f704,f705,f706,f707,f708,f709,f710,f711,f712,f713,f714,f715,f716,f717,f718,f719,f720,f721,f722,f723,f724,f725,f726,f727,f728,f729,f730,f731,f732,f733,f734,f735,f736,f737,f738,f739,f740,f741,f742,f743,f744,f745,f746,f747,f748,f749,f750,f751,f752,f753,f754,f755,f756,f757,f758,f759,f760,f761,f762,f763,f764,f765,f766,f767,f768,f769,f770,f771,f772,f773,f774,f775,f776,f777,f778,f779,f780,f781,f782,f783,f784,f785,f786,f787,f788,f789,f790,f791,f792,f793,f794,f795,f796,f797,f798,f799,f800,f801,f802,f803,f804,f805,f806,f807,f808,f809,f810,f811,f812,f813,f814,f815,f816,f817,f818,f819,f820,f821,f822,f823,f824,f825,f826,f827,f828,f829,f830,f831,f832,f833,f834,f835,f836,f837,f838,f839,f840,f841,f842,f843,f844,f845,f846,f847,f848,f849,f850,f851,f852,f853,f854,f855,f856,f857,f858,f859,f860,f861,f862,f863,f864,f865,f866,f867,f868,f869,f870,f871,f872,f873,f874,f875,f876,f877,f878,f879,f880,f881,f882,f883,f884,f885,f886,f887,f888,f889,f890,f891,f892,f893,f894,f895,f896,f897,f898,f899,f900,f901,f902,f903,f904,f905,f906,f907,f908,f909,f910,f911,f912,f913,f914,f915,f916,f917,f918,f919,f920,f921,f922,f923,f924,f925,f926,f927,f928,f929,f930,f931,f932,f933,f934,f935,f936,f937,f938,f939,f940,f941,f942,f943,f944,f945,f946,f947,f948,f949,f950,f951,f952,f953,f954,f955,f956,f957,f958,f959,f960,f961,f962,f963,f964,f965,f966,f967,f968,f969,f970,f971,f972,f973,f974,f975,f976,f977,f978,f979,f980,f981,f982,f983,f984,f985,f986,f987,f988,f989,f990,f991,f992,f993,f994,f995,f996,f997,f998,f999,f1000,f1001,f1002,f1003,f1004,f1005,f1006,f1007,f1008,f1009,f1010,f1011,f1012,f1013,f1014,f1015,f1016,f1017,f1018,f1019,f1020,f1021,f1022,f1023}
)
