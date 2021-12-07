import { WorkshopPropertiesI } from './resources/Workshop/WorkshopTypes';

import {
	createWorkshop,
	getWorkshop,
	getWorkshops,
	deleteWorkshop,
	updateWorkshop,
} from './resources/Workshop/WorkshopModels';

async function main() {
	const user = await createWorkshop({
		firstname: 'Tom',
		lastname: 'Chauveau',
		email: 'tom.chauveau@gmail.com',
		password: 'Chuuut',
	});

	console.log('Create Get');
	let users = await getWorkshops();
	console.log(users);

	// await Promise.all(users.map(async (elem) => deleteWorkshop(elem.id)));

	console.log('getOne');
	let user1 = await getWorkshop(users[0].id);
	console.log(user1);

	if (!user1) {
		return;
	}
	console.log('update');
	user1 = await updateWorkshop(user1.id, { email: 'Tom ' });

	users = await getWorkshops();
	console.log(users);

	if (!user1) {
		console.log(user1);
		return;
	}

	console.log('Delete');
	await deleteWorkshop(user1.id);

	users = await getWorkshops();
	console.log(users);

	await Promise.all(
		users.map(async (us) => {
			await deleteWorkshop(us.id);
		}),
	);
	users = await getWorkshops();
	console.log(users);
}

main();
