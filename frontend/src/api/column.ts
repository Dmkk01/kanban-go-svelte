import { API_URL, authHeader, authHeaderWithJSON, handleResponse } from '.'

const getColumn = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<ColumnBoard>(res))
}

const getTasks = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}/task`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<Task[]>(res))
}

const editColumn = async (column_id: number, emoji: string, name: string) => {
  return await fetch(`${API_URL}/column/${column_id}`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ emoji, name }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const deleteColumn = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}`, {
    method: 'DELETE',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const ColumnAPI = {
  getColumn,
  getTasks,
  editColumn,
  deleteColumn
}

export default ColumnAPI
