export const API_URL = import.meta.env.MODE === 'development' ? 'http://localhost:8000' : 'https://api.example.com'


export const authHeader = () => {
    const token = localStorage.getItem('token')
    if (token) {
        return `Bearer ${token}`
    }
    return ''
}