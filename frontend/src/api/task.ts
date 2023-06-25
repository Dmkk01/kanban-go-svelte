import { API_URL, authHeader } from '.'

const getTask = async (task_id: number) => {
  return await fetch(`${API_URL}/task/${task_id}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      Authorization: authHeader(),
    },
  }).then(async (res) => {
    const data = await res.json()
    if (res.ok) {
      return data as Task
    }

    throw new Error(data.message)
  })
}

const taskUpdatePosition = async (data: { task_id: number; position: number; column_id: number }[]) => {
  return await fetch(`${API_URL}/task/position`, {
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

const TaskAPI = {
  getTask,
  taskUpdatePosition,
}

export default TaskAPI
