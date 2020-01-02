export class HttpClient {
	constructor($axios) {
		this.$axios = $axios;
	}

	static transformResponse(response) {
		return response.data;
	}
}
