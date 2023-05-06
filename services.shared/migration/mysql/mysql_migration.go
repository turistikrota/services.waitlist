package mysql_migration

import (
	"github.com/jmoiron/sqlx"
)

type Migration struct {
	db    *sqlx.DB
	items []MigrationItem
}

type MigrationItem interface {
	Up(db *sqlx.DB) error
	Down(db *sqlx.DB) error
	GetDatabase() string
}

func New(db *sqlx.DB) *Migration {
	return &Migration{db: db, items: []MigrationItem{}}
}

func (m *Migration) Add(item MigrationItem) {
	m.items = append(m.items, item)
}

func (m *Migration) Up() error {
	query := `CREATE TABLE IF NOT EXISTS migrations ( id INT AUTO_INCREMENT PRIMARY KEY, dbName VARCHAR(255) NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`
	_, err := m.db.Exec(query)
	if err != nil {
		return err
	}

	query = `SELECT dbName FROM migrations`
	rows, err := m.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	migrated := []string{}
	for rows.Next() {
		var database string
		err := rows.Scan(&database)
		if err != nil {
			return err
		}
		migrated = append(migrated, database)
	}

	for _, item := range m.items {
		if !contains(migrated, item.GetDatabase()) {
			err := item.Up(m.db)
			if err != nil {
				return err
			}

			query = `INSERT INTO migrations (dbName) VALUES (?)`
			_, err = m.db.Exec(query, item.GetDatabase())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *Migration) Down() error {
	query := `SELECT dbName FROM migrations`
	rows, err := m.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	migrated := []string{}
	for rows.Next() {
		var database string
		err := rows.Scan(&database)
		if err != nil {
			return err
		}
		migrated = append(migrated, database)
	}

	for _, item := range m.items {
		if contains(migrated, item.GetDatabase()) {
			err := item.Down(m.db)
			if err != nil {
				return err
			}

			query = `DELETE FROM migrations WHERE dbName = ?`
			_, err = m.db.Exec(query, item.GetDatabase)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
