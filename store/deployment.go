package store

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Deployment struct {
	ID        int       `db:"id"`
	RepoURL   string    `db:"repo_url"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type DeploymentStore struct {
	db *sqlx.DB
}

func NewDeploymentStore(db *sqlx.DB) *DeploymentStore {
	return &DeploymentStore{db: db}
}

func (s *DeploymentStore) Create(repoURL string) (Deployment, error) {
	var d Deployment
	query := "INSERT INTO deployments (repo_url) VALUES ($1) RETURNING *"
	row := s.db.QueryRowx(query, repoURL)
	err := row.StructScan(&d)
	if err != nil {
		return Deployment{}, err
	}
	return d, nil
}
func (s *DeploymentStore) GetByID(id int) (Deployment, error) {

	var d Deployment

	query := "SELECT * FROM deployments WHERE id = $1"
	err := s.db.Get(&d, query, id)
	if err != nil {
		return Deployment{}, err
	}
	return d, nil
}
