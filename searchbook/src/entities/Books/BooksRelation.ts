import { QueryBuilder } from 'neogma';

import { queryRunner } from '../../appDbInstance';

import { AuthorWithBooksPropertiesI, Books } from './BooksEntity';
import { Authors } from '../Authors/AuthorsEntity';

export async function getAuthorBooks(id: string): Promise<AuthorWithBooksPropertiesI | null> {
	const result = await new QueryBuilder()
		.match({
			related: [
				{
					identifier: 'books',
					model: Books,
				},
				{
					direction: 'out',
					name: 'WriteBy',
				},
				{
					identifier: 'authors',
					model: Authors,
					where: { id },
				},
			],
		})
		.return('books, authors')
		.run(queryRunner);

	if (result.records.length === 0) {
		return null;
	}

	const nodes = result.records.map((record) => record.toObject());
	return {
		...nodes[0].authors.properties,
		books: nodes.map((node) => node.books.properties),
	};
}

export async function linkBookToAuthor(bookId: string, authorId: string): Promise<AuthorWithBooksPropertiesI | null> {
	const nbBooksUpdated = await Books.relateTo({
		alias: 'WriteBy',
		where: {
			source: { id: bookId },
			target: { id: authorId },
		},
	});

	if (nbBooksUpdated !== 1) {
		return null;
	}
	return getAuthorBooks(authorId);
}

export async function detachBookFromAuthor(bookId: string, authorId: string): Promise<boolean> {
	const result = await new QueryBuilder()
		.match({
			related: [
				{
					identifier: 'book',
					model: Books,
					where: { id: bookId },
				},
				{
					direction: 'out',
					name: 'WriteBy',
					identifier: 'relation',
				},
				{
					identifier: 'author',
					model: Authors,
					where: { id: authorId },
				},
			],
		})
		.delete('relation')
		.return('book')
		.run(queryRunner);

	if (result.records.length === 0) {
		return false;
	}
	return true;
}
