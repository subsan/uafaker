package uafaker

import (
	"math/rand"
	"time"
)

const (
	FilterDesktop = 1 << iota // 0b001
	FilterMobile
	FilterOS
	FilterWindows
	FilterMacOS
	FilterLinux
	FilterIOS
	FilterAndroid
	FilterChrome
	FilterFirefox
	FilterSafari
	FilterIE
	FilterEdge
	FilterOpera
	FilterYandex
)

type UAFaker struct {
	filters []uint64
}

func Desktop() *UAFaker {
	return &UAFaker{filters: []uint64{FilterDesktop}}

}
func (f *UAFaker) Desktop() *UAFaker {
	f.filters = append(f.filters, FilterDesktop)
	return f
}
func Mobile() *UAFaker {
	return &UAFaker{filters: []uint64{FilterMobile}}

}
func (f *UAFaker) Mobile() *UAFaker {
	f.filters = append(f.filters, FilterMobile)
	return f
}
func OS() *UAFaker {
	return &UAFaker{filters: []uint64{FilterOS}}

}
func (f *UAFaker) OS() *UAFaker {
	f.filters = append(f.filters, FilterOS)
	return f
}
func Windows() *UAFaker {
	return &UAFaker{filters: []uint64{FilterWindows}}

}
func (f *UAFaker) Windows() *UAFaker {
	f.filters = append(f.filters, FilterWindows)
	return f
}
func MacOS() *UAFaker {
	return &UAFaker{filters: []uint64{FilterMacOS}}

}
func (f *UAFaker) MacOS() *UAFaker {
	f.filters = append(f.filters, FilterMacOS)
	return f
}
func Linux() *UAFaker {
	return &UAFaker{filters: []uint64{FilterLinux}}

}
func (f *UAFaker) Linux() *UAFaker {
	f.filters = append(f.filters, FilterLinux)
	return f
}
func IOS() *UAFaker {
	return &UAFaker{filters: []uint64{FilterIOS}}

}
func (f *UAFaker) IOS() *UAFaker {
	f.filters = append(f.filters, FilterIOS)
	return f
}
func Android() *UAFaker {
	return &UAFaker{filters: []uint64{FilterAndroid}}

}
func (f *UAFaker) Android() *UAFaker {
	f.filters = append(f.filters, FilterAndroid)
	return f
}
func Chrome() *UAFaker {
	return &UAFaker{filters: []uint64{FilterChrome}}

}
func (f *UAFaker) Chrome() *UAFaker {
	f.filters = append(f.filters, FilterChrome)
	return f
}
func Firefox() *UAFaker {
	return &UAFaker{filters: []uint64{FilterFirefox}}

}
func (f *UAFaker) Firefox() *UAFaker {
	f.filters = append(f.filters, FilterFirefox)
	return f
}
func Safari() *UAFaker {
	return &UAFaker{filters: []uint64{FilterSafari}}

}
func (f *UAFaker) Safari() *UAFaker {
	f.filters = append(f.filters, FilterSafari)
	return f
}
func IE() *UAFaker {
	return &UAFaker{filters: []uint64{FilterIE}}

}
func (f *UAFaker) IE() *UAFaker {
	f.filters = append(f.filters, FilterIE)
	return f
}
func Edge() *UAFaker {
	return &UAFaker{filters: []uint64{FilterEdge}}

}
func (f *UAFaker) Edge() *UAFaker {
	f.filters = append(f.filters, FilterEdge)
	return f
}
func Opera() *UAFaker {
	return &UAFaker{filters: []uint64{FilterOpera}}

}
func (f *UAFaker) Opera() *UAFaker {
	f.filters = append(f.filters, FilterOpera)
	return f
}
func Yandex() *UAFaker {
	return &UAFaker{filters: []uint64{FilterYandex}}

}
func (f *UAFaker) Yandex() *UAFaker {
	f.filters = append(f.filters, FilterYandex)
	return f
}

func Random() string {
	rand.Seed(time.Now().UTC().UnixNano())

	return UserAgentStorage[rand.Intn(len(UserAgentStorage))].userAgent
}
func (f *UAFaker) Random() (str string) {
	rand.Seed(time.Now().UTC().UnixNano())

	var filteredBrowsersIndex []int
	var filter uint64

	if len(f.filters) == 0 {
		return UserAgentStorage[rand.Intn(len(UserAgentStorage))].userAgent
	}

	for _, s := range f.filters {
		filter = filter | s
	}

	for i, s := range UserAgentStorage {
		if s.bitmap&filter == filter {
			filteredBrowsersIndex = append(filteredBrowsersIndex, i)
		}
	}

	if len(filteredBrowsersIndex) == 0 {
		return
	}

	return UserAgentStorage[filteredBrowsersIndex[rand.Intn(len(filteredBrowsersIndex))]].userAgent
}
