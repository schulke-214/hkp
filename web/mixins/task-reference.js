export default {
	props: {
		taskId: {
			type: String,
			required: true
		}
	},
	computed: {
		task() {
			return this.$store.getters['tasks/byId'](this.taskId);
		}
	}
};
