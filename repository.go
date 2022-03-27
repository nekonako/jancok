package main

import "database/sql"

type repository interface {
	Create(code string) error
	GetLastCode() (int, error)
}

type repositoryImp struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return &repositoryImp{
		db: db,
	}
}

func (r repositoryImp) Create(code string) error {
	_, err := r.db.Exec(createCode, code)
	if err != nil {
		return err
	}
	return nil
}

func (r repositoryImp) GetLastCode() (int, error) {
	var totalCode int
	rows, err := r.db.Query(getTotalCode)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&totalCode)
		if err != nil {
			return -1, err
		}
	}
	return totalCode, nil
}
