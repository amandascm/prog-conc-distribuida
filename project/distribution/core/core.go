package core

import (
	"test/project/distribution/marshaller"
	"test/project/distribution/requestor"
	"test/project/infrastructure/crh"
)

type Core struct {
	_Requestor  requestor.Requestor
	_Crh        crh.CRH
	_Marshaller marshaller.Marshaller
}

func NewCore() Core {
	r := Core{}
	/*	r._Requestor = NewRequestor()
		r._Crh = crh.NewCRH()
		r._Marshaller = marshaller.Marshaller{}
	*/
	return r
}
