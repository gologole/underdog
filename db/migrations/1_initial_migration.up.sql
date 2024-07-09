
CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        age INT
);


CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT
);

CREATE TABLE work_logs (
                           id SERIAL PRIMARY KEY,
                           user_id INT NOT NULL REFERENCES people(id),
                           task_id INT NOT NULL REFERENCES tasks(id),
                           start_time TIMESTAMP NOT NULL,
                           end_time TIMESTAMP,
                           duration INTERVAL
);
