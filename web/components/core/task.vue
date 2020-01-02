<template>
	<card :title="task.name">
		<template #header>
			<span
				style="cursor: pointer; display: flex; flex-direction: column; justify-content: center; align-items: center; padding: 0 0.75rem;"
				@click="$emit('delete')"
			>
				<b-icon icon="delete" type="is-danger" />
			</span>
		</template>

		<template>
			<!-- <ul>
				<li
					v-for="entry in entries"
					:key="entry.id"
					style="cursor: pointer;"
					@click="removeById(entry.id)"
				>{{ entry.user }}</li>
			</ul>-->

			<ul>
				<li v-for="name in sortedNames" :key="name" v-text="name + ': ' + occurrence[name]" />
			</ul>
			<br />
			<strong>Als Nächstes: {{ next }}</strong>
		</template>

		<template #footer>
			<div class="buttons" style="width: 100%">
				<b-button
					type="is-primary"
					icon-left="plus"
					expanded
					@click="add"
					:disabled="!canAdd"
				>Hinzufügen</b-button>
				<b-button
					type="is-danger"
					icon-left="delete"
					expanded
					@click="removeFirst"
					:disabled="!canRemoveFirst"
				>Entfernen</b-button>
			</div>
		</template>
	</card>
</template>

<script>
import { mapActions } from 'vuex';

import taskReference from '~/mixins/task-reference';
import Card from '~/components/ui/card';

export default {
	mixins: [taskReference],
	components: {
		Card
	},
	computed: {
		entries() {
			return this.$store.getters['entries/byTask'](this.task.id);
		},

		sortedNames() {
			return Object.keys(this.occurrence).sort((a, b) => {
				if (this.occurrence[a] > this.occurrence[b]) return -1;
				else return 1;
			});
		},

		occurrence() {
			const names = this.entries.map(e => e.user);
			const initial = names.reduce(
				(accum, name) => ({
					...accum,
					[name]: 0
				}),
				{}
			);

			const occurrence = this.entries.reduce(
				(accum, entry) => {
					return {
						...accum,
						[entry.user]: accum[entry.user] + 1
					};
				},
				{ ...initial }
			);

			return occurrence;
		},

		next() {
			const names = Object.keys(this.occurrence);

			const smallest = names.reduce((accum, name) => {
				if (this.occurrence[name] < accum) return this.occurrence[name];
				else return accum;
			}, Infinity);

			return names.find(name => this.occurrence[name] === smallest);
		},

		canAdd() {
			return !!this.$store.state.user.name.length;
		},

		canRemoveFirst() {
			return !!this.occurrence[this.$store.state.user.name];
		}
	},
	methods: {
		...mapActions({
			create: 'entries/create',
			delete: 'entries/delete'
		}),

		add() {
			this.create({
				id: this.task.id,
				name: this.$store.state.user.name
			});
		},

		removeFirst() {
			const { name } = this.$store.state.user;

			const match = this.entries.find(entry => entry.user === name);

			if (!match) return console.warn('ERR');

			this.delete({ task: this.task.id, entry: match.id });
		},

		removeById(id) {
			this.delete({ task: this.task.id, entry: id });
		}
	}
};
</script>

<style></style>
