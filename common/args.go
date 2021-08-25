package common

import (
	"github.com/alexflint/go-arg"
)

func Args() * arg.Parser  {

	return arg.MustParse(GetStruct())


}