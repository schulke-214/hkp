<template>
	<div>
		<nav class="navbar header has-shadow is-primary" role="navigation" aria-label="main navigation">
			<div class="navbar-brand">
				<a class="navbar-item" href="/">
					<img src="~assets/buefy.png" alt="Buefy" height="28" />
				</a>

				<div class="navbar-burger">
					<span />
					<span />
					<span />
				</div>
			</div>
		</nav>

		<section class="main-content columns">
			<aside class="column is-2 section">
				<p class="menu-label is-hidden-touch">General</p>
				<ul class="menu-list">
					<li v-for="(item, key) of items" :key="key">
						<nuxt-link :to="item.to" exact-active-class="is-active">
							<b-icon :icon="item.icon" /> {{ item.title }}
						</nuxt-link>
					</li>
				</ul>
			</aside>

			<div class="container column is-10">
				<nuxt />
			</div>
		</section>
	</div>
</template>

<script>
export default {};
</script>

<style lang="scss">
@import '~bulma/sass/utilities/_all';

// Set your colors
$primary: #8c67ef;
$primary-invert: findColorInvert($primary);
$twitter: #4099ff;
$twitter-invert: findColorInvert($twitter);

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
	),
	'twitter': (
		$twitter,
		$twitter-invert
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
