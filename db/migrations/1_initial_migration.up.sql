CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        surname VARCHAR(255) NOT NULL,
                        name VARCHAR(255) NOT NULL,
                        patronymic VARCHAR(255),
                        address VARCHAR(255) NOT NULL,
                        passport_number VARCHAR(255) NOT NULL
);

INSERT INTO people (surname, name, patronymic, address, passport_number) VALUES
                                                                             ('Иванов', 'Иван', 'Иванович', 'г. Москва, ул. Ленина, д. 1, кв. 1', '1234567890'),
                                                                             ('Петров', 'Петр', 'Петрович', 'г. Санкт-Петербург, ул. Невский, д. 2, кв. 2', '2234567890'),
                                                                             ('Сидоров', 'Сидор', 'Сидорович', 'г. Казань, ул. Баумана, д. 3, кв. 3', '3234567890'),
                                                                             ('Смирнов', 'Алексей', 'Алексеевич', 'г. Новосибирск, ул. Красный проспект, д. 4, кв. 4', '4234567890'),
                                                                             ('Кузнецов', 'Михаил', 'Михайлович', 'г. Екатеринбург, ул. Ленина, д. 5, кв. 5', '5234567890'),
                                                                             ('Попов', 'Андрей', 'Андреевич', 'г. Нижний Новгород, ул. Большая Покровская, д. 6, кв. 6', '6234567890'),
                                                                             ('Васильев', 'Сергей', 'Сергеевич', 'г. Самара, ул. Ленинградская, д. 7, кв. 7', '7234567890'),
                                                                             ('Зайцев', 'Антон', 'Антонович', 'г. Омск, ул. Ленина, д. 8, кв. 8', '8234567890'),
                                                                             ('Волков', 'Николай', 'Николаевич', 'г. Челябинск, ул. Кирова, д. 9, кв. 9', '9234567890'),
                                                                             ('Федоров', 'Дмитрий', 'Дмитриевич', 'г. Ростов-на-Дону, ул. Садовая, д. 10, кв. 10', '1034567890');


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
