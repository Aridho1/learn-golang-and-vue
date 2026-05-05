import { useQuery } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import { UserResponse } from "../../types/user";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";

export const useUserById = (id: number) => {
    return useQuery<UserResponse, Error>({
        queryKey: ["user", id],
        queryFn: async () => {
            const token = Cookies.get("token");

            const response = await Api.get<ApiSuccessResponse<UserResponse>>(`/api/post/${id}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data.data;
        },
    });
};
