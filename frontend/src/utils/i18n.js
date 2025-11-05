import { createI18n } from "vue-i18n";
import en from "./locals/en.json";
import ar from "./locals/ar.json";
const language = navigator.language;

export const i18n = createI18n({
	locale: language,
	fallbackLocale: "en",
	messages: {
		en,
		ar,
	},
});
