import { API_URL, authHeader } from '.'

const gettingStarted = async (app_name: string, app_emoji: string) => {
    return await fetch(`${API_URL}/user/getting-started`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: authHeader(),
        },
        body: JSON.stringify({ app_name, app_emoji }),
    })
}

const getUserSettings = async () => {
    return await fetch(`${API_URL}/user/me/settings`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            Authorization: authHeader(),
        },
    })
}

const UserAPI = {
    gettingStarted,
    getUserSettings
}

export default UserAPI