import { API_URL, authHeader } from '.'

const getColumn = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as ColumnBoard
    }

    throw new Error(data.message)
  })
}

const getTasks = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}/task`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as Task[]
    }

    throw new Error(data.message)
  })
}

const ColumnAPI = {
  getColumn,
  getTasks,
}

export default ColumnAPI
