package repository

import (
	"final-project/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *entity.User) error
	GetAllUser() ([]entity.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) GetAllUser() ([]entity.User, error) {
	users := []entity.User{}
	err := u.db.Find(&users)
	fmt.Println(err)
	return err
}

func (u *userRepo) RegisterUser(user *entity.User) error {
	err := u.db.Create(&user).Error
	return err
}

// func RegisterUser(w http.ResponseWriter, r *http.Request) {
// 	var user entity.User

// 	db := config.Connect()
// 	defer db.Close()

// 	w.Header().Set("Content-Type", "application/json")
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&user)
// 	if err != nil {
// 		userResponseReturn(501, "Error Decoding JSON", []entity.User{}, w)
// 	}

// 	bs, _ := helper.HashPassword(user.Password)

// 	em, _ := helper.ValidMailAddress(user.Email)
// 	// fmt.Println(em)
// 	rows, err := db.Exec("INSERT INTO user(username,email,password,age,created_at) VALUES (?,?,?,?,?)", user.Username, em, bs, user.Age, time.Now())
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	insertUser, _ := rows.LastInsertId()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// fmt.Println(insertUser)
// 	user.Id = int(insertUser)
// 	user.CreatedAt = time.Now()
// 	userResponseReturn(201, "", []entity.User{user}, w)
// }

// func LoginUser(w http.ResponseWriter, r *http.Request) {
// 	var creds entity.Credentials
// 	var user entity.User
// 	var jwtKey = []byte("my_secret_key")
// 	db := config.Connect()
// 	defer db.Close()

// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	rows, err := db.Query("SELECT email,password FROM user WHERE email=?", creds.Email)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&user.Email, &user.Password)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	checkPass := helper.CheckPassword(creds.Password, user.Password)
// 	if !checkPass {
// 		userResponseReturn(500, "Invalid Password", "", w)
// 	}
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims := &entity.Claims{
// 		Email: creds.Email,
// 		StandardClaims: jwt.StandardClaims{
// 			// In JWT, the expiry time is expressed as unix milliseconds
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtKey)
// 	if err != nil {
// 		// If there is an error in creating the JWT return an internal server error
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	userResponseReturn(200, "", tokenString, w)

// }

// func userResponseReturn(status int, message string, data interface{}, w http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	json.NewEncoder(w).Encode(entity.Response{Status: status, Message: message, Data: data})
// }
