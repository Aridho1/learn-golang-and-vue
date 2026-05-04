import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";

import { VueQueryPlugin } from "@tanstack/vue-query";
import router from "./routes";

console.log("Hello World");
createApp(App).use(VueQueryPlugin).use(router).mount("#app");
console.log("Hello Worlds");
