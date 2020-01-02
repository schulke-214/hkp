import { HttpClient } from '~/models/http-client';

export class TaskClient extends HttpClient {
	async all() {
		return HttpClient.transformResponse(await this.$axios.get('/tasks/'));
	}

	async create(name) {
		return HttpClient.transformResponse(await this.$axios.post('/tasks/', { name }));
	}

	async update() {
		console.warn('Update Task');
	}

	async delete(id) {
		return HttpClient.transformResponse(await this.$axios.delete(`/tasks/${id}`));
	}
}
