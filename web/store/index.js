export const actions = {
	async nuxtServerInit({ dispatch }) {
		await dispatch('entries/fetch', null, { root: true });
		await dispatch('tasks/fetch', null, { root: true });
	}
};
