BEGIN;

DROP TABLE IF EXISTS sets;
DROP TABLE IF EXISTS workouts;

CREATE TABLE workouts (
    workout_id SERIAL PRIMARY KEY,
    timestamp BIGINT
);

CREATE TABLE sets (
    workout_id INT REFERENCES workouts (workout_id),
    set_id INT,
    text VARCHAR(64),
    name VARCHAR(32),
    weights VARCHAR(32),
    reps VARCHAR(32),
    PRIMARY KEY (workout_id, set_id)
);

INSERT INTO workouts (timestamp)
VALUES (1704776488);

INSERT INTO sets (workout_id, set_id, text, name, weights, reps)
VALUES (1, 1, 'Dumbbell Curls 8x60/8x60/6x65/6x65', 'Dumbbell Curls', '8/8/6/6', '60/60/65/65');

INSERT INTO sets (workout_id, set_id, text, name, weights, reps)
VALUES (1, 2, 'Side Raises 8x20/8x20/6x25/4x25', 'Side Raises', '8/8/6/4', '20/20/25/25');

COMMIT;