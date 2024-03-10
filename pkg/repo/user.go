package repo

import (
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
	result, err := r.db.Exec(`
    INSERT INTO users (first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		user.FirstName, user.LastName, user.EncryptedPass, user.AvatarLink, user.Gender, user.Age, user.PhoneNum, user.ResidenceCity, user.Description, user.Email, user.IsEmailVerified,
	)
	if err != nil {
		log.Printf("Error at the stage of data Inserting CreateUser(repo): %s\n", err.Error())
		return err
	}
	user.ID, err = result.LastInsertId()
	if err != nil {
		log.Printf("Error at the stage of data Inserting userID(repo): %s\n", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByID(ID int64) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.QueryRow(`
    SELECT id, first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified
    FROM users
    WHERE id = ?`,
		ID,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)

	if err != nil {
		log.Printf("Error while querying row and scanning user to GetUserByID(repo): %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.QueryRow(`
    SELECT id, first_name, last_name, password, avatar_link, gender, age, phone_number, city_of_residence, description, email, is_email_verified
    FROM users
    WHERE email = ?`,
		email,
	).Scan(&user.ID, &user.FirstName, &user.LastName, &user.EncryptedPass, &user.AvatarLink, &user.Gender, &user.Age, &user.PhoneNum, &user.ResidenceCity, &user.Description, &user.Email, &user.IsEmailVerified)

	if err != nil {
		log.Printf("Error while querying row and scanning GetUserByEmail(repo): %s\n", err.Error())
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateUser(user *entity.User) error {
	_, err := r.db.Exec(`
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
    WHERE id = ?`,
		user.FirstName, user.LastName, user.EncryptedPass, user.AvatarLink, user.Gender, user.Age, user.PhoneNum, user.ResidenceCity, user.Description, user.Email, user.IsEmailVerified, user.ID,
	)
	if err != nil {
		log.Printf("Error while executing UpdateUser(repo) by id: %s\n", err.Error())
		return err
	}
	return nil
}
