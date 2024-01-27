/* pgmigrate-encoding: utf-8 */

ALTER TABLE IF NOT EXISTS gamelib.t_games
    ADD COLUMN hltb_id BIGINT DEFAULT 0;

ALTER TABLE IF NOT EXISTS gamelib.t_games
    ADD COLUMN hltb_main_time BIGINT DEFAULT 0;

ALTER TABLE IF NOT EXISTS gamelib.t_games
    ADD COLUMN hltb_full_time BIGINT DEFAULT 0;