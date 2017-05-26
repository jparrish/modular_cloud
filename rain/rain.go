package rain

import (
	"github.com/pearish/modular_cloud/common"
)

type RainCloud interface {
	common.Cloud
}

type rainImpl struct {
	common.CloudImpl
	sound string
}

func (cloud *rainImpl) my_c() string {
	return cloud.sound
}

func NewRainCloud(i, j string) RainCloud {
	cloud := &rainImpl{
		CloudImpl: common.CloudImpl{
			I: i,
			J: j,
		},
		sound: "pitter, patter, ....",
	}
	// This overrides C() in all cases: both direct calls and indirect
	// calls through the base class will be correctly overridden.
	cloud.CFunc = cloud.my_c
	return cloud
}

