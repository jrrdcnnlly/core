package logging

// ANSI code value for escape sequences.
type ANSICode uint8

const (
	ANSIReset               ANSICode = 0
	ANSIBold                ANSICode = 1
	ANSIFaint               ANSICode = 2
	ANSIDim                          = ANSIFaint
	ANSIItalic              ANSICode = 3
	ANSIUnderline           ANSICode = 4
	ANSISlowBlink           ANSICode = 5
	ANSIRapidBlink          ANSICode = 6
	ANSIReverse             ANSICode = 7
	ANSIInvert                       = ANSIReverse
	ANSIConceal             ANSICode = 8
	ANSIHide                         = ANSIConceal
	ANSICrossedOut          ANSICode = 9
	ANSIStrike                       = ANSICrossedOut
	ANSIPrimaryFont         ANSICode = 10
	ANSIDefaultFont                  = ANSIPrimaryFont
	ANSIAlternativeFont1    ANSICode = 11
	ANSIAlternativeFont2    ANSICode = 12
	ANSIAlternativeFont3    ANSICode = 13
	ANSIAlternativeFont4    ANSICode = 14
	ANSIAlternativeFont5    ANSICode = 15
	ANSIAlternativeFont6    ANSICode = 16
	ANSIAlternativeFont7    ANSICode = 17
	ANSIAlternativeFont8    ANSICode = 18
	ANSIAlternativeFont9    ANSICode = 19
	ANSIGothicFont          ANSICode = 20
	ANSIDoubleUnderline     ANSICode = 21
	ANSINotBold                      = ANSIDoubleUnderline
	ANSINormalIntensity     ANSICode = 22
	ANSINotItalic           ANSICode = 23
	ANSINotUnderlined       ANSICode = 24
	ANSINotBlinking         ANSICode = 25
	ANSIProportionalSpacing ANSICode = 26
	ANSINotReversed         ANSICode = 27
	ANSIReveal              ANSICode = 28
	ANSINotCrossedOut       ANSICode = 29

	ANSIForegroundBlack   ANSICode = 30
	ANSIForegroundRed     ANSICode = 31
	ANSIForegroundGreen   ANSICode = 32
	ANSIForegroundYellow  ANSICode = 33
	ANSIForegroundBlue    ANSICode = 34
	ANSIForegroundMagenta ANSICode = 35
	ANSIForegroundCyan    ANSICode = 36
	ANSIForegroundWhite   ANSICode = 37
	ANSIForegroundDefault ANSICode = 39

	ANSIBackgroundBlack   ANSICode = 40
	ANSIBackgroundRed     ANSICode = 41
	ANSIBackgroundGreen   ANSICode = 42
	ANSIBackgroundYellow  ANSICode = 43
	ANSIBackgroundBlue    ANSICode = 44
	ANSIBackgroundMagenta ANSICode = 45
	ANSIBackgroundCyan    ANSICode = 46
	ANSIBackgroundWhite   ANSICode = 47
	ANSIBackgroundDefault ANSICode = 49

	ANSIForegroundBrightBlack   ANSICode = 90
	ANSIForegroundGray                   = ANSIForegroundBrightBlack
	ANSIForegroundBrightRed     ANSICode = 91
	ANSIForegroundBrightGreen   ANSICode = 92
	ANSIForegroundBrightYellow  ANSICode = 93
	ANSIForegroundBrightBlue    ANSICode = 94
	ANSIForegroundBrightMagenta ANSICode = 95
	ANSIForegroundBrightCyan    ANSICode = 96
	ANSIForegroundBrightWhite   ANSICode = 97

	ANSIBackgroundBrightBlack   ANSICode = 100
	ANSIBackgroundGray                   = ANSIBackgroundBrightBlack
	ANSIBackgroundBrightRed     ANSICode = 101
	ANSIBackgroundBrightGreen   ANSICode = 102
	ANSIBackgroundBrightYellow  ANSICode = 103
	ANSIBackgroundBrightBlue    ANSICode = 104
	ANSIBackgroundBrightMagenta ANSICode = 105
	ANSIBackgroundBrightCyan    ANSICode = 106
	ANSIBackgroundBrightWhite   ANSICode = 107
)
