export const API_URL = import.meta.env.MODE === 'development' ? 'http://localhost:8000' : 'https://api.example.com'

export const authHeader = () => {
  const token = localStorage.getItem('token')
  if (token) {
    return `Bearer ${token}`
  }
  return ''
}

export const authHeaderWithJSON = () => {
  return {
    'Content-Type': 'application/json',
    Authorization: authHeader(),
  }
}

export const handleResponse = async <T>(res: Response) => {
  const data = await res.json()
  if (res.ok) {
    return data as T
  }

  throw new Error(data.message)
}
