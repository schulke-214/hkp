<template>
	<div>
		<section class="section">
			<div class="columns" style="flex-wrap: wrap;">
				<task v-for="task in tasks" :key="task.id" :task-id="task.id" @delete="remove(task.id)" />
				<b-message v-if="!tasks.length">Sieht aus als wären da noch keine aufgaben...</b-message>
			</div>
		</section>
	</div>
</template>

<script>
import { mapState } from 'vuex';

import Task from '~/components/core/task';
import TaskForm from '~/components/core/task-form';

export default {
	name: 'index',
	components: {
		Task,
		TaskForm
	},
	methods: {
		async remove(id) {
			await this.$store.dispatch('tasks/delete', { id }, { root: true });
		}
	},
	computed: {
		tasks() {
			return this.$store.state.tasks.data;
		}
	}
};
</script>
