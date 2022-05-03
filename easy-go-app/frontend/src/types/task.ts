/**
 * // Task is the model entity for the Task schema.
 * type Task struct {
 * 	config `json:"-"`
 * 	// ID of the ent.
 * 	ID uuid.UUID `json:"id,omitempty"`
 * 	// Title holds the value of the "title" field.
 * 	Title string `json:"title,omitempty"`
 * 	// Description holds the value of the "description" field.
 * 	Description string `json:"description,omitempty"`
 * 	// CreatedAt holds the value of the "created_at" field.
 * 	CreatedAt time.Time `json:"created_at,omitempty"`
 * 	// UpdatedAt holds the value of the "updated_at" field.
 * 	UpdatedAt time.Time `json:"updated_at,omitempty"`
 * 	// Status holds the value of the "status" field.
 * 	Status task.Status `json:"status,omitempty"`
 * }
 */
export enum EStatus {

}

export interface Task {
	id: string;
	title: string;
	description: string;
	created_at: Date;
	updated_at: Date;
};