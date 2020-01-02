<template>
	<div>
		<b-navbar>
			<template #brand>
				<b-navbar-item tag="nuxt-link" to="/">
					HKP
					<!-- poll-h -->
				</b-navbar-item>
				<b-navbar-item tag="div">{{ greeting }}</b-navbar-item>
			</template>

			<template #end>
				<b-navbar-item tag="div">
					<div class="buttons">
						<b-button type="is-primary" @click="modal.task = true">Create Task</b-button>
						<b-button type="is-light" @click="modal.name = true">Set Name</b-button>
					</div>
				</b-navbar-item>
			</template>
		</b-navbar>

		<b-modal :active.sync="modal.name">
			<card title :header="false" :footer="false">
				<name-form @submit="modal.name = false" />
			</card>
		</b-modal>

		<b-modal :active.sync="modal.task">
			<card title :header="false" :footer="false">
				<task-form @submit="modal.task = false" />
			</card>
		</b-modal>

		<section class="main-content columns">
			<div class="container column is-10">
				<nuxt />
			</div>
		</section>
	</div>
</template>

<script>
import Card from '~/components/ui/card';
import NameForm from '~/components/core/name-form';
import TaskForm from '~/components/core/task-form';

export default {
	name: 'layout',
	components: {
		Card,
		NameForm,
		TaskForm
	},
	data() {
		return {
			modal: {
				name: this.$store.state.user.name.length === 0,
				task: false
			}
		};
	},
	computed: {
		greeting() {
			const { name } = this.$store.state.user;

			if (name) {
				return `Hello, ${name}`;
			}

			return 'Please set a name...';
		}
	}
};
</script>

<style lang="scss">
// Import Bulma's core
@import '~bulma/sass/utilities/_all';

// Set your colors
$primary: #242259;
$primary-invert: findColorInvert($primary);

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
	'white': (
		$white,
		$black
	),
	'black': (
		$black,
		$white
	),
	'light': (
		$light,
		$light-invert
	),
	'dark': (
		$dark,
		$dark-invert
	),
	'primary': (
		$primary,
		$primary-invert
	),
	'info': (
		$info,
		$info-invert
	),
	'success': (
		$success,
		$success-invert
	),
	'warning': (
		$warning,
		$warning-invert
	),
	'danger': (
		$danger,
		$danger-invert
	)
);

// Links
$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

// Import Bulma and Buefy styles
@import '~bulma';
@import '~buefy/src/scss/buefy';
</style>
