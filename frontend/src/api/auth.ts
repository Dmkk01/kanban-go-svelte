import { API_URL } from '.'

const login = async (email: string, password: string) => {
  return fetch(`${API_URL}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as TokenResponse
    }

    throw new Error(data.message)
  })
}

const register = async (email: string, username: string, password: string) => {
  return fetch(`${API_URL}/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password, username }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as StatusResponse
    }

    throw new Error(data.message)
  })
}

const AuthAPI = {
  login,
  register,
}

export default AuthAPI
