export interface User {
    id: number;
    name: string;
    username: string;
    email: string;
}

export interface UserRequest {
    name: string;
    username: string;
    password: string;
    email: string;
}

export interface UserResponse extends User {
    // id: number;
    // name: string;
    // username: string;
    // email: string;
    created_at: string;
    updated_at: string;
    token?: string;
}
