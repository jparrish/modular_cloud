package common

import (
	"fmt"
)

type Cloud interface {
	A() string
	B() string
	C() string
	D() string
	E() string
}

// If you need to create an abstract base class with
// default implementation, you have do it manually in Go.
// Function pointers are used to enable "derived classes"
// to correctly override default implementations.
// All methods bound to the CloudImpl receiver become the
// default implementation. If they need to be overridden,
// they will need a corresponding func entry in the struct
// and be initialized in the constructor.

type CloudImpl struct {
	I     string
	J     string
	CFunc func() string
}

func (cloud *CloudImpl) A() string {
	return cloud.I
}

func (cloud *CloudImpl) B() string {
	return cloud.J
}

func (cloud *CloudImpl) C() string {
	return cloud.CFunc()
}

func (cloud *CloudImpl) D() string {
	return cloud.C()
}

func (cloud *CloudImpl) E() string {
	return cloud.B()
}

func (cloud *CloudImpl) default_c() string {
	return fmt.Sprintf("%s & %s", cloud.I, cloud.J)
}

func NewCloud(i, j string) Cloud {
	cloud := &CloudImpl{I: i, J: j}
	cloud.CFunc = cloud.default_c
	return cloud
}

