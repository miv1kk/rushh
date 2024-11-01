package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// Строка подключения к базе данных PostgreSQL
	connectionString := "тут короче строка подключения"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Проверяем, что соединение установлено успешно
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Возвращаем объект соединения
	return db, nil
}
