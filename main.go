package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

const (
	HOST     = "localhost"
	PORT     = "5432"
	USERNAME = "postgres"
	PASSWORD = "postgres"
	DATABASE = "master"
)

func openConnection() (*sql.DB, error) {
	// Строка подключения
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USERNAME, PASSWORD, DATABASE)

	// Открыть БД
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open(): ", err)
		return nil, err
	}
	return db, nil
}

func exists(username string) int {
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userID := -1
	statement := fmt.Sprintf(`SELECT UserID FROM Userdata WHERE Username = '%s'`, username)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println("Query: ", err)
		return -1
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Scan: ", err)
			return -1
		}
		userID = id
	}

	return userID
}

func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userID := exists(d.Username)
	if userID != -1 {
		fmt.Println("User already exists:", d.Username)
		return -1
	}
	statement := `INSERT INTO Userdata (Username, Name, Surname, Description) VALUES ($1, $2, $3, $4) RETURNING UserID`
	err = db.QueryRow(statement, d.Username, d.Name, d.Surname, d.Description).Scan(&userID)
	if err != nil {
		fmt.Println("QueryRow: ", err)
		return -1
	}
	return userID
}

func DeleteUser(id int) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	statement := `DELETE FROM Userdata WHERE UserID = $1`
	_, err = db.Exec(statement, id)
	if err != nil {
		return err
	}

	return nil
}

func ListUsers() ([]Userdata, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	userData := []Userdata{}
	statement := `SELECT * FROM Userdata`
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user Userdata
		err = rows.Scan(&user.ID, &user.Username, &user.Name, &user.Surname, &user.Description)
		if err != nil {
			fmt.Println("Scan: ", err)
			return nil, err
		}
		userData = append(userData, user)
	}
	defer rows.Close()

	return userData, nil
}

func UpdateUser(d Userdata) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userName := strings.ToLower(d.Username)
	if exists(userName) == -1 {
		return fmt.Errorf("user does not  exist")
	}

	statement := `UPDATE Userdata SET Name=$1, Surname=$2, Description=$3 WHERE UserID=$4`
	_, err = db.Exec(statement, d.Name, d.Surname, d.Description, d.ID)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Существует ли пользователь: ", exists("jane_smith"))
	fmt.Println("Добавление нового пользователя: ", AddUser(Userdata{
		Username:    "NEW_TEST",
		Name:        "NEW_TEST",
		Surname:     "NEW_TEST",
		Description: "NEW_TEST",
	}))
	//fmt.Println("Удаление пользователя: ", DeleteUser(1))
	fmt.Println("Обновление пользователя: ", UpdateUser(Userdata{
		ID:          7,
		Username:    "james_miller",
		Name:        "James_update",
		Surname:     "Miller_update",
		Description: "Laboris nisi ut aliquip ex ea commodo consequat",
	}))
	fmt.Println(ListUsers())
}
