import { API_URL, authHeaderWithJSON, handleResponse } from '.'

const getUser = async () => {
  return await fetch(`${API_URL}/user/me`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<User>(res))
}

const gettingStarted = async (app_name: string, app_emoji: string) => {
  return await fetch(`${API_URL}/user/getting-started`, {
    method: 'POST',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ app_name, app_emoji }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const getUserSettings = async () => {
  return await fetch(`${API_URL}/user/me/settings`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<UserSettings>(res))
}

const updateUserSettings = async (settings: Omit<Omit<UserSettings, 'primary_board_id'>, 'user_id'>) => {
  return await fetch(`${API_URL}/user/me/settings`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify(settings),
  }).then(async (res) => handleResponse<MessageResponse>(res))
}

const updatePrimaryBoard = async (primary_board_id: number) => {
  return await fetch(`${API_URL}/user/me/settings/primary-board`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ primary_board_id }),
  }).then(async (res) => handleResponse<UserSettings>(res))
}

const updatePassword = async (oldPassword: string, newPassword: string) => {
  return await fetch(`${API_URL}/user/me/password`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ oldPassword, newPassword }),
  }).then((res) => handleResponse<MessageResponse>(res))
}

const updateUserName = async (username: string) => {
  return await fetch(`${API_URL}/user/me/username`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ username }),
  }).then(async (res) => handleResponse<MessageResponse>(res))
}

const updateEmail = async (email: string) => {
  return await fetch(`${API_URL}/user/me/email`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ email }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as MessageResponse
    }

    throw new Error(data.message)
  })
}

const UserAPI = {
  getUser,
  gettingStarted,
  getUserSettings,
  updateUserSettings,
  updatePrimaryBoard,
  updatePassword,
  updateUserName,
  updateEmail,
}

export default UserAPI
