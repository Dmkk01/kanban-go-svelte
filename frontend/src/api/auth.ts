import { API_URL, handleResponse } from '.'

const login = async (email: string, password: string) => {
  localStorage.setItem('email', email)
  localStorage.setItem('password', password)
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

const refresh = () => {
  const email = localStorage.getItem('email')
  const password = localStorage.getItem('password')

  if (email && password) {
    return login(email, password)
  }

  return Promise.reject('No email or password')
}

const AuthAPI = {
  login,
  register,
  refresh,
}

export default AuthAPI
