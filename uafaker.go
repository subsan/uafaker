package uafaker

import (
	"math/rand"
	"time"
)

const (
	filterDesktop = 1 << iota // 0b001
	filterMobile
	filterOS
	filterWindows
	filterMacOS
	filterLinux
	filterIOS
	filterAndroid
	filterChrome
	filterFirefox
	filterSafari
	filterIE
	filterEdge
	filterOpera
	filterYandex
)

const (
	maskDevice   = 0b000000000000011
	maskOS       = 0b000000011111000
	maskBrowsers = 0b111111100000100
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type UAFaker struct {
	filter uint64
	mask   uint64
}

func Desktop() *UAFaker {
	return &UAFaker{
		filter: filterDesktop,
		mask:   maskDevice,
	}
}
func (f *UAFaker) Desktop() *UAFaker {
	f.filter = f.filter | filterDesktop
	f.mask = f.mask | maskDevice
	return f
}
func Mobile() *UAFaker {
	return &UAFaker{
		filter: filterMobile,
		mask:   maskDevice,
	}
}
func (f *UAFaker) Mobile() *UAFaker {
	f.filter = f.filter | filterMobile
	f.mask = f.mask | maskDevice
	return f
}
func OS() *UAFaker {
	return &UAFaker{
		filter: filterOS,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) OS() *UAFaker {
	f.filter = f.filter | filterOS
	f.mask = f.mask | maskBrowsers
	return f
}
func Windows() *UAFaker {
	return &UAFaker{
		filter: filterWindows,
		mask:   maskOS,
	}
}
func (f *UAFaker) Windows() *UAFaker {
	f.filter = f.filter | filterWindows
	f.mask = f.mask | maskOS
	return f
}
func MacOS() *UAFaker {
	return &UAFaker{
		filter: filterMacOS,
		mask:   maskOS,
	}
}
func (f *UAFaker) MacOS() *UAFaker {
	f.filter = f.filter | filterMacOS
	f.mask = f.mask | maskOS
	return f
}
func Linux() *UAFaker {
	return &UAFaker{
		filter: filterLinux,
		mask:   maskOS,
	}
}
func (f *UAFaker) Linux() *UAFaker {
	f.filter = f.filter | filterLinux
	f.mask = f.mask | maskOS
	return f
}
func IOS() *UAFaker {
	return &UAFaker{
		filter: filterIOS,
		mask:   maskOS,
	}
}
func (f *UAFaker) IOS() *UAFaker {
	f.filter = f.filter | filterIOS
	f.mask = f.mask | maskOS
	return f
}
func Android() *UAFaker {
	return &UAFaker{
		filter: filterAndroid,
		mask:   maskOS,
	}
}
func (f *UAFaker) Android() *UAFaker {
	f.filter = f.filter | filterAndroid
	f.mask = f.mask | maskOS
	return f
}
func Chrome() *UAFaker {
	return &UAFaker{
		filter: filterChrome,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Chrome() *UAFaker {
	f.filter = f.filter | filterChrome
	f.mask = f.mask | maskBrowsers
	return f
}
func Firefox() *UAFaker {
	return &UAFaker{
		filter: filterFirefox,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Firefox() *UAFaker {
	f.filter = f.filter | filterFirefox
	f.mask = f.mask | maskBrowsers
	return f
}
func Safari() *UAFaker {
	return &UAFaker{
		filter: filterSafari,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Safari() *UAFaker {
	f.filter = f.filter | filterSafari
	f.mask = f.mask | maskBrowsers
	return f
}
func IE() *UAFaker {
	return &UAFaker{
		filter: filterIE,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) IE() *UAFaker {
	f.filter = f.filter | filterIE
	f.mask = f.mask | maskBrowsers
	return f
}
func Edge() *UAFaker {
	return &UAFaker{
		filter: filterEdge,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Edge() *UAFaker {
	f.filter = f.filter | filterEdge
	f.mask = f.mask | maskBrowsers
	return f
}
func Opera() *UAFaker {
	return &UAFaker{
		filter: filterOpera,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Opera() *UAFaker {
	f.filter = f.filter | filterOpera
	f.mask = f.mask | maskBrowsers
	return f
}
func Yandex() *UAFaker {
	return &UAFaker{
		filter: filterYandex,
		mask:   maskBrowsers,
	}
}
func (f *UAFaker) Yandex() *UAFaker {
	f.filter = f.filter | filterYandex
	f.mask = f.mask | maskBrowsers
	return f
}

func Random() string {
	return userAgentStorage[rand.Intn(len(userAgentStorage))].userAgent
}
func (f *UAFaker) Random() (str string) {
	var filteredBrowsersIndex []int

	if f.mask == 0 {
		return Random()
	}

	for i, s := range userAgentStorage {
		if (s.bitmap^f.filter)&s.bitmap&f.mask == 0b0 {
			filteredBrowsersIndex = append(filteredBrowsersIndex, i)
		}
	}

	if len(filteredBrowsersIndex) == 0 {
		return
	}

	return userAgentStorage[filteredBrowsersIndex[rand.Intn(len(filteredBrowsersIndex))]].userAgent
}
