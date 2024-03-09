package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Game struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Genre     string `json:"genre"`
	Platform  string `json:"platform"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GameModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m GameModel) GetAll() ([]Game, error) {
	query := `SELECT * FROM books`

	var games []Game
	var game Game
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	rows, err := m.DB.Query(query)
	if err != nil {
		m.ErrorLog.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&game.Id, &game.Name, &game.Genre, &game.Platform, &game.CreatedAt, &game.UpdatedAt)
		if err != nil {
			m.ErrorLog.Println(err)
			return nil, err
		}
		m.InfoLog.Println(game)
		games = append(games, game)
	}

	return games, nil
}

func (m GameModel) Get(id int) (*Game, error) {
	query := `
		SELECT id, name, genre, platform, created_at, updated_at
		FROM games
		WHERE id = $1
		`

	var game Game

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&game.Id, &game.Name, &game.Genre, &game.Platform, &game.CreatedAt, &game.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (m GameModel) Insert(game *Game) error {
	query := `
		INSERT INTO gamess (name, genre, platform) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
	`

	args := []interface{}{game.Name, game.Genre, game.Platform}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&game.Id, &game.CreatedAt, &game.UpdatedAt)
}

func (m GameModel) Delete(id int) error {
	query := `
		DELETE FROM games
		WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func (m GameModel) Update(game *Game) error {
	query := `
		UPDATE games
		SET name = $1, genre = $2, platform = $3
		WHERE id = $4
		RETURNING updated_at
		`
	args := []interface{}{game.Name, game.Genre, game.Platform, game.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&game.UpdatedAt)
}
