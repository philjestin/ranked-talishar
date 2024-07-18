package matchmaking

import (
	"testing"

	"github.com/philjestin/ranked-talishar/schemas"
	"github.com/stretchr/testify/require"
)

type MockPlayer struct {
	UserName string
	Elo      int64
}

func TestNewMatchmakingPool(t *testing.T) {
	pool := NewMatchMakingPool()

	require.NotNil(t, pool, "Expected a non-nil MatchmakingPool")
	require.NotNil(t, pool.Players, "Players map should be initialized")
	require.NotNil(t, pool.mutex, "mutex should be initialized")
}

func TestAddPlayer(t *testing.T) {
	pool := NewMatchMakingPool()
	player := &schemas.MatchmakingUser{
		UserName: "philjestin",
		Wins:     1,
		Losses:   2,
		Elo:      1475,
	}

	pool.AddPlayer(player)
	require.Equal(t, player, pool.Players[player.UserName], "Player should be in the pool")

	player2 := &schemas.MatchmakingUser{
		UserName: "codemanjack",
		Wins:     1,
		Losses:   2,
		Elo:      1475,
	}
	pool.AddPlayer(player2)
	require.Equal(t, player2, pool.Players[player2.UserName], "Player should be in the pool")
	require.Equal(t, 2, len(pool.Players), "The length of the map of players in the pool should be 2")
}

func TestRemovePlayer(t *testing.T) {
	pool := NewMatchMakingPool()

	player := &schemas.MatchmakingUser{
		UserName: "philjestin",
		Wins:     1,
		Losses:   2,
		Elo:      1475,
	}

	pool.AddPlayer(player)
	removedPlayer := pool.RemovePlayer("philjestin")
	require.Equal(t, player, removedPlayer, "Expected player to be removed from the pool.")
	require.Equal(t, 0, len(pool.Players), "The length of the map of players in the pool should be 0")

	removedPlayer2 := pool.RemovePlayer("billyjoel")
	require.Nil(t, removedPlayer2, "Player doesn't exist and RemovePlayer should return nil")
}

func TestGetPlayer(t *testing.T) {
	pool := NewMatchMakingPool()

	player := &schemas.MatchmakingUser{
		UserName: "philjestin",
		Wins:     1,
		Losses:   2,
		Elo:      1475,
	}

	pool.AddPlayer(player)
	retrievedPlayer := pool.GetPlayer("philjestin")
	require.Equal(t, player, retrievedPlayer, "Expected player to be retrieved from the pool.")
	require.Equal(t, 1, len(pool.Players), "The length of the map of players in the pool should remain unchanged.")
}

func TestFindOpponent_InitialRange(t *testing.T) {
	pool := NewMatchMakingPool()
	players := []*schemas.MatchmakingUser{
		{UserName: "player1", Elo: 1500, Wins: 1, Losses: 2},
		{UserName: "player2", Elo: 1490, Wins: 1, Losses: 2},
		{UserName: "player3", Elo: 1450, Wins: 1, Losses: 2},
		{UserName: "player3", Elo: 1600, Wins: 1, Losses: 2},
	}

	for _, player := range players {
		pool.AddPlayer(player)
	}

	opponent, searchRange := pool.FindOpponent(players[0])
	require.NotNil(t, opponent, "Opponent should be found within initial target elo difference")
	require.Equal(t, opponent, players[1], "Opponent should be player within elo range")
	require.Equal(t, 15.0, searchRange, "Search range should be initial target elo range")
}

func TestFindOpponent_ExponentialBackoff(t *testing.T) {
	pool := NewMatchMakingPool()
	players := []*schemas.MatchmakingUser{
		{UserName: "player1", Elo: 1500, Wins: 1, Losses: 2},
		{UserName: "player2", Elo: 1200, Wins: 1, Losses: 8},
		{UserName: "player3", Elo: 2000, Wins: 15, Losses: 3},
		{UserName: "player4", Elo: 2100, Wins: 19, Losses: 1},
	}

	for _, player := range players {
		pool.AddPlayer(player)
	}

	opponent, searchRange := pool.FindOpponent(players[0])
	require.NotNil(t, opponent, "Opponent should be found with exponential backoff.")
	require.Equal(t, opponent, players[1], "Opponent should be found within elo range during backoff")
	require.Equal(t, 480.0, searchRange, "SearchRange would have hit 15 x the initial backoff")
}
