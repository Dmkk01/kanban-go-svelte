import { API_URL, handleResponse } from '.'

const login = async (email: string, password: string) => {
  return fetch(`${API_URL}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  }).then(async (res) => handleResponse<TokenResponse>(res))
}

const register = async (email: string, username: string, password: string) => {
  return fetch(`${API_URL}/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password, username }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const AuthAPI = {
  login,
  register,
}

export default AuthAPI
