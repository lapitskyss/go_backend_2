package sharding

import (
	"database/sql"
	"math/rand"
	"time"
)

type Sharding struct {
	M *Manager
	P *Pool
	r *rand.Rand
}

func NewSharding(m *Manager, p *Pool) *Sharding {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return &Sharding{
		M: m,
		P: p,
		r: r,
	}
}

func (s *Sharding) Conn(entityId int) (*sql.DB, error) {
	shard, err := s.M.ShardById(entityId)
	if err != nil {
		return nil, err
	}

	return s.P.Connection(shard.Address)
}

// Master connection
func (s *Sharding) Master(entityId int) (*sql.DB, error) {
	return s.Conn(entityId)
}

// Slave random connection
func (s *Sharding) Slave(entityId int) (*sql.DB, error) {
	shard, err := s.M.ShardById(entityId)
	if err != nil {
		return nil, err
	}

	l := len(shard.Slaves)
	if l == 0 {
		return nil, ErrorSlaveNotFound
	}

	return s.P.Connection(shard.Slaves[s.r.Intn(l)])
}

// Random master/slave connection
func (s *Sharding) Random(entityId int) (*sql.DB, error) {
	if s.r.Intn(2) == 0 {
		return s.Master(entityId)
	}

	return s.Slave(entityId)
}
