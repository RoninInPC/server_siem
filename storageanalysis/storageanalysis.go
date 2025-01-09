package storageanalysis

import (
	"server_siem/entity/subject/notification/receivernotification"
)

type AddAnalysis receivernotification.Notification

type ArrAddAnalysis []AddAnalysis

func (a ArrAddAnalysis) Len() int {
	return len(a)
}

func (a ArrAddAnalysis) Less(i, j int) bool {
	return a[i].GetTime().Before(a[j].GetTime())
}

func (a ArrAddAnalysis) Swap(i, j int) {
	cop := a[i]
	a[i] = a[j]
	a[j] = cop
}

type StorageAnalysis interface {
	Add(AddAnalysis) bool
	GetAllAndDelete() ArrAddAnalysis
}
