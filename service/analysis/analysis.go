package analysis

import (
	"server_siem/entity/subject"
	"server_siem/entity/subject/notification/receivernotification"
	"server_siem/hostinfo"
	"server_siem/sender"
	"server_siem/storageanalysis"
	"server_siem/storageservers"
	"server_siem/storagesubject"
	"sort"
	"time"
)

type Who struct {
	Who     *subject.User
	Process *subject.Process
}

type AnalysisService struct {
	Storage         storageanalysis.StorageAnalysis
	StorageServers  storageservers.StorageServers
	StorageSubjects storagesubject.StorageSubjects
	Sender          sender.Sender
	Channel         chan receivernotification.Notification
	Duration        time.Duration
}

func (a AnalysisService) Work() {
	go func() {
		for input := range a.Channel {
			a.Storage.Add(input)
		}
	}()
	for {
		all := a.Storage.GetAllAndDelete()
		sort.Sort(all)
		for _, receiver := range a.StorageServers.GetType(storageservers.Receiver) {
			for _, info := range all {
				codeName := a.StorageServers.Get(storageservers.Receiver, receiver).CodeName
				a.Sender.Send(MakeAddressHost(receiver, codeName+"/api/alert"), subject.InitMessage(
					"",
					"send_receiver",
					"send_receiver",
					hostinfo.GetHostInfo(),
					info,
					"",
					""))
			}
		}
		time.Sleep(a.Duration)
	}
}

func MakeAddressHost(address, specifically string) string {
	return "https://" + address + "/" + specifically
}
