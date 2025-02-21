
DROP TRIGGER IF EXISTS enforce_target_limit ON targets;
DROP TRIGGER IF EXISTS block_notes_update ON notes;


DROP FUNCTION IF EXISTS check_target_count;
DROP FUNCTION IF EXISTS prevent_notes_update;


DROP TABLE IF EXISTS notes;
DROP TABLE IF EXISTS targets;
DROP TABLE IF EXISTS missions;
DROP TABLE IF EXISTS spy_cats;
