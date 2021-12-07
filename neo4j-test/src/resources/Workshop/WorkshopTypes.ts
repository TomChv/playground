import { ModelFactory, NeogmaInstance } from 'neogma';

import neogma from '../../appDatabase';
import generateUuid from '../../helpers/generateUuid';

export type WorkshopPropertiesI = {
	id: string;
	email: string;
	firstname: string;
	lastname: string;
	password: string;
};

// eslint-disable-next-line @typescript-eslint/no-empty-interface
export interface WorkshopsRelatedNodesI {}

export type WorkshopInstance = NeogmaInstance<WorkshopPropertiesI, WorkshopsRelatedNodesI>;

export const Workshop = ModelFactory<WorkshopPropertiesI, WorkshopsRelatedNodesI>(
	{
		label: 'Workshop',
		primaryKeyField: 'email',
		schema: {
			id: {
				type: 'string',
				required: true,
				default: generateUuid(),
			},
			email: {
				type: 'string',
				required: true,
			},
			firstname: {
				type: 'string',
				required: true,
			},
			lastname: {
				type: 'string',
				required: true,
			},
			password: {
				type: 'string',
				required: true,
			},
		},
	},
	neogma,
);
