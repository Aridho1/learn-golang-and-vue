<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";

import { ValidationErrors } from "../../types/validatioError";
import { useLogin } from "../../composable/auth/useLogin";
import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";

import Cookies from "js-cookie";
import { AxiosError } from "axios";

const router = useRouter();

const form = reactive({
    username: "",
    password: "",
});

const errors = ref<ValidationErrors>();

const { mutate, isPending } = useLogin();

const handleLogin = (e: Event) => {
    e.preventDefault();

    mutate(
        {
            ...form,
        },
        {
            onSuccess: (data: ApiSuccessResponse<UserResponse>) => {
                const user = structuredClone(data.data);

                Cookies.set("token", user.token ?? "");

                Cookies.set(
                    "user",
                    JSON.stringify({
                        id: user.id,
                        name: user.name,
                        username: user.username,
                        email: user.email,
                    }),
                );

                router.push("/admin/dashboard");
            },
            onError: (err: AxiosError<ApiErrorResponse>) => {
                console.log(err);
                errors.value = err.response?.data.errors || {};
            },
        },
    );
};
</script>

<template>
    <div class="row justify-content-center mt-5">
        <div class="col-md-4">
            <div class="card border-0 rounded-4 shadow-sm">
                <div class="card-body">
                    <h4 class="fw-bold text-center">LOGIN</h4>
                    <hr />
                    <div v-if="errors?.error" class="alert alert-danger mt-2 rounded-4">Username or Password is incorrect</div>
                    <form @submit="handleLogin">
                        <div class="form-group mb-3">
                            <label class="mb-1 fw-bold">Username</label>
                            <input v-model="form.username" type="text" class="form-control" placeholder="Username" />
                            <div v-if="errors?.username" class="alert alert-danger mt-2 rounded-4">
                                {{ errors?.username }}
                            </div>
                        </div>

                        <div class="form-group mb-3">
                            <label class="mb-1 fw-bold">Password</label>
                            <input v-model="form.password" type="password" class="form-control" placeholder="Password" />
                            <div v-if="errors?.password" class="alert alert-danger mt-2 rounded-4">
                                {{ errors?.password }}
                            </div>
                        </div>

                        <button type="submit" class="btn btn-primary w-100 rounded-4" :disabled="isPending">
                            {{ isPending ? "Loading..." : "LOGIN" }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>
