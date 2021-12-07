import { Neogma } from 'neogma';

import { dbConfig } from './config';

const neogma = new Neogma(
	{
		url: `bolt://${dbConfig.host}:${dbConfig.port}`,
		username: dbConfig.user,
		password: dbConfig.password,
	},
	{
		//logger: console.debug,
	},
);

export default neogma;
