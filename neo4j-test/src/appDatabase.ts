import { Neogma, neo4jDriver } from 'neogma';

const neogma = new Neogma({
	url: 'bolt://localhost:7687',
	username: 'neo4j',
	password: 'vasektommy',
});

export default neogma;
