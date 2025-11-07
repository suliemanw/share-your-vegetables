<template>
	<div class="flex flex-col items-center justify-center text-center p-6">
		<h1 class="text-4xl font-bold text-black-700">Admin Panel</h1>
	</div>
	<div
		class="height h-screen flex flex-col items-center justify-center text-center p-6"
	>
		<div class="bg-white shadow-lg rounded-2xl w-full max-w-md p-8">
			<button
				class="bg-green-600 text-white px-6 py-3 rounded hover:bg-green-700 transition font-semibold"
				@click="clearUploads"
			>
				Clear Uploads
			</button>
		</div>
	</div>
</template>

<script setup>
import api from "../api";

const clearUploads = async () => {
	await api.post("/clear-uploads");
	alert("Uploads cleared!");
};
</script>

<style scoped>
.height {
	max-height: fit-content;
}
</style>
