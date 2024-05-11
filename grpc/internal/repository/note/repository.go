package note

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Zasedatelev/service_grpc/internal/model"
	"github.com/Zasedatelev/service_grpc/internal/repository"
	"github.com/Zasedatelev/service_grpc/internal/repository/note/converter"

	modelRepo "github.com/Zasedatelev/service_grpc/internal/repository/note/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName = "note"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.NoteRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.NoteInfo) (int64, error) {

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		Values(info.Title, info.Content).
		Suffix("RETURNING id")

	quary, args, err := builder.ToSql()

	if err != nil {
		return 0, err
	}

	var id int64

	err = r.db.QueryRow(ctx, quary, args...).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Note, error) {

	builder := sq.Select(idColumn, titleColumn, contentColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	quary, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	var note modelRepo.Note
	err = r.db.QueryRow(ctx, quary, args...).Scan(&note.ID, &note.Info.Title, &note.Info.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToNoteFromRepo(&note), nil
}
