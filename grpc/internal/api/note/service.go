package note

import (
	"github.com/Zasedatelev/service_grpc/internal/service"
	desc "github.com/Zasedatelev/service_grpc/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
