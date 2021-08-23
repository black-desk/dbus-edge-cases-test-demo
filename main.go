package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type STRUCT struct {
	Field1, Field2, Field3, Field4, Field5, Field6, Field7, Field8, Field9, Field10, Field11, Field12, Field13, Field14, Field15, Field16, Field17, Field18, Field19, Field20, Field21, Field22, Field23, Field24, Field25, Field26, Field27, Field28, Field29, Field30, Field31, Field32, Field33, Field34, Field35, Field36, Field37, Field38, Field39, Field40, Field41, Field42, Field43, Field44, Field45, Field46, Field47, Field48, Field49, Field50, Field51, Field52, Field53, Field54, Field55, Field56, Field57, Field58, Field59, Field60, Field61, Field62, Field63, Field64, Field65, Field66, Field67, Field68, Field69, Field70, Field71, Field72, Field73, Field74, Field75, Field76, Field77, Field78, Field79, Field80, Field81, Field82, Field83, Field84, Field85, Field86, Field87, Field88, Field89, Field90, Field91, Field92, Field93, Field94, Field95, Field96, Field97, Field98, Field99, Field100, Field101, Field102, Field103, Field104, Field105, Field106, Field107, Field108, Field109, Field110, Field111, Field112, Field113, Field114, Field115, Field116, Field117, Field118, Field119, Field120, Field121, Field122, Field123, Field124, Field125, Field126, Field127, Field128, Field129, Field130, Field131, Field132, Field133, Field134, Field135, Field136, Field137, Field138, Field139, Field140, Field141, Field142, Field143, Field144, Field145, Field146, Field147, Field148, Field149, Field150, Field151, Field152, Field153, Field154, Field155, Field156, Field157, Field158, Field159, Field160, Field161, Field162, Field163, Field164, Field165, Field166, Field167, Field168, Field169, Field170, Field171, Field172, Field173, Field174, Field175, Field176, Field177, Field178, Field179, Field180, Field181, Field182, Field183, Field184, Field185, Field186, Field187, Field188, Field189, Field190, Field191, Field192, Field193, Field194, Field195, Field196, Field197, Field198, Field199, Field200, Field201, Field202, Field203, Field204, Field205, Field206, Field207, Field208, Field209, Field210, Field211, Field212, Field213, Field214, Field215, Field216, Field217, Field218, Field219, Field220, Field221, Field222, Field223, Field224, Field225, Field226, Field227, Field228, Field229, Field230, Field231, Field232, Field233, Field234, Field235, Field236, Field237, Field238, Field239, Field240, Field241, Field242, Field243, Field244, Field245, Field246, Field247, Field248, Field249, Field250, Field251, Field252, Field253, Field254 int
}

type Struct1 struct{ Field int }
type Struct2 struct{ Field Struct1 }
type Struct3 struct{ Field Struct2 }
type Struct4 struct{ Field Struct3 }
type Struct5 struct{ Field Struct4 }
type Struct6 struct{ Field Struct5 }
type Struct7 struct{ Field Struct6 }
type Struct8 struct{ Field Struct7 }
type Struct9 struct{ Field Struct8 }
type Struct10 struct{ Field Struct9 }
type Struct11 struct{ Field Struct10 }
type Struct12 struct{ Field Struct11 }
type Struct13 struct{ Field Struct12 }
type Struct14 struct{ Field Struct13 }
type Struct15 struct{ Field Struct14 }
type Struct16 struct{ Field Struct15 }
type Struct17 struct{ Field Struct16 }
type Struct18 struct{ Field Struct17 }
type Struct19 struct{ Field Struct18 }
type Struct20 struct{ Field Struct19 }
type Struct21 struct{ Field Struct20 }
type Struct22 struct{ Field Struct21 }
type Struct23 struct{ Field Struct22 }
type Struct24 struct{ Field Struct23 }
type Struct25 struct{ Field Struct24 }
type Struct26 struct{ Field Struct25 }
type Struct27 struct{ Field Struct26 }
type Struct28 struct{ Field Struct27 }
type Struct29 struct{ Field Struct28 }
type Struct30 struct{ Field Struct29 }
type Struct31 struct{ Field Struct30 }
type Struct32 struct{ Field Struct31 }
type Struct33 struct{ Field Struct32 }

type map1 map[string]string
type map2 map[string]map1
type map3 map[string]map2
type map4 map[string]map3
type map5 map[string]map4
type map6 map[string]map5
type map7 map[string]map6
type map8 map[string]map7
type map9 map[string]map8
type map10 map[string]map9
type map11 map[string]map10
type map12 map[string]map11
type map13 map[string]map12
type map14 map[string]map13
type map15 map[string]map14
type map16 map[string]map15
type map17 map[string]map16
type map18 map[string]map17
type map19 map[string]map18
type map20 map[string]map19
type map21 map[string]map20
type map22 map[string]map21
type map23 map[string]map22
type map24 map[string]map23
type map25 map[string]map24
type map26 map[string]map25
type map27 map[string]map26
type map28 map[string]map27
type map29 map[string]map28
type map30 map[string]map29
type map31 map[string]map30
type map32 map[string]map31
type map33 map[string]map32

