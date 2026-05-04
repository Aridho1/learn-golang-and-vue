export interface UserResponse {
    id: number;
    name: string;
    username: string;
    email: string;
    created_at: string;
    updated_at: string;
    token?: string;
}
