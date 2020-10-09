package services

import (
	"time"

	"github.com/ibamaulana/golang-kreditplus/model"
	"github.com/jinzhu/gorm"
)

type PlayerServiceContract interface {
	Get() ([]*model.Player, error)
	GetMax() ([]*model.Player, error)
	Create(player *model.Player, tx *gorm.DB) (*model.Player, error)
}

type playerContractService struct {
	db *gorm.DB
}

func NewPlayerServiceContract(db *gorm.DB) PlayerServiceContract {
	return &playerContractService{db}
}

func (srv *playerContractService) Get() ([]*model.Player, error) {
	var user []*model.Player
	var err error

	err = srv.db.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *playerContractService) GetMax() ([]*model.Player, error) {
	var user []*model.Player
	var err error

	err = srv.db.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *playerContractService) Create(player *model.Player, tx *gorm.DB) (*model.Player, error) {
	row := new(model.Player)
	player.Created = time.Now()
	d := tx.Create(&player).Scan(&row)
	if d.Error != nil {
		return nil, d.Error
	}

	return row, nil
}

