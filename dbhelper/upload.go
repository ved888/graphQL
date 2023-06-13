package dbhelper

import (
	"github.com/jmoiron/sqlx"
	"grapgQL/graph/model"
)

type DAO struct {
	DB *sqlx.DB
}

func (d DAO) UploadImage(user model.Image) (*string, error) {
	// SQL=language
	SQL := `INSERT INTO upload(
                      bucket_name,
                      path,
                      user_id
                      )    
                VALUES($1,$2,$3)
                returning id::uuid
                                  `
	arg := []interface{}{
		user.BucketName,
		user.Path,
		user.UserId,
	}
	var uploadId string
	err := d.DB.QueryRowx(SQL, arg...).Scan(&uploadId)
	if err != nil {
		return nil, err
	}
	return &uploadId, nil
}
