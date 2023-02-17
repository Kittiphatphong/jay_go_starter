package services

import "go_starter/repositories"

type Service interface{}

type service struct {
	repository repositories.Repository
}

func NewService(
	repository repositories.Repository,
	//repo
) Service {
	return &service{
		repository: repository,
		//repo
	}

}
