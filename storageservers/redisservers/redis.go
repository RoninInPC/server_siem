package redisservers

import (
	red "github.com/go-redis/redis"
	"server_siem/hash"
	"server_siem/hostinfo"
	"server_siem/storageservers"
)

type RedisDB struct {
	client *red.Client
	hash   hash.Hash
}

func Init(address string, passwd string, db int, h hash.Hash) RedisDB {
	client := red.NewClient(
		&red.Options{
			Addr:     address,
			Password: passwd,
			DB:       db})
	return RedisDB{client: client, hash: h}
}

func (s RedisDB) Add(info hostinfo.HostInfo, host storageservers.TypeHost) bool {
	return nil == s.client.HSet(string(host), info.HostName, info.JSON())
}

func (s RedisDB) Exists(info hostinfo.HostInfo) (storageservers.TypeHost, bool) {
	b := s.client.HGet(string(storageservers.Server), info.HostName).String() != ""
	if b {
		return storageservers.Server, b
	}
	b1 := s.client.HGet(string(storageservers.Receiver), info.HostName).String() != ""
	if b1 {
		return storageservers.Receiver, b
	}
	return storageservers.Nope, false
}

func (s RedisDB) Update(info hostinfo.HostInfo) bool {
	t, _ := s.Exists(info)
	if t != storageservers.Nope {
		return s.Add(info, t)
	}
	return false
}

func (s RedisDB) Delete(info hostinfo.HostInfo) bool {
	t, _ := s.Exists(info)
	if t != storageservers.Nope {
		return s.client.HDel(string(t), info.HostName).Err() == nil
	}
	return false
}

func (s RedisDB) GetType(host storageservers.TypeHost) []string {
	return s.client.HKeys(string(host)).Val()
}

func (s RedisDB) Compare(info hostinfo.HostInfo) (storageservers.TypeHost, bool) {
	b := s.client.HGet(string(storageservers.Server), info.HostName).String() == info.Hash(s.hash)
	if b {
		return storageservers.Server, b
	}
	b1 := s.client.HGet(string(storageservers.Receiver), info.HostName).String() == info.Hash(s.hash)
	if b1 {
		return storageservers.Receiver, b
	}
	return storageservers.Nope, false
}
