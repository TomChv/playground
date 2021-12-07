import { get } from 'env-var';

const env = (name: string, required= true) => get(name).required(required);

export const API_URL = env('REACT_APP_API_URL').asString();
