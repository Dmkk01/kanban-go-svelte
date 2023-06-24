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
  emojis: {
    isOpen: false,
    activeKey: null,
    keys: [],
  },
}

const store = writable<IStore>(initial)

export default store
