import { API_URL, authHeader } from '.'

const gettingStarted = async (app_name: string, app_emoji: string) => {
  return await fetch(`${API_URL}/user/getting-started`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
    body: JSON.stringify({ app_name, app_emoji }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as StatusResponse
    }

    throw new Error(data.message)
  })
}

const getUserSettings = async () => {
  return await fetch(`${API_URL}/user/me/settings`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = (await res.json()) as UserSettings
    return data
  })
}

const UserAPI = {
  gettingStarted,
  getUserSettings,
}

export default UserAPI
