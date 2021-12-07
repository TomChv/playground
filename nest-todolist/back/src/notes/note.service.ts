import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';

import { NoteRepository } from './note.repository';
import { Note } from './note.entity';
import { NoteDto } from './note.dto';

@Injectable()
export class NoteService {
	constructor(
		@InjectRepository(NoteRepository)
		private readonly noteRepository: NoteRepository,
	) {}

	async getNotes(): Promise<Note[]> {
		return this.noteRepository.getNotes();
	}

	async getNote(title: string): Promise<Note> {
		const note = await this.noteRepository.getNote(title);

		if (!note) {
			throw new NotFoundException('Note not found');
		}

		return note;
	}

	async createNote(noteDto: NoteDto): Promise<Note> {
		const note = await this.noteRepository.getNote(noteDto.title);

		if (note) {
			throw new NotFoundException('Note already exist');
		}
		return this.noteRepository.createNote(noteDto);
	}

	async updateNote(noteDto: Partial<NoteDto>): Promise<Note> {
		const note = await this.noteRepository.getNote(noteDto.title);

		if (!note) {
			throw new NotFoundException('Note already exist');
		}
		return this.noteRepository.updateNote(noteDto);
	}

	async deleteNote(title: string): Promise<string> {
		const note = await this.noteRepository.getNote(title);

		if (!note) {
			throw new NotFoundException(`Note '${title}' doesn't exist`);
		}
		await this.noteRepository.deleteNote(title);
		return `Note '${title}' has been delete`;
	}
}
