import { API_URL, authHeader } from '.'

const getBoardById = async (board_id: number) => {
  return await fetch(`${API_URL}/board/${board_id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as Board
    }

    throw new Error(data.message)
  })
}

const getColumns = async (board_id: number) => {
  return await fetch(`${API_URL}/board/${board_id}/column`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as ColumnBoard[]
    }

    throw new Error(data.message)
  })
}

const createColumn = async (board_id: number, name: string, emoji: string, position: number) => {
  return await fetch(`${API_URL}/board/${board_id}/column`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
    body: JSON.stringify({
      name: name,
      emoji: emoji,
      position: position,
    }),
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as StatusResponse
    }

    throw new Error(data.message)
  })
}

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

const updateColumnPosition = async (board_id: number, data: { id: number; position: number }[]) => {
  return await fetch(`${API_URL}/board/${board_id}/column/position`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
    body: JSON.stringify(data),
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
  getColumns,
  getBoardById,
  createNewBoard,
  createColumn,
  updateColumnPosition,
}

export default BoardsAPI
