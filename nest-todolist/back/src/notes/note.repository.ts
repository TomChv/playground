import { DeleteResult, EntityRepository, Repository } from 'typeorm';

import { Note } from './note.entity';
import { NoteDto } from './note.dto';

@EntityRepository(Note)
export class NoteRepository extends Repository<Note> {
	async createNote(note: NoteDto): Promise<Note> {
		return this.save(note);
	}

	async getNotes(): Promise<Note[]> {
		return this.find();
	}

	async getNote(title: string): Promise<Note | undefined> {
		return this.findOne({ title });
	}

	async updateNote(noteDto: Partial<NoteDto>): Promise<Note> {
		await this.update({ title: noteDto.title }, noteDto);

		return this.getNote(noteDto.title);
	}

	async deleteNote(title: string): Promise<DeleteResult> {
		return this.delete({ title });
	}
}
