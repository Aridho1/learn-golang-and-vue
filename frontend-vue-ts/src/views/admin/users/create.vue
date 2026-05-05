<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { useUserCreate } from "../../../composable/user/useUserCreate";
// import { AxiosError } from "axios";
// import { ApiErrorResponse } from "../../../types/api";
import SidebarMenu from "../../../components/SidebarMenu.vue";
import { ValidationErrors } from "../../../types/validatioError";
import { AxiosError } from "axios";
import { ApiErrorResponse } from "../../../types/api";

const router = useRouter();

const form = reactive({
    name: "",
    username: "",
    password: "",
    email: "",
});

const errors = ref<ValidationErrors>({});

const { mutate, isPending } = useUserCreate();

const storeUser = (e: Event) => {
    e.preventDefault();

    mutate(
        { ...form },
        {
            onSuccess: () => {
                router.push("/admin/users");
            },
            onError: (err: AxiosError<ApiErrorResponse>) => {
                errors.value = err.response?.data.errors || {};
            },
        },
    );
};
</script>

<template>
    <div class="container mt-5 mb-5">
        <div class="row">
            <div class="col-md-3">
                <SidebarMenu />
            </div>
            <div class="col-md-9">
                <div class="card border-0 rounded-4 shadow-sm">
                    <div class="card-header">ADD USER</div>
                    <div class="card-body">
                        <form @submit="storeUser">
                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Full Name</label>
                                <input type="text" v-model="form.name" class="form-control" placeholder="Full Name" />
                                <div v-if="errors.name" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.name }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Username</label>
                                <input type="text" v-model="form.username" class="form-control" placeholder="Username" />
                                <div v-if="errors.username" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.username }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Email address</label>
                                <input type="email" v-model="form.email" class="form-control" placeholder="Email Address" />
                                <div v-if="errors.email" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.email }}
                                </div>
                            </div>

                            <div class="form-group mb-3">
                                <label class="mb-1 fw-bold">Password</label>
                                <input type="password" v-model="form.password" class="form-control" placeholder="Password" />
                                <div v-if="errors.password" class="alert alert-danger mt-2 rounded-4">
                                    {{ errors.password }}
                                </div>
                            </div>

                            <button type="submit" class="btn btn-md btn-primary rounded-4 shadow-sm border-0" :disabled="isPending">
                                {{ isPending ? "Saving..." : "Save" }}
                            </button>

                            <router-link to="/admin/users" class="btn btn-md btn-secondary rounded-4 shadow-sm border-0 ms-2">Cancel</router-link>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
