import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';

import { dbConfig } from '../appConfig';

import NoteController from './note.controller';
import { Note } from './note.entity';
import { NoteRepository } from './note.repository';
import { NoteService } from './note.service';

@Module({
	imports: [
		TypeOrmModule.forRoot({
			type: 'postgres',
			...dbConfig,
			entities: [Note],
			synchronize: true,
		}),
		TypeOrmModule.forFeature([Note, NoteRepository]),
	],
	controllers: [NoteController],
	providers: [NoteService],
})
export default class NoteModule {}
