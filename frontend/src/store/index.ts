import { writable } from 'svelte/store'

interface IStore {
  isSidebarOpen: boolean
  boardDrawer: {
    isOpen: boolean
    boardID: number | null
  }
  columnDrawer: {
    isOpen: boolean
    columnID: number | null
    boardID: number | null
  }
  taskDrawer: {
    isOpen: boolean
    ids: {
      board: number | null
      column: number | null
      task: number | null
    }
  },
  taskDrawerView: {
    isOpen: boolean
    taskID: number | null
  },
  emojis: {
    isOpen: boolean
    activeKey: string | null
    keys: {
      key: string
      emoji: string
    }[]
  }
}

const initial: IStore = {
  isSidebarOpen: false,
  boardDrawer: {
    isOpen: false,
    boardID: null,
  },
  columnDrawer: {
    isOpen: false,
    columnID: null,
    boardID: null,
  },
  taskDrawer: {
    isOpen: false,
    ids: {
      board: null,
      column: null,
      task: null,
    },
  },
  taskDrawerView: {
    isOpen: false,
    taskID: null,
  },
  emojis: {
    isOpen: false,
    activeKey: null,
    keys: [],
  },
}

const store = writable<IStore>(initial)

export default store
