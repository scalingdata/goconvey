package executor

import "github.com/scalingdata/goconvey/web/server/contract"

type Parser interface {
	Parse([]*contract.Package)
}

type Tester interface {
	SetBatchSize(batchSize int)
	TestAll(folders []*contract.Package)
}
