export default {
	props: {
		entryId: {
			type: String,
			required: true
		}
	},
	computed: {
		task() {
			return this.$store.getters['entries/byId'](this.entryId);
		}
	}
};
