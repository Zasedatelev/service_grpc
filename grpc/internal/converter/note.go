package converter

import (
	"github.com/Zasedatelev/service_grpc/internal/model"
	desc "github.com/Zasedatelev/service_grpc/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteFromService(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:       note.ID,
		Info:     ToNoteInfoFromService(note.Info),
		CreateAt: timestamppb.New(note.CreatedAt),
		UpdateAt: updatedAt,
	}
}

func ToNoteInfoFromDesc(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteInfoFromService(info model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
