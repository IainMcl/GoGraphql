package users

import (
    "log"
    "database/sql"
    "iainmcl/gographql/internal/pkg/db/postgresql"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES(?, ?)")
	print(statement)
	if err != nil {
		log.Fatal(err)
	}
    hashedPassword, err := HashPassword(user.Password)
    _, err = statement.Exec(user.Username, hashedPassword)
    if err != nil{
        log.Fatal(err)
    }
}

func CheckPasswordHash(password, hash string) bool{
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GetUserIdByUsername(username string) (int, error){
    statement, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username = ?")
    if err != nil{
        log.Fatal(err)
    }
    row := statement.QueryRow(username)

    var Id int
    err = row.Scan(&Id)
    if err != nil{
        if err != sql.ErrNoRows{
            log.Print(err)
        }
        return 0, err
    }
    return Id, nil
}
