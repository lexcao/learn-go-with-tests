package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecodingWinsAndRetrievingThem(t *testing.T) {
	database, clean := createTempFile(t, "")
	defer clean()

	store := NewFileSystemPlayerStore(database)
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{player, 3},
		}

		assertLeague(t, got, want)
	})
}
