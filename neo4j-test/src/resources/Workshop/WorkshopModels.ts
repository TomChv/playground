import { NeogmaInstance } from 'neogma';

import { Workshop, WorkshopInstance, WorkshopPropertiesI } from './WorkshopTypes';
import generateUuid from '../../helpers/generateUuid';

export function convertInstanceToWorkshop(workshop: WorkshopInstance): WorkshopPropertiesI {
	return workshop.dataValues;
}

export function convertInstanceToWorkshops(workshops: Array<WorkshopInstance>): Array<WorkshopPropertiesI> {
	return workshops.map((workshop) => convertInstanceToWorkshop(workshop));
}

export async function getWorkshops(): Promise<WorkshopPropertiesI[]> {
	const workshops = await Workshop.findMany({
		session: null,
	});
	return convertInstanceToWorkshops(workshops);
}

export async function getWorkshop(id: string): Promise<WorkshopPropertiesI | null> {
	const workshop = await Workshop.findOne({
		where: {
			id,
		},
		session: null,
	});
	if (workshop == null) {
		return null;
	}
	return convertInstanceToWorkshop(workshop);
}

export async function createWorkshop(workshop: Omit<WorkshopPropertiesI, 'id'>): Promise<WorkshopPropertiesI> {
	const newWorkshop = await Workshop.createOne(
		{ ...workshop, id: generateUuid() },
		{
			session: null,
		},
	);
	return convertInstanceToWorkshop(newWorkshop);
}

export async function updateWorkshop(
	id: string,
	workshop: Partial<WorkshopPropertiesI>,
): Promise<WorkshopPropertiesI | null> {
	const user = await Workshop.update(workshop, {
		where: {
			id,
		},
		return: true,
	});
	if (user[0].length !== 0) {
		return convertInstanceToWorkshop(user[0][0]);
	}
	return null;
}

export async function deleteWorkshop(id: string): Promise<number> {
	return Workshop.delete({
		where: {
			id,
		},
		detach: true,
	});
}
