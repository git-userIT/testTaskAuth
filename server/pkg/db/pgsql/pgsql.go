package pgsql

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	host   = os.Getenv("PGSQL_HOST")
	port   = os.Getenv("PGSQL_PORT")
	user   = os.Getenv("PGSQL_USER")
	dbname = os.Getenv("PGSQL_DBNAME")
	pswPg  = os.Getenv("PGSQL_PSW")
)

type User struct {
	Username string
	Password string
	Email    string
}

type DataUser struct {
	IDUser   int
	Username string
	Email    string
}

var mu sync.Mutex

func selDataUser(userData User) (DataUser, bool) {

	var us DataUser
	ctx := context.Background()

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pswPg, dbname)
	conn, err := pgx.Connect(ctx, connString)

	query := `select id_user, email from users where username = $1;`
	errr := conn.QueryRow(ctx, query, userData.Username).Scan(&us.IDUser, &us.Email)

	us.Username = userData.Username
	conn.Close(ctx)

	if err != nil || errr != nil {
		fmt.Println("Ошибка при запросе данных пользователя", err, errr)
		return us, false
	} else {
		return us, true
	}
}

func SelDataUser(userData User) (DataUser, bool) {
	mu.Lock()
	ud, res := selDataUser(userData)
	mu.Unlock()
	return ud, res
}

func chUserExist(userData User) bool {
	var check bool
	ctx := context.Background()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pswPg, dbname)
	conn, err := pgx.Connect(ctx, connStr)

	query := `SELECT user_exists($1, $2);`
	errr := conn.QueryRow(ctx, query, userData.Username, userData.Password).Scan(&check)
	conn.Close(ctx)

	if err != nil || errr != nil {
		fmt.Println("Ошибка при проверке пользователя в БД", err, errr)
		return check
	} else {
		return check
	}
}

func ChUserExist(u User) bool {
	mu.Lock()
	res := chUserExist(u)
	mu.Unlock()
	return res
}

func AddNewUser(u User) bool {
	mu.Lock()
	res := addNewUser(u)
	mu.Unlock()
	return res
}

func addNewUser(userData User) bool {
	var in int
	ctx := context.Background()

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pswPg, dbname)
	conn, err := pgx.Connect(ctx, connString)

	query := `INSERT INTO users (username, pass, email) VALUES ($1, $2, $3) RETURNING id_user;`
	row := conn.QueryRow(ctx, query, userData.Username, userData.Password, userData.Email).Scan(&in)

	conn.Close(ctx)

	if err != nil || in == 0 {
		fmt.Println("Ошибка при добавлении пользователя в таблицу", err, row)
		return false
	} else {
		return true
	}
}
