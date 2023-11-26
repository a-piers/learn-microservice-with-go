package server

import (
	"lib/models"
	"sync"
)

func (srv *Server) MakeMigrations() error {
	migration_models := []interface{}{
		&models.UserModel{},
	}

	l := len(migration_models)
	wg := sync.WaitGroup{}
	wg.Add(l)
	lock := sync.Mutex{}

	for _, model := range migration_models {
		go func(model interface{}) {
			defer wg.Done()
			lock.Lock()
			if err := srv.Storage.AutoMigration(model); err != nil {
				panic(err)
			}
			lock.Unlock()
		}(model)
	}

	wg.Wait()

	return nil
}
