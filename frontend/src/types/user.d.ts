
type User = {
    id: number;
    username: string;
    email: string;
    password: string;
    created_at: string;
    updated_at: string;
    inactive_status: boolean;
    role: 'ADMIN' | 'USER';
}

type UserSettings = {
    user_id: number;
    primary_board_id: number;
    app_name: string;
    app_emoji: string;
}