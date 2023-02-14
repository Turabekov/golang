package controller

import (
	"app/models"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id int, err error) {

	id, err = c.store.User.Create(req)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Controller) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {

	users, err := c.store.User.GetList(req)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	return users, nil
}
