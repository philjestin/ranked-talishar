// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	AddParticipant(ctx context.Context, arg AddParticipantParams) error
	CreateContact(ctx context.Context, arg CreateContactParams) (Contact, error)
	CreateConversation(ctx context.Context) (CreateConversationRow, error)
	CreateFormat(ctx context.Context, arg CreateFormatParams) (Format, error)
	CreateGame(ctx context.Context, gameName string) (Game, error)
	CreateHero(ctx context.Context, arg CreateHeroParams) (Hero, error)
	CreateMatch(ctx context.Context, arg CreateMatchParams) (Match, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteContact(ctx context.Context, contactID uuid.UUID) error
	DeleteFormat(ctx context.Context, formatID uuid.UUID) error
	DeleteGame(ctx context.Context, gameID uuid.UUID) error
	DeleteHero(ctx context.Context, heroID uuid.UUID) error
	DeleteMatch(ctx context.Context, matchID uuid.UUID) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetContactById(ctx context.Context, contactID uuid.UUID) (Contact, error)
	GetConversationsByUser(ctx context.Context, userID uuid.NullUUID) ([]Conversation, error)
	GetFormatById(ctx context.Context, formatID uuid.UUID) (Format, error)
	GetGameByID(ctx context.Context, gameID uuid.UUID) (Game, error)
	GetHeroById(ctx context.Context, heroID uuid.UUID) (Hero, error)
	GetMatchById(ctx context.Context, matchID uuid.UUID) (Match, error)
	GetMatchPlayers(ctx context.Context, matchID uuid.UUID) ([]GetMatchPlayersRow, error)
	GetMessagesByConversation(ctx context.Context, conversationID sql.NullInt32) ([]Message, error)
	GetRefreshTokenByUserID(ctx context.Context, userID uuid.UUID) (RefreshToken, error)
	GetUser(ctx context.Context, userName string) (User, error)
	GetUserById(ctx context.Context, userID uuid.UUID) (User, error)
	IncrementLosses(ctx context.Context, arg IncrementLossesParams) error
	IncrementWins(ctx context.Context, arg IncrementWinsParams) error
	ListContacts(ctx context.Context, arg ListContactsParams) ([]Contact, error)
	ListFormats(ctx context.Context, arg ListFormatsParams) ([]Format, error)
	ListGames(ctx context.Context, arg ListGamesParams) ([]Game, error)
	ListHeroes(ctx context.Context, arg ListHeroesParams) ([]Hero, error)
	ListMatches(ctx context.Context, arg ListMatchesParams) ([]Match, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	SendMessage(ctx context.Context, arg SendMessageParams) (Message, error)
	UpdateContact(ctx context.Context, arg UpdateContactParams) (Contact, error)
	UpdateFormat(ctx context.Context, arg UpdateFormatParams) (Format, error)
	UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error)
	UpdateHero(ctx context.Context, arg UpdateHeroParams) (Hero, error)
	UpdateMatch(ctx context.Context, arg UpdateMatchParams) (Match, error)
	UpdatePlayerRating(ctx context.Context, arg UpdatePlayerRatingParams) error
	UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (RefreshToken, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
