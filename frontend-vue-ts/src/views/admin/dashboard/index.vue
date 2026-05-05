<script setup lang="ts">
//import sidebar component
import SidebarMenu from "../../../components/SidebarMenu.vue";
</script>

<template>
    <div class="container mt-5 mb-5">
        <div class="row">
            <div class="col-md-3">
                <SidebarMenu />
            </div>
            <div class="col-md-9">
                <div class="card border-0 rounded-4 shadow-sm">
                    <div class="card-header">DASHBOARD</div>
                    <div class="card-body">
                        Selamat Datang,
                        <strong></strong>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
