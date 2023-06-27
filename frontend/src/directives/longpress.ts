export const longpress = (node: HTMLElement, data: { category: string }) => {
  const TIME_MS = 500
  let timeoutPtr: number

  const handleMouseDown = (e: MouseEvent) => {
    e.stopPropagation()
    window.addEventListener('mousemove', handleMoveBeforeLong)
    timeoutPtr = window.setTimeout(() => {
      window.removeEventListener('mousemove', handleMoveBeforeLong)
      node.dispatchEvent(
        new CustomEvent('long', {
          detail: {
            target: e.target,
            category: data.category,
          },
        })
      )
      window.setTimeout(() => node.dispatchEvent(e), 0)
    }, TIME_MS)
  }

  const handleMoveBeforeLong = (e: MouseEvent) => {
    window.clearTimeout(timeoutPtr)
    window.removeEventListener('mousemove', handleMoveBeforeLong)
  }

  const handleMouseUp = (e: MouseEvent) => {
    window.clearTimeout(timeoutPtr)
    window.removeEventListener('mousemove', handleMoveBeforeLong)
  }

  const handleTouchStart = (e: TouchEvent) => {
    e.stopPropagation()
    window.addEventListener('touchmove', handleTouchMove)
    timeoutPtr = window.setTimeout(() => {
      window.removeEventListener('touchmove', handleTouchMove)
      node.dispatchEvent(
        new CustomEvent('long', {
          detail: {
            target: e.target,
            category: data.category,
          },
        })
      )
      window.setTimeout(() => node.dispatchEvent(e), 0)
    }, TIME_MS)
  }

  const handleTouchMove = (e: TouchEvent) => {
    window.clearTimeout(timeoutPtr)
    window.removeEventListener('touchmove', handleTouchMove)
  }

  const handleTouchEnd = (e: TouchEvent) => {
    window.clearTimeout(timeoutPtr)
    window.removeEventListener('touchmove', handleTouchMove)
  }

  node.addEventListener('mousedown', handleMouseDown)
  node.addEventListener('mouseup', handleMouseUp)
  node.addEventListener('touchstart', handleTouchStart)
  node.addEventListener('touchend', handleTouchEnd)

  return {
    destroy: () => {
      node.removeEventListener('mousedown', handleMouseDown)
      node.removeEventListener('mouseup', handleMouseUp)
      node.removeEventListener('touchstart', handleTouchStart)
      node.removeEventListener('touchend', handleTouchEnd)
    },
  }
}
