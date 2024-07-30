package main

import (
	"context"
	"database/sql"
	_"fmt"
	"time"
	_ "modernc.org/sqlite"
)

// начало решения

// SQLMap представляет карту, которая хранится в SQL-базе данных
type SQLMap struct{
	db *sql.DB
	getStmt *sql.Stmt
	setStmt *sql.Stmt
	delStmt *sql.Stmt
	timeout time.Duration
}

// NewSQLMap создает новую SQL-карту в указанной базе
// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	db.Exec("create table if not exists map(key text primary key, val blob)")
	getStmt, err := db.Prepare("select val from map where key = ?")
	if err != nil {
		return nil, err
	}
	setStmt, err := db.Prepare(`
	insert into map(key, val) values (?, ?)
	on conflict (key) do update set val = excluded.val
	`)
	if err != nil {
		return nil, err
	}
	delStmt, err := db.Prepare(`delete from map where key = ?`)
	if err != nil {
		return nil, err
	}
	return &SQLMap{db, getStmt, setStmt, delStmt, 60 * time.Millisecond}, nil
}


// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	res := m.getStmt.QueryRowContext(ctx, key)
	var result any
	err := res.Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	//defer m.setCancelFunc()
	_, err := m.setStmt.ExecContext(ctx, key, val)
	if err != nil {
		return err
	}
	return nil
}

// SetItems устанавливает значения указанных ключей.
func (m *SQLMap) SetItems(items map[string]any) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt := tx.Stmt(m.setStmt)
	for k, v := range items {
		_, err := stmt.ExecContext(ctx, k, v)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// Delete удаляет запись карты с указанным ключом.
// Если такого ключа нет - ничего не делает (это не считается ошибкой).
func (m *SQLMap) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	_, err := m.delStmt.ExecContext(ctx, key)
	if err != nil {
		return err
	}
	return nil
}

// Close освобождает ресурсы, занятые картой в базе.
func (m *SQLMap) Close() error {
	err := m.delStmt.Close()
	if err != nil {
		return err
	}
	err = m.getStmt.Close()
	if err != nil {
		return err
	}
	err = m.setStmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (m *SQLMap) SetTimeout(timeout time.Duration) {
	m.timeout = timeout
}

// конец решения

func main() {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.SetTimeout(10 * time.Millisecond)

	m.Set("name", "Alice")
	m.Get("name")
}