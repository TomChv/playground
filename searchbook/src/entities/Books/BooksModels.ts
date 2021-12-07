import { v4 as uuidv4 } from 'uuid';
import { QueryBuilder } from 'neogma';

import { Books, BooksInstance, BooksPropertiesI } from './BooksEntity';
import { queryRunner } from '../../appDbInstance';
import { Authors } from '../Authors/AuthorsEntity';

export function convertInstanceToProps(book: BooksInstance): BooksPropertiesI {
	return book.dataValues;
}

export async function createBook(book: Omit<BooksPropertiesI, 'id'>): Promise<BooksPropertiesI> {
	const id = uuidv4().toString();
	const newBook = await Books.createOne({ id, ...book });
	return convertInstanceToProps(newBook);
}

export async function getBooks(): Promise<BooksPropertiesI[]> {
	const books = await Books.findMany();
	return books.map((book) => convertInstanceToProps(book));
}

export async function getBook(id: string): Promise<BooksPropertiesI | null> {
	const book = await Books.findOne({
		where: {
			id,
		},
	});
	if (!book) {
		return null;
	}
	return convertInstanceToProps(book);
}

export async function updateBook(id: string, book: Partial<BooksPropertiesI>): Promise<BooksPropertiesI | null> {
	const updatedBook = await Books.update(book, {
		where: {
			id,
		},
		return: true,
	});
	if (updatedBook[0].length === 0) {
		return null;
	}
	return convertInstanceToProps(updatedBook[0][0]);
}

export async function deleteBook(id: string): Promise<boolean> {
	const nbBooksDeleted = await Books.delete({
		where: {
			id,
		},
		detach: true,
	});
	return nbBooksDeleted <= 0;
}