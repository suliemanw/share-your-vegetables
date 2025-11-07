<template>
	<div
		class="height h-screen flex flex-col items-center justify-center text-center p-6"
	>
		<h1 class="text-3xl font-bold text-green-700 mb-6">
			Choose a file to share
		</h1>
		<div class="bg-white shadow-lg rounded-2xl w-full max-w-md p-6">
			<h2 class="text-xl font-semibold mb-4">Send Files</h2>

			<input
				type="file"
				multiple
				@change="handleFiles"
				class="block w-full border border-gray-300 rounded-lg p-2 mb-4 cursor-pointer"
			/>

			<ul class="divide-y divide-gray-200 mb-4">
				<li
					v-for="(file, index) in files"
					:key="index"
					class="flex justify-between items-center py-2"
				>
					<span class="text-gray-800 truncate w-3/4">{{ file.name }}</span>
					<button
						class="text-red-500 hover:text-red-700 font-medium"
						@click="removeFile(index)"
					>
						Remove
					</button>
				</li>
			</ul>

			<button
				class="bg-green-600 hover:bg-green-700 text-white font-semibold py-2 px-4 rounded-lg w-full"
				@click="uploadFiles"
			>
				Upload
			</button>
		</div>
	</div>
</template>

<script setup>
import { ref } from "vue";
import api from "../api.js";
const files = ref([]);

const handleFiles = (e) => {
	files.value = Array.from(e.target.files);
};

const removeFile = (index) => {
	files.value.splice(index, 1);
};

const uploadFiles = async () => {
	if (files.value.length === 0) {
		alert("Select at least one file!");
		return;
	}

	const formData = new FormData();
	files.value.forEach((file) => formData.append("files", file));

	await api.post("/upload", formData, {
		headers: { "Content-Type": "multipart/form-data" },
	});

	alert("Files uploaded!");
	files.value = [];
};
</script>

<style scoped>
.height {
	max-height: fit-content;
}
</style>
