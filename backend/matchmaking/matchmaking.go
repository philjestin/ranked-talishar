package matchmaking

import (
	"math"
	"sync"
	"time"

	"github.com/philjestin/ranked-talishar/schemas"
)

type Matchmaking interface {
	AddPlayer(player *schemas.MatchmakingUser)
	RemovePlayer(username string) *schemas.MatchmakingUser
	GetPlayer(username string) *schemas.MatchmakingUser
	FindOpponent(player *schemas.MatchmakingUser) (*schemas.MatchmakingUser, float64)
}

type MatchmakingPool struct {
	Players map[string]*schemas.MatchmakingUser
	mutex   *sync.Mutex
}

var defaultTargetEloDifference = 15.0

var CCPool MatchmakingPool
var BlitzPool MatchmakingPool

func init() {
	CCPool = *NewMatchMakingPool()
	BlitzPool = *NewMatchMakingPool()
}

func NewMatchMakingPool() *MatchmakingPool {
	return &MatchmakingPool{
		Players: make(map[string]*schemas.MatchmakingUser),
		mutex:   &sync.Mutex{},
	}
}

func (pool *MatchmakingPool) AddPlayer(player *schemas.MatchmakingUser) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	player.QueuedSince = time.Now()
	pool.Players[player.UserName] = player
}

func (pool *MatchmakingPool) RemovePlayer(username string) *schemas.MatchmakingUser {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	player, ok := pool.Players[username]
	if ok {
		delete(pool.Players, username)
		player.QueueStopped = time.Now()
		return player
	}
	return nil
}

func (pool *MatchmakingPool) GetPlayer(username string) *schemas.MatchmakingUser {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	return pool.Players[username]
}

// Iterates through players in the pool
func (pool *MatchmakingPool) FindOpponent(player *schemas.MatchmakingUser) (*schemas.MatchmakingUser, float64) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	searchRange := defaultTargetEloDifference

	// Search for opponents until an opponent is found or the search reaches its limit.
	for {
		opponent := pool.findOpponentWithinRange(player, searchRange)
		if opponent != nil {
			return opponent, searchRange
		}
		searchRange *= 2
		// Prevent extremely wide searches.
		if searchRange > 500 {
			// Restart
			searchRange = defaultTargetEloDifference
		}
	}

}

// Skip checking yourself
// Calculate the elo difference between two players
// if the eloDiff is less than or equal to the targetDiff, the opponent is a potential match and is returned.
// If no opponent is found within the targetDiff the function returns nil.
func (pool *MatchmakingPool) findOpponentWithinRange(player *schemas.MatchmakingUser, targetEloDifference float64) *schemas.MatchmakingUser {
	for username, opponent := range pool.Players {
		if username == player.UserName {
			continue // Can't play yourself.
		}
		eloDiff := math.Abs(float64(player.Elo) - float64(opponent.Elo))
		if eloDiff <= targetEloDifference {
			return opponent
		}
	}
	return nil
}
