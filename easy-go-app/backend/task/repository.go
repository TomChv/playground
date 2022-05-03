package task

import (
	"context"

	"github.com/google/uuid"
	"izi-go.com/ent"
	"izi-go.com/ent/task"
)

type Repository struct {
	db *ent.Client
}

func NewRepo(db *ent.Client) *Repository {
	return &Repository{db: db}
}

// Create a new task
func (r *Repository) Create(ctx context.Context, data *CreateData) (*ent.Task, error) {
	return r.db.Task.Create().
		SetTitle(data.Title).
		SetDescription(data.Description).
		Save(ctx)
}

// List retrieve all task
func (r *Repository) List(ctx context.Context) ([]*ent.Task, error) {
	return r.db.Task.Query().All(ctx)
}

// Get retrieve a task by his identifier
func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	return r.db.Task.Get(ctx, id)
}

// Update data of a task
func (r *Repository) Update(ctx context.Context, id uuid.UUID, data *UpdateData) (*ent.Task, error) {
	uTask := r.db.Task.UpdateOneID(id)

	if data.Title != "" {
		uTask.SetTitle(data.Title)
	}

	if data.Description != "" {
		uTask.SetDescription(data.Description)
	}

	return uTask.Save(ctx)
}

// Delete remove a task
func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Task.DeleteOneID(id).Exec(ctx)
}

// Start update task status to 'in progress'
func (r *Repository) Start(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	return r.db.Task.UpdateOneID(id).SetStatus(task.StatusInProgress).Save(ctx)
}

// Done update task status to 'done'
func (r *Repository) Done(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
	return r.db.Task.UpdateOneID(id).SetStatus(task.StatusDone).Save(ctx)
}
