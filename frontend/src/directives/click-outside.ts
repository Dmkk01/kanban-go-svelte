export const clickOutside = (node: HTMLElement) => {
  const handleClick = (event: MouseEvent) => {
    if (node && !node.contains(event.target as Node)) {
      node.dispatchEvent(new CustomEvent<HTMLElement>('click_outside', { detail: node }))
    }
  }

  document.addEventListener('click', handleClick)

  return {
    destroy() {
      document.removeEventListener('click', handleClick)
    },
  }
}
