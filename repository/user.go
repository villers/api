package repository

import (
	"errors"
	"github.com/villers/api/datamodel"
	"sync"
	"time"
)

type UserRepository interface {
	GetBy(query func(user datamodel.User) bool) (user datamodel.User, found bool)
	GetByID(id int64) (user datamodel.User, found bool)
	GetAll() (results map[int64]datamodel.User)
	InsertOrUpdate(user datamodel.User) (updatedUser datamodel.User, err error)
	DeleteByID(id int64) (deleted bool)
}

func NewUserRepository(source map[int64]datamodel.User) UserRepository {
	return &userMemoryRepository{source: source}
}

type userMemoryRepository struct {
	source map[int64]datamodel.User
	mu     sync.RWMutex
}

func (r *userMemoryRepository) GetBy(query func(user datamodel.User) bool) (user datamodel.User, found bool) {
	r.mu.RLock()
	for _, user = range r.source {
		found = query(user)
		if found {
			break
		}
	}
	r.mu.RUnlock()
	return
}

func (r *userMemoryRepository) GetByID(id int64) (user datamodel.User, found bool) {
	return r.GetBy(func(user datamodel.User) bool {
		return user.ID == id
	})
}

func (r *userMemoryRepository) GetAll() (results map[int64]datamodel.User) {
	r.mu.RLock()
	results = r.source
	r.mu.RUnlock()
	return
}

func (r *userMemoryRepository) InsertOrUpdate(user datamodel.User) (updatedUser datamodel.User, err error) {
	// update
	if id := user.ID; id > 0 {
		_, found := r.GetByID(id)
		if !found {
			return user, errors.New("ID should be zero or a valid one that maps to an existing User")
		}
		r.mu.Lock()
		r.source[id] = user
		r.mu.Unlock()
		return user, nil
	}

	// insert
	id := r.getLastID() + 1
	user.ID = id
	user.Companies = []*datamodel.Company{}
	user.CreatedAt = time.Now()
	r.mu.Lock()
	r.source[id] = user
	r.mu.Unlock()

	return user, nil
}

func (r *userMemoryRepository) DeleteByID(id int64) (deleted bool) {
	r.mu.RLock()
	user, found := r.GetByID(id)
	if found {
		delete(r.source, user.ID)
		deleted = true
	}
	r.mu.RUnlock()

	return
}

func (r *userMemoryRepository) getLastID() (lastID int64) {
	r.mu.RLock()
	for id := range r.source {
		if id > lastID {
			lastID = id
		}
	}
	r.mu.RUnlock()

	return lastID
}
