import { useMutation } from "@tanstack/vue-query";
import Api from "../../services/api";

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

                const response = await Api.post("/api/register", data);

                // console.log("SUCCESS:", response);

                return response.data;
            } catch (error) {
                // console.log("ERROR:", error);

                throw error;
            }
        },
    });
};
