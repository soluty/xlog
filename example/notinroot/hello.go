package notinroot

import (
	"github.com/soluty/xlog"
)

func Hello() {
	xlog.Debug("debug in not inroot")
}

// GOROOT=C:\Go #gosetup
// GOPATH=F:\badgo\badgo\uwproj\master\server #gosetup
// C:\Go\bin\go.exe build -ldflags "-X 'ifgs/gamesrv/game.GAME_VERSION=abc'" -tags "-tags etcd consul" -o F:\badgo\badgo\uwproj\master\server\bin\sfsrv.exe . #gosetup
// F:\badgo\badgo\uwproj\master\server\bin\sfsrv.exe -c -w #gosetup

// GOROOT=C:\Go #gosetup
// GOPATH=C:\Users\soluty\go #gosetup
// C:\Go\bin\go.exe build -o C:\Users\soluty\AppData\Local\Temp\___2xlog.exe github.com/soluty/xlog/example/notinroot/cmd #gosetup
// C:\Users\soluty\AppData\Local\Temp\___2xlog.exe #gosetup

// GOROOT=C:\Go #gosetup
// GOPATH=C:\Users\soluty\go #gosetup
// C:\Go\bin\go.exe build -o C:\Users\soluty\AppData\Local\Temp\___2xlog.exe . #gosetup
// C:\Users\soluty\AppData\Local\Temp\___2xlog.exe #gosetup
