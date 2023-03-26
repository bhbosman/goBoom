package test

import (
	"fmt"
	"github.com/bhbosman/goBoom/model"
	"io"
)

type Model00004Data struct {
	UserDefinition UserDefinition
	UserAddress    UserAddress
}

func intValue(a, b, c, d int) int {
	return a + b + c + d
}
func newUserDefinitionForTest00004() model.Seq[UserDefinitionView] {
	dataModel := Model00004Data{
		UserDefinition: UserDefinition{},
		UserAddress: model.InnerRelationship[UserDefinition, UserAddress](
			UserAddress{},
			func(from UserDefinition, to UserAddress) bool {
				return from.Id == to.Id && from.Id*4 == 400
			},
		),
	}
	seq := model.Query[Model00004Data](dataModel)
	result02 := model.Filter[Model00004Data](seq, func(v Model00004Data) bool {
		return v.UserDefinition.Surname == "Bosman"
	})
	result03 := model.Map[Model00004Data, UserDefinitionView](result02, func(input Model00004Data) UserDefinitionView {
		WA := true
		switch WA {
		case true:
			break
		case false:
			break
		}
		WB := "sadsas"
		switch WB {
		case "true":
			WB = "adsd"
		default:
			WB = "sfsssssssd"
		}
		// sfsdfdsjh

		var ob interface{} = &UserAddress{}

		switch v := ob.(type) {
		case io.ReadCloser:
			println(v)
		case io.Closer:
			println(v)
		default:
			println(v)
		}

		m := make(map[string]int)
		ss := 2333
		ss++
		ss--
		m["a"] = 23
		m["b"] = 233
		m["d"] = ss

		if a, ok := m["ddd"]; ok {
			println(a)
		}

		combined := "A" + "B"
		combined = combined + input.UserDefinition.Name
		combined = fmt.Sprintf("adsaas%v", combined)
		for i, i2 := range combined {
			fmt.Println(i, i2)
		}
		for true {
			break
		}

		dd := make([]int, 10)
		dd[0] = 123 + 44
		dd[0] += 23
		dd[1] = intValue(1, 2, 3, 4)

		for i := 0; i < 100; i++ {
			println(i)
		}

		id := func() int {
			if combined != "DDDD" {
				return 56 * 3334444 / 10 * 23 / 2
			}
			return 33
		}()

		return UserDefinitionView{
			Id:         id,
			Name:       combined,
			Surname:    input.UserDefinition.Surname,
			Line01:     input.UserAddress.Line01,
			Line02:     input.UserAddress.Line02,
			Line03:     input.UserAddress.Line03,
			Line04:     input.UserAddress.Line04,
			Line05:     input.UserAddress.Line05,
			Province:   input.UserAddress.Province,
			Country:    input.UserAddress.Country,
			PostalCode: input.UserAddress.PostalCode,
		}
	})
	return result03
}

func init() {
	model.RegisterSelect[UserDefinitionView](
		func(data model.Seq[UserDefinitionView]) {
			newUserDefinitionForTest00004()
		},
	)
}
