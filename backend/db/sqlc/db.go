// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createContactStmt, err = db.PrepareContext(ctx, createContact); err != nil {
		return nil, fmt.Errorf("error preparing query CreateContact: %w", err)
	}
	if q.createFormatStmt, err = db.PrepareContext(ctx, createFormat); err != nil {
		return nil, fmt.Errorf("error preparing query CreateFormat: %w", err)
	}
	if q.createGameStmt, err = db.PrepareContext(ctx, createGame); err != nil {
		return nil, fmt.Errorf("error preparing query CreateGame: %w", err)
	}
	if q.createHeroStmt, err = db.PrepareContext(ctx, createHero); err != nil {
		return nil, fmt.Errorf("error preparing query CreateHero: %w", err)
	}
	if q.createMatchStmt, err = db.PrepareContext(ctx, createMatch); err != nil {
		return nil, fmt.Errorf("error preparing query CreateMatch: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteContactStmt, err = db.PrepareContext(ctx, deleteContact); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteContact: %w", err)
	}
	if q.deleteFormatStmt, err = db.PrepareContext(ctx, deleteFormat); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteFormat: %w", err)
	}
	if q.deleteGameStmt, err = db.PrepareContext(ctx, deleteGame); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteGame: %w", err)
	}
	if q.deleteHeroStmt, err = db.PrepareContext(ctx, deleteHero); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteHero: %w", err)
	}
	if q.deleteMatchStmt, err = db.PrepareContext(ctx, deleteMatch); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteMatch: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getContactByIdStmt, err = db.PrepareContext(ctx, getContactById); err != nil {
		return nil, fmt.Errorf("error preparing query GetContactById: %w", err)
	}
	if q.getFormatByIdStmt, err = db.PrepareContext(ctx, getFormatById); err != nil {
		return nil, fmt.Errorf("error preparing query GetFormatById: %w", err)
	}
	if q.getGameByIDStmt, err = db.PrepareContext(ctx, getGameByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetGameByID: %w", err)
	}
	if q.getHeroByIdStmt, err = db.PrepareContext(ctx, getHeroById); err != nil {
		return nil, fmt.Errorf("error preparing query GetHeroById: %w", err)
	}
	if q.getMatchByIdStmt, err = db.PrepareContext(ctx, getMatchById); err != nil {
		return nil, fmt.Errorf("error preparing query GetMatchById: %w", err)
	}
	if q.getMatchPlayersStmt, err = db.PrepareContext(ctx, getMatchPlayers); err != nil {
		return nil, fmt.Errorf("error preparing query GetMatchPlayers: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.incrementLossesStmt, err = db.PrepareContext(ctx, incrementLosses); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementLosses: %w", err)
	}
	if q.incrementWinsStmt, err = db.PrepareContext(ctx, incrementWins); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementWins: %w", err)
	}
	if q.listContactsStmt, err = db.PrepareContext(ctx, listContacts); err != nil {
		return nil, fmt.Errorf("error preparing query ListContacts: %w", err)
	}
	if q.listFormatsStmt, err = db.PrepareContext(ctx, listFormats); err != nil {
		return nil, fmt.Errorf("error preparing query ListFormats: %w", err)
	}
	if q.listGamesStmt, err = db.PrepareContext(ctx, listGames); err != nil {
		return nil, fmt.Errorf("error preparing query ListGames: %w", err)
	}
	if q.listHeroesStmt, err = db.PrepareContext(ctx, listHeroes); err != nil {
		return nil, fmt.Errorf("error preparing query ListHeroes: %w", err)
	}
	if q.listMatchesStmt, err = db.PrepareContext(ctx, listMatches); err != nil {
		return nil, fmt.Errorf("error preparing query ListMatches: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateContactStmt, err = db.PrepareContext(ctx, updateContact); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateContact: %w", err)
	}
	if q.updateFormatStmt, err = db.PrepareContext(ctx, updateFormat); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateFormat: %w", err)
	}
	if q.updateGameStmt, err = db.PrepareContext(ctx, updateGame); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateGame: %w", err)
	}
	if q.updateHeroStmt, err = db.PrepareContext(ctx, updateHero); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateHero: %w", err)
	}
	if q.updateMatchStmt, err = db.PrepareContext(ctx, updateMatch); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateMatch: %w", err)
	}
	if q.updatePlayerRatingStmt, err = db.PrepareContext(ctx, updatePlayerRating); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePlayerRating: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createContactStmt != nil {
		if cerr := q.createContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createContactStmt: %w", cerr)
		}
	}
	if q.createFormatStmt != nil {
		if cerr := q.createFormatStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createFormatStmt: %w", cerr)
		}
	}
	if q.createGameStmt != nil {
		if cerr := q.createGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createGameStmt: %w", cerr)
		}
	}
	if q.createHeroStmt != nil {
		if cerr := q.createHeroStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createHeroStmt: %w", cerr)
		}
	}
	if q.createMatchStmt != nil {
		if cerr := q.createMatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createMatchStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteContactStmt != nil {
		if cerr := q.deleteContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteContactStmt: %w", cerr)
		}
	}
	if q.deleteFormatStmt != nil {
		if cerr := q.deleteFormatStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteFormatStmt: %w", cerr)
		}
	}
	if q.deleteGameStmt != nil {
		if cerr := q.deleteGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteGameStmt: %w", cerr)
		}
	}
	if q.deleteHeroStmt != nil {
		if cerr := q.deleteHeroStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteHeroStmt: %w", cerr)
		}
	}
	if q.deleteMatchStmt != nil {
		if cerr := q.deleteMatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteMatchStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getContactByIdStmt != nil {
		if cerr := q.getContactByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getContactByIdStmt: %w", cerr)
		}
	}
	if q.getFormatByIdStmt != nil {
		if cerr := q.getFormatByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getFormatByIdStmt: %w", cerr)
		}
	}
	if q.getGameByIDStmt != nil {
		if cerr := q.getGameByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getGameByIDStmt: %w", cerr)
		}
	}
	if q.getHeroByIdStmt != nil {
		if cerr := q.getHeroByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getHeroByIdStmt: %w", cerr)
		}
	}
	if q.getMatchByIdStmt != nil {
		if cerr := q.getMatchByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMatchByIdStmt: %w", cerr)
		}
	}
	if q.getMatchPlayersStmt != nil {
		if cerr := q.getMatchPlayersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getMatchPlayersStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.incrementLossesStmt != nil {
		if cerr := q.incrementLossesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementLossesStmt: %w", cerr)
		}
	}
	if q.incrementWinsStmt != nil {
		if cerr := q.incrementWinsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementWinsStmt: %w", cerr)
		}
	}
	if q.listContactsStmt != nil {
		if cerr := q.listContactsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listContactsStmt: %w", cerr)
		}
	}
	if q.listFormatsStmt != nil {
		if cerr := q.listFormatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listFormatsStmt: %w", cerr)
		}
	}
	if q.listGamesStmt != nil {
		if cerr := q.listGamesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listGamesStmt: %w", cerr)
		}
	}
	if q.listHeroesStmt != nil {
		if cerr := q.listHeroesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listHeroesStmt: %w", cerr)
		}
	}
	if q.listMatchesStmt != nil {
		if cerr := q.listMatchesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listMatchesStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateContactStmt != nil {
		if cerr := q.updateContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateContactStmt: %w", cerr)
		}
	}
	if q.updateFormatStmt != nil {
		if cerr := q.updateFormatStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateFormatStmt: %w", cerr)
		}
	}
	if q.updateGameStmt != nil {
		if cerr := q.updateGameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateGameStmt: %w", cerr)
		}
	}
	if q.updateHeroStmt != nil {
		if cerr := q.updateHeroStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateHeroStmt: %w", cerr)
		}
	}
	if q.updateMatchStmt != nil {
		if cerr := q.updateMatchStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateMatchStmt: %w", cerr)
		}
	}
	if q.updatePlayerRatingStmt != nil {
		if cerr := q.updatePlayerRatingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePlayerRatingStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	createContactStmt      *sql.Stmt
	createFormatStmt       *sql.Stmt
	createGameStmt         *sql.Stmt
	createHeroStmt         *sql.Stmt
	createMatchStmt        *sql.Stmt
	createUserStmt         *sql.Stmt
	deleteContactStmt      *sql.Stmt
	deleteFormatStmt       *sql.Stmt
	deleteGameStmt         *sql.Stmt
	deleteHeroStmt         *sql.Stmt
	deleteMatchStmt        *sql.Stmt
	deleteUserStmt         *sql.Stmt
	getContactByIdStmt     *sql.Stmt
	getFormatByIdStmt      *sql.Stmt
	getGameByIDStmt        *sql.Stmt
	getHeroByIdStmt        *sql.Stmt
	getMatchByIdStmt       *sql.Stmt
	getMatchPlayersStmt    *sql.Stmt
	getUserByIdStmt        *sql.Stmt
	incrementLossesStmt    *sql.Stmt
	incrementWinsStmt      *sql.Stmt
	listContactsStmt       *sql.Stmt
	listFormatsStmt        *sql.Stmt
	listGamesStmt          *sql.Stmt
	listHeroesStmt         *sql.Stmt
	listMatchesStmt        *sql.Stmt
	listUsersStmt          *sql.Stmt
	updateContactStmt      *sql.Stmt
	updateFormatStmt       *sql.Stmt
	updateGameStmt         *sql.Stmt
	updateHeroStmt         *sql.Stmt
	updateMatchStmt        *sql.Stmt
	updatePlayerRatingStmt *sql.Stmt
	updateUserStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		createContactStmt:      q.createContactStmt,
		createFormatStmt:       q.createFormatStmt,
		createGameStmt:         q.createGameStmt,
		createHeroStmt:         q.createHeroStmt,
		createMatchStmt:        q.createMatchStmt,
		createUserStmt:         q.createUserStmt,
		deleteContactStmt:      q.deleteContactStmt,
		deleteFormatStmt:       q.deleteFormatStmt,
		deleteGameStmt:         q.deleteGameStmt,
		deleteHeroStmt:         q.deleteHeroStmt,
		deleteMatchStmt:        q.deleteMatchStmt,
		deleteUserStmt:         q.deleteUserStmt,
		getContactByIdStmt:     q.getContactByIdStmt,
		getFormatByIdStmt:      q.getFormatByIdStmt,
		getGameByIDStmt:        q.getGameByIDStmt,
		getHeroByIdStmt:        q.getHeroByIdStmt,
		getMatchByIdStmt:       q.getMatchByIdStmt,
		getMatchPlayersStmt:    q.getMatchPlayersStmt,
		getUserByIdStmt:        q.getUserByIdStmt,
		incrementLossesStmt:    q.incrementLossesStmt,
		incrementWinsStmt:      q.incrementWinsStmt,
		listContactsStmt:       q.listContactsStmt,
		listFormatsStmt:        q.listFormatsStmt,
		listGamesStmt:          q.listGamesStmt,
		listHeroesStmt:         q.listHeroesStmt,
		listMatchesStmt:        q.listMatchesStmt,
		listUsersStmt:          q.listUsersStmt,
		updateContactStmt:      q.updateContactStmt,
		updateFormatStmt:       q.updateFormatStmt,
		updateGameStmt:         q.updateGameStmt,
		updateHeroStmt:         q.updateHeroStmt,
		updateMatchStmt:        q.updateMatchStmt,
		updatePlayerRatingStmt: q.updatePlayerRatingStmt,
		updateUserStmt:         q.updateUserStmt,
	}
}
