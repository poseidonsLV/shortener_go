package Shortener

import (
	"database/sql"
	"fmt"
	"time"
)

type UsersPaths struct {
	FromPath   string    `json:"FromPath"`
	ToPath     string    `json:"ToPath"`
	ExpireDate time.Time `json:"ExpireDate"`
}

func (up UsersPaths) GetToPath() string {
	return up.ToPath
}

func GetPathForRedirect(db *sql.DB, path string, userToken string) UsersPaths {
	rows, err := db.Query(fmt.Sprintf(`
		SELECT 
			from_path AS FromPath,
			to_path AS ToPath,
			expire_date AS ExpireDate
		FROM user_path
		WHERE address = '%s'
			AND from_path = '%s'
			AND expire_date >= NOW()
		LIMIT 1
		`, userToken, path))
	if err != nil {
		panic(err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)

	var paths UsersPaths

	for rows.Next() {
		var path UsersPaths

		err = rows.Scan(&path.FromPath, &path.ToPath, &path.ExpireDate)

		if err != nil {
			panic(err.Error())
		}

		paths = path
	}

	return paths
}

func GetPaths(db *sql.DB, userToken string) []UsersPaths {
	rows, err := db.Query(fmt.Sprintf(`
		SELECT from_path AS FromPath,
		       to_path AS ToPath,
		       expire_date AS ExpireDate
		FROM user_path
		WHERE address = '%s'
		`, userToken))
	if err != nil {
		panic(err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err.Error())
		}
	}(rows)

	var paths []UsersPaths

	for rows.Next() {
		var path UsersPaths

		err = rows.Scan(&path.FromPath, &path.ToPath, &path.ExpireDate)

		if err != nil {
			panic(err.Error())
		}

		paths = append(paths, path)
	}

	return paths
}

func CreatPath(db *sql.DB, data UsersPaths, userToken string) (bool, error) {
	rows, err := db.Query(fmt.Sprintf(
		`
			INSERT INTO user_path (address, to_path, from_path, expire_date)
			VALUES ('%s', '%s', '%s', '%s')
		`,
		userToken,
		data.ToPath,
		data.FromPath,
		data.ExpireDate.Format("2006-1-2 15:4:5"),
	))
	if err != nil {
		return false, err
	}

	err = rows.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}
