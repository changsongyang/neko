package desktop

import (
	"time"
	"demodesk/neko/internal/desktop/drop"
)

const (
	DELAY = 100 * time.Millisecond
)

func (manager *DesktopManagerCtx) DropFiles(x int, y int, files []string) {
	go drop.DragWindow(files)

	// TODO: Find a bettter way.
	time.Sleep(DELAY)
	manager.Move(10, 10)
	manager.ButtonDown(1)
	manager.Move(x, y)
	time.Sleep(DELAY)
	manager.Move(x, y)
	time.Sleep(DELAY)
	manager.ButtonUp(1)
}
