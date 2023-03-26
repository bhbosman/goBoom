package internal

import "github.com/bhbosman/goBoom/model"

type Model01Data struct {
	UserDefinition UserDefinition
	UserAddress    UserAddress
}

func dd() model.Seq[UserDefinitionView] {
	data := Model01Data{
		UserDefinition: UserDefinition{},
		UserAddress: model.InnerRelationship[UserDefinition, UserAddress](
			UserAddress{},
			func(from UserDefinition, to UserAddress) bool {
				return from.Id == to.Id && from.Id*4 == 400
			},
		),
	}
	seq := model.Query[Model01Data](data)
	result02 := model.Filter[Model01Data](
		seq,
		func(v Model01Data) bool {
			return v.UserDefinition.Surname == "Bosman"
		},
	)
	result03 := model.Map[Model01Data, UserDefinitionView](
		result02,
		func(input Model01Data) UserDefinitionView {
			return UserDefinitionView{
				Id:         input.UserDefinition.Id,
				Name:       input.UserDefinition.Name,
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
		},
	)

	return result03
}

func init() {
	model.BoomSelect[UserDefinitionView](dd())
}
