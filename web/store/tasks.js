const matchTasks = task => compared => task.id === compared.id;

export const state = () => ({
	data: []
});

export const mutations = {
	delete(state, id) {
		state.data = state.data.filter(item => item.id !== id);
	},
	edit(state, payload) {
		const index = state.data.findIndex(item => item.id === payload.id);

		if (index < 0) {
			return console.warn(`vuex/tasks/edit -> tried to edit nonexisting task`);
		}

		// keep vues reactivity
		Vue.set(state.data, index, payload);
	},
	store(state, payload) {
		const storeSingle = task => {
			const index = state.data.findIndex(matchTasks(task));

			if (index < 0) {
				return state.data.push(task);
			}

			state.data[index] = task;
		};

		if (Array.isArray(payload)) {
			return payload.map(storeSingle);
		}

		storeSingle(payload);
	}
};

export const actions = {
	store({ commit }, payload) {
		commit('store', payload);
	},
	async fetch({ commit }) {
		commit('loading/start', null, { root: true });

		const tasks = await this.$api.tasks.all();
		commit('store', tasks);

		commit('loading/stop', null, { root: true });
	},

	async edit({ commit }, { id, changes }) {
		if (!id || !changes) return;

		commit('loading/start', null, { root: true });
		const task = await this.$api.tasks.update(id, changes);
		commit('loading/stop', null, { root: true });

		commit('edit', task);
	},

	async create({ commit }, { name }) {
		commit('loading/start', null, { root: true });
		const task = await this.$api.tasks.create(name);
		commit('loading/stop', null, { root: true });

		commit('store', task);
	},

	async delete({ commit }, { id }) {
		commit('loading/start', null, { root: true });
		await this.$api.tasks.delete(id);
		commit('loading/stop', null, { root: true });

		commit('delete', id);
	}
};

export const getters = {
	byId: state => id => state.data.find(el => el.id === id),
	manyById: state => ids => state.data.filter(el => ids.includes(el.id))
};
