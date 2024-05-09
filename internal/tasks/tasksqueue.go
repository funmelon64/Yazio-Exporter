package tasks

import (
	"sync"
)

type SyncTasksQueue[Ttask comparable, Tresult any] struct {
	mx      sync.Mutex
	Tasks   []Ttask
	Results map[Ttask]Tresult
}

func (que *SyncTasksQueue[Ttask, Tresult]) PopTask() (task *Ttask) {
	que.mx.Lock()
	defer que.mx.Unlock()
	if len(que.Tasks) > 0 {
		task = &que.Tasks[0]
		que.Tasks = que.Tasks[1:]
	}
	return
}

func (que *SyncTasksQueue[Ttask, Tresult]) PushResult(task Ttask, result Tresult) {
	que.mx.Lock()
	defer que.mx.Unlock()
	que.Results[task] = result
}