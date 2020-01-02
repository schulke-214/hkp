<template>
	<b-loading :is-full-page="true" :active.sync="loading" v-show="false"></b-loading>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex';

export default {
	computed: mapGetters({
		loading: 'loading/status'
	}),
	methods: mapMutations({
		start: 'loading/start',
		finish: 'loading/stop'
	})
};
</script>

<style></style>
