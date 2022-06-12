package Shortener

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"playground/MainUrlShortener/app/Modules/Shortener"
)

func Route(db *sql.DB) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		userToken, errToken := req.Cookie("SHRTTOKEN")
		if errToken != nil {
			_, err := res.Write([]byte("Token is not found"))
			if err != nil {
				panic(err.Error())
			}
			return
		}

		data := Shortener.GetPathForRedirect(db, req.URL.Path, userToken.Value)

		if len(data.GetToPath()) == 0 {
			_, err := res.Write([]byte("Path does not exist"))
			if err != nil {
				panic(err.Error())
			}
			return
		}

		http.Redirect(res, req, "https://www."+data.GetToPath(), http.StatusSeeOther)
		return
	})

	http.HandleFunc("/api/getPaths", func(res http.ResponseWriter, req *http.Request) {
		userToken, errToken := req.Cookie("SHRTTOKEN")
		if errToken != nil {
			_, err := res.Write([]byte("Token is not found"))
			if err != nil {
				panic(err.Error())
			}
			return
		}

		data := Shortener.GetPaths(db, userToken.Value)

		res.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(res).Encode(data); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/api/createPath", func(res http.ResponseWriter, req *http.Request) {
		var up Shortener.UsersPaths
		var cookie http.Cookie
		var tokenValue string

		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&up)

		if err != nil {
			panic(err)
		}

		userToken, errCookieVal := req.Cookie("SHRTTOKEN")
		if errCookieVal != nil {
			chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "l"}
			generatedToken := ""
			tokenLength := 15

			for i := 0; i < tokenLength; i++ {
				generatedToken += chars[rand.Intn(len(chars)-0)+0]
			}

			cookie = http.Cookie{
				Name:     "SHRTTOKEN",
				Value:    generatedToken,
				HttpOnly: true,
			}

			tokenValue = generatedToken
		} else {
			tokenValue = userToken.Value
		}

		_, err = Shortener.CreatPath(db, up, tokenValue)
		if err != nil {
			panic(err.Error())
		}

		data := map[string]string{
			"ErrorCode":    "0",
			"ErrorMessage": "Created Path",
		}

		res.Header().Set("Content-Type", "application/json")

		if len(cookie.Value) > 0 {
			http.SetCookie(res, &cookie)
		}

		if err := json.NewEncoder(res).Encode(data); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
