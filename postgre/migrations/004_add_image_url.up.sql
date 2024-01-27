/* pgmigrate-encoding: utf-8 */

ALTER TABLE IF NOT EXISTS gamelib.t_games
    ADD COLUMN image_url TEXT;
