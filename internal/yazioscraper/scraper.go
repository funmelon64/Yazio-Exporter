package yazioscraper

import (
	"YazioExporter/internal/tasks"
	"YazioExporter/pkg/yzapi"
	"log"
	"sync"
)

type yazioJsonsScraper[Ttask comparable] struct {
	tasks        []Ttask
	workersCount int
	yzFactory    yzapi.ClientFactory
}

func NewYazioJsonsScraper[Ttask comparable](tasks []Ttask, workersCount int, yzFactory yzapi.ClientFactory) yazioJsonsScraper[Ttask] {
	return yazioJsonsScraper[Ttask]{tasks, workersCount, yzFactory}
}

func (scrapper *yazioJsonsScraper[Ttask]) Scrape(getJsonFromYazio func(yzapi.Client, Ttask) (string, error),
	saveJsons func(map[Ttask]string)) {
	jsonTasks := tasks.SyncTasksQueue[Ttask, string]{Tasks: scrapper.tasks, Results: map[Ttask]string{}}

	runWorkers(scrapper.workersCount, func() func() {
		client := scrapper.yzFactory.NewClient()
		return func() {
			jsonsScrapeWorker(&jsonTasks, func(task Ttask) (string, error) {
				res, err := getJsonFromYazio(client, task)
				if err != nil {
					log.Printf("fail to get json from Yazio for %v: %v", task, err)
				}
				return res, err
			})
		}
	})

	saveJsons(jsonTasks.Results)
}

func jsonsScrapeWorker[Ttask comparable](que *tasks.SyncTasksQueue[Ttask, string], getJson func(Ttask) (string, error)) {
	for task := que.PopTask(); task != nil; task = que.PopTask() {
		consumedJson, err := getJson(*task)
		if err != nil {
			continue
		}

		que.PushResult(*task, consumedJson)
	}
}

func runWorkers(workersCount int, workersFactory func() func()) {
	wg := sync.WaitGroup{}
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			workersFactory()()
			wg.Done()
		}()
	}
	wg.Wait()
}
