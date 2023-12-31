-- Создание таблицы "users"
CREATE TABLE users (
  -- Уникальный идентификатор пользователя, создается автоматически при вставке записи в таблицу
  id SERIAL PRIMARY KEY,
  -- Имя пользователя в приложении
  name VARCHAR(50) NOT NULL,
  -- Электронная почта пользователя, уникальное поле (на одну почту - только один пользователь)
  email VARCHAR(50) UNIQUE,
  -- Хэш от пароля пользователя
  password VARCHAR(50)
);

-- Создание таблицы "postcards"
CREATE TABLE postcards (
  -- Уникальный идентификатор пожелания, создается автоматически при вставке записи в таблицу
  id SERIAL PRIMARY KEY,
  -- Текст пожелания
  text TEXT NOT NULL,
  -- Уникальный идентификатор автора пожелания
  author_id INT REFERENCES users(id)
);
-- Пока под вопросом эта часть бд:
    -- Создание перечисления "reaction"
    -- CREATE TYPE mood AS ENUM ('thanks', 'like', 'favourite', 'lol');

    -- Создание таблицы "reactions"
    -- CREATE TABLE reactions (
    -- Уникальный идентификатор реакции, создается автоматически при вставке записи в таблицу
    --   id SERIAL PRIMARY KEY,    
    -- Уникальный идентификатор автора пожелания
    --   author_id INT REFERENCES users(id),
    -- Собственно реакция на пожелание
    --   reaction VARCHAR(128)
    -- );
    -- 