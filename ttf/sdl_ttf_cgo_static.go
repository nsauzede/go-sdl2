// +build static

package ttf

//#cgo CFLAGS: -I${SRCDIR}/../.go-sdl2-libs/include -I${SRCDIR}/../.go-sdl2-libs/include/SDL2
//#cgo LDFLAGS: -L${SRCDIR}/../.go-sdl2-libs
//#cgo linux,386 LDFLAGS: -lSDL2_ttf_linux_386 -Wl,--no-undefined -lfreetype_linux_386 -lz_linux_386 -lSDL2_linux_386 -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,amd64 LDFLAGS: -lSDL2_ttf_linux_amd64 -Wl,--no-undefined -lfreetype_linux_amd64 -lz_linux_amd64 -lSDL2_linux_amd64 -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo windows,386 LDFLAGS: -lSDL2_ttf_windows_386 -Wl,--no-undefined -lfreetype_windows_386 -lz_windows_386 -lSDL2_windows_386 -lSDL2main_windows_386 -mwindows -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -static-libgcc
//#cgo windows,amd64 LDFLAGS: -lSDL2_ttf_windows_amd64 -Wl,--no-undefined -lfreetype_windows_amd64 -lz_windows_amd64 -lSDL2_windows_amd64 -lSDL2main_windows_amd64 -mwindows -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -static-libgcc
//#cgo darwin,amd64 LDFLAGS: -lSDL2_ttf_darwin_amd64 -lm -liconv -lfreetype_darwin_amd64 -lz_darwin_amd64 -lSDL2_darwin_amd64 -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit
import "C"
