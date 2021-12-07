import neogma from './appDatabase';

import { createBook, deleteBook, getBooks } from './entities/Books/BooksModels';
import { createAuthor, deleteAuthor, getAuthors } from './entities/Authors/AuthorsModels';
import { Country } from './entities/Authors/AuthorsEntity';
import { detachBookFromAuthor, getAuthorBooks, linkBookToAuthor } from './entities/Books/BooksRelation';

async function clearDb(): Promise<void> {
	const books = await getBooks();
	await Promise.all(books.map(async (book) => deleteBook(book.id)));

	const authors = await getAuthors();
	await Promise.all(authors.map(async (author) => deleteAuthor(author.id)));
}

async function main() {
	const author = await createAuthor({
		firstname: 'test',
		lastname: 'Test',
		pseudonym: 'pute',
		age: 13,
		country: Country.FRA,
	});

	const book1 = await createBook({
		title: 'test',
		summary: 'Test',
		rate: 0,
		price: 13,
		pages: 24,
	});

	const book2 = await createBook({
		title: 'test',
		summary: 'Test',
		rate: 0,
		price: 13,
		pages: 24,
	});

	const book3 = await createBook({
		title: 'test',
		summary: 'Test',
		rate: 0,
		price: 13,
		pages: 24,
	});

	await linkBookToAuthor(book1.id, author.id);
	await linkBookToAuthor(book2.id, author.id);
	await linkBookToAuthor(book3.id, author.id);

	console.log(await getAuthorBooks(author.id));

	await detachBookFromAuthor(book2.id, 'tt');

	console.log(await getAuthorBooks(author.id));
	await clearDb();
}

main();
