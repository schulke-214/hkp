export const state = () => ({
	name: ''
});

export const mutations = {
	set: (state, { key, value }) => (state[key] = value)
};
