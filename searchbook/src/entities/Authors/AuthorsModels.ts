import { v4 as uuidv4 } from 'uuid';

import { Authors, AuthorsInstance, AuthorsPropertiesI } from './AuthorsEntity';

export function convertInstanceToProps(authors: AuthorsInstance): AuthorsPropertiesI {
	return authors.dataValues;
}

export async function createAuthor(authors: Omit<AuthorsPropertiesI, 'id'>): Promise<AuthorsPropertiesI> {
	const id = uuidv4().toString();
	const newAuthors = await Authors.createOne({ id, ...authors });
	return convertInstanceToProps(newAuthors);
}

export async function getAuthors(): Promise<AuthorsPropertiesI[]> {
	const authors = await Authors.findMany();
	return authors.map((author) => convertInstanceToProps(author));
}

export async function getAuthor(id: string): Promise<AuthorsPropertiesI | null> {
	const authors = await Authors.findOne({
		where: {
			id,
		},
	});
	if (!authors) {
		return null;
	}
	return convertInstanceToProps(authors);
}

export async function updateAuthors(id: string, authors: Partial<AuthorsPropertiesI>): Promise<AuthorsPropertiesI | null> {
	const updatedAuthors = await Authors.update(authors, {
		where: {
			id,
		},
		return: true,
	});
	if (updatedAuthors[0].length === 0) {
		return null;
	}
	return convertInstanceToProps(updatedAuthors[0][0]);
}

export async function deleteAuthor(id: string): Promise<boolean> {
	const nbAuthorsDeleted = await Authors.delete({
		where: {
			id,
		},
		detach: true,
	});
	return nbAuthorsDeleted <= 0;
}
