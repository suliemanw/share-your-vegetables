<template>
	<div>
		<div v-if="loading" class="text-gray-500">Loading...</div>
		<ul v-else>
			<li
				v-for="f in files"
				:key="f.name"
				class="flex justify-between items-center py-2"
			>
				<strong class="text-gray-800">{{ f.name }}</strong> —
				{{ Math.round(f.size / 1024) }} KB
				<button
					v-if="showDownload"
					@click="download(f.name)"
					class="text-blue-500 hover:text-blue-700"
				>
					Download
				</button>
				<button
					v-if="showDelete"
					@click="del(f.name)"
					class="text-red-500 hover:text-red-700"
				>
					Delete
				</button>
			</li>
		</ul>
	</div>
</template>

<script>
import api from "../api";

export default {
	name: "FileList",
	props: {
		showDownload: { type: Boolean, default: true },
		showDelete: { type: Boolean, default: true },
	},
	data() {
		return {
			files: [],
			loading: true,
		};
	},
	created() {
		this.load();
		this.$root.$on("refresh-list", this.load);
	},
	methods: {
		async load() {
			this.loading = true;
			try {
				const res = await api.get("/files");
				this.files = res.data;
			} catch (err) {
				console.error(err);
				this.files = [];
			} finally {
				this.loading = false;
			}
		},
		download(name) {
			const base =
				import.meta.env.VITE_API_URL ||
				`http://${window.location.hostname}:8081`;
			window.open(`${base}/download/${encodeURIComponent(name)}`, "_blank");
		},
		async del(name) {
			if (!confirm("Delete " + name + "?")) return;
			try {
				await api.post("/delete/" + encodeURIComponent(name));
				this.load();
				this.$emit("deleteFile");
			} catch (err) {
				alert("Delete failed");
			}
		},
	},
};
</script>
