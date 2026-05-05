import { useMutation } from "@tanstack/vue-query";
import Cookies from "js-cookie";
import Api from "../../services/api";
import { ApiSuccessResponse } from "../../types/api";

export const useUserRemove = () => {
    return useMutation({
        mutationFn: async (id: number) => {
            const token = Cookies.get("token");

            const response = await Api.delete<ApiSuccessResponse<void>>(`/api/users/${id}`, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            return response.data;
        },
    });
};
