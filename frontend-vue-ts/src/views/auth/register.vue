<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useRegister } from "../../composable/auth/userRegister";

interface ValidationErrors {
    [key: string]: string;
}

const router = useRouter();

const name = ref<string>("");
const username = ref<string>("");
const email = ref<string>("");
const password = ref<string>("");

const errors = ref<ValidationErrors>({});

const { mutate, isPending } = useRegister();

const handleRegister = (e: Event) => {
    e.preventDefault();

    mutate(
        {
            name: name.value,
            username: username.value,
            email: email.value,
            password: password.value,
        },
        {
            onSuccess: () => {
                router.push("/login");
                // alert("Register Berhasil");
            },
            onError: (error: any) => {
                errors.value = error?.response?.data?.errors || {};
                // console.error(error);
                // Object.assign(errors, error?.response?.data?.errors || {});
            },
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
                                    <input v-model="name" type="text" class="form-control" placeholder="Full Name" />
                                    <div v-if="errors.Name" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Name }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Username</label>
                                    <input v-model="username" type="text" class="form-control" placeholder="Username" />
                                    <div v-if="errors.Username" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Username }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <div class="row">
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Email address</label>
                                    <input v-model="email" type="email" class="form-control" placeholder="Email Address" />
                                    <div v-if="errors.Email" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Email }}
                                    </div>
                                </div>
                            </div>
                            <div class="col-md-6 mb-3">
                                <div class="form-group">
                                    <label class="mb-1 fw-bold">Password</label>
                                    <input v-model="password" type="password" class="form-control" placeholder="Password" />
                                    <div v-if="errors.Password" class="alert alert-danger mt-2 rounded-4">
                                        {{ errors.Password }}
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
