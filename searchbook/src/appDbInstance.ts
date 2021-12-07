import { QueryRunner } from 'neogma';

import neogma from './appDatabase';

export const queryRunner = new QueryRunner({
	driver: neogma.driver,
});
