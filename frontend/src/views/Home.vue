<template>
	<div
		class="height h-screen flex flex-col items-center justify-center text-center p-6"
	>
		<h1 class="text-4xl font-bold text-green-700 mb-8">Choose an action</h1>
		<div class="flex gap-6">
			<RouterLink
				to="/send"
				class="bg-green-600 hover:bg-green-700 text-white px-8 py-3 rounded-xl font-semibold transition"
			>
				Send Files
			</RouterLink>
			<RouterLink
				to="/receive"
				class="bg-blue-600 hover:bg-blue-700 text-white px-8 py-3 rounded-xl font-semibold transition"
			>
				Receive Files
			</RouterLink>
		</div>
	</div>
</template>

<style scoped>
.height {
	max-height: fit-content;
}
</style>
