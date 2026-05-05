import { useMutation } from "@tanstack/vue-query";
import Api from "../../services/api";
import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";
import { AxiosError } from "axios";

interface LoginRequest {
    username: string;
    password: string;
}

export const useLogin = () => {
    return useMutation<ApiSuccessResponse<UserResponse>, AxiosError<ApiErrorResponse>, LoginRequest>({
        mutationFn: async (data: LoginRequest) => {
            const response = await Api.post<ApiSuccessResponse<UserResponse>>("/api/login", data);

            return response.data;
        },
    });
};
