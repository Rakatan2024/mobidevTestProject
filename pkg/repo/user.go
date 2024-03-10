package repo

import (
	"database/sql"
	"log"
	"mobidevtestProject/pkg/entity"
)

type UserInterface interface {
	CreateUser(*entity.User) error
	GetUserByID(int64) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)
	UpdateUser(*entity.User) error
}

func (r *Repository) CreateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Printf("Error at the stage of preparing data to Insert CreateUser(repo): %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error at the stage of closing stmt CreateUser(repo): %s\n", err.Error())
		}
	}(stmt)

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.EncryptedPass,
		user.AvatarLink,
		user.Gender,
		user.Age,
		user.PhoneNum,
		user.ResidenceCity,
		user.Description,
		user.Email,
		user.IsEmailVerified,
	)
	if err != nil {
		log.Printf("Error at the stage of data Inserting CreateUser(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByID(ID int64) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id, first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified FROM users WHERE id = ?")
	if err != nil {
		log.Printf("Error while preparing data to GetUserByID(repo) from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error at the stage of closing stmt GetUserByID(repo): %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)
	if err != nil {
		log.Printf("Error while querying row and scanning user to GetUserByID(repo): %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	stmt, err := r.db.Prepare("SELECT id, first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified FROM users WHERE email = ?")
	if err != nil {
		log.Printf("Error while preparing data to GetUserByEmail(repo) from user table: %s\n", err.Error())
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error at the stage of closing stmt in GetUserByEmail(repo): %s\n", err.Error())
		}
	}(stmt)
	user := &entity.User{}
	err = stmt.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)
	if err != nil {
		log.Printf("Error while querying row and scanning GetUserByEmail(repo): %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`
    UPDATE users
    SET 
        first_name = ?,
        last_name = ?,
        password = ?,
        avatar_link = ?,
        gender = ?,
        age = ?,
        phone_number = ?,
        city_of_residence = ?,
        description = ?,
        email = ?,
        is_email_verified = ?
    WHERE id = ?
`)
	if err != nil {
		log.Printf("Error while preparing data to UpdateUser(repo) by id: %s\n", err.Error())
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error at the stage of closing stmt in UpdateUser(repo): %s\n", err.Error())
		}
	}(stmt)

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.EncryptedPass,
		user.AvatarLink,
		user.Gender,
		user.Age,
		user.PhoneNum,
		user.ResidenceCity,
		user.Description,
		user.Email,
		user.IsEmailVerified,
		user.ID,
	)
	if err != nil {
		log.Printf("Error while executing UpdateUser(repo) by id: %s\n", err.Error())
		return err
	}
	return nil
}
