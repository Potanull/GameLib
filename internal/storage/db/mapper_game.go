package db

import (
	"fmt"
	"strings"

	"gamelib/internal/entities"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	GamesTabler = "gamelib.t_games"

	IDCol   = "id"
	DoneCol = "done"
	NameCol = "name"

	CreateDTCol = "create_dt"
	UpdateDTCol = "update_dt"
)

var GamesBaseCols = []string{NameCol, DoneCol}
var GamesAllCols = append([]string{IDCol, CreateDTCol, UpdateDTCol}, GamesBaseCols...)

func GetGame(_ *gin.Context, id int64, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&game.ID,
			&game.CreateDt,
			&game.UpdateDt,
			&game.Name,
			&game.Done,
		); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &game, nil
}

func CreateGame(_ *gin.Context, createGame *entities.CreateGame, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Insert(GamesTabler).
		Columns(NameCol).
		Values(createGame.Name).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ","))).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()

	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func PutGame(_ *gin.Context, id int64, updateGame *entities.UpdateGame, repo *sqlx.DB) (*entities.Game, error) {
	query := sq.Update(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")))

	if updateGame.Done != nil {
		query = query.Set(DoneCol, *updateGame.Done)
	}

	if len(updateGame.Name) > 0 {
		query = query.Set(NameCol, updateGame.Name)
	}

	rows, err := query.PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func DeleteGame(_ *gin.Context, id int64, repo *sqlx.DB) (*entities.Game, error) {
	query := sq.Delete(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		Suffix(fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")))

	rows, err := query.PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	if err != nil {
		return nil, err
	}

	var result entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
			&result.CreateDt,
			&result.UpdateDt,
			&result.Name,
			&result.Done,
		); err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// ============================================================================================================

func GetGameByName(_ *gin.Context, name string, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		Where(sq.Eq{NameCol: name}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(
			&game.ID,
			&game.CreateDt,
			&game.UpdateDt,
			&game.Name,
			&game.Done,
		); err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &game, nil
}

func GetAllGames(_ *gin.Context, repo *sqlx.DB) ([]*entities.Game, error) {
	rows, err := sq.Select(GamesAllCols...).
		From(GamesTabler).
		OrderBy(DoneCol, fmt.Sprintf("LOWER(%v)", NameCol)).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gameList []*entities.Game
	for rows.Next() {
		var tmp entities.Game
		if err := rows.Scan(
			&tmp.ID,
			&tmp.CreateDt,
			&tmp.UpdateDt,
			&tmp.Name,
			&tmp.Done,
		); err != nil {
			return gameList, err
		}
		gameList = append(gameList, &tmp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gameList, nil
}

func GetRandomGame(_ *gin.Context, done bool, repo *sqlx.DB) (*entities.Game, error) {
	rows, err := sq.Select(GamesBaseCols...).
		From(GamesTabler).
		Where(sq.Eq{DoneCol: done}).
		OrderBy("random()").
		Limit(1).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var game entities.Game
	for rows.Next() {
		if err := rows.Scan(&game.Name, &game.Done); err != nil {
			return nil, err
		}
	}

	return &game, nil
}

func GetRandomListGames(_ *gin.Context, done bool, repo *sqlx.DB) ([]*entities.Game, error) {
	rows, err := sq.Select(GamesBaseCols...).
		From(GamesTabler).
		Where(sq.Eq{DoneCol: done}).
		OrderBy("random()").
		Limit(30).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var gameList []*entities.Game
	for rows.Next() {
		var tmp entities.Game
		if err := rows.Scan(&tmp.Name, &tmp.Done); err != nil {
			return gameList, err
		}
		gameList = append(gameList, &tmp)
	}

	if err = rows.Err(); err != nil {
		return gameList, err
	}
	return gameList, nil
}

func CheckGame(_ *gin.Context, id int64, repo *sqlx.DB) (bool, *entities.Game, error) {
	rows, err := sq.Select(GamesBaseCols...).
		From(GamesTabler).
		Where(sq.Eq{IDCol: id}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return false, nil, err
	}

	check := false
	var game entities.Game
	for rows.Next() {
		check = true
		if err := rows.Scan(
			&game.Name,
			&game.Done,
		); err != nil {
			return false, nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return false, nil, err
	}

	return check, &game, nil
}

func CheckGameByName(_ *gin.Context, name string, repo *sqlx.DB) (bool, *entities.Game, error) {
	rows, err := sq.Select(GamesBaseCols...).
		From(GamesTabler).
		Where(sq.Eq{NameCol: name}).
		PlaceholderFormat(sq.Dollar).
		RunWith(repo.DB).Query()
	defer rows.Close()

	if err != nil {
		return false, nil, err
	}

	check := false
	var game entities.Game
	for rows.Next() {
		check = true
		if err := rows.Scan(
			&game.Name,
			&game.Done,
		); err != nil {
			return false, nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return false, nil, err
	}

	return check, &game, nil
}

func ReverseDoneStatus(_ *gin.Context, id int64, repo *sqlx.DB) error {
	_, err := repo.DB.Exec("UPDATE gamelib.t_games SET DONE = NOT DONE WHERE id = $1 "+
		fmt.Sprintf("RETURNING %s", strings.Join(GamesAllCols, ",")), id)

	if err != nil {
		return err
	}
	return nil
}

//
//func (s *Storage) DeleteGameRequest(ctx *gin.Context, game string) error {
//	if _, err := s.DataBase.Exec("DELETE FROM gamelib.t_games WHERE LOWER(name) = LOWER($1)", game); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *Storage) UpdateGameDoneRequest(ctx *gin.Context, game string) error {
//	if _, err := s.DataBase.Exec("UPDATE gamelib.t_games SET DONE = NOT DONE WHERE name = $1", game); err != nil {
//		return err
//	}
//	return nil
//}
