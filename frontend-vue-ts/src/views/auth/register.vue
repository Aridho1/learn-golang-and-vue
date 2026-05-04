<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { useRegister } from "../../composable/auth/useRegister";
import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";
import { AxiosError } from "axios";

interface ValidationErrors {
    [key: string]: string;
}

const router = useRouter();

const form = reactive({
    name: "",
    username: "",
    email: "",
    password: "",
});

const errors = ref<ValidationErrors>({});

const { mutate, isPending } = useRegister();

const handleRegister = (e: Event) => {
    e.preventDefault();

    mutate(
        { ...form },
        {
            onSuccess: (data: ApiSuccessResponse<UserResponse>) => {
                router.push("/login");

                console.log("DATA:", data);
                // console.log("token", data?.data)
                // console.log("token", data?.data.token);
                // alert("Register Berhasil");
            },
            // onError: (error: any) => {
            // errors.value = error?.response?.data?.errors || {};
            //     // console.error(error);
            //     // Object.assign(errors, error?.response?.data?.errors || {});
            // },

            onError: (error: AxiosError<ApiErrorResponse>) => {
                errors.value = error.response?.data.errors || {};
            },

            // onError: (error: AxiosError<ApiErrorResponse>) => {
            //     errors.value = error?.response?.data?.errors || {};
            //     // console.error(error);
            //     // Object.assign(errors, error?.response?.data?.errors || {});
            // },
        },
    );
};
</script>

<template>
    <div class="row justify-content-center">
        <div class="col-md-6">
            <div class="card border-0 rounded-4 shadow-sm">
                <div class="card-body">
                    <h4 class="fw-bold text-center">REGISTER</h4>
                    <hr />
                    <form @submit.prevent="handleRegister">
                        <div class="row">
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Full Name</label>
                                    <input v-model="form.name" type="text" class="form-control" placeholder="Full Name" />
                                    <div v-if="errors.name" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.name }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Username</label>
                                    <input v-model="form.username" type="text" class="form-control" placeholder="Username" />
                                    <div v-if="errors.username" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.username }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Email address</label>
                                    <input v-model="form.email" type="email" class="form-control" placeholder="Email Address" />
                                    <div v-if="errors.email" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.email }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Password</label>
                                    <input v-model="form.password" type="password" class="form-control" placeholder="Password" />
                                    <div v-if="errors.password" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.password }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <button type="submit" class="btn btn-primary w-100 rounded-4" :disabled="isPending">
                            {{ isPending ? "Loading..." : "REGISTER" }}
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>
