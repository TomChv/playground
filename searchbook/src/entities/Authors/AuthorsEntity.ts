import { ModelFactory, NeogmaInstance } from 'neogma';
import neogma from '../../appDatabase';

export enum Country {
	FRA = 'France',
	UK = 'United Kingdom',
	ITA = 'Italy',
	RUS = 'Russia',
	US = 'United State',
	GER = 'Germany',
}

export type AuthorsPropertiesI = {
	id: string;
	firstname: string;
	lastname: string;
	pseudonym: string;
	age: number;
	country: Country;
};

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface AuthorsRelationNodesI {}

export type AuthorsInstance = NeogmaInstance<AuthorsPropertiesI, AuthorsRelationNodesI>;

export const Authors = ModelFactory<AuthorsPropertiesI, AuthorsRelationNodesI>(
	{
		label: 'Authors',
		primaryKeyField: 'id',
		schema: {
			id: {
				type: 'string',
				required: true,
			},
			firstname: {
				type: 'string',
				required: true,
				minLength: 1,
			},
			lastname: {
				type: 'string',
				required: true,
				minLength: 1,
			},
			pseudonym: {
				type: 'string',
				required: true,
				minLength: 1,
			},
			age: {
				type: 'number',
				minimum: 0,
				maximum: 100,
				required: true,
			},
			country: {
				type: 'string',
				required: true,
			},
		},
	},
	neogma,
);
