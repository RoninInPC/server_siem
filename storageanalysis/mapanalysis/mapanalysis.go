package mapanalysis

import (
	"server_siem/storageanalysis"
	"slices"
	"sync"
)

type MapAnalysis struct {
	local []storageanalysis.AddAnalysis
	sync.Mutex
}

func Init() *MapAnalysis {
	return &MapAnalysis{local: make([]storageanalysis.AddAnalysis, 0)}
}

func (m *MapAnalysis) Add(analysis storageanalysis.AddAnalysis) bool {
	m.Lock()
	defer m.Unlock()
	m.local = append(m.local, analysis)
	return true
}

func (m *MapAnalysis) GetAllAndDelete() storageanalysis.ArrAddAnalysis {
	m.Lock()
	defer m.Unlock()
	answer := slices.Clone(m.local)
	m.local = make([]storageanalysis.AddAnalysis, 0)
	return answer
}
