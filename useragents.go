package uafaker

type userAgentStorageItem struct {
	userAgent string
	bitmap    uint64
}

var userAgentStorage = []userAgentStorageItem{
	// chrome/windows
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
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0",
		bitmap:    0b1000001001,
	},
	{
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:86.0) Gecko/20100101 Firefox/86.0",
		bitmap:    0b1000001001,
	},
	// firefox/macOS
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
		userAgent: "Mozilla/5.0 (Android 11; Mobile; rv:68.0) Gecko/68.0 Firefox/86.0",
		bitmap:    0b1010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Android 11; Mobile; LG-M255; rv:86.0) Gecko/86.0 Firefox/86.0",
		bitmap:    0b1010000010,
	},
	// safari/macOS
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
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.1.10 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	{
		userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 EdgiOS/46.1.10 Mobile/15E148 Safari/605.1.15",
		bitmap:    0b1000001000010,
	},
	// edge/android
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
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.86 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
	{
		userAgent: "Mozilla/5.0 (Linux; arm_64; Android 11; SM-G965F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.72 YaBrowser/20.11.5.113 Mobile Safari/537.36",
		bitmap:    0b100000010000010,
	},
}
