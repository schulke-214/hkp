<template>
	<div class="column" style="min-width: 200px;">
		<div class="card">
			<header v-if="header" class="card-header">
				<p class="card-header-title">{{ title }}</p>
				<slot name="header" />
			</header>
			<div class="card-content">
				<slot />
			</div>
			<footer v-if="footer" class="card-footer">
				<div class="card-footer-item">
					<slot name="footer" />
				</div>
			</footer>
		</div>
	</div>
</template>

<script>
export default {
	props: {
		title: {
			type: String,
			required: true
		},
		header: {
			type: Boolean,
			default: true
		},
		footer: {
			type: Boolean,
			default: true
		}
	}
};
</script>
