import { ApiProperty } from '@nestjs/swagger';
import { IsString, Length } from 'class-validator';

export class NoteDto {
	@ApiProperty()
	@Length(1, 25)
	@IsString()
	title: string;

	@ApiProperty()
	@IsString()
	content: string;
}
