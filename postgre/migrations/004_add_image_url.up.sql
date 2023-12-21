/* pgmigrate-encoding: utf-8 */

ALTER TABLE gamelib.t_games
    ADD COLUMN image_url TEXT;
