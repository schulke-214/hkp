const matchEntry = entry => compared => entry.id === compared.id;

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
			return console.warn(`vuex/entries/edit -> tried to edit nonexisting entry`);
		}

		// keep vues reactivity
		Vue.set(state.data, index, payload);
	},
	store(state, payload) {
		const storeSingle = entry => {
			const index = state.data.findIndex(matchEntry(entry));

			if (index < 0) {
				return state.data.push(entry);
			}

			state.data[index] = entry;
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

		const entries = await this.$api.entries.all();
		commit('store', entries);

		commit('loading/stop', null, { root: true });
	},

	async create({ commit }, { id, name }) {
		commit('loading/start', null, { root: true });
		const entry = await this.$api.entries.create(id, name);
		commit('loading/stop', null, { root: true });

		commit('store', entry);
	},

	async delete({ commit }, { task, entry }) {
		commit('loading/start', null, { root: true });
		await this.$api.entries.delete(task, entry);
		commit('loading/stop', null, { root: true });

		commit('delete', entry);
	}
};

export const getters = {
	byId: state => id => state.data.find(el => el.id === id),
	manyById: state => ids => state.data.filter(el => ids.includes(el.id)),
	byTask: state => id => state.data.filter(el => el.task_id === id)
};
