package note

import (
	"github.com/Zasedatelev/service_grpc/internal/repository"
	"github.com/Zasedatelev/service_grpc/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
	}
}
