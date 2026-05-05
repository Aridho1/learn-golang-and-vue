import Cookies from "js-cookie";

import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";

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
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

const getToken = () => Cookies.get("token");

router.beforeEach((to, _from, next) => {
    const token = getToken();

    if (!token && to.matched.some((record) => record?.meta?.requiresAuth)) {
        next({ name: "login" });
    } else if (token && (to.name == "login" || to.name == "register")) {
        next({ name: "dashboard" });
    } else {
        next();
    }
});

export default router;
