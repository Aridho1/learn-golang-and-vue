import { useQuery } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";

export interface User {
    id: number;
    name: string;
    username: string;
    email: string;
}

export const useUser = () => {
    return useQuery<User[], Error>({
        queryKey: ["users"],
        queryFn: async () => {
            const token = Cookies.get("token");

            const response = await Api.get<ApiSuccessResponse<User[]>>("/api/users", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data.data;
        },
    });
};
