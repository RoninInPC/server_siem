package service

import (
	"server_siem/config"
	"server_siem/entity/subject/notification/receivernotification"
	"server_siem/hash"
	"server_siem/hostinfo"
	"server_siem/sender"
	"server_siem/service/analysis"
	"server_siem/storageanalysis/mapanalysis"
	"server_siem/storagepids/redispids"
	"server_siem/storageservers/redisservers"
	"server_siem/storagesubject/mongosubject"
	"time"
)

type Program struct {
	ApiService      ApiService
	AnalysisService analysis.AnalysisService
}

func InitProgram(fileName string) *Program {
	hostinfo.HostInfoInit()
	conf, err := config.ReadFromFile(fileName)
	if err != nil {
		panic(err)
	}

	redisPIDs := redispids.Init(conf.RedisPIDs.Address, conf.RedisPIDs.Password, conf.RedisPIDs.DB)
	redisServers := redisservers.Init(conf.RedisServers.Address, conf.RedisServers.Password, conf.RedisServers.DB, hash.ToMD5)
	mongoSubjects := mongosubject.Init(conf.MongoSubject.Address, conf.MongoSubject.Username, conf.MongoSubject.Password)
	notChannel := make(chan receivernotification.Notification)

	apiService := InitApiService(conf.Server.AddressUp,
		redisPIDs,
		redisServers,
		mongoSubjects,
		notChannel)

	storageAnalysis := mapanalysis.Init()
	analysisService := analysis.AnalysisService{
		Storage:         storageAnalysis,
		StorageServers:  redisServers,
		StorageSubjects: mongoSubjects,
		Sender:          sender.InitJWTSender(),
		Channel:         notChannel,
		Duration:        time.Minute * 10,
	}
	return &Program{apiService, analysisService}
}

func (p *Program) Work() {
	go p.ApiService.Work()
	p.AnalysisService.Work()
}
