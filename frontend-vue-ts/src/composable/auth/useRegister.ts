import { useMutation } from "@tanstack/vue-query";
import Api from "../../services/api";

import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";
import { AxiosError } from "axios";

interface RegisterRequest {
    name: string;
    username: string;
    email: string;
    password: string;
}

export const useRegister = () => {
    return useMutation<ApiSuccessResponse<UserResponse>, AxiosError<ApiErrorResponse>, RegisterRequest>({
        mutationFn: async (data: RegisterRequest) => {
            try {
                // console.log("DATA:", data);
                // console.log("JSON:", JSON.stringify(data));

                const response = await Api.post<ApiSuccessResponse<UserResponse>>("/api/register", data);

                // console.log("SUCCESS:", response);

                return response.data;
            } catch (error) {
                // console.log("ERROR:", error);

                throw error;
            }
        },
    });
};
