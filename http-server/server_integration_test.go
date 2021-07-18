package http_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
	s.scores[name]++
}

func TestRecodingWinsAndRetrievingThem(t *testing.T) {
	store := &InMemoryPlayerStore{map[string]int{}}
	server := &PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
