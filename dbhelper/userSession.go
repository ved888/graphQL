package dbhelper

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"grapgQL/database"
	"grapgQL/graph/model"
	"grapgQL/utils"
	"time"
)

type UserSession struct {
	UserID     uuid.UUID `json:"userId"`
	SessionKey string    `json:"sessionKey"`
	ExpiryAt   time.Time `json:"expiryAt"`
}

const expiryTime = time.Hour * 12

func (us UserSession) CreateUserSession(db sqlx.Ext, userSession *model.UserSession) (string, error) {
	SQL := `INSERT INTO user_session (
		user_id,
		session_key,
		expiry_at
	) VALUES ($1, $2, $3)`

	// Convert UUID to a string
	uuidString := userSession.UserID.String()

	//token := utils.HashString(time.Now().String()) + userSession.UserID
	token := utils.HashString(time.Now().String()) + uuidString

	userSession.SessionKey = token
	arg := []interface{}{
		uuidString,
		token,
		userSession.ExpiryAt.Add(expiryTime),
	}

	_, err := db.Exec(SQL, arg...)
	if err != nil {
		logrus.Println("CreateUserSession: failed to generate token", err)
	}
	return token, err
}

func VerifySession(sessionKey *string) (*model.User, error) {
	// SQL=language
	SQL := `SELECT 
                  u.id,
                  u.first_name,
                  u.last_name,
                  u.dob,
                  u.phone,
                  u.email,
                  u.password
             From user_session us
             join users u ON u.id=us.user_id
                WHERE us.session_key = $1
				AND us.archived_at IS NULL
 	            AND u.archived_at IS NULL
 	            AND us.expiry_at > now()
                                          `

	var user model.User

	err := database.DB.Get(&user, SQL, sessionKey)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func (d DAO) ValidateSession(sessionKey *string) (*model.User, string, error) {
	// SQL query
	SQL := `SELECT 
		u.id,
		u.first_name,
		u.last_name,
		u.dob,
		u.phone,
		u.email,
		u.password
	FROM user_session us
	JOIN users u ON u.id = us.user_id
	WHERE us.session_key = $1
		AND us.archived_at IS NULL
		AND u.archived_at IS NULL
		AND us.expiry_at > now()`

	var user model.User
	var userID string

	err := d.DB.Get(&user, SQL, sessionKey)
	if err != nil && err != sql.ErrNoRows {
		return nil, "", err
	}
	if err == sql.ErrNoRows {
		return nil, "", nil
	}

	//userID = user.ID.String()

	return &user, userID, nil
}
