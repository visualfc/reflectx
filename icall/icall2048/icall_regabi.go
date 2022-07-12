//go:build (go1.17 && goexperiment.regabireflect) || (go1.19 && goexperiment.regabiargs) || (go1.18 && amd64) || (go1.19 && arm64) || (go1.19 && ppc64) || (go1.19 && ppc64le)
// +build go1.17,goexperiment.regabireflect go1.19,goexperiment.regabiargs go1.18,amd64 go1.19,arm64 go1.19,ppc64 go1.19,ppc64le

package icall

import (
	"reflect"
	"unsafe"

	"github.com/goplus/reflectx"
)

const capacity = 2048

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
func f1024()
func f1025()
func f1026()
func f1027()
func f1028()
func f1029()
func f1030()
func f1031()
func f1032()
func f1033()
func f1034()
func f1035()
func f1036()
func f1037()
func f1038()
func f1039()
func f1040()
func f1041()
func f1042()
func f1043()
func f1044()
func f1045()
func f1046()
func f1047()
func f1048()
func f1049()
func f1050()
func f1051()
func f1052()
func f1053()
func f1054()
func f1055()
func f1056()
func f1057()
func f1058()
func f1059()
func f1060()
func f1061()
func f1062()
func f1063()
func f1064()
func f1065()
func f1066()
func f1067()
func f1068()
func f1069()
func f1070()
func f1071()
func f1072()
func f1073()
func f1074()
func f1075()
func f1076()
func f1077()
func f1078()
func f1079()
func f1080()
func f1081()
func f1082()
func f1083()
func f1084()
func f1085()
func f1086()
func f1087()
func f1088()
func f1089()
func f1090()
func f1091()
func f1092()
func f1093()
func f1094()
func f1095()
func f1096()
func f1097()
func f1098()
func f1099()
func f1100()
func f1101()
func f1102()
func f1103()
func f1104()
func f1105()
func f1106()
func f1107()
func f1108()
func f1109()
func f1110()
func f1111()
func f1112()
func f1113()
func f1114()
func f1115()
func f1116()
func f1117()
func f1118()
func f1119()
func f1120()
func f1121()
func f1122()
func f1123()
func f1124()
func f1125()
func f1126()
func f1127()
func f1128()
func f1129()
func f1130()
func f1131()
func f1132()
func f1133()
func f1134()
func f1135()
func f1136()
func f1137()
func f1138()
func f1139()
func f1140()
func f1141()
func f1142()
func f1143()
func f1144()
func f1145()
func f1146()
func f1147()
func f1148()
func f1149()
func f1150()
func f1151()
func f1152()
func f1153()
func f1154()
func f1155()
func f1156()
func f1157()
func f1158()
func f1159()
func f1160()
func f1161()
func f1162()
func f1163()
func f1164()
func f1165()
func f1166()
func f1167()
func f1168()
func f1169()
func f1170()
func f1171()
func f1172()
func f1173()
func f1174()
func f1175()
func f1176()
func f1177()
func f1178()
func f1179()
func f1180()
func f1181()
func f1182()
func f1183()
func f1184()
func f1185()
func f1186()
func f1187()
func f1188()
func f1189()
func f1190()
func f1191()
func f1192()
func f1193()
func f1194()
func f1195()
func f1196()
func f1197()
func f1198()
func f1199()
func f1200()
func f1201()
func f1202()
func f1203()
func f1204()
func f1205()
func f1206()
func f1207()
func f1208()
func f1209()
func f1210()
func f1211()
func f1212()
func f1213()
func f1214()
func f1215()
func f1216()
func f1217()
func f1218()
func f1219()
func f1220()
func f1221()
func f1222()
func f1223()
func f1224()
func f1225()
func f1226()
func f1227()
func f1228()
func f1229()
func f1230()
func f1231()
func f1232()
func f1233()
func f1234()
func f1235()
func f1236()
func f1237()
func f1238()
func f1239()
func f1240()
func f1241()
func f1242()
func f1243()
func f1244()
func f1245()
func f1246()
func f1247()
func f1248()
func f1249()
func f1250()
func f1251()
func f1252()
func f1253()
func f1254()
func f1255()
func f1256()
func f1257()
func f1258()
func f1259()
func f1260()
func f1261()
func f1262()
func f1263()
func f1264()
func f1265()
func f1266()
func f1267()
func f1268()
func f1269()
func f1270()
func f1271()
func f1272()
func f1273()
func f1274()
func f1275()
func f1276()
func f1277()
func f1278()
func f1279()
func f1280()
func f1281()
func f1282()
func f1283()
func f1284()
func f1285()
func f1286()
func f1287()
func f1288()
func f1289()
func f1290()
func f1291()
func f1292()
func f1293()
func f1294()
func f1295()
func f1296()
func f1297()
func f1298()
func f1299()
func f1300()
func f1301()
func f1302()
func f1303()
func f1304()
func f1305()
func f1306()
func f1307()
func f1308()
func f1309()
func f1310()
func f1311()
func f1312()
func f1313()
func f1314()
func f1315()
func f1316()
func f1317()
func f1318()
func f1319()
func f1320()
func f1321()
func f1322()
func f1323()
func f1324()
func f1325()
func f1326()
func f1327()
func f1328()
func f1329()
func f1330()
func f1331()
func f1332()
func f1333()
func f1334()
func f1335()
func f1336()
func f1337()
func f1338()
func f1339()
func f1340()
func f1341()
func f1342()
func f1343()
func f1344()
func f1345()
func f1346()
func f1347()
func f1348()
func f1349()
func f1350()
func f1351()
func f1352()
func f1353()
func f1354()
func f1355()
func f1356()
func f1357()
func f1358()
func f1359()
func f1360()
func f1361()
func f1362()
func f1363()
func f1364()
func f1365()
func f1366()
func f1367()
func f1368()
func f1369()
func f1370()
func f1371()
func f1372()
func f1373()
func f1374()
func f1375()
func f1376()
func f1377()
func f1378()
func f1379()
func f1380()
func f1381()
func f1382()
func f1383()
func f1384()
func f1385()
func f1386()
func f1387()
func f1388()
func f1389()
func f1390()
func f1391()
func f1392()
func f1393()
func f1394()
func f1395()
func f1396()
func f1397()
func f1398()
func f1399()
func f1400()
func f1401()
func f1402()
func f1403()
func f1404()
func f1405()
func f1406()
func f1407()
func f1408()
func f1409()
func f1410()
func f1411()
func f1412()
func f1413()
func f1414()
func f1415()
func f1416()
func f1417()
func f1418()
func f1419()
func f1420()
func f1421()
func f1422()
func f1423()
func f1424()
func f1425()
func f1426()
func f1427()
func f1428()
func f1429()
func f1430()
func f1431()
func f1432()
func f1433()
func f1434()
func f1435()
func f1436()
func f1437()
func f1438()
func f1439()
func f1440()
func f1441()
func f1442()
func f1443()
func f1444()
func f1445()
func f1446()
func f1447()
func f1448()
func f1449()
func f1450()
func f1451()
func f1452()
func f1453()
func f1454()
func f1455()
func f1456()
func f1457()
func f1458()
func f1459()
func f1460()
func f1461()
func f1462()
func f1463()
func f1464()
func f1465()
func f1466()
func f1467()
func f1468()
func f1469()
func f1470()
func f1471()
func f1472()
func f1473()
func f1474()
func f1475()
func f1476()
func f1477()
func f1478()
func f1479()
func f1480()
func f1481()
func f1482()
func f1483()
func f1484()
func f1485()
func f1486()
func f1487()
func f1488()
func f1489()
func f1490()
func f1491()
func f1492()
func f1493()
func f1494()
func f1495()
func f1496()
func f1497()
func f1498()
func f1499()
func f1500()
func f1501()
func f1502()
func f1503()
func f1504()
func f1505()
func f1506()
func f1507()
func f1508()
func f1509()
func f1510()
func f1511()
func f1512()
func f1513()
func f1514()
func f1515()
func f1516()
func f1517()
func f1518()
func f1519()
func f1520()
func f1521()
func f1522()
func f1523()
func f1524()
func f1525()
func f1526()
func f1527()
func f1528()
func f1529()
func f1530()
func f1531()
func f1532()
func f1533()
func f1534()
func f1535()
func f1536()
func f1537()
func f1538()
func f1539()
func f1540()
func f1541()
func f1542()
func f1543()
func f1544()
func f1545()
func f1546()
func f1547()
func f1548()
func f1549()
func f1550()
func f1551()
func f1552()
func f1553()
func f1554()
func f1555()
func f1556()
func f1557()
func f1558()
func f1559()
func f1560()
func f1561()
func f1562()
func f1563()
func f1564()
func f1565()
func f1566()
func f1567()
func f1568()
func f1569()
func f1570()
func f1571()
func f1572()
func f1573()
func f1574()
func f1575()
func f1576()
func f1577()
func f1578()
func f1579()
func f1580()
func f1581()
func f1582()
func f1583()
func f1584()
func f1585()
func f1586()
func f1587()
func f1588()
func f1589()
func f1590()
func f1591()
func f1592()
func f1593()
func f1594()
func f1595()
func f1596()
func f1597()
func f1598()
func f1599()
func f1600()
func f1601()
func f1602()
func f1603()
func f1604()
func f1605()
func f1606()
func f1607()
func f1608()
func f1609()
func f1610()
func f1611()
func f1612()
func f1613()
func f1614()
func f1615()
func f1616()
func f1617()
func f1618()
func f1619()
func f1620()
func f1621()
func f1622()
func f1623()
func f1624()
func f1625()
func f1626()
func f1627()
func f1628()
func f1629()
func f1630()
func f1631()
func f1632()
func f1633()
func f1634()
func f1635()
func f1636()
func f1637()
func f1638()
func f1639()
func f1640()
func f1641()
func f1642()
func f1643()
func f1644()
func f1645()
func f1646()
func f1647()
func f1648()
func f1649()
func f1650()
func f1651()
func f1652()
func f1653()
func f1654()
func f1655()
func f1656()
func f1657()
func f1658()
func f1659()
func f1660()
func f1661()
func f1662()
func f1663()
func f1664()
func f1665()
func f1666()
func f1667()
func f1668()
func f1669()
func f1670()
func f1671()
func f1672()
func f1673()
func f1674()
func f1675()
func f1676()
func f1677()
func f1678()
func f1679()
func f1680()
func f1681()
func f1682()
func f1683()
func f1684()
func f1685()
func f1686()
func f1687()
func f1688()
func f1689()
func f1690()
func f1691()
func f1692()
func f1693()
func f1694()
func f1695()
func f1696()
func f1697()
func f1698()
func f1699()
func f1700()
func f1701()
func f1702()
func f1703()
func f1704()
func f1705()
func f1706()
func f1707()
func f1708()
func f1709()
func f1710()
func f1711()
func f1712()
func f1713()
func f1714()
func f1715()
func f1716()
func f1717()
func f1718()
func f1719()
func f1720()
func f1721()
func f1722()
func f1723()
func f1724()
func f1725()
func f1726()
func f1727()
func f1728()
func f1729()
func f1730()
func f1731()
func f1732()
func f1733()
func f1734()
func f1735()
func f1736()
func f1737()
func f1738()
func f1739()
func f1740()
func f1741()
func f1742()
func f1743()
func f1744()
func f1745()
func f1746()
func f1747()
func f1748()
func f1749()
func f1750()
func f1751()
func f1752()
func f1753()
func f1754()
func f1755()
func f1756()
func f1757()
func f1758()
func f1759()
func f1760()
func f1761()
func f1762()
func f1763()
func f1764()
func f1765()
func f1766()
func f1767()
func f1768()
func f1769()
func f1770()
func f1771()
func f1772()
func f1773()
func f1774()
func f1775()
func f1776()
func f1777()
func f1778()
func f1779()
func f1780()
func f1781()
func f1782()
func f1783()
func f1784()
func f1785()
func f1786()
func f1787()
func f1788()
func f1789()
func f1790()
func f1791()
func f1792()
func f1793()
func f1794()
func f1795()
func f1796()
func f1797()
func f1798()
func f1799()
func f1800()
func f1801()
func f1802()
func f1803()
func f1804()
func f1805()
func f1806()
func f1807()
func f1808()
func f1809()
func f1810()
func f1811()
func f1812()
func f1813()
func f1814()
func f1815()
func f1816()
func f1817()
func f1818()
func f1819()
func f1820()
func f1821()
func f1822()
func f1823()
func f1824()
func f1825()
func f1826()
func f1827()
func f1828()
func f1829()
func f1830()
func f1831()
func f1832()
func f1833()
func f1834()
func f1835()
func f1836()
func f1837()
func f1838()
func f1839()
func f1840()
func f1841()
func f1842()
func f1843()
func f1844()
func f1845()
func f1846()
func f1847()
func f1848()
func f1849()
func f1850()
func f1851()
func f1852()
func f1853()
func f1854()
func f1855()
func f1856()
func f1857()
func f1858()
func f1859()
func f1860()
func f1861()
func f1862()
func f1863()
func f1864()
func f1865()
func f1866()
func f1867()
func f1868()
func f1869()
func f1870()
func f1871()
func f1872()
func f1873()
func f1874()
func f1875()
func f1876()
func f1877()
func f1878()
func f1879()
func f1880()
func f1881()
func f1882()
func f1883()
func f1884()
func f1885()
func f1886()
func f1887()
func f1888()
func f1889()
func f1890()
func f1891()
func f1892()
func f1893()
func f1894()
func f1895()
func f1896()
func f1897()
func f1898()
func f1899()
func f1900()
func f1901()
func f1902()
func f1903()
func f1904()
func f1905()
func f1906()
func f1907()
func f1908()
func f1909()
func f1910()
func f1911()
func f1912()
func f1913()
func f1914()
func f1915()
func f1916()
func f1917()
func f1918()
func f1919()
func f1920()
func f1921()
func f1922()
func f1923()
func f1924()
func f1925()
func f1926()
func f1927()
func f1928()
func f1929()
func f1930()
func f1931()
func f1932()
func f1933()
func f1934()
func f1935()
func f1936()
func f1937()
func f1938()
func f1939()
func f1940()
func f1941()
func f1942()
func f1943()
func f1944()
func f1945()
func f1946()
func f1947()
func f1948()
func f1949()
func f1950()
func f1951()
func f1952()
func f1953()
func f1954()
func f1955()
func f1956()
func f1957()
func f1958()
func f1959()
func f1960()
func f1961()
func f1962()
func f1963()
func f1964()
func f1965()
func f1966()
func f1967()
func f1968()
func f1969()
func f1970()
func f1971()
func f1972()
func f1973()
func f1974()
func f1975()
func f1976()
func f1977()
func f1978()
func f1979()
func f1980()
func f1981()
func f1982()
func f1983()
func f1984()
func f1985()
func f1986()
func f1987()
func f1988()
func f1989()
func f1990()
func f1991()
func f1992()
func f1993()
func f1994()
func f1995()
func f1996()
func f1997()
func f1998()
func f1999()
func f2000()
func f2001()
func f2002()
func f2003()
func f2004()
func f2005()
func f2006()
func f2007()
func f2008()
func f2009()
func f2010()
func f2011()
func f2012()
func f2013()
func f2014()
func f2015()
func f2016()
func f2017()
func f2018()
func f2019()
func f2020()
func f2021()
func f2022()
func f2023()
func f2024()
func f2025()
func f2026()
func f2027()
func f2028()
func f2029()
func f2030()
func f2031()
func f2032()
func f2033()
func f2034()
func f2035()
func f2036()
func f2037()
func f2038()
func f2039()
func f2040()
func f2041()
func f2042()
func f2043()
func f2044()
func f2045()
func f2046()
func f2047()

