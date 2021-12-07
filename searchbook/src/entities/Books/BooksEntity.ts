import { ModelFactory, ModelRelatedNodesI, NeogmaInstance } from 'neogma';
import neogma from '../../appDatabase';

import { Authors, AuthorsPropertiesI, AuthorsInstance } from '../Authors/AuthorsEntity';

export type BooksPropertiesI = {
	id: string;
	title: string;
	summary: string;
	price: number;
	rate: number;
	pages: number;
};

export type BooksWithAuthorsPropertiesI = BooksPropertiesI & { author: AuthorsPropertiesI };

export type AuthorWithBooksPropertiesI = AuthorsPropertiesI & { books: BooksPropertiesI[] };

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface BooksRelationNodesI {
	WriteBy: ModelRelatedNodesI<typeof Authors, AuthorsInstance>;
}

export type BooksInstance = NeogmaInstance<BooksPropertiesI, BooksRelationNodesI>;

export const Books = ModelFactory<BooksPropertiesI, BooksRelationNodesI>(
	{
		label: 'Books',
		primaryKeyField: 'id',
		schema: {
			id: {
				type: 'string',
				required: true,
			},
			title: {
				type: 'string',
				required: true,
				minLength: 1,
			},
			summary: {
				type: 'string',
				required: true,
				minLength: 1,
			},
			price: {
				type: 'number',
				minimum: 0,
				required: true,
			},
			rate: {
				type: 'number',
				minimum: 0,
				maximum: 10,
				required: true,
			},
			pages: {
				type: 'number',
				minimum: 0,
				required: true,
			},
		},
		relationships: {
			WriteBy: {
				model: Authors,
				direction: 'out',
				name: 'WriteBy',
			},
		},
	},
	neogma,
);
