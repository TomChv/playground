import {
	Body,
	Controller,
	Delete,
	Get,
	Param,
	Post,
	Put,
} from '@nestjs/common';

import { NoteDto } from './note.dto';
import { Note } from './note.entity';

import { NoteService } from './note.service';

@Controller('notes')
export default class NoteController {
	constructor(private readonly noteService: NoteService) {}

	@Get()
	async getNotes(): Promise<Note[]> {
		return this.noteService.getNotes();
	}

	@Get(':title')
	async getNote(@Param('title') title: string): Promise<Note> {
		return this.noteService.getNote(title);
	}

	@Post()
	async postNote(@Body() noteDto: NoteDto): Promise<Note> {
		return this.noteService.createNote(noteDto);
	}

	@Put()
	async putNote(@Body() noteDto: Partial<NoteDto>): Promise<Note> {
		return this.noteService.updateNote(noteDto);
	}

	@Delete(':title')
	async deleteNote(@Param('title') title: string): Promise<string> {
		return this.noteService.deleteNote(title);
	}
}
