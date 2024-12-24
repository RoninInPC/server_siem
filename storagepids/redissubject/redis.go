package redissubject

import (
	red "github.com/go-redis/redis"
	"time"
)

var (
	Temporal    = "Temporal"
	NotTemporal = "NotTemporal"
)

type RedisDB struct {
	client *red.Client
}

func Init(address string, passwd string, db int) RedisDB {
	client := red.NewClient(
		&red.Options{
			Addr:     address,
			Password: passwd,
			DB:       db})
	return RedisDB{client: client}
}

func (r RedisDB) AddTemporalPID(host string, pid string, d time.Duration) bool {
	return r.client.HSet(host+Temporal, pid, time.Now().Add(d).String()).Err() == nil
}
func (r RedisDB) AddPID(host string, pid string) bool {
	return r.client.HSet(host+NotTemporal, pid, "").Err() == nil
}
func (r RedisDB) DeletePID(host string, pid string) bool {
	return r.client.HDel(host+NotTemporal, pid).Err() == nil
}
func (r RedisDB) DeleteTemporalPID(host string, pid string) bool {
	return r.client.HDel(host+Temporal, pid).Err() == nil
}
func (r RedisDB) GetTemporalPIDs(host string) []string {
	answer := make([]string, 0)
	for pid, t := range r.client.HGetAll(host).Val() {
		if time.Now().String() > t {
			r.DeleteTemporalPID(host, pid)
		}
		answer = append(answer, pid)
	}
	return answer
}
func (r RedisDB) GetPIDs(host string) []string {
	return r.client.HKeys(host).Val()
}

func (r RedisDB) GetAllPIDs(host string) []string {
	return append(r.GetTemporalPIDs(host), r.GetPIDs(host)...)
}
