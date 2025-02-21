CREATE TABLE spy_cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    experience_years INTEGER NOT NULL CHECK (experience_years >= 0),
    breed VARCHAR(25) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL CHECK (salary >= 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE missions (
    id SERIAL PRIMARY KEY,
    cat_id INTEGER UNIQUE NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (cat_id) REFERENCES spy_cats(id) ON DELETE RESTRICT
);

CREATE TABLE targets (
    id SERIAL PRIMARY KEY,
    mission_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(50) NOT NULL,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (mission_id) REFERENCES missions(id) ON DELETE CASCADE
);

CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    target_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (target_id) REFERENCES targets(id) ON DELETE CASCADE
);


CREATE FUNCTION check_target_count() RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM targets WHERE mission_id = NEW.mission_id) >= 3 THEN
        RAISE EXCEPTION 'A mission can have a maximum of 3 targets';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER enforce_target_limit
BEFORE INSERT ON targets
FOR EACH ROW
EXECUTE FUNCTION check_target_count();

CREATE FUNCTION prevent_notes_update() RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (SELECT 1 FROM targets WHERE id = NEW.target_id AND is_completed = TRUE) THEN
        RAISE EXCEPTION 'Cannot update notes for a completed target';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER block_notes_update
BEFORE UPDATE ON notes
FOR EACH ROW
EXECUTE FUNCTION prevent_notes_update();

