package storagepids

import "time"

type StoragePIDs interface {
	AddTemporalPID(string, string, time.Duration) bool
	AddPID(string, string) bool
	DeletePID(string, string) bool
	DeleteTemporalPID(string, string) bool
	GetTemporalPIDs(string) []string
	GetPIDs(string) []string
	GetAllPIDs(string) []string
}
