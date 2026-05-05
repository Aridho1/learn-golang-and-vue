import { useMutation } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse, UserCreateRequest } from "../../types/user";
import { AxiosError } from "axios";

export const useUserCreate = () => {
    return useMutation<ApiSuccessResponse<UserResponse>, AxiosError<ApiErrorResponse>, UserCreateRequest>({
        mutationFn: async (request: UserCreateRequest) => {
            try {
                const token = Cookies.get("token");

                const response = await Api.post<ApiSuccessResponse<UserResponse>>("/api/users", request, {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });

                return response.data;
            } catch (err) {
                throw err;
            }
        },
    });
};
