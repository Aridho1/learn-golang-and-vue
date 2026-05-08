import Cookies from "js-cookie";

import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import { useAuthStore } from "../composable/auth/useAuthStore";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "home",
        component: () => import("../views/home/index.vue"),
    },
    {
        path: "/register",
        name: "register",
        component: () => import("../views/auth/register.vue"),
    },
    {
        path: "/login",
        name: "login",
        component: () => import("../views/auth/login.vue"),
    },
    {
        path: "/admin/dashboard",
        name: "dashboard",
        component: () => import("../views/admin/dashboard/index.vue"),
        meta: { requiresAuth: true },
    },
    {
        path: "/admin/users",
        name: "admin.users.index",
        component: () => import("../views/admin/users/index.vue"),
        meta: { requiresAuth: true },
    },
    {
        path: "/admin/users/create",
        name: "admin.users.create",
        component: () => import("../views/admin/users/create.vue"),
        meta: { requiresAuth: true },
    },
    {
        path: "/admin/users/edit/:id",
        name: "admin.users.edit",
        component: () => import("../views/admin/users/edit.vue"),
        meta: { requiresAuth: true },
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// const getToken = () => Cookies.get("token");

router.beforeEach(async (to, _from, next) => {
    // const token = getToken();

    // if (!token && to.matched.some((record) => record?.meta?.requiresAuth)) {
    //     next({ name: "login" });
    // } else if (token && (to.name == "login" || to.name == "register")) {
    //     next({ name: "dashboard" });
    // } else {
    //     next();
    // }

    const authStore = useAuthStore();

    if (!authStore.user && Cookies.get("token")) {
        await authStore.getUser();
    }

    if (to.meta.requiresAuth && !authStore.isAuthenticated) {
        return next({ name: "login" });
    }

    next();
});

export default router;
