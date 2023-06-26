import { API_URL, authHeader, authHeaderWithJSON, handleResponse } from '.'

const getBoardById = async (board_id: number) => {
  return await fetch(`${API_URL}/board/${board_id}`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<Board>(res))
}

const getColumns = async (board_id: number) => {
  return await fetch(`${API_URL}/board/${board_id}/column`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<ColumnBoard[]>(res))
}

const createColumn = async (board_id: number, name: string, emoji: string, position: number) => {
  return await fetch(`${API_URL}/board/${board_id}/column`, {
    method: 'POST',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({
      name: name,
      emoji: emoji,
      position: position,
    }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const getBoards = async () => {
  return await fetch(`${API_URL}/board`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<Board[]>(res))
}

const getBoardFull = async (board_id: number) => {
  return await fetch(`${API_URL}/board/${board_id}/full`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<BoardFull>(res))
}

const createNewBoard = async (board_name: string, board_emoji: string, columns: { name: string; emoji: string; position: number }[]) => {
  return await fetch(`${API_URL}/board`, {
    method: 'POST',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({
      emoji: board_emoji,
      name: board_name,
      columns: columns,
    }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const updateColumnPosition = async (board_id: number, data: { id: number; position: number }[]) => {
  return await fetch(`${API_URL}/board/${board_id}/column/position`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify(data),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const BoardsAPI = {
  getBoards,
  getColumns,
  getBoardById,
  createNewBoard,
  createColumn,
  updateColumnPosition,
  getBoardFull,
}

export default BoardsAPI
