import { defineStore } from "pinia";

export const useMainStore = defineStore("main", {
  state: () => {
    return { message: "hello" };
  },

  actions: {},
});
