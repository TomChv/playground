import axios from 'axios';

import { NoteDto, NoteRo } from '../dto/apiDto';
import { API_URL } from '../AppConfig';

/**
 * Note API SDK
 *
 * - sendNote - Send note to APi
 * - getNotes - get All notes
 */
class Requester {
	/**
	 * Axios requester to communicate with the notes API
	 * @private
	 */
	private requester = axios.create({
		baseURL: API_URL,
		timeout: 1000
	});

	/**
	 * Save notes on API
	 *
	 * @method
	 * @param dto - NoteDto
	 * @return NoteRo - Note result
	 */
	async sendNote(dto: NoteDto) {
		return this.requester.post<NoteRo>('/notes', dto)
	}

	/**
	 * Get all notes
	 *
	 * @method
	 * @return An array of all notes
	 */
	async getNotes() {
		return this.requester.get<NoteRo[]>('/notes')
	}
}

export default new Requester();
