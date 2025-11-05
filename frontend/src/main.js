import { createApp } from "vue";
import { createPinia } from "pinia";
import AOS from "aos";

import { faInit } from "./utils/fontawesome";
import { i18n } from "./utils/i18n";

import router from "./router";

import "aos/dist/aos.css";
import "./style.css";
import App from "./App.vue";

const app = createApp(App);
const pinia = createPinia();

faInit(app);

app.use(
	AOS.init({
		duration: 1000,
		offset: 120,
		once: true
	})
);

app.use(router);
app.use(i18n);
app.use(pinia);
app.mount("#app");
