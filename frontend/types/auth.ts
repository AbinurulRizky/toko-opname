export interface LoginResponse {
    token: string;
    role: string;
    message: string;
}

export interface ApiError {
    error: string;
}