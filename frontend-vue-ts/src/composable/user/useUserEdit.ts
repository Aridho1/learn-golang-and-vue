import { useMutation } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";
import { UserResponse, UserEditRequest } from "../../types/user";

export const useUserEdit = () => {
    return useMutation({
        mutationFn: async ({ id, request }: { id: number; request: UserEditRequest }) => {
            const token = Cookies.get("token");

            const response = await Api.put<ApiSuccessResponse<UserResponse>>(`/api/users/${id}`, request, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        },
    });
};
