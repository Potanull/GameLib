/* pgmigrate-encoding: utf-8 */

-- name: create-idx-t-game-name
CREATE INDEX IF NOT EXISTS idx_t_game_name ON gamelib.t_games USING BTREE (name)