type Board = {
  id: number
  user_id: number
  name: string
  emoji: string
  created_at: string
  updated_at: string
}

type ColumnBoard = {
  id: number
  board_id: number
  created_at: string
  updated_at: string
  emoji: string
  name: string
  position: number
}

type Task = {
  id: number
  column_id: number
  board_id: number
  title: string
  description: string
  time_needed: number
  due_date: string
  position: number
  created_at: string
  updated_at: string
  sub_tasks: SubTask[]
  links: LinkTask[]
}

type SubTask = {
  id: number
  task_id: number
  title: string
  completed: boolean
  created_at: string
  updated_at: string
}

type LinkTask = {
  id: number
  task_id: number
  title: string
  url: string
  emoji: string
  created_at: string
  updated_at: string
}
