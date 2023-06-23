import { writable } from 'svelte/store'

interface IStore {
  isSidebarOpen: boolean
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
  emojis: {
    isOpen: false,
    activeKey: null,
    keys: [],
  },
}

const store = writable<IStore>(initial)

export default store
