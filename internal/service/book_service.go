package service

import (
	"errors"
	"proyectoapi/internal/model"
	"proyectoapi/internal/store"
)

type Logger interface {
	Log(msg, error string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store:  s,
		logger: nil,
	}
}

func (s *Service) ObtieneTodosLibros() ([]*model.Libro, error) {
	if s.logger != nil {
		s.logger.Log("obteniendo libros", "")
	}
	libros, err := s.store.GetAll()
	if err != nil {
		if s.logger != nil {
			s.logger.Log("el error es %v\n", err.Error())
		}
		return nil, err
	}

	return libros, nil
}

func (s *Service) ObtieneLibroID(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}

func (s *Service) CrearLibro(libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("el titulo es necesario")
	}
	return s.store.Create(&libro)
}

func (s *Service) ActualizaLibro(id int, libro model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("el titulo es necesario")
	}
	return s.store.Update(id, &libro)
}

func (s *Service) BorrarLibro(id int) error {
	return s.store.Delete(id)
}
