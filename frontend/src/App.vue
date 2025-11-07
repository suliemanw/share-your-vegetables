<template>
	<div class="app">
		<header style="padding: 15px; background: #2d6a4f; color: white">
			<h1 class="text-2xl font-bold text-center">🥦 Share Your Vegetables</h1>
		</header>
		<main class="bg-green-100">
			<router-view />
		</main>
		<footer class="p-3 text-center mt-auto bg-[#2d6a4f] text-white">
			<h3 class="text-2xl">Local LAN file sharing — no internet needed.</h3>
			<p class="mt-3 text-2xl">
				by
				<a
					href="https://github.com/suliemanw/share-your-vegetables"
					target="_blank"
					rel="noopener noreferrer"
					>Suliemanw</a
				>
				|
				<a
					href="https://dragonware.ly"
					target="_blank"
					rel="noopener noreferrer"
					>Dragonware</a
				>
			</p>
		</footer>
	</div>
</template>

<script>
export default {
	name: "App",
};
</script>

<style>
body {
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
		Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
	margin: 0;
}
.app {
	min-height: 100vh;
	display: flex;
	flex-direction: column;
}
main {
	flex: 1;
	max-height: 80vh;
	align-content: center;
	overflow-y: auto;
}
</style>
