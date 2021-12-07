import { get } from 'env-var';

const env = (key: string, required = true) => get(key).required(required);

export const dbConfig = {
	user: env('NEO4J_USER').asString(),
	password: env('NEO4J_PASSWORD').asString(),
	host: env('NEO4J_HOST').asString(),
	port: env('NEO4J_PORT').asPortNumber(),
};
