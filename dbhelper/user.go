package dbhelper

import (
	"database/sql"
	"grapgQL/graph/model"
	"log"
)

type User struct {
	ID        string `json:"Id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Dob       string `json:"dob" db:"dob"`
	Phone     string `json:"phone" db:"phone"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

// CreateUser create the entry for the user with different given fields
func (d DAO) CreateUser(users *model.User) (*string, error) {
	// SQL=language
	SQL := `Insert into users(
                  first_name,
                  last_name,
                  dob,
                  phone,
                  email,
                  password
                  )
            VALUES ($1,$2,$3,$4,$5,$6)
             returning id::uuid`

	arg := []interface{}{
		users.FirstName,
		users.LastName,
		users.Dob,
		users.Phone,
		users.Email,
		users.Password,
	}
	var id string
	err := d.DB.QueryRow(SQL, arg...).Scan(&id)
	if err != nil {
		log.Println("failed to insert data in database", err)
	}
	return &id, nil
}

// UpdateUser update user entry for given fields.
func (d DAO) UpdateUser(users *model.User, userId string) error {
	// SQL=language
	SQL := `update users 
                  set
                     first_name=COALESCE($1,first_name),
                     last_name=COALESCE($2,last_name),
                     dob=COALESCE($3,dob),
                     phone=COALESCE($4,phone),
                     email=COALESCE($5,email),
                     updated_at=now()
                where id=$6::uuid
                and archived_at IS NULL
                      `

	arg := []interface{}{
		users.FirstName,
		users.LastName,
		users.Dob,
		users.Phone,
		users.Email,
		userId,
	}
	_, err := d.DB.Exec(SQL, arg...)
	return err

}

// GetUserById get user entry by user id if it is exists
func (d DAO) GetUserById(user model.User, userId string) (*model.User, error) {
	// SQL:=language
	SQL := `select
		        first_name,
			    last_name,
			    dob,
		    	phone,
			    email
		   from users
		   where id=$1::uuid`

	//var users model.User
	err := d.DB.Get(&user, SQL, &userId)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// GetUserId get user entry by user id if it is exists
func (d DAO) GetUserId(userId string) (*model.User, error) {
	// SQL:=language
	SQL := `select
                 id,  
		        first_name,
			    last_name,
			    dob,
		    	phone,
			    email
		   from users
		   where id=$1::uuid`

	var user model.User
	//var users model.User
	err := d.DB.Get(&user, SQL, &userId)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// GetUserByEmail get the user profile entry if it exists
func (d DAO) GetUserByEmail(email string) (*model.User, error) {
	// SQL=language
	SQL := `select
                id,
		        first_name,
			    last_name,
			    dob,
		    	phone,
			    email,
			    password
		   from users
		   where email=$1`
	var user model.User
	err := d.DB.Get(&user, SQL, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, err
}

// GetAllUser get all user entry if it exists
func (d DAO) GetAllUser(users []*model.User) ([]*model.User, error) {
	// SQL=language
	SQL := `select 
                   id,
                   first_name,
                   last_name,
                   dob,
                   phone,
                   email
             from users
`
	//var user []*model.User
	err := d.DB.Select(&users, SQL)
	if err != nil {
		return nil, err
	}
	return users, err

}

// DeleteUser delete user entry by id if it exists
func (d DAO) DeleteUser(userId string) error {
	//SQL=language
	SQL := `update users 
              set 
                  archived_at=now()
              where id=$1::uuid
                               `

	_, err := d.DB.Exec(SQL, userId)
	if err != nil {
		return err
	}
	return err
}
