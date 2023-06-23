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

const createNewBoard = async (board_name: string, board_emoji: string, columns: { name: string; emoji: string; position: number }[]) => {
  return await fetch(`${API_URL}/board`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
    body: JSON.stringify({
      emoji: board_emoji,
      name: board_name,
      columns: columns,
    }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as StatusResponse
    }

    throw new Error(data.message)
  })
}

const BoardsAPI = {
  getBoards,
  createNewBoard,
}

export default BoardsAPI
