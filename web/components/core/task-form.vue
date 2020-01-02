<template>
	<validation-observer v-slot="{ invalid }" ref="observer">
		<form @submit.prevent="() => !invalid && create()">
			<b-field label="Neue Aufgabe">
				<validation-provider name="Name" rules="required">
					<b-input v-model="name" />
				</validation-provider>
			</b-field>
			<b-button
				type="is-success"
				icon-left="plus"
				:disabled="invalid"
				expanded
				@click="create"
			>Erstellen</b-button>
		</form>
	</validation-observer>
</template>

<script>
export default {
	data: () => ({
		name: ''
	}),
	methods: {
		async create() {
			await this.$store.dispatch(
				'tasks/create',
				{ name: this.name },
				{ root: true }
			);
			this.name = '';
			this.$refs.observer.reset();
			this.$emit('submit');
		}
	}
};
</script>

<style lang="scss" scoped>
</style>