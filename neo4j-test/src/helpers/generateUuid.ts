import { v4 } from 'uuid';

export default function generateUuid(): string {
	return v4().toString();
}