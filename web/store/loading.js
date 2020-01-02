export const state = () => ({
	count: 0
});

export const getters = {
	status: state => state.count > 0
};

export const mutations = {
	start: state => state.count++,
	stop: state => state.count - 1 > -1 && state.count--
};
