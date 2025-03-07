package repo

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/rarebek/web-playground/models"
	"golang.org/x/crypto/bcrypt"
)

type Repo struct {
	DB      *sql.DB
	builder sq.StatementBuilderType
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *Repo) InsertUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query, args, err := r.builder.Insert("users").
		Columns("id", "email", "password").
		Values(uuid.NewString(), user.Username, hashedPassword).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