var (
	Map1  = map1{"": ""}
	Map2  = map2{"": Map1}
	Map3  = map3{"": Map2}
	Map4  = map4{"": Map3}
	Map5  = map5{"": Map4}
	Map6  = map6{"": Map5}
	Map7  = map7{"": Map6}
	Map8  = map8{"": Map7}
	Map9  = map9{"": Map8}
	Map10 = map10{"": Map9}
	Map11 = map11{"": Map10}
	Map12 = map12{"": Map11}
	Map13 = map13{"": Map12}
	Map14 = map14{"": Map13}
	Map15 = map15{"": Map14}
	Map16 = map16{"": Map15}
	Map17 = map17{"": Map16}
	Map18 = map18{"": Map17}
	Map19 = map19{"": Map18}
	Map20 = map20{"": Map19}
	Map21 = map21{"": Map20}
	Map22 = map22{"": Map21}
	Map23 = map23{"": Map22}
	Map24 = map24{"": Map23}
	Map25 = map25{"": Map24}
	Map26 = map26{"": Map25}
	Map27 = map27{"": Map26}
	Map28 = map28{"": Map27}
	Map29 = map29{"": Map28}
	Map30 = map30{"": Map29}
	Map31 = map31{"": Map30}
	Map32 = map32{"": Map31}
	Map33 = map33{"": Map32}
)

func dbusMakeVariant64(v interface{}) dbus.Variant {
	return dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(dbus.MakeVariant(
		v,
	)))))))), // 1*8
	)))))))), // 2*8
	)))))))), // 3*8
	)))))))), // 4*8
	)))))))), // 5*8
	)))))))), // 6*8
	)))))))), // 7*8
	)))))))) //   64

}

type foo struct{}

func (f foo) Func1() (string, *dbus.Error) {
	println("Func1 called")
	defer println("invalid string, lose connection")
	enc := simplifiedchinese.GB18030.NewEncoder()
	ret, _ := enc.String("一些汉字") // ret is a STRING in GB18030
	return ret, nil
}

func (f foo) Func2() (Struct32, *dbus.Error) {
	println("Func2 called")
	defer println("should work fine")
	return Struct32{}, nil
}

func (f foo) Func3() (struct{}, *dbus.Error) {
	println("Func3 called")
	defer println("empty struct, lose connection")
	return struct{}{}, nil
}

func (f foo) Func4() (dbus.Variant, *dbus.Error) {
	println("Func4 called")
	defer println("should work fine")
	return dbus.MakeVariant(Struct32{}), nil
}

func (f foo) Func5() (Struct33, *dbus.Error) {
	println("Func5 called")
	defer println("invalid signature (exceeded maximum struct recursion), lose connection")
	return Struct33{}, nil
}

func (f foo) Func6() (dbus.Variant, *dbus.Error) {
	println("Func6 called")
	defer println("should work fine")
	a := dbusMakeVariant64(1)
	return a, nil
}

func (f foo) Func7() (dbus.Variant, *dbus.Error) {
	println("Func7 called")
	defer println("this is a valid message, but will cause this demo lose its connection")
	return dbusMakeVariant64(map32{}), nil
}

func (f foo) Func8() (dbus.Variant, *dbus.Error) {
	println("Func8 called")
	defer println("message exceeds depth limitation, lose connection")
	return dbusMakeVariant64(Map1), nil
}

func (f foo) Func9() (dbus.Variant, *dbus.Error) {
	println("Func9 called")
	defer println("message exceeds depth limitation, lose connection")
	return dbus.MakeVariant(dbusMakeVariant64(1)), nil
}

func (f foo) Func10() (map32, *dbus.Error) {
	println("Func10 called")
	defer println("should work fine")
	return Map32, nil
}

func (f foo) Func11() ([]map31, *dbus.Error) {
	println("Foo11 called")
	defer println("should work fine")
	return []map31{Map31}, nil
}

func (f foo) Func12() ([]map32, *dbus.Error) {
	println("Func12 called")
	defer println("invalid signature (exceeded maximum array recursion), lose connection")
	return []map32{}, nil
}

func (f foo) Func13() (STRUCT, *dbus.Error) {
	println("Func13 called")
	defer println("invalid signature (longer than 255), lose connection")
	return STRUCT{}, nil
}

func main() {
	println("before my patch , if a error occur during encoding, it will lose its dbus connection.")
	println("type command like `qdbus --literal com.github.blackdesk.Demo /com/github/blackdesk/Demo com.github.blackdesk.Demo.Func3` to test this demo")

	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	f := foo{}
	conn.Export(f, "/com/github/blackdesk/Demo", "com.github.blackdesk.Demo")
	reply, err := conn.RequestName("com.github.blackdesk.Demo",
		dbus.NameFlagDoNotQueue)
	if err != nil {
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		fmt.Fprintln(os.Stderr, "name already taken")
		os.Exit(1)
	}
	fmt.Println("start listening on com.github.blackdesk.Demo /com/github/blackdesk/Demo com.github.blackdesk.Demo")
	select {}
}
