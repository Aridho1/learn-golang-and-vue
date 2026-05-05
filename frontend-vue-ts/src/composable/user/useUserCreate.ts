import { useMutation } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";

interface UserRequest {
    name: string;
    username: string;
    password: string;
    email: string;
}

export const useUserCreate = () => {
    return useMutation({
        mutationFn: async (request: UserRequest) => {
            const token = Cookies.get("token");

            const response = await Api.post<ApiSuccessResponse<UserResponse>>("/api/users", request, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        },
    });
};
