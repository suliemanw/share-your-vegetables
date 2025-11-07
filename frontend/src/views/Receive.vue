<template>
	<div class="h-screen flex flex-col items-center justify-center text-center p-6">
		<h1 class="text-3xl font-bold text-blue-700 mb-6">📥 Receive Files</h1>
		<div class="bg-white shadow-lg rounded-2xl w-full max-w-md p-6">
			<button
				class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-lg w-full mb-4"
				@click="fetchFiles"
			>
				Show Available Files
			</button>

			<ul v-if="files.length" class="divide-y divide-gray-200">
				<li
					v-for="(file, index) in files"
					:key="index"
					class="flex justify-between items-center py-2"
				>
					<span class="text-gray-800 truncate w-3/4">{{ file.name }}</span>
					<a
						:href="downloadUrl(file.name)"
						class="text-blue-600 hover:underline font-medium"
						download
					>
						Download
					</a>
				</li>
			</ul>

			<p v-else class="text-gray-500">No files available</p>
		</div>
	</div>
</template>

<script setup>
import { ref } from "vue";
import api from "../api.js";

const files = ref([]);

const fetchFiles = async () => {
	try {
		const res = await api.get("/files");
		files.value = res.data || [];
	} catch (err) {
		console.error("Failed to fetch files:", err);
	}
};

// Generate correct download URL using current host
const downloadUrl = (name) => {
	const host = window.location.hostname;
	return `http://${host}:8081/download/${encodeURIComponent(name)}`;
};
</script>

<style scoped>
.height {
	max-height: fit-content;
}
</style>
