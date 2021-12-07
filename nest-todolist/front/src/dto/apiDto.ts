export interface NoteDto {
	title: string;
	content: string;
}

export interface NoteRo extends NoteDto {
	publishAt: Date;
}
