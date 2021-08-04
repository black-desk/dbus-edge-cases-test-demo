package main

import (
	"fmt"
	"os"

	"github.com/godbus/dbus/v5"
	"golang.org/x/text/encoding/simplifiedchinese"
)

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
	defer println("Func1 should fail while encode: invalid string")
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
	defer println("empty struct, should panic while generate signature")
	return struct{}{}, nil
}

func (f foo) Func4() (dbus.Variant, *dbus.Error) {
	println("Func4 called")
	defer println("should work fine")
	return dbus.MakeVariant(Struct32{}), nil
}

func (f foo) Func5() (Struct33, *dbus.Error) {
	println("Func5 called")
	defer println("should panic invalid signature (exceeded maximum struct recursion)")
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
	defer println("this message has a empty array, which will not be check depth, so this is a valid message")
	return dbusMakeVariant64(map32{}), nil
}

func (f foo) Func8() (dbus.Variant, *dbus.Error) {
	println("Func8 called")
	defer println("message exceeds depth limitation")
	return dbusMakeVariant64(Map1), nil
}

func (f foo) Func9() (dbus.Variant, *dbus.Error) {
	println("Func9 called")
	defer println("message exceeds depth limitation")
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
	defer println("should panic invalid signature (exceeded maximum array recursion)")
	return []map32{}, nil
}

func main() {

	println("when error occur during encode, my modified godbus will send this error to dbus as the method_return's dbus.Error.")
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
