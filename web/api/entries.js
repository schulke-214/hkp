import { HttpClient } from '~/models/http-client';

export class TaskEntryClient extends HttpClient {
	async all() {
		return HttpClient.transformResponse(await this.$axios.get('/entries/'));
	}

	async byTask(id) {
		return HttpClient.transformResponse(await this.$axios.get(`/tasks/${id}/entries/`));
	}

	async create(id, user) {
		return HttpClient.transformResponse(await this.$axios.post(`/tasks/${id}/entries/`, { user }));
	}

	async delete(taskId, entryId) {
		return HttpClient.transformResponse(await this.$axios.delete(`/tasks/${taskId}/entries/${entryId}`));
	}
}
