export interface User {
    username: string;
    password: string;
}

export interface UserResponse {
    id: number,
    username: string,
    email: string,
    role: string,
    accessToken: string,
    refreshToken: string,
}