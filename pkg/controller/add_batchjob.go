package controller

import (
	"github.com/wangpy1489/DNative/pkg/controller/batchjob"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, batchjob.Add)
}
