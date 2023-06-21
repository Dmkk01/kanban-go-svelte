import { API_URL, authHeader } from '.'

const getBoards = async () => {
  return await fetch(`${API_URL}/board`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as Board[]
    }

    throw new Error(data.message)
  })
}

const BoardsAPI = {
  getBoards,
}

export default BoardsAPI
