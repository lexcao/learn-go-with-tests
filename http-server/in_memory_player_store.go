package main

type InMemoryPlayerStore struct {
	scores map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (s *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range s.scores {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.scores[name]++
}
