import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

import { VueQueryPlugin } from "@tanstack/vue-query";
import router from "./routes";
import { createPinia } from "pinia";

console.log("Hello World");
createApp(App).use(createPinia()).use(VueQueryPlugin).use(router).mount("#app");
console.log("Hello Worlds");
