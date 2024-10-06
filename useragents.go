package uafaker

type userAgentStorageItem struct {
	userAgent string
	bitmap    uint64
}

var userAgentStorage = []userAgentStorageItem{
	// chrome/windows
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36",
		bitmap:    0b100001001,
	},
	// chrome/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		bitmap:    0b100010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36",
		bitmap:    0b100010001,
	},
	// chrome/linux
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		bitmap:    0b100100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36",
		bitmap:    0b100100001,
	},
	// chrome/iOS
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.69 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.69 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.69 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/129.0.6668.46 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.98 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.98 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.98 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.34 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.34 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/128.0.6613.34 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.107 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.107 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.107 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.56 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.56 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/127.0.6533.56 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.190 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.190 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.190 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.108 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.108 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.108 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.54 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.54 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.54 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.35 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.35 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/126.0.6478.35 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.145 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.145 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.145 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.80 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/125.0.6422.33 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.111 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.111 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.111 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.88 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.88 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.88 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.71 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.71 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.71 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.68 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.68 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/124.0.6367.68 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/123.0.6312.52 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.89 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.62 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.62 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.62 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/122.0.6261.51 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/118.0.5993.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/118.0.5993.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/118.0.5993.92 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/113.0.5672.121 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/113.0.5672.121 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/113.0.5672.121 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/111.0.5563.101 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/111.0.5563.101 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/111.0.5563.101 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.114 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.114 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.114 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.83 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.83 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/110.0.5481.83 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/109.0.5414.112 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/109.0.5414.112 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/109.0.5414.112 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/103.0.5060.63 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/103.0.5060.63 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/103.0.5060.63 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/102.0.5005.87 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/102.0.5005.87 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/102.0.5005.87 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/101.0.4951.44 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/101.0.4951.44 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/101.0.4951.44 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/98.0.4758.97 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/98.0.4758.97 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/98.0.4758.97 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.163 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.163 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.163 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/87.0.4280.77 Mobile/15E148 Safari/604.1",
		bitmap:    0b101000010,
	},
	// chrome/android
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-A102U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q720) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-X420) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; LM-Q710(FGN)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36",
		bitmap:    0b110000010,
	},
	// firefox/windows
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000001001,
	},
	// firefox/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.7; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.7; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.7; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.6; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.6; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.6; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.5; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.4; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.3; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.3; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.3; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.0; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14.0; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.3; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.3; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.3; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.2; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.2; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13.2; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.4; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.4; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.4; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.2; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.2; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11.2; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000010001,
	},
	// firefox/linux
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:131.0) Gecko/20100101 Firefox/131.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:130.0) Gecko/20100101 Firefox/130.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:129.0) Gecko/20100101 Firefox/129.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:128.0) Gecko/20100101 Firefox/128.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:127.0) Gecko/20100101 Firefox/127.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:126.0) Gecko/20100101 Firefox/126.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:125.0) Gecko/20100101 Firefox/125.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:124.0) Gecko/20100101 Firefox/124.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:123.0) Gecko/20100101 Firefox/123.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:122.0) Gecko/20100101 Firefox/122.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:115.0) Gecko/20100101 Firefox/115.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:109.0) Gecko/20100101 Firefox/118.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:109.0) Gecko/20100101 Firefox/113.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:111.0) Gecko/20100101 Firefox/111.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:110.0) Gecko/20100101 Firefox/110.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:109.0) Gecko/20100101 Firefox/109.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:102.0) Gecko/20100101 Firefox/102.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:101.0) Gecko/20100101 Firefox/101.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:100.0) Gecko/20100101 Firefox/100.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:99.0) Gecko/20100101 Firefox/99.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:91.0) Gecko/20100101 Firefox/91.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:97.0) Gecko/20100101 Firefox/97.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:87.0) Gecko/20100101 Firefox/87.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux i686; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000100001,
	},
	// firefox/iOS
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/131.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/131.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/131.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_6_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/130.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_6_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/129.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/128.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/127.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/127.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/127.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/126.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/125.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/125.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/125.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/124.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_3_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/123.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/119.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/119.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/119.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/118.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/118.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/118.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/113.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/113.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/113.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/111.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/111.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/111.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 13_2_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/110.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 13_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/110.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/110.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 13_2 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/109.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 13_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/109.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/109.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/102.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/102.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/102.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/101.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/101.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/101.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/99.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/99.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/99.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 12_2_1 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/97.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 12_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/97.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_2_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/97.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 11_2_3 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/33.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 11_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/33.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/33.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 11_2_3 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 11_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 11_2_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 11_2_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone OS 11_2_2 like Mac OS X) AppleWebKit/604.5.6 (KHTML, like Gecko) FxiOS/32.0 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1001000010,
	},
	// firefox/android
	{
		userAgent: "Mozilla/5.0 (Android 15; Mobile; LG-M255; rv:131.0) Gecko/131.0 Firefox/131.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 15; Mobile; rv:131.0) Gecko/131.0 Firefox/131.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 15; Mobile; LG-M255; rv:130.0) Gecko/130.0 Firefox/130.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 15; Mobile; rv:130.0) Gecko/130.0 Firefox/130.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:129.0) Gecko/129.0 Firefox/129.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:129.0) Gecko/129.0 Firefox/129.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:128.0) Gecko/128.0 Firefox/128.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:128.0) Gecko/128.0 Firefox/128.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:127.0) Gecko/127.0 Firefox/127.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:127.0) Gecko/127.0 Firefox/127.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:126.0) Gecko/126.0 Firefox/126.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:126.0) Gecko/126.0 Firefox/126.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:125.0) Gecko/125.0 Firefox/125.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:125.0) Gecko/125.0 Firefox/125.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:124.0) Gecko/124.0 Firefox/124.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:124.0) Gecko/124.0 Firefox/124.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:123.0) Gecko/123.0 Firefox/123.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:123.0) Gecko/123.0 Firefox/123.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:122.0) Gecko/122.0 Firefox/122.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:122.0) Gecko/122.0 Firefox/122.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; LG-M255; rv:118.0) Gecko/118.0 Firefox/118.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 14; Mobile; rv:109.0) Gecko/118.0 Firefox/118.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; LG-M255; rv:113.0) Gecko/113.0 Firefox/113.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; rv:109.0) Gecko/113.0 Firefox/113.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; LG-M255; rv:111.0) Gecko/111.0 Firefox/111.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; rv:68.0) Gecko/68.0 Firefox/111.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; LG-M255; rv:110.0) Gecko/110.0 Firefox/110.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; rv:68.0) Gecko/68.0 Firefox/110.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; LG-M255; rv:109.0) Gecko/109.0 Firefox/109.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 13; Mobile; rv:68.0) Gecko/68.0 Firefox/109.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; LG-M255; rv:102.0) Gecko/102.0 Firefox/102.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/102.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; LG-M255; rv:101.0) Gecko/101.0 Firefox/101.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/101.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; LG-M255; rv:100.0) Gecko/100.0 Firefox/100.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/100.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; LG-M255; rv:99.0) Gecko/99.0 Firefox/99.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/99.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; LG-M255; rv:97.0) Gecko/97.0 Firefox/97.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 12; Mobile; rv:68.0) Gecko/68.0 Firefox/97.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 11; Mobile; rv:68.0) Gecko/68.0 Firefox/86.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 11; Mobile; LG-M255; rv:86.0) Gecko/86.0 Firefox/86.0",
		bitmap:    0b1010000010,
	},
	// safari/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_3_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_3_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
		bitmap:    0b10000010001,
	},
	// safari/iOS
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1",
		bitmap:    0b10001000010,
	},
	// ie/windows
	{
		userAgent: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; WOW64; Trident/4.0;)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 6.1)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2)",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 6.2; Trident/7.0; rv:11.0) like Gecko",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
		bitmap:    0b100000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
		bitmap:    0b100000001001,
	},
	// edge/windows
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.2792.79",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.2792.65",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/128.0.2739.90",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/128.0.2739.79",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.79",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.67",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.63",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.54",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/127.0.2651.105",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.105",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.98",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.86",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/126.0.2592.113",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.113",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.102",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.87",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.81",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.68",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/125.0.2535.92",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.92",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.85",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.79",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.67",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/124.0.2478.109",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.105",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.97",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.80",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.67",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/123.0.2420.97",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.97",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.81",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.65",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.113",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.106",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.92",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.92",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.80",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.66",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.63",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/121.0.2277.128",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.2277.128",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.61",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.57",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/109.0.1518.78",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.78",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 Edg/103.0.1264.44",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 Edg/103.0.1264.37",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edge/44.18363.8131",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/100.0.1185.39",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/100.0.1185.39",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/97.0.1072.69",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.41",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 Edg/89.0.774.75",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 Edg/89.0.774.68",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 Edg/89.0.774.63",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.63",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.57",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.50",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 Edg/89.0.774.48",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows Mobile 10; Android 10.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Mobile Safari/537.36 Edge/40.15254.603",
		bitmap:    0b1000000001010,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 Edg/89.0.774.45",
		bitmap:    0b1000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 Edg/88.0.705.81",
		bitmap:    0b1000000001001,
	},
	// edge/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.2792.79",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.2792.65",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/128.0.2739.90",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/128.0.2739.79",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.79",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.67",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.63",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.2739.54",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/127.0.2651.105",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.105",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.98",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.2651.86",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/126.0.2592.113",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.113",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.102",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.87",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.81",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/126.0.2592.68",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 Edg/125.0.2535.92",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.92",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.85",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.79",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.2535.67",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/124.0.2478.109",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.105",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.97",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.80",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.2478.67",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/123.0.2420.97",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.97",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.81",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.2420.65",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.113",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.106",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/122.0.2365.92",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.92",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.80",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.66",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.2365.63",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/121.0.2277.128",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.2277.128",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.61",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.57",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.42",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.54",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/109.0.1518.78",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 Edg/109.0.1518.78",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 Edg/103.0.1264.44",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 Edg/103.0.1264.37",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 Edg/102.0.1245.33",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/100.0.1185.39",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/100.0.1185.39",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 Edg/97.0.1072.69",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 Edg/90.0.818.39",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 Edg/89.0.774.63",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.63",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.57",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.48",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 Edg/89.0.774.48",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 Edg/89.0.774.45",
		bitmap:    0b1000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 Edg/88.0.705.81",
		bitmap:    0b1000000010001,
	},
	// edge/iOS
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 EdgiOS/129.2792.58 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 EdgiOS/128.2739.82 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/128.2739.82 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/128.2739.72 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/128.2739.72 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/128.2739.60 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/128.2739.42 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/127.2651.111 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/127.2651.102 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/127.2651.96 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/127.2651.81 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/127.2651.74 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.120 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.106 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.100 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.86 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.78 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.67 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/126.2592.56 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/125.2535.96 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/125.2535.87 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/125.2535.72 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/125.2535.60 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.105 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.105 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.89 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.89 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.71 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/124.2478.50 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/123.2420.108 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/123.2420.90 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/123.2420.74 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/123.2420.74 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/123.2420.56 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.99 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.99 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.86 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.78 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.71 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/122.2365.56 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/121.2277.133 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/121.2277.107 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/118.2088.60 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 EdgiOS/118.2088.52 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 EdgiOS/113.1774.42 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 EdgiOS/111.1661.58 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 EdgiOS/110.1587.54 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 EdgiOS/109.1518.80 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 EdgiOS/109.1518.72 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 EdgiOS/100.1185.50 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 EdgiOS/100.1185.50 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 EdgiOS/97.1072.69 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.3.7 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.1.10 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.1.10 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.1.10 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	// edge/android
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/129.0.2792.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/128.0.2739.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.111",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.90",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.90",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.90",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.90",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/127.0.2651.82",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.117",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.106",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.80",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.80",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.80",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/126.0.2592.80",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 EdgA/125.0.2535.96",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 EdgA/125.0.2535.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 EdgA/125.0.2535.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 EdgA/125.0.2535.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 EdgA/125.0.2535.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 EdgA/125.0.2535.72",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 EdgA/125.0.2535.72",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 EdgA/125.0.2535.72",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 EdgA/125.0.2535.72",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 EdgA/124.0.2478.104",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.87",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.71",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/124.0.2478.62",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.102",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.74",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/123.0.2420.61",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.86",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.76",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/122.0.2365.56",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 EdgA/121.0.2277.133",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 EdgA/121.0.2277.105",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 EdgA/121.0.2277.105",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 EdgA/121.0.2277.105",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 EdgA/121.0.2277.105",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/118.0.2088.58",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/118.0.2088.58",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/118.0.2088.58",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/118.0.2088.58",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/117.0.2045.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/117.0.2045.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/117.0.2045.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 EdgA/117.0.2045.65",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 EdgA/113.0.1774.38",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 EdgA/113.0.1774.38",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 EdgA/113.0.1774.38",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 EdgA/113.0.1774.38",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 EdgA/110.0.1587.66",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 EdgA/110.0.1587.66",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 EdgA/110.0.1587.66",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 EdgA/110.0.1587.66",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 EdgA/110.0.1587.54",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 EdgA/110.0.1587.54",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 EdgA/110.0.1587.54",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 EdgA/110.0.1587.54",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 EdgA/109.0.1518.70",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 EdgA/100.0.1185.50",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 EdgA/97.0.1072.69",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 EdgA/97.0.1072.69",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 EdgA/97.0.1072.69",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 EdgA/97.0.1072.69",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 EdgA/46.3.4.5155",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 EdgA/46.3.4.5155",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 EdgA/46.3.4.5155",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 EdgA/46.3.4.5155",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; HD1913) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; Pixel 3 XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 EdgA/46.1.2.5140",
		bitmap:    0b1000010000010,
	},
	// opera/windows
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/115.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/115.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/114.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/114.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 OPR/107.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 OPR/107.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 OPR/103.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 OPR/103.0.0.0",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 OPR/98.0.4759.39",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 OPR/98.0.4759.39",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.4719.17",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.4719.17",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 OPR/88.0.4412.40",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 OPR/88.0.4412.40",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 OPR/87.0.4390.36",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 OPR/87.0.4390.36",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 OPR/83.0.4254.27",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 OPR/83.0.4254.27",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 OPR/75.0.3969.171",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 OPR/75.0.3969.171",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/75.0.3969.149",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/75.0.3969.149",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 OPR/74.0.3911.160",
		bitmap:    0b10000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 OPR/74.0.3911.160",
		bitmap:    0b10000000001001,
	},
	// opera/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/115.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/114.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 OPR/107.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 OPR/103.0.0.0",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 OPR/98.0.4759.39",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.4719.17",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 OPR/88.0.4412.40",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 OPR/87.0.4390.36",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 OPR/83.0.4254.27",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 OPR/75.0.3969.171",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/75.0.3969.149",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 OPR/74.0.3911.160",
		bitmap:    0b10000000010001,
	},
	// opera/linux
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/115.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/114.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/113.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/112.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/111.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/110.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 OPR/108.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 OPR/107.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 OPR/103.0.0.0",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 OPR/98.0.4759.39",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.4719.17",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36 OPR/95.0.4635.25",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36 OPR/88.0.4412.40",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36 OPR/87.0.4390.36",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 OPR/86.0.4363.23",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 OPR/83.0.4254.27",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 Safari/537.36 OPR/75.0.3969.171",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/75.0.3969.149",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.232",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36 OPR/74.0.3911.203",
		bitmap:    0b10000000100001,
	},
	{
		userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Safari/537.36 OPR/74.0.3911.160",
		bitmap:    0b10000000100001,
	},
	// opera/android
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 Mobile Safari/537.36 OPR/76.2.4027.73374",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 Mobile Safari/537.36 OPR/63.3.3216.58675",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 OPR/63.0.3216.58473",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 OPR/63.0.3216.58473",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 Mobile Safari/537.36 OPR/63.0.3216.58473",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 OPR/62.3.3146.57763",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 OPR/62.3.3146.57763",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 Mobile Safari/537.36 OPR/62.3.3146.57763",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; VOG-L29) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-G970F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; Android 10; SM-N975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 Mobile Safari/537.36 OPR/61.2.3076.56749",
		bitmap:    0b10000010000010,
	},
	// yandex/windows
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.6.970 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.6.949 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.3.1248 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.3.1225 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.1.1144 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1144 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1116 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1069 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.7.1.1069 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.7.1.1004 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.6.3.785 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.6.3.745 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.6.3.745 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.4.5.506 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.4.5.506 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.4.4.1160 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.1.5.825 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.1.5.825 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 YaBrowser/23.9.0.2293 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 YaBrowser/23.9.0.2293 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.2 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.2 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 YaBrowser/22.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 YaBrowser/22.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 YaBrowser/22.3.3 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 YaBrowser/22.3.3 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 YaBrowser/22.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 YaBrowser/22.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000001001,
	},
	// yandex/macOS
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.6.970 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.6.949 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.3.1248 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.3.1225 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 YaBrowser/24.7.1.1144 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1144 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1116 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 YaBrowser/24.7.1.1069 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.7.1.1069 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.7.1.1004 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.6.3.785 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 YaBrowser/24.6.3.745 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.6.3.745 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 YaBrowser/24.4.5.506 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.4.5.506 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.4.4.1160 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 YaBrowser/24.1.5.825 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.1.5.825 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 YaBrowser/23.9.0.2325 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 YaBrowser/23.9.0.2293 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 YaBrowser/23.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.2 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 YaBrowser/23.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 YaBrowser/22.5.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 YaBrowser/22.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_3_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 YaBrowser/22.3.3 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 YaBrowser/22.1.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.72 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.3.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	{
		userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 YaBrowser/21.2.0 Yowser/2.5 Safari/537.36",
		bitmap:    0b100000000010001,
	},
	// yandex/iOS
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.218 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.218 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.218 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.194 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.194 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.9.0.194 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.430 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.430 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.430 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.389 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.389 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.9.389 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.8.420 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.8.420 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.8.420 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.7.799 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.7.799 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 YaBrowser/24.7.7.799 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.799 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.799 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.799 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.783 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.756 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.756 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.756 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.689 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.689 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.7.689 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.6.701 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.6.701 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.6 YaBrowser/24.7.6.701 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.701 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.701 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.701 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.601 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.601 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.6.601 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.583 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.583 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.583 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.565 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.565 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.565 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.505 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.505 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.5.505 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.819 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.819 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.819 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.792 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.792 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.4.792 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.3.656 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.3.656 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.3.656 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.802 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.802 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.802 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.724 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.724 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.2.724 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.1.771 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.1.771 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.1.771 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.7.0.2235 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.7.10 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.7.10 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.7.10 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.6.47 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.6.47 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.6.47 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.3.424 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.3.424 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.5 YaBrowser/24.6.3.424 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.424 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.424 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.424 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.341 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.341 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.3.341 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.2.364 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.2.364 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.2.364 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.1.342 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.1.342 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.1.342 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.0.1544 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.0.1544 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.6.0.1544 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.6.98 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.6.98 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.6.98 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.5.383 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.4.857 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.4.857 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.4.857 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.3.637 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.3.637 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.3.637 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.837 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.837 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.837 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.763 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.763 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.2.763 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.1.816 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.1.816 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.1.816 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.0.2362 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.0.2362 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.4.0.2362 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.4.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.1.178 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.0.174 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.0.174 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.2.0.174 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.9.283 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.8.486 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.8.486 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.8.486 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.7.499 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.7.499 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.7.499 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.6.455 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.5.438 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.5.438 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 YaBrowser/24.1.5.438 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 YaBrowser/24.1.4.440 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 YaBrowser/24.1.4.440 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 YaBrowser/24.1.4.440 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 YaBrowser/23.9.6.1142 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 YaBrowser/23.9.6.1142 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 YaBrowser/23.9.6.1142 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.3.330 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.2.420 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.2.420 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.4 YaBrowser/23.3.2.420 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.8.63 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.8.63 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.8.63 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.4.429 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.4.429 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 YaBrowser/23.1.4.429 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.1.569 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.1.569 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.1.569 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.0.1651 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.0.1651 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.7.0.1651 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.6.544 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.6.544 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.6.544 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.4.623 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.4.623 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.5.4.623 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.3.7.56 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.3.7.56 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 YaBrowser/22.3.7.56 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 YaBrowser/22.3.7.56 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 YaBrowser/22.3.7.56 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 YaBrowser/22.3.7.56 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 YaBrowser/22.1.6.583 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 YaBrowser/22.1.6.583 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 YaBrowser/22.1.6.583 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.6.356 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.6.356 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.6.356 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.5.713 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.5.713 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.5.713 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.4.913 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.4.913 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.4.913 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.952 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.952 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.952 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.914 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.914 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.3.914 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.2.778 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.2.778 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.2.778 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.1.806 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.1.806 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.1.806 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.0.2413 Mobile/15E148 Safari/604.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPad; CPU OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.0.2413 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPod touch; CPU iPhone 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 YaBrowser/21.2.0.2413 Mobile/15E148 Safari/605.1",
		bitmap:    0b100000001000010,
	},
	// yandex/android
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.81 YaBrowser/24.9.0.48 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 YaBrowser/24.7.9.61 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.69 YaBrowser/24.7.8.165 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 YaBrowser/24.7.8.165 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.6668.54 YaBrowser/24.7.8.148 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 YaBrowser/24.7.6.41 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.127 YaBrowser/24.7.5.124 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 YaBrowser/24.7.5.124 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 15; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 YaBrowser/24.7.4.147 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.99 YaBrowser/24.7.4.147 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 YaBrowser/24.7.4.142 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 YaBrowser/24.7.3.124 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.88 YaBrowser/24.7.2.116 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 YaBrowser/24.7.2.116 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.103 YaBrowser/24.7.0.309 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 YaBrowser/24.7.0.303 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.84 YaBrowser/24.6.6.26 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.64 YaBrowser/24.6.6.26 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 YaBrowser/24.6.5.46 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.186 YaBrowser/24.6.4.98 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 YaBrowser/24.6.4.98 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 YaBrowser/24.6.3.103 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.122 YaBrowser/24.6.1.82 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.110 YaBrowser/24.6.1.82 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.4.6478.71 YaBrowser/24.6.0.233 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.6478.50 YaBrowser/24.4.7.39 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.165 YaBrowser/24.4.7.39 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.147 YaBrowser/24.4.6.49 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 YaBrowser/24.4.6.49 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.113 YaBrowser/24.4.5.105 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.72 YaBrowser/24.4.4.99 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.53 YaBrowser/24.4.4.99 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.179 YaBrowser/24.4.3.96 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.171 YaBrowser/24.4.3.96 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.159 YaBrowser/24.4.3.96 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 YaBrowser/24.4.3.96 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.113 YaBrowser/24.4.0.308 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 YaBrowser/24.4.0.308 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.82 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.54 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.118 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.99 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.80 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.6312.40 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.119 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.105 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.90 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.64 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.6167.178 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 14; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.5993.80 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5672.76 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.5563.115 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.153 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.5481.63 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 13; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.5414.117 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.71 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.70 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.53 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.99 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 12; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.101 YaBrowser/21.3.4.59 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.66 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.105 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
}
