import { useMutation } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiErrorResponse, ApiSuccessResponse } from "../../types/api";
import { UserResponse, UserEditRequest } from "../../types/user";
import { AxiosError } from "axios";

export const useUserEdit = () => {
    return useMutation<ApiSuccessResponse<UserResponse>, AxiosError<ApiErrorResponse>, { id: number; request: UserEditRequest }>({
        mutationFn: async ({ id, request }: { id: number; request: UserEditRequest }) => {
            const token = Cookies.get("token");

            console.log("request", request);

            const response = await Api.put<ApiSuccessResponse<UserResponse>>(`/api/users/${id}`, request, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        },
    });
};
