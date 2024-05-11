package app

import (
	"context"
	"log"

	"github.com/Zasedatelev/service_grpc/internal/api/note"
	"github.com/Zasedatelev/service_grpc/internal/closer"
	"github.com/Zasedatelev/service_grpc/internal/config"
	"github.com/Zasedatelev/service_grpc/internal/repository"
	noteRepository "github.com/Zasedatelev/service_grpc/internal/repository/note"
	"github.com/Zasedatelev/service_grpc/internal/service"
	noteService "github.com/Zasedatelev/service_grpc/internal/service/note"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool *pgxpool.Pool

	noteRepository repository.NoteRepository
	noteService    service.NoteService

	noteImpl *note.Implementation
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig != nil {
		ctf, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = ctf
	}

	return s.pgConfig

}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig != nil {
		grpc, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.grpcConfig = grpc
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetPgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool

}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepository.NewRepository(s.GetPgPool(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}

func (s *serviceProvider) GetNoteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.GetNoteService(ctx))
	}

	return s.noteImpl
}
