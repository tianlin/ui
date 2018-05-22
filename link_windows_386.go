// 13 december 2015

package ui

// #cgo LDFLAGS:  ${SRCDIR}/libui_windows_386.a ${SRCDIR}/libui_windows_386.res.o ${SRCDIR}/libwinapi_d2d1.a ${SRCDIR}/libwinapi_dwrite.a ${SRCDIR}/libwinapi_comctl32.a
// /* note the order; also note the lack of uuid */
// #cgo LDFLAGS: -luser32 -lkernel32 -lusp10 -lgdi32  -luxtheme -lmsimg32 -lcomdlg32 -lole32 -loleaut32 -loleacc -static -static-libgcc -static-libstdc++
import "C"

func ensureMainThread() {
	// do nothing; Windows doesn't care which thread we're on so long as we don't change it after starting
}
