package config_test

import (
	"testing"
	"time"

	"github.com/go-trellis/config"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	jsonFile  = "example.json"
	yamlFile  = "example.yml"
	wrongFile = "wrong_file"
)

func TestNewConfig(t *testing.T) {
	Convey("get config", t, func() {
		Convey("failed get config", func() {
			c, err := config.NewConfig(wrongFile)
			So(err, ShouldBeError)
			So(c, ShouldBeNil)
			c, err = config.NewConfig("")
			So(err, ShouldBeError)
			So(c, ShouldBeNil)

			err = config.ReadJSONFile(wrongFile, nil)
			So(err, ShouldBeError)
		})

		Convey("json config", func() {
			c, err := config.NewConfig(jsonFile)
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)

			Convey("json checker", func() {
				So(c.GetString("a"), ShouldEqual, "Easy!")
				So(c.GetMap("a"), ShouldBeNil)
				So(c.GetInterface("b.c.cn.a"), ShouldEqual, "test")
				c.SetKeyValue("b.c.cn.a", "value")

				newC := c.Copy()
				newC.SetKeyValue("b.c.cn.a", "joking")
				So(newC.GetInterface("b.c.cn.a"), ShouldEqual, "joking")

				So(c.GetInterface("b.c.cn.a"), ShouldEqual, "value")
				So(c.GetInterface("n.a"), ShouldEqual, "value")

				intList := c.GetIntList("b.c.cbd")
				So(intList[0], ShouldEqual, 3)
				So(intList[1], ShouldEqual, 4)
				So(c.GetString("a.b.c", "example"), ShouldEqual, "example")
				So(c.GetString("b.c.e"), ShouldEqual, "Just Do it")
				So(c.GetFloat("h"), ShouldEqual, 1.01)
				So(c.GetInt("b.c.f", 100), ShouldEqual, 2)
				So(c.GetInt("b.c.e"), ShouldEqual, 0)
				So(c.GetBoolean("b.c.g"), ShouldBeTrue)
				So(c.GetBoolean("b.c.x", true), ShouldBeTrue)
				So(c.GetConfig("b"), ShouldNotBeNil)
				faceList := c.GetList("b.d")
				So(faceList[0], ShouldEqual, "3")
				So(faceList[1], ShouldEqual, "4")
				So(c.GetTimeDuration("b.c.t"), ShouldEqual, time.Hour*24)

				c.SetKeyValue("a.b.c", "Correct")
				c.SetKeyValue("b.c.e", "Correct")
				c.SetKeyValue("b.c.d", "d")
				c.SetKeyValue("b.c.g", false)
				c.SetKeyValue("b.d", []int{1, 2, 3, 4})

				So(c.GetString("a", "example"), ShouldEqual, "example")
				So(c.GetInterface("a", "example"), ShouldNotBeNil)
				So(c.GetString("a.b.c", "example"), ShouldEqual, "Correct")
				So(c.GetString("b.c.e", "example"), ShouldEqual, "Correct")
				So(c.GetBoolean("b.c.g", true), ShouldBeFalse)
				So(c.GetString("b.c.d", "example"), ShouldEqual, "d")

				c.SetKeyValue("a", "Difficult!")
				c.SetKeyValue("h.a", []bool{false, true, false})
				c.SetKeyValue("h.f", []float64{1.2, 2.3, 3.4})
				c.SetKeyValue("h.b", "10T")

				So(c.GetString("a.b.c", "example"), ShouldEqual, "example")
				So(c.GetString("a", "example"), ShouldEqual, "Difficult!")
				So(c.GetList("a"), ShouldBeNil)
				So(c.GetBooleanList("h.a"), ShouldNotBeNil)
				So(c.GetFloatList("h.f"), ShouldNotBeNil)
				So(c.GetFloat("h"), ShouldEqual, 0)

				hb := c.GetByteSize("h.b")
				So(hb.Int64(), ShouldEqual, 10995116277760)
				So(c.GetString("b.d", "example"), ShouldEqual, "example")
				So(c.GetList("b.d"), ShouldNotBeNil)
				So(c.GetStringList("b.d"), ShouldBeNil)
				So(c.GetIntList("b.d"), ShouldNotBeNil)

				c.SetKeyValue("b.d", []string{"1", "2", "3"})

				stringList := c.GetStringList("b.d")
				So(stringList[0], ShouldEqual, "1")
				So(stringList[1], ShouldEqual, "2")
				So(stringList[2], ShouldEqual, "3")

				So(c.GetKeys(), ShouldNotBeNil)
				bs, _ := c.Dump()
				So(bs, ShouldNotBeNil)
			})
		})

		Convey("yaml config", func() {
			c, err := config.NewConfig(yamlFile)
			So(err, ShouldBeNil)
			So(c, ShouldNotBeNil)

			Convey("yaml checker", func() {
				So(c.GetString("a"), ShouldEqual, "Easy!")
				So(c.GetMap("a"), ShouldBeNil)
				So(c.GetInterface("b.c.cn.a"), ShouldEqual, "test")
				c.SetKeyValue("b.c.cn.a", "value")

				newC := c.Copy()
				newC.SetKeyValue("b.c.cn.a", "joking")
				So(newC.GetInterface("b.c.cn.a"), ShouldEqual, "joking")

				So(c.GetInterface("b.c.cn.a"), ShouldEqual, "value")
				So(c.GetInterface("n.a"), ShouldEqual, "value")

				intList := c.GetIntList("b.c.cbd")
				So(intList[0], ShouldEqual, 3)
				So(intList[1], ShouldEqual, 4)
				So(c.GetString("a.b.c", "example"), ShouldEqual, "example")
				So(c.GetString("b.c.e"), ShouldEqual, "Just Do it")
				So(c.GetFloat("h"), ShouldEqual, 1.01)
				So(c.GetInt("b.c.f", 100), ShouldEqual, 2)
				So(c.GetInt("b.c.e"), ShouldEqual, 0)
				So(c.GetBoolean("b.c.g"), ShouldBeTrue)
				So(c.GetBoolean("b.c.x", true), ShouldBeTrue)
				So(c.GetConfig("b"), ShouldNotBeNil)
				faceList := c.GetList("b.d")
				So(faceList[0], ShouldEqual, 3)
				So(faceList[1], ShouldEqual, 4)
				So(c.GetTimeDuration("b.c.t"), ShouldEqual, time.Hour*24)

				c.SetKeyValue("a.b.c", "Correct")
				c.SetKeyValue("b.c.e", "Correct")
				c.SetKeyValue("b.c.d", "d")
				c.SetKeyValue("b.c.g", false)
				c.SetKeyValue("b.d", []int{1, 2, 3, 4})

				So(c.GetString("a", "example"), ShouldEqual, "example")
				So(c.GetInterface("a", "example"), ShouldNotBeNil)
				So(c.GetString("a.b.c", "example"), ShouldEqual, "Correct")
				So(c.GetString("b.c.e", "example"), ShouldEqual, "Correct")
				So(c.GetBoolean("b.c.g", true), ShouldBeFalse)
				So(c.GetString("b.c.d", "example"), ShouldEqual, "d")

				c.SetKeyValue("a", "Difficult!")
				c.SetKeyValue("h.a", []bool{false, true, false})
				c.SetKeyValue("h.f", []float64{1.2, 2.3, 3.4})
				c.SetKeyValue("h.b", "10T")

				So(c.GetString("a.b.c", "example"), ShouldEqual, "example")
				So(c.GetString("a", "example"), ShouldEqual, "Difficult!")
				So(c.GetList("a"), ShouldBeNil)
				So(c.GetBooleanList("h.a"), ShouldNotBeNil)
				So(c.GetFloatList("h.f"), ShouldNotBeNil)
				So(c.GetFloat("h"), ShouldEqual, 0)

				hb := c.GetByteSize("h.b")
				So(hb.Int64(), ShouldEqual, 10995116277760)
				So(c.GetString("b.d", "example"), ShouldEqual, "example")
				So(c.GetList("b.d"), ShouldNotBeNil)
				So(c.GetStringList("b.d"), ShouldBeNil)
				So(c.GetIntList("b.d"), ShouldNotBeNil)

				c.SetKeyValue("b.d", []string{"1", "2", "3"})

				stringList := c.GetStringList("b.d")
				So(stringList[0], ShouldEqual, "1")
				So(stringList[1], ShouldEqual, "2")
				So(stringList[2], ShouldEqual, "3")

				So(c.GetKeys(), ShouldNotBeNil)
				bs, _ := c.Dump()
				So(bs, ShouldNotBeNil)
			})
		})
	})
}
