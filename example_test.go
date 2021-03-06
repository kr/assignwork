package assignwork_test

import (
	"github.com/kr/assignwork"
)

const (
	NDyno          = 5
	Dyno           = 2
	RecheckWorkers = 1
)

var self = assignwork.NewMember(Dyno, NDyno)

func own(s string) bool {
	return self.Owns(s)
}

func shouldRecheck(s, t string) bool {
	excl := self.Pool.Owners(s)[0]
	owners := self.Pool.OwnersExcluding(s+t, excl)
	return self.In(owners[:RecheckWorkers])
}
