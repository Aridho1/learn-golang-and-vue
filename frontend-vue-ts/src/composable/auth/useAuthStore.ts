import { defineStore } from "pinia";
import Cookies from "js-cookie";
import Api from "../../services/api";

export const useAuthStore = defineStore("auth", {
    state: () => ({
        user: null,
        isAuthenticated: false,
    }),

    actions: {
        async getUser() {
            try {
                const token = Cookies.get("token");

                const response = await Api.get("/me", {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                this.user = response.data;
                this.isAuthenticated = true;
            } catch {
                this.user = null;
                this.isAuthenticated = false;

                Cookies.remove("token");
            }
        },
    },
});
