package adapters

import (
	"github.com/chiefy/gometal/gometal"
)

func Drafts() gometal.Adapter {
	var mw gometal.Adapter
	mw = func(gm *gometal.GoMetal) error {
		return nil
	}
	return mw
}
