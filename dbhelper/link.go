package dbhelper

import (
	"database/sql"
	"grapgQL/graph/model"
	"log"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    string
}

func (d DAO) CreateLink(links *model.Link) (*string, error) {

	// SQL=language
	SQL := `Insert into links(
                            title,
                            address,
                            user_id
                            )
                      values ($1,$2,$3)
                      returning id
                                  `
	args := []interface{}{
		links.Title,
		links.Address,
		links.User,
	}
	var linkId string
	err := d.DB.QueryRow(SQL, args...).Scan(&linkId)
	if err != nil {
		log.Println(err)
	}
	return &linkId, nil
}

func (d DAO) GetLinkById(link model.Link, linkId string) (*model.Link, error) {
	// SQL=language
	SQL := `SELECT
                title,
                address,
                user_id
          FROM links
          where id=$1::uuid
                           `

	err := d.DB.Get(&link, SQL, linkId)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &link, err
}

func (d DAO) GetAllLink(links []*model.Link) ([]*model.Link, error) {
	// SQL=language
	SQL := `SELECT
                id,
                title,
                address,
                user_id
           FROM links
                  `

	err := d.DB.Select(&links, SQL)
	if err != nil {
		return nil, err
	}
	return links, nil
}
