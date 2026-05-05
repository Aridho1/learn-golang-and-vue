import { useQuery } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";
import { UserResponse } from "../../types/user";

export const useUser = () => {
    return useQuery<UserResponse[], Error>({
        queryKey: ["users"],
        queryFn: async () => {
            const token = Cookies.get("token");

            const response = await Api.get<ApiSuccessResponse<UserResponse[]>>("/api/users", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data.data;
        },
    });
};