var (
	icall_fn = []func(){f0,f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13,f14,f15,f16,f17,f18,f19,f20,f21,f22,f23,f24,f25,f26,f27,f28,f29,f30,f31,f32,f33,f34,f35,f36,f37,f38,f39,f40,f41,f42,f43,f44,f45,f46,f47,f48,f49,f50,f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f62,f63,f64,f65,f66,f67,f68,f69,f70,f71,f72,f73,f74,f75,f76,f77,f78,f79,f80,f81,f82,f83,f84,f85,f86,f87,f88,f89,f90,f91,f92,f93,f94,f95,f96,f97,f98,f99,f100,f101,f102,f103,f104,f105,f106,f107,f108,f109,f110,f111,f112,f113,f114,f115,f116,f117,f118,f119,f120,f121,f122,f123,f124,f125,f126,f127,f128,f129,f130,f131,f132,f133,f134,f135,f136,f137,f138,f139,f140,f141,f142,f143,f144,f145,f146,f147,f148,f149,f150,f151,f152,f153,f154,f155,f156,f157,f158,f159,f160,f161,f162,f163,f164,f165,f166,f167,f168,f169,f170,f171,f172,f173,f174,f175,f176,f177,f178,f179,f180,f181,f182,f183,f184,f185,f186,f187,f188,f189,f190,f191,f192,f193,f194,f195,f196,f197,f198,f199,f200,f201,f202,f203,f204,f205,f206,f207,f208,f209,f210,f211,f212,f213,f214,f215,f216,f217,f218,f219,f220,f221,f222,f223,f224,f225,f226,f227,f228,f229,f230,f231,f232,f233,f234,f235,f236,f237,f238,f239,f240,f241,f242,f243,f244,f245,f246,f247,f248,f249,f250,f251,f252,f253,f254,f255,f256,f257,f258,f259,f260,f261,f262,f263,f264,f265,f266,f267,f268,f269,f270,f271,f272,f273,f274,f275,f276,f277,f278,f279,f280,f281,f282,f283,f284,f285,f286,f287,f288,f289,f290,f291,f292,f293,f294,f295,f296,f297,f298,f299,f300,f301,f302,f303,f304,f305,f306,f307,f308,f309,f310,f311,f312,f313,f314,f315,f316,f317,f318,f319,f320,f321,f322,f323,f324,f325,f326,f327,f328,f329,f330,f331,f332,f333,f334,f335,f336,f337,f338,f339,f340,f341,f342,f343,f344,f345,f346,f347,f348,f349,f350,f351,f352,f353,f354,f355,f356,f357,f358,f359,f360,f361,f362,f363,f364,f365,f366,f367,f368,f369,f370,f371,f372,f373,f374,f375,f376,f377,f378,f379,f380,f381,f382,f383,f384,f385,f386,f387,f388,f389,f390,f391,f392,f393,f394,f395,f396,f397,f398,f399,f400,f401,f402,f403,f404,f405,f406,f407,f408,f409,f410,f411,f412,f413,f414,f415,f416,f417,f418,f419,f420,f421,f422,f423,f424,f425,f426,f427,f428,f429,f430,f431,f432,f433,f434,f435,f436,f437,f438,f439,f440,f441,f442,f443,f444,f445,f446,f447,f448,f449,f450,f451,f452,f453,f454,f455,f456,f457,f458,f459,f460,f461,f462,f463,f464,f465,f466,f467,f468,f469,f470,f471,f472,f473,f474,f475,f476,f477,f478,f479,f480,f481,f482,f483,f484,f485,f486,f487,f488,f489,f490,f491,f492,f493,f494,f495,f496,f497,f498,f499,f500,f501,f502,f503,f504,f505,f506,f507,f508,f509,f510,f511,f512,f513,f514,f515,f516,f517,f518,f519,f520,f521,f522,f523,f524,f525,f526,f527,f528,f529,f530,f531,f532,f533,f534,f535,f536,f537,f538,f539,f540,f541,f542,f543,f544,f545,f546,f547,f548,f549,f550,f551,f552,f553,f554,f555,f556,f557,f558,f559,f560,f561,f562,f563,f564,f565,f566,f567,f568,f569,f570,f571,f572,f573,f574,f575,f576,f577,f578,f579,f580,f581,f582,f583,f584,f585,f586,f587,f588,f589,f590,f591,f592,f593,f594,f595,f596,f597,f598,f599,f600,f601,f602,f603,f604,f605,f606,f607,f608,f609,f610,f611,f612,f613,f614,f615,f616,f617,f618,f619,f620,f621,f622,f623,f624,f625,f626,f627,f628,f629,f630,f631,f632,f633,f634,f635,f636,f637,f638,f639,f640,f641,f642,f643,f644,f645,f646,f647,f648,f649,f650,f651,f652,f653,f654,f655,f656,f657,f658,f659,f660,f661,f662,f663,f664,f665,f666,f667,f668,f669,f670,f671,f672,f673,f674,f675,f676,f677,f678,f679,f680,f681,f682,f683,f684,f685,f686,f687,f688,f689,f690,f691,f692,f693,f694,f695,f696,f697,f698,f699,f700,f701,f702,f703,f704,f705,f706,f707,f708,f709,f710,f711,f712,f713,f714,f715,f716,f717,f718,f719,f720,f721,f722,f723,f724,f725,f726,f727,f728,f729,f730,f731,f732,f733,f734,f735,f736,f737,f738,f739,f740,f741,f742,f743,f744,f745,f746,f747,f748,f749,f750,f751,f752,f753,f754,f755,f756,f757,f758,f759,f760,f761,f762,f763,f764,f765,f766,f767,f768,f769,f770,f771,f772,f773,f774,f775,f776,f777,f778,f779,f780,f781,f782,f783,f784,f785,f786,f787,f788,f789,f790,f791,f792,f793,f794,f795,f796,f797,f798,f799,f800,f801,f802,f803,f804,f805,f806,f807,f808,f809,f810,f811,f812,f813,f814,f815,f816,f817,f818,f819,f820,f821,f822,f823,f824,f825,f826,f827,f828,f829,f830,f831,f832,f833,f834,f835,f836,f837,f838,f839,f840,f841,f842,f843,f844,f845,f846,f847,f848,f849,f850,f851,f852,f853,f854,f855,f856,f857,f858,f859,f860,f861,f862,f863,f864,f865,f866,f867,f868,f869,f870,f871,f872,f873,f874,f875,f876,f877,f878,f879,f880,f881,f882,f883,f884,f885,f886,f887,f888,f889,f890,f891,f892,f893,f894,f895,f896,f897,f898,f899,f900,f901,f902,f903,f904,f905,f906,f907,f908,f909,f910,f911,f912,f913,f914,f915,f916,f917,f918,f919,f920,f921,f922,f923,f924,f925,f926,f927,f928,f929,f930,f931,f932,f933,f934,f935,f936,f937,f938,f939,f940,f941,f942,f943,f944,f945,f946,f947,f948,f949,f950,f951,f952,f953,f954,f955,f956,f957,f958,f959,f960,f961,f962,f963,f964,f965,f966,f967,f968,f969,f970,f971,f972,f973,f974,f975,f976,f977,f978,f979,f980,f981,f982,f983,f984,f985,f986,f987,f988,f989,f990,f991,f992,f993,f994,f995,f996,f997,f998,f999,f1000,f1001,f1002,f1003,f1004,f1005,f1006,f1007,f1008,f1009,f1010,f1011,f1012,f1013,f1014,f1015,f1016,f1017,f1018,f1019,f1020,f1021,f1022,f1023,f1024,f1025,f1026,f1027,f1028,f1029,f1030,f1031,f1032,f1033,f1034,f1035,f1036,f1037,f1038,f1039,f1040,f1041,f1042,f1043,f1044,f1045,f1046,f1047,f1048,f1049,f1050,f1051,f1052,f1053,f1054,f1055,f1056,f1057,f1058,f1059,f1060,f1061,f1062,f1063,f1064,f1065,f1066,f1067,f1068,f1069,f1070,f1071,f1072,f1073,f1074,f1075,f1076,f1077,f1078,f1079,f1080,f1081,f1082,f1083,f1084,f1085,f1086,f1087,f1088,f1089,f1090,f1091,f1092,f1093,f1094,f1095,f1096,f1097,f1098,f1099,f1100,f1101,f1102,f1103,f1104,f1105,f1106,f1107,f1108,f1109,f1110,f1111,f1112,f1113,f1114,f1115,f1116,f1117,f1118,f1119,f1120,f1121,f1122,f1123,f1124,f1125,f1126,f1127,f1128,f1129,f1130,f1131,f1132,f1133,f1134,f1135,f1136,f1137,f1138,f1139,f1140,f1141,f1142,f1143,f1144,f1145,f1146,f1147,f1148,f1149,f1150,f1151,f1152,f1153,f1154,f1155,f1156,f1157,f1158,f1159,f1160,f1161,f1162,f1163,f1164,f1165,f1166,f1167,f1168,f1169,f1170,f1171,f1172,f1173,f1174,f1175,f1176,f1177,f1178,f1179,f1180,f1181,f1182,f1183,f1184,f1185,f1186,f1187,f1188,f1189,f1190,f1191,f1192,f1193,f1194,f1195,f1196,f1197,f1198,f1199,f1200,f1201,f1202,f1203,f1204,f1205,f1206,f1207,f1208,f1209,f1210,f1211,f1212,f1213,f1214,f1215,f1216,f1217,f1218,f1219,f1220,f1221,f1222,f1223,f1224,f1225,f1226,f1227,f1228,f1229,f1230,f1231,f1232,f1233,f1234,f1235,f1236,f1237,f1238,f1239,f1240,f1241,f1242,f1243,f1244,f1245,f1246,f1247,f1248,f1249,f1250,f1251,f1252,f1253,f1254,f1255,f1256,f1257,f1258,f1259,f1260,f1261,f1262,f1263,f1264,f1265,f1266,f1267,f1268,f1269,f1270,f1271,f1272,f1273,f1274,f1275,f1276,f1277,f1278,f1279,f1280,f1281,f1282,f1283,f1284,f1285,f1286,f1287,f1288,f1289,f1290,f1291,f1292,f1293,f1294,f1295,f1296,f1297,f1298,f1299,f1300,f1301,f1302,f1303,f1304,f1305,f1306,f1307,f1308,f1309,f1310,f1311,f1312,f1313,f1314,f1315,f1316,f1317,f1318,f1319,f1320,f1321,f1322,f1323,f1324,f1325,f1326,f1327,f1328,f1329,f1330,f1331,f1332,f1333,f1334,f1335,f1336,f1337,f1338,f1339,f1340,f1341,f1342,f1343,f1344,f1345,f1346,f1347,f1348,f1349,f1350,f1351,f1352,f1353,f1354,f1355,f1356,f1357,f1358,f1359,f1360,f1361,f1362,f1363,f1364,f1365,f1366,f1367,f1368,f1369,f1370,f1371,f1372,f1373,f1374,f1375,f1376,f1377,f1378,f1379,f1380,f1381,f1382,f1383,f1384,f1385,f1386,f1387,f1388,f1389,f1390,f1391,f1392,f1393,f1394,f1395,f1396,f1397,f1398,f1399,f1400,f1401,f1402,f1403,f1404,f1405,f1406,f1407,f1408,f1409,f1410,f1411,f1412,f1413,f1414,f1415,f1416,f1417,f1418,f1419,f1420,f1421,f1422,f1423,f1424,f1425,f1426,f1427,f1428,f1429,f1430,f1431,f1432,f1433,f1434,f1435,f1436,f1437,f1438,f1439,f1440,f1441,f1442,f1443,f1444,f1445,f1446,f1447,f1448,f1449,f1450,f1451,f1452,f1453,f1454,f1455,f1456,f1457,f1458,f1459,f1460,f1461,f1462,f1463,f1464,f1465,f1466,f1467,f1468,f1469,f1470,f1471,f1472,f1473,f1474,f1475,f1476,f1477,f1478,f1479,f1480,f1481,f1482,f1483,f1484,f1485,f1486,f1487,f1488,f1489,f1490,f1491,f1492,f1493,f1494,f1495,f1496,f1497,f1498,f1499,f1500,f1501,f1502,f1503,f1504,f1505,f1506,f1507,f1508,f1509,f1510,f1511,f1512,f1513,f1514,f1515,f1516,f1517,f1518,f1519,f1520,f1521,f1522,f1523,f1524,f1525,f1526,f1527,f1528,f1529,f1530,f1531,f1532,f1533,f1534,f1535,f1536,f1537,f1538,f1539,f1540,f1541,f1542,f1543,f1544,f1545,f1546,f1547,f1548,f1549,f1550,f1551,f1552,f1553,f1554,f1555,f1556,f1557,f1558,f1559,f1560,f1561,f1562,f1563,f1564,f1565,f1566,f1567,f1568,f1569,f1570,f1571,f1572,f1573,f1574,f1575,f1576,f1577,f1578,f1579,f1580,f1581,f1582,f1583,f1584,f1585,f1586,f1587,f1588,f1589,f1590,f1591,f1592,f1593,f1594,f1595,f1596,f1597,f1598,f1599,f1600,f1601,f1602,f1603,f1604,f1605,f1606,f1607,f1608,f1609,f1610,f1611,f1612,f1613,f1614,f1615,f1616,f1617,f1618,f1619,f1620,f1621,f1622,f1623,f1624,f1625,f1626,f1627,f1628,f1629,f1630,f1631,f1632,f1633,f1634,f1635,f1636,f1637,f1638,f1639,f1640,f1641,f1642,f1643,f1644,f1645,f1646,f1647,f1648,f1649,f1650,f1651,f1652,f1653,f1654,f1655,f1656,f1657,f1658,f1659,f1660,f1661,f1662,f1663,f1664,f1665,f1666,f1667,f1668,f1669,f1670,f1671,f1672,f1673,f1674,f1675,f1676,f1677,f1678,f1679,f1680,f1681,f1682,f1683,f1684,f1685,f1686,f1687,f1688,f1689,f1690,f1691,f1692,f1693,f1694,f1695,f1696,f1697,f1698,f1699,f1700,f1701,f1702,f1703,f1704,f1705,f1706,f1707,f1708,f1709,f1710,f1711,f1712,f1713,f1714,f1715,f1716,f1717,f1718,f1719,f1720,f1721,f1722,f1723,f1724,f1725,f1726,f1727,f1728,f1729,f1730,f1731,f1732,f1733,f1734,f1735,f1736,f1737,f1738,f1739,f1740,f1741,f1742,f1743,f1744,f1745,f1746,f1747,f1748,f1749,f1750,f1751,f1752,f1753,f1754,f1755,f1756,f1757,f1758,f1759,f1760,f1761,f1762,f1763,f1764,f1765,f1766,f1767,f1768,f1769,f1770,f1771,f1772,f1773,f1774,f1775,f1776,f1777,f1778,f1779,f1780,f1781,f1782,f1783,f1784,f1785,f1786,f1787,f1788,f1789,f1790,f1791,f1792,f1793,f1794,f1795,f1796,f1797,f1798,f1799,f1800,f1801,f1802,f1803,f1804,f1805,f1806,f1807,f1808,f1809,f1810,f1811,f1812,f1813,f1814,f1815,f1816,f1817,f1818,f1819,f1820,f1821,f1822,f1823,f1824,f1825,f1826,f1827,f1828,f1829,f1830,f1831,f1832,f1833,f1834,f1835,f1836,f1837,f1838,f1839,f1840,f1841,f1842,f1843,f1844,f1845,f1846,f1847,f1848,f1849,f1850,f1851,f1852,f1853,f1854,f1855,f1856,f1857,f1858,f1859,f1860,f1861,f1862,f1863,f1864,f1865,f1866,f1867,f1868,f1869,f1870,f1871,f1872,f1873,f1874,f1875,f1876,f1877,f1878,f1879,f1880,f1881,f1882,f1883,f1884,f1885,f1886,f1887,f1888,f1889,f1890,f1891,f1892,f1893,f1894,f1895,f1896,f1897,f1898,f1899,f1900,f1901,f1902,f1903,f1904,f1905,f1906,f1907,f1908,f1909,f1910,f1911,f1912,f1913,f1914,f1915,f1916,f1917,f1918,f1919,f1920,f1921,f1922,f1923,f1924,f1925,f1926,f1927,f1928,f1929,f1930,f1931,f1932,f1933,f1934,f1935,f1936,f1937,f1938,f1939,f1940,f1941,f1942,f1943,f1944,f1945,f1946,f1947,f1948,f1949,f1950,f1951,f1952,f1953,f1954,f1955,f1956,f1957,f1958,f1959,f1960,f1961,f1962,f1963,f1964,f1965,f1966,f1967,f1968,f1969,f1970,f1971,f1972,f1973,f1974,f1975,f1976,f1977,f1978,f1979,f1980,f1981,f1982,f1983,f1984,f1985,f1986,f1987,f1988,f1989,f1990,f1991,f1992,f1993,f1994,f1995,f1996,f1997,f1998,f1999,f2000,f2001,f2002,f2003,f2004,f2005,f2006,f2007,f2008,f2009,f2010,f2011,f2012,f2013,f2014,f2015,f2016,f2017,f2018,f2019,f2020,f2021,f2022,f2023,f2024,f2025,f2026,f2027,f2028,f2029,f2030,f2031,f2032,f2033,f2034,f2035,f2036,f2037,f2038,f2039,f2040,f2041,f2042,f2043,f2044,f2045,f2046,f2047}
)
