<template>
	<validation-observer v-slot="{ invalid }" ref="observer">
		<form @submit.prevent="() => !invalid && submit()">
			<b-field label="Dein Name">
				<validation-provider name="Name" rules="required">
					<b-input v-model="username" />
				</validation-provider>
			</b-field>
			<b-button type="is-success" expanded :disabled="invalid" @click="submit">Bestätigen</b-button>
		</form>
	</validation-observer>
</template>

<script>
export default {
	data() {
		return {
			username: this.$store.state.user.name
		};
	},
	methods: {
		submit() {
			this.$store.commit('user/set', {
				key: 'name',
				value: this.username
			});
			this.$emit('submit');
		}
	}
};
</script>

<style>
</style>