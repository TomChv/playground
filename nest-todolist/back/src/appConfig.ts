import { get } from 'env-var';

const env = (name: string, required = true) => get(name).required(required);

export const dbConfig = {
	username: env('DB_USER').asString(),
	password: env('DB_PASS').asString(),
	host: env('DB_HOST').asString(),
	port: env('DB_PORT').asPortNumber(),
	database: env('DB_NAME').asString(),
};

export const apiConfig = {
	port: env('API_PORT').asPortNumber(),
	host: env('API_HOST').asString(),
};
