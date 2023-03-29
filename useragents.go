package uafaker

type userAgentStorageItem struct {
	userAgent string
	bitmap    uint64
}

var userAgentStorage = []userAgentStorageItem{
	// chrome/windows
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
