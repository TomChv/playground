// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"izi-go.com/ent/schema"
	"izi-go.com/ent/task"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescTitle is the schema descriptor for title field.
	taskDescTitle := taskFields[1].Descriptor()
	// task.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	task.TitleValidator = taskDescTitle.Validators[0].(func(string) error)
	// taskDescDescription is the schema descriptor for description field.
	taskDescDescription := taskFields[2].Descriptor()
	// task.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	task.DescriptionValidator = taskDescDescription.Validators[0].(func(string) error)
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[3].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescUpdatedAt is the schema descriptor for updated_at field.
	taskDescUpdatedAt := taskFields[4].Descriptor()
	// task.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	task.DefaultUpdatedAt = taskDescUpdatedAt.Default.(func() time.Time)
	// task.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	task.UpdateDefaultUpdatedAt = taskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskFields[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() uuid.UUID)
}