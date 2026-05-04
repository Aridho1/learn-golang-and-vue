import { useMutation } from "@tanstack/vue-query";
import Api from "../../services/api";

import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";

interface RegisterRequest {
    name: string;
    username: string;
    email: string;
    password: string;
}

export const useRegister = () => {
    return useMutation({
        mutationFn: async (data: RegisterRequest) => {
            try {
                // console.log("DATA:", data);
                // console.log("JSON:", JSON.stringify(data));

                const response = await Api.post<ApiSuccessResponse<UserResponse> | ApiErrorResponse>("/api/register", data);

                // console.log("SUCCESS:", response);

                return response.data as ApiSuccessResponse<UserResponse> | ApiErrorResponse;
            } catch (error) {
                // console.log("ERROR:", error);

                throw error;
            }
        },
    });
};
