package thunder

import (
	"github.com/pearish/modular_cloud/common"
)

type ThunderCloud interface {
	common.Cloud
}

type thunderImpl struct {
	common.CloudImpl
	sound string
}

// This will only override B() in the "base class", i.e. Cloud
// for direct calls. Calls to B() from the base class will bypass the override.
func (cloud *thunderImpl) B() string {
	return "stronger than a common cloud"
}

func (cloud *thunderImpl) my_c() string {
	return cloud.sound
}

func NewThunderCloud(i, j string) ThunderCloud {
	cloud := &thunderImpl{
		CloudImpl: common.CloudImpl{
			I: i,
			J: j,
		},
		sound: "CRRRRAAAKKK!",
	}
	// This overrides C() in all cases: both direct calls and indirect
	// calls through the base class will be correctly overridden.
	cloud.CFunc = cloud.my_c
	return cloud
}
