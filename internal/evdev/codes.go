package evdev

var eventTypes = [...]uint16{
	EvSyn, EvKey, EvRel, EvAbs, EvMsc, EvSw, EvLed, EvSnd, EvRep, EvFf, EvFfStatus,
}

const ( // event types
	EvSyn      = 0x00
	EvKey      = 0x01
	EvRel      = 0x02
	EvAbs      = 0x03
	EvMsc      = 0x04
	EvSw       = 0x05
	EvLed      = 0x11
	EvSnd      = 0x12
	EvRep      = 0x14
	EvFf       = 0x15
	EvPwr      = 0x16
	EvFfStatus = 0x17
)

const ( // event types max
	EvMax  = 0x1f
	SynMax = 0xf
	KeyMax = 0x2ff
	RelMax = 0x0f
	AbsMax = 0x3f
	SwMax  = 0x0f
	MscMax = 0x07
	LedMax = 0x0f
	RepMax = 0x01
	SndMax = 0x07
)

const ( // keys
	KeyReserved                = 0
	KeyEsc                     = 1
	Key1                       = 2
	Key2                       = 3
	Key3                       = 4
	Key4                       = 5
	Key5                       = 6
	Key6                       = 7
	Key7                       = 8
	Key8                       = 9
	Key9                       = 10
	Key0                       = 11
	KeyMinus                   = 12
	KeyEqual                   = 13
	KeyBackspace               = 14
	KeyTab                     = 15
	KeyQ                       = 16
	KeyW                       = 17
	KeyE                       = 18
	KeyR                       = 19
	KeyT                       = 20
	KeyY                       = 21
	KeyU                       = 22
	KeyI                       = 23
	KeyO                       = 24
	KeyP                       = 25
	KeyLeftbrace               = 26
	KeyRightbrace              = 27
	KeyEnter                   = 28
	KeyLeftctrl                = 29
	KeyA                       = 30
	KeyS                       = 31
	KeyD                       = 32
	KeyF                       = 33
	KeyG                       = 34
	KeyH                       = 35
	KeyJ                       = 36
	KeyK                       = 37
	KeyL                       = 38
	KeySemicolon               = 39
	KeyApostrophe              = 40
	KeyGrave                   = 41
	KeyLeftshift               = 42
	KeyBackslash               = 43
	KeyZ                       = 44
	KeyX                       = 45
	KeyC                       = 46
	KeyV                       = 47
	KeyB                       = 48
	KeyN                       = 49
	KeyM                       = 50
	KeyComma                   = 51
	KeyDot                     = 52
	KeySlash                   = 53
	KeyRightshift              = 54
	KeyKpasterisk              = 55
	KeyLeftalt                 = 56
	KeySpace                   = 57
	KeyCapslock                = 58
	KeyF1                      = 59
	KeyF2                      = 60
	KeyF3                      = 61
	KeyF4                      = 62
	KeyF5                      = 63
	KeyF6                      = 64
	KeyF7                      = 65
	KeyF8                      = 66
	KeyF9                      = 67
	KeyF10                     = 68
	KeyNumlock                 = 69
	KeyScrolllock              = 70
	KeyKp7                     = 71
	KeyKp8                     = 72
	KeyKp9                     = 73
	KeyKpminus                 = 74
	KeyKp4                     = 75
	KeyKp5                     = 76
	KeyKp6                     = 77
	KeyKpplus                  = 78
	KeyKp1                     = 79
	KeyKp2                     = 80
	KeyKp3                     = 81
	KeyKp0                     = 82
	KeyKpdot                   = 83
	KeyZenkakuhankaku          = 85
	Key102nd                   = 86
	KeyF11                     = 87
	KeyF12                     = 88
	KeyRo                      = 89
	KeyKatakana                = 90
	KeyHiragana                = 91
	KeyHenkan                  = 92
	KeyKatakanahiragana        = 93
	KeyMuhenkan                = 94
	KeyKpjpcomma               = 95
	KeyKpenter                 = 96
	KeyRightctrl               = 97
	KeyKpslash                 = 98
	KeySysrq                   = 99
	KeyRightalt                = 100
	KeyLinefeed                = 101
	KeyHome                    = 102
	KeyUp                      = 103
	KeyPageup                  = 104
	KeyLeft                    = 105
	KeyRight                   = 106
	KeyEnd                     = 107
	KeyDown                    = 108
	KeyPagedown                = 109
	KeyInsert                  = 110
	KeyDelete                  = 111
	KeyMacro                   = 112
	KeyMute                    = 113
	KeyVolumedown              = 114
	KeyVolumeup                = 115
	KeyPower                   = 116
	KeyKpequal                 = 117
	KeyKpplusminus             = 118
	KeyPause                   = 119
	KeyScale                   = 120
	KeyKpcomma                 = 121
	KeyHangeul                 = 122
	KeyHanja                   = 123
	KeyYen                     = 124
	KeyLeftmeta                = 125
	KeyRightmeta               = 126
	KeyCompose                 = 127
	KeyStop                    = 128
	KeyAgain                   = 129
	KeyProps                   = 130
	KeyUndo                    = 131
	KeyFront                   = 132
	KeyCopy                    = 133
	KeyOpen                    = 134
	KeyPaste                   = 135
	KeyFind                    = 136
	KeyCut                     = 137
	KeyHelp                    = 138
	KeyMenu                    = 139
	KeyCalc                    = 140
	KeySetup                   = 141
	KeySleep                   = 142
	KeyWakeup                  = 143
	KeyFile                    = 144
	KeySendfile                = 145
	KeyDeletefile              = 146
	KeyXfer                    = 147
	KeyProg1                   = 148
	KeyProg2                   = 149
	KeyWww                     = 150
	KeyMsdos                   = 151
	KeyCoffee                  = 152
	KeyRotateDisplay           = 153
	KeyCyclewindows            = 154
	KeyMail                    = 155
	KeyBookmarks               = 156
	KeyComputer                = 157
	KeyBack                    = 158
	KeyForward                 = 159
	KeyClosecd                 = 160
	KeyEjectcd                 = 161
	KeyEjectclosecd            = 162
	KeyNextsong                = 163
	KeyPlaypause               = 164
	KeyPrevioussong            = 165
	KeyStopcd                  = 166
	KeyRecord                  = 167
	KeyRewind                  = 168
	KeyPhone                   = 169
	KeyIso                     = 170
	KeyConfig                  = 171
	KeyHomepage                = 172
	KeyRefresh                 = 173
	KeyExit                    = 174
	KeyMove                    = 175
	KeyEdit                    = 176
	KeyScrollup                = 177
	KeyScrolldown              = 178
	KeyKpleftparen             = 179
	KeyKprightparen            = 180
	KeyNew                     = 181
	KeyRedo                    = 182
	KeyF13                     = 183
	KeyF14                     = 184
	KeyF15                     = 185
	KeyF16                     = 186
	KeyF17                     = 187
	KeyF18                     = 188
	KeyF19                     = 189
	KeyF20                     = 190
	KeyF21                     = 191
	KeyF22                     = 192
	KeyF23                     = 193
	KeyF24                     = 194
	KeyPlaycd                  = 200
	KeyPausecd                 = 201
	KeyProg3                   = 202
	KeyProg4                   = 203
	KeyDashboard               = 204
	KeySuspend                 = 205
	KeyClose                   = 206
	KeyPlay                    = 207
	KeyFastforward             = 208
	KeyBassboost               = 209
	KeyPrint                   = 210
	KeyHp                      = 211
	KeyCamera                  = 212
	KeySound                   = 213
	KeyQuestion                = 214
	KeyEmail                   = 215
	KeyChat                    = 216
	KeySearch                  = 217
	KeyConnect                 = 218
	KeyFinance                 = 219
	KeySport                   = 220
	KeyShop                    = 221
	KeyAlterase                = 222
	KeyCancel                  = 223
	KeyBrightnessdown          = 224
	KeyBrightnessup            = 225
	KeyMedia                   = 226
	KeySwitchvideomode         = 227
	KeyKbdillumtoggle          = 228
	KeyKbdillumdown            = 229
	KeyKbdillumup              = 230
	KeySend                    = 231
	KeyReply                   = 232
	KeyForwardmail             = 233
	KeySave                    = 234
	KeyDocuments               = 235
	KeyBattery                 = 236
	KeyBluetooth               = 237
	KeyWlan                    = 238
	KeyUwb                     = 239
	KeyUnknown                 = 240
	KeyVideoNext               = 241
	KeyVideoPrev               = 242
	KeyBrightnessCycle         = 243
	KeyBrightnessAuto          = 244
	KeyDisplayOff              = 245
	KeyWwan                    = 246
	KeyRfkill                  = 247
	KeyMicmute                 = 248
	BtnMisc                    = 0x100
	Btn0                       = 0x100
	Btn1                       = 0x101
	Btn2                       = 0x102
	Btn3                       = 0x103
	Btn4                       = 0x104
	Btn5                       = 0x105
	Btn6                       = 0x106
	Btn7                       = 0x107
	Btn8                       = 0x108
	Btn9                       = 0x109
	BtnMouse                   = 0x110
	BtnLeft                    = 0x110
	BtnRight                   = 0x111
	BtnMiddle                  = 0x112
	BtnSide                    = 0x113
	BtnExtra                   = 0x114
	BtnForward                 = 0x115
	BtnBack                    = 0x116
	BtnTask                    = 0x117
	BtnJoystick                = 0x120
	BtnTrigger                 = 0x120
	BtnThumb                   = 0x121
	BtnThumb2                  = 0x122
	BtnTop                     = 0x123
	BtnTop2                    = 0x124
	BtnPinkie                  = 0x125
	BtnBase                    = 0x126
	BtnBase2                   = 0x127
	BtnBase3                   = 0x128
	BtnBase4                   = 0x129
	BtnBase5                   = 0x12a
	BtnBase6                   = 0x12b
	BtnDead                    = 0x12f
	BtnGamepad                 = 0x130
	BtnSouth                   = 0x130
	BtnEast                    = 0x131
	BtnC                       = 0x132
	BtnNorth                   = 0x133
	BtnWest                    = 0x134
	BtnZ                       = 0x135
	BtnTl                      = 0x136
	BtnTr                      = 0x137
	BtnTl2                     = 0x138
	BtnTr2                     = 0x139
	BtnSelect                  = 0x13a
	BtnStart                   = 0x13b
	BtnMode                    = 0x13c
	BtnThumbl                  = 0x13d
	BtnThumbr                  = 0x13e
	BtnDigi                    = 0x140
	BtnToolPen                 = 0x140
	BtnToolRubber              = 0x141
	BtnToolBrush               = 0x142
	BtnToolPencil              = 0x143
	BtnToolAirbrush            = 0x144
	BtnToolFinger              = 0x145
	BtnToolMouse               = 0x146
	BtnToolLens                = 0x147
	BtnToolQuinttap            = 0x148
	BtnTouch                   = 0x14a
	BtnStylus                  = 0x14b
	BtnStylus2                 = 0x14c
	BtnToolDoubletap           = 0x14d
	BtnToolTripletap           = 0x14e
	BtnToolQuadtap             = 0x14f
	BtnWheel                   = 0x150
	BtnGearDown                = 0x150
	BtnGearUp                  = 0x151
	KeyOk                      = 0x160
	KeySelect                  = 0x161
	KeyGoto                    = 0x162
	KeyClear                   = 0x163
	KeyPower2                  = 0x164
	KeyOption                  = 0x165
	KeyInfo                    = 0x166
	KeyTime                    = 0x167
	KeyVendor                  = 0x168
	KeyArchive                 = 0x169
	KeyProgram                 = 0x16a
	KeyChannel                 = 0x16b
	KeyFavorites               = 0x16c
	KeyEpg                     = 0x16d
	KeyPvr                     = 0x16e
	KeyMhp                     = 0x16f
	KeyLanguage                = 0x170
	KeyTitle                   = 0x171
	KeySubtitle                = 0x172
	KeyAngle                   = 0x173
	KeyZoom                    = 0x174
	KeyMode                    = 0x175
	KeyKeyboard                = 0x176
	KeyScreen                  = 0x177
	KeyPc                      = 0x178
	KeyTv                      = 0x179
	KeyTv2                     = 0x17a
	KeyVcr                     = 0x17b
	KeyVcr2                    = 0x17c
	KeySat                     = 0x17d
	KeySat2                    = 0x17e
	KeyCd                      = 0x17f
	KeyTape                    = 0x180
	KeyRadio                   = 0x181
	KeyTuner                   = 0x182
	KeyPlayer                  = 0x183
	KeyText                    = 0x184
	KeyDvd                     = 0x185
	KeyAux                     = 0x186
	KeyMp3                     = 0x187
	KeyAudio                   = 0x188
	KeyVideo                   = 0x189
	KeyDirectory               = 0x18a
	KeyList                    = 0x18b
	KeyMemo                    = 0x18c
	KeyCalendar                = 0x18d
	KeyRed                     = 0x18e
	KeyGreen                   = 0x18f
	KeyYellow                  = 0x190
	KeyBlue                    = 0x191
	KeyChannelup               = 0x192
	KeyChanneldown             = 0x193
	KeyFirst                   = 0x194
	KeyLast                    = 0x195
	KeyAb                      = 0x196
	KeyNext                    = 0x197
	KeyRestart                 = 0x198
	KeySlow                    = 0x199
	KeyShuffle                 = 0x19a
	KeyBreak                   = 0x19b
	KeyPrevious                = 0x19c
	KeyDigits                  = 0x19d
	KeyTeen                    = 0x19e
	KeyTwen                    = 0x19f
	KeyVideophone              = 0x1a0
	KeyGames                   = 0x1a1
	KeyZoomin                  = 0x1a2
	KeyZoomout                 = 0x1a3
	KeyZoomreset               = 0x1a4
	KeyWordprocessor           = 0x1a5
	KeyEditor                  = 0x1a6
	KeySpreadsheet             = 0x1a7
	KeyGraphicseditor          = 0x1a8
	KeyPresentation            = 0x1a9
	KeyDatabase                = 0x1aa
	KeyNews                    = 0x1ab
	KeyVoicemail               = 0x1ac
	KeyAddressbook             = 0x1ad
	KeyMessenger               = 0x1ae
	KeyDisplaytoggle           = 0x1af
	KeySpellcheck              = 0x1b0
	KeyLogoff                  = 0x1b1
	KeyDollar                  = 0x1b2
	KeyEuro                    = 0x1b3
	KeyFrameback               = 0x1b4
	KeyFrameforward            = 0x1b5
	KeyContextMenu             = 0x1b6
	KeyMediaRepeat             = 0x1b7
	Key10channelsup            = 0x1b8
	Key10channelsdown          = 0x1b9
	KeyImages                  = 0x1ba
	KeyDelEol                  = 0x1c0
	KeyDelEos                  = 0x1c1
	KeyInsLine                 = 0x1c2
	KeyDelLine                 = 0x1c3
	KeyFn                      = 0x1d0
	KeyFnEsc                   = 0x1d1
	KeyFnF1                    = 0x1d2
	KeyFnF2                    = 0x1d3
	KeyFnF3                    = 0x1d4
	KeyFnF4                    = 0x1d5
	KeyFnF5                    = 0x1d6
	KeyFnF6                    = 0x1d7
	KeyFnF7                    = 0x1d8
	KeyFnF8                    = 0x1d9
	KeyFnF9                    = 0x1da
	KeyFnF10                   = 0x1db
	KeyFnF11                   = 0x1dc
	KeyFnF12                   = 0x1dd
	KeyFn1                     = 0x1de
	KeyFn2                     = 0x1df
	KeyFnD                     = 0x1e0
	KeyFnE                     = 0x1e1
	KeyFnF                     = 0x1e2
	KeyFnS                     = 0x1e3
	KeyFnB                     = 0x1e4
	KeyBrlDot1                 = 0x1f1
	KeyBrlDot2                 = 0x1f2
	KeyBrlDot3                 = 0x1f3
	KeyBrlDot4                 = 0x1f4
	KeyBrlDot5                 = 0x1f5
	KeyBrlDot6                 = 0x1f6
	KeyBrlDot7                 = 0x1f7
	KeyBrlDot8                 = 0x1f8
	KeyBrlDot9                 = 0x1f9
	KeyBrlDot10                = 0x1fa
	KeyNumeric0                = 0x200
	KeyNumeric1                = 0x201
	KeyNumeric2                = 0x202
	KeyNumeric3                = 0x203
	KeyNumeric4                = 0x204
	KeyNumeric5                = 0x205
	KeyNumeric6                = 0x206
	KeyNumeric7                = 0x207
	KeyNumeric8                = 0x208
	KeyNumeric9                = 0x209
	KeyNumericStar             = 0x20a
	KeyNumericPound            = 0x20b
	KeyNumericA                = 0x20c
	KeyNumericB                = 0x20d
	KeyNumericC                = 0x20e
	KeyNumericD                = 0x20f
	KeyCameraFocus             = 0x210
	KeyWpsButton               = 0x211
	KeyTouchpadToggle          = 0x212
	KeyTouchpadOn              = 0x213
	KeyTouchpadOff             = 0x214
	KeyCameraZoomin            = 0x215
	KeyCameraZoomout           = 0x216
	KeyCameraUp                = 0x217
	KeyCameraDown              = 0x218
	KeyCameraLeft              = 0x219
	KeyCameraRight             = 0x21a
	KeyAttendantOn             = 0x21b
	KeyAttendantOff            = 0x21c
	KeyAttendantToggle         = 0x21d
	KeyLightsToggle            = 0x21e
	BtnDpadUp                  = 0x220
	BtnDpadDown                = 0x221
	BtnDpadLeft                = 0x222
	BtnDpadRight               = 0x223
	KeyAlsToggle               = 0x230
	KeyButtonconfig            = 0x240
	KeyTaskmanager             = 0x241
	KeyJournal                 = 0x242
	KeyControlpanel            = 0x243
	KeyAppselect               = 0x244
	KeyScreensaver             = 0x245
	KeyVoicecommand            = 0x246
	KeyBrightnessMin           = 0x250
	KeyBrightnessMax           = 0x251
	KeyKbdinputassistPrev      = 0x260
	KeyKbdinputassistNext      = 0x261
	KeyKbdinputassistPrevgroup = 0x262
	KeyKbdinputassistNextgroup = 0x263
	KeyKbdinputassistAccept    = 0x264
	KeyKbdinputassistCancel    = 0x265
	KeyRightUp                 = 0x266
	KeyRightDown               = 0x267
	KeyLeftUp                  = 0x268
	KeyLeftDown                = 0x269
	KeyRootMenu                = 0x26a
	KeyMediaTopMenu            = 0x26b
	KeyNumeric11               = 0x26c
	KeyNumeric12               = 0x26d
	KeyAudioDesc               = 0x26e
	Key3dMode                  = 0x26f
	KeyNextFavorite            = 0x270
	KeyStopRecord              = 0x271
	KeyPauseRecord             = 0x272
	KeyVod                     = 0x273
	KeyUnmute                  = 0x274
	KeyFastreverse             = 0x275
	KeySlowreverse             = 0x276
	KeyData                    = 0x275
	BtnTriggerHappy            = 0x2c0
	BtnTriggerHappy1           = 0x2c0
	BtnTriggerHappy2           = 0x2c1
	BtnTriggerHappy3           = 0x2c2
	BtnTriggerHappy4           = 0x2c3
	BtnTriggerHappy5           = 0x2c4
	BtnTriggerHappy6           = 0x2c5
	BtnTriggerHappy7           = 0x2c6
	BtnTriggerHappy8           = 0x2c7
	BtnTriggerHappy9           = 0x2c8
	BtnTriggerHappy10          = 0x2c9
	BtnTriggerHappy11          = 0x2ca
	BtnTriggerHappy12          = 0x2cb
	BtnTriggerHappy13          = 0x2cc
	BtnTriggerHappy14          = 0x2cd
	BtnTriggerHappy15          = 0x2ce
	BtnTriggerHappy16          = 0x2cf
	BtnTriggerHappy17          = 0x2d0
	BtnTriggerHappy18          = 0x2d1
	BtnTriggerHappy19          = 0x2d2
	BtnTriggerHappy20          = 0x2d3
	BtnTriggerHappy21          = 0x2d4
	BtnTriggerHappy22          = 0x2d5
	BtnTriggerHappy23          = 0x2d6
	BtnTriggerHappy24          = 0x2d7
	BtnTriggerHappy25          = 0x2d8
	BtnTriggerHappy26          = 0x2d9
	BtnTriggerHappy27          = 0x2da
	BtnTriggerHappy28          = 0x2db
	BtnTriggerHappy29          = 0x2dc
	BtnTriggerHappy30          = 0x2dd
	BtnTriggerHappy31          = 0x2de
	BtnTriggerHappy32          = 0x2df
	BtnTriggerHappy33          = 0x2e0
	BtnTriggerHappy34          = 0x2e1
	BtnTriggerHappy35          = 0x2e2
	BtnTriggerHappy36          = 0x2e3
	BtnTriggerHappy37          = 0x2e4
	BtnTriggerHappy38          = 0x2e5
	BtnTriggerHappy39          = 0x2e6
	BtnTriggerHappy40          = 0x2e7
)

var KeyMap = map[uint16]string{
	KeyReserved:                "Reserved",
	KeyEsc:                     "Esc",
	Key1:                       "1",
	Key2:                       "2",
	Key3:                       "3",
	Key4:                       "4",
	Key5:                       "5",
	Key6:                       "6",
	Key7:                       "7",
	Key8:                       "8",
	Key9:                       "9",
	Key0:                       "0",
	KeyMinus:                   "Minus",
	KeyEqual:                   "Equal",
	KeyBackspace:               "Backspace",
	KeyTab:                     "Tab",
	KeyQ:                       "Q",
	KeyW:                       "W",
	KeyE:                       "E",
	KeyR:                       "R",
	KeyT:                       "T",
	KeyY:                       "Y",
	KeyU:                       "U",
	KeyI:                       "I",
	KeyO:                       "O",
	KeyP:                       "P",
	KeyLeftbrace:               "Leftbrace",
	KeyRightbrace:              "Rightbrace",
	KeyEnter:                   "Enter",
	KeyLeftctrl:                "Leftctrl",
	KeyA:                       "A",
	KeyS:                       "S",
	KeyD:                       "D",
	KeyF:                       "F",
	KeyG:                       "G",
	KeyH:                       "H",
	KeyJ:                       "J",
	KeyK:                       "K",
	KeyL:                       "L",
	KeySemicolon:               "Semicolon",
	KeyApostrophe:              "Apostrophe",
	KeyGrave:                   "Grave",
	KeyLeftshift:               "Leftshift",
	KeyBackslash:               "Backslash",
	KeyZ:                       "Z",
	KeyX:                       "X",
	KeyC:                       "C",
	KeyV:                       "V",
	KeyB:                       "B",
	KeyN:                       "N",
	KeyM:                       "M",
	KeyComma:                   "Comma",
	KeyDot:                     "Dot",
	KeySlash:                   "Slash",
	KeyRightshift:              "Rightshift",
	KeyKpasterisk:              "Kpasterisk",
	KeyLeftalt:                 "Leftalt",
	KeySpace:                   "Space",
	KeyCapslock:                "Capslock",
	KeyF1:                      "F1",
	KeyF2:                      "F2",
	KeyF3:                      "F3",
	KeyF4:                      "F4",
	KeyF5:                      "F5",
	KeyF6:                      "F6",
	KeyF7:                      "F7",
	KeyF8:                      "F8",
	KeyF9:                      "F9",
	KeyF10:                     "F10",
	KeyNumlock:                 "Numlock",
	KeyScrolllock:              "Scrolllock",
	KeyKp7:                     "Kp7",
	KeyKp8:                     "Kp8",
	KeyKp9:                     "Kp9",
	KeyKpminus:                 "Kpminus",
	KeyKp4:                     "Kp4",
	KeyKp5:                     "Kp5",
	KeyKp6:                     "Kp6",
	KeyKpplus:                  "Kpplus",
	KeyKp1:                     "Kp1",
	KeyKp2:                     "Kp2",
	KeyKp3:                     "Kp3",
	KeyKp0:                     "Kp0",
	KeyKpdot:                   "Kpdot",
	KeyZenkakuhankaku:          "Zenkakuhankaku",
	Key102nd:                   "102nd",
	KeyF11:                     "F11",
	KeyF12:                     "F12",
	KeyRo:                      "Ro",
	KeyKatakana:                "Katakana",
	KeyHiragana:                "Hiragana",
	KeyHenkan:                  "Henkan",
	KeyKatakanahiragana:        "Katakanahiragana",
	KeyMuhenkan:                "Muhenkan",
	KeyKpjpcomma:               "Kpjpcomma",
	KeyKpenter:                 "Kpenter",
	KeyRightctrl:               "Rightctrl",
	KeyKpslash:                 "Kpslash",
	KeySysrq:                   "Sysrq",
	KeyRightalt:                "Rightalt",
	KeyLinefeed:                "Linefeed",
	KeyHome:                    "Home",
	KeyUp:                      "Up",
	KeyPageup:                  "Pageup",
	KeyLeft:                    "Left",
	KeyRight:                   "Right",
	KeyEnd:                     "End",
	KeyDown:                    "Down",
	KeyPagedown:                "Pagedown",
	KeyInsert:                  "Insert",
	KeyDelete:                  "Delete",
	KeyMacro:                   "Macro",
	KeyMute:                    "Mute",
	KeyVolumedown:              "Volumedown",
	KeyVolumeup:                "Volumeup",
	KeyPower:                   "Power",
	KeyKpequal:                 "Kpequal",
	KeyKpplusminus:             "Kpplusminus",
	KeyPause:                   "Pause",
	KeyScale:                   "Scale",
	KeyKpcomma:                 "Kpcomma",
	KeyHangeul:                 "Hangeul",
	KeyHanja:                   "Hanja",
	KeyYen:                     "Yen",
	KeyLeftmeta:                "Leftmeta",
	KeyRightmeta:               "Rightmeta",
	KeyCompose:                 "Compose",
	KeyStop:                    "Stop",
	KeyAgain:                   "Again",
	KeyProps:                   "Props",
	KeyUndo:                    "Undo",
	KeyFront:                   "Front",
	KeyCopy:                    "Copy",
	KeyOpen:                    "Open",
	KeyPaste:                   "Paste",
	KeyFind:                    "Find",
	KeyCut:                     "Cut",
	KeyHelp:                    "Help",
	KeyMenu:                    "Menu",
	KeyCalc:                    "Calc",
	KeySetup:                   "Setup",
	KeySleep:                   "Sleep",
	KeyWakeup:                  "Wakeup",
	KeyFile:                    "File",
	KeySendfile:                "Sendfile",
	KeyDeletefile:              "Deletefile",
	KeyXfer:                    "Xfer",
	KeyProg1:                   "Prog1",
	KeyProg2:                   "Prog2",
	KeyWww:                     "Www",
	KeyMsdos:                   "Msdos",
	KeyCoffee:                  "Coffee",
	KeyRotateDisplay:           "RotateDisplay",
	KeyCyclewindows:            "Cyclewindows",
	KeyMail:                    "Mail",
	KeyBookmarks:               "Bookmarks",
	KeyComputer:                "Computer",
	KeyBack:                    "Back",
	KeyForward:                 "Forward",
	KeyClosecd:                 "Closecd",
	KeyEjectcd:                 "Ejectcd",
	KeyEjectclosecd:            "Ejectclosecd",
	KeyNextsong:                "Nextsong",
	KeyPlaypause:               "Playpause",
	KeyPrevioussong:            "Previoussong",
	KeyStopcd:                  "Stopcd",
	KeyRecord:                  "Record",
	KeyRewind:                  "Rewind",
	KeyPhone:                   "Phone",
	KeyIso:                     "Iso",
	KeyConfig:                  "Config",
	KeyHomepage:                "Homepage",
	KeyRefresh:                 "Refresh",
	KeyExit:                    "Exit",
	KeyMove:                    "Move",
	KeyEdit:                    "Edit",
	KeyScrollup:                "Scrollup",
	KeyScrolldown:              "Scrolldown",
	KeyKpleftparen:             "Kpleftparen",
	KeyKprightparen:            "Kprightparen",
	KeyNew:                     "New",
	KeyRedo:                    "Redo",
	KeyF13:                     "F13",
	KeyF14:                     "F14",
	KeyF15:                     "F15",
	KeyF16:                     "F16",
	KeyF17:                     "F17",
	KeyF18:                     "F18",
	KeyF19:                     "F19",
	KeyF20:                     "F20",
	KeyF21:                     "F21",
	KeyF22:                     "F22",
	KeyF23:                     "F23",
	KeyF24:                     "F24",
	KeyPlaycd:                  "Playcd",
	KeyPausecd:                 "Pausecd",
	KeyProg3:                   "Prog3",
	KeyProg4:                   "Prog4",
	KeyDashboard:               "Dashboard",
	KeySuspend:                 "Suspend",
	KeyClose:                   "Close",
	KeyPlay:                    "Play",
	KeyFastforward:             "Fastforward",
	KeyBassboost:               "Bassboost",
	KeyPrint:                   "Print",
	KeyHp:                      "Hp",
	KeyCamera:                  "Camera",
	KeySound:                   "Sound",
	KeyQuestion:                "Question",
	KeyEmail:                   "Email",
	KeyChat:                    "Chat",
	KeySearch:                  "Search",
	KeyConnect:                 "Connect",
	KeyFinance:                 "Finance",
	KeySport:                   "Sport",
	KeyShop:                    "Shop",
	KeyAlterase:                "Alterase",
	KeyCancel:                  "Cancel",
	KeyBrightnessdown:          "Brightnessdown",
	KeyBrightnessup:            "Brightnessup",
	KeyMedia:                   "Media",
	KeySwitchvideomode:         "Switchvideomode",
	KeyKbdillumtoggle:          "Kbdillumtoggle",
	KeyKbdillumdown:            "Kbdillumdown",
	KeyKbdillumup:              "Kbdillumup",
	KeySend:                    "Send",
	KeyReply:                   "Reply",
	KeyForwardmail:             "Forwardmail",
	KeySave:                    "Save",
	KeyDocuments:               "Documents",
	KeyBattery:                 "Battery",
	KeyBluetooth:               "Bluetooth",
	KeyWlan:                    "Wlan",
	KeyUwb:                     "Uwb",
	KeyUnknown:                 "Unknown",
	KeyVideoNext:               "VideoNext",
	KeyVideoPrev:               "VideoPrev",
	KeyBrightnessCycle:         "BrightnessCycle",
	KeyBrightnessAuto:          "BrightnessAuto",
	KeyDisplayOff:              "DisplayOff",
	KeyWwan:                    "Wwan",
	KeyRfkill:                  "Rfkill",
	KeyMicmute:                 "Micmute",
	KeyOk:                      "Ok",
	KeySelect:                  "Select",
	KeyGoto:                    "Goto",
	KeyClear:                   "Clear",
	KeyPower2:                  "Power2",
	KeyOption:                  "Option",
	KeyInfo:                    "Info",
	KeyTime:                    "Time",
	KeyVendor:                  "Vendor",
	KeyArchive:                 "Archive",
	KeyProgram:                 "Program",
	KeyChannel:                 "Channel",
	KeyFavorites:               "Favorites",
	KeyEpg:                     "Epg",
	KeyPvr:                     "Pvr",
	KeyMhp:                     "Mhp",
	KeyLanguage:                "Language",
	KeyTitle:                   "Title",
	KeySubtitle:                "Subtitle",
	KeyAngle:                   "Angle",
	KeyZoom:                    "Zoom",
	KeyMode:                    "Mode",
	KeyKeyboard:                "Keyboard",
	KeyScreen:                  "Screen",
	KeyPc:                      "Pc",
	KeyTv:                      "Tv",
	KeyTv2:                     "Tv2",
	KeyVcr:                     "Vcr",
	KeyVcr2:                    "Vcr2",
	KeySat:                     "Sat",
	KeySat2:                    "Sat2",
	KeyCd:                      "Cd",
	KeyTape:                    "Tape",
	KeyRadio:                   "Radio",
	KeyTuner:                   "Tuner",
	KeyPlayer:                  "Player",
	KeyText:                    "Text",
	KeyDvd:                     "Dvd",
	KeyAux:                     "Aux",
	KeyMp3:                     "Mp3",
	KeyAudio:                   "Audio",
	KeyVideo:                   "Video",
	KeyDirectory:               "Directory",
	KeyList:                    "List",
	KeyMemo:                    "Memo",
	KeyCalendar:                "Calendar",
	KeyRed:                     "Red",
	KeyGreen:                   "Green",
	KeyYellow:                  "Yellow",
	KeyBlue:                    "Blue",
	KeyChannelup:               "Channelup",
	KeyChanneldown:             "Channeldown",
	KeyFirst:                   "First",
	KeyLast:                    "Last",
	KeyAb:                      "Ab",
	KeyNext:                    "Next",
	KeyRestart:                 "Restart",
	KeySlow:                    "Slow",
	KeyShuffle:                 "Shuffle",
	KeyBreak:                   "Break",
	KeyPrevious:                "Previous",
	KeyDigits:                  "Digits",
	KeyTeen:                    "Teen",
	KeyTwen:                    "Twen",
	KeyVideophone:              "Videophone",
	KeyGames:                   "Games",
	KeyZoomin:                  "Zoomin",
	KeyZoomout:                 "Zoomout",
	KeyZoomreset:               "Zoomreset",
	KeyWordprocessor:           "Wordprocessor",
	KeyEditor:                  "Editor",
	KeySpreadsheet:             "Spreadsheet",
	KeyGraphicseditor:          "Graphicseditor",
	KeyPresentation:            "Presentation",
	KeyDatabase:                "Database",
	KeyNews:                    "News",
	KeyVoicemail:               "Voicemail",
	KeyAddressbook:             "Addressbook",
	KeyMessenger:               "Messenger",
	KeyDisplaytoggle:           "Displaytoggle",
	KeySpellcheck:              "Spellcheck",
	KeyLogoff:                  "Logoff",
	KeyDollar:                  "Dollar",
	KeyEuro:                    "Euro",
	KeyFrameback:               "Frameback",
	KeyFrameforward:            "Frameforward",
	KeyContextMenu:             "ContextMenu",
	KeyMediaRepeat:             "MediaRepeat",
	Key10channelsup:            "10channelsup",
	Key10channelsdown:          "10channelsdown",
	KeyImages:                  "Images",
	KeyDelEol:                  "DelEol",
	KeyDelEos:                  "DelEos",
	KeyInsLine:                 "InsLine",
	KeyDelLine:                 "DelLine",
	KeyFn:                      "Fn",
	KeyFnEsc:                   "FnEsc",
	KeyFnF1:                    "FnF1",
	KeyFnF2:                    "FnF2",
	KeyFnF3:                    "FnF3",
	KeyFnF4:                    "FnF4",
	KeyFnF5:                    "FnF5",
	KeyFnF6:                    "FnF6",
	KeyFnF7:                    "FnF7",
	KeyFnF8:                    "FnF8",
	KeyFnF9:                    "FnF9",
	KeyFnF10:                   "FnF10",
	KeyFnF11:                   "FnF11",
	KeyFnF12:                   "FnF12",
	KeyFn1:                     "Fn1",
	KeyFn2:                     "Fn2",
	KeyFnD:                     "FnD",
	KeyFnE:                     "FnE",
	KeyFnF:                     "FnF",
	KeyFnS:                     "FnS",
	KeyFnB:                     "FnB",
	KeyBrlDot1:                 "BrlDot1",
	KeyBrlDot2:                 "BrlDot2",
	KeyBrlDot3:                 "BrlDot3",
	KeyBrlDot4:                 "BrlDot4",
	KeyBrlDot5:                 "BrlDot5",
	KeyBrlDot6:                 "BrlDot6",
	KeyBrlDot7:                 "BrlDot7",
	KeyBrlDot8:                 "BrlDot8",
	KeyBrlDot9:                 "BrlDot9",
	KeyBrlDot10:                "BrlDot10",
	KeyNumeric0:                "Numeric0",
	KeyNumeric1:                "Numeric1",
	KeyNumeric2:                "Numeric2",
	KeyNumeric3:                "Numeric3",
	KeyNumeric4:                "Numeric4",
	KeyNumeric5:                "Numeric5",
	KeyNumeric6:                "Numeric6",
	KeyNumeric7:                "Numeric7",
	KeyNumeric8:                "Numeric8",
	KeyNumeric9:                "Numeric9",
	KeyNumericStar:             "NumericStar",
	KeyNumericPound:            "NumericPound",
	KeyNumericA:                "NumericA",
	KeyNumericB:                "NumericB",
	KeyNumericC:                "NumericC",
	KeyNumericD:                "NumericD",
	KeyCameraFocus:             "CameraFocus",
	KeyWpsButton:               "WpsButton",
	KeyTouchpadToggle:          "TouchpadToggle",
	KeyTouchpadOn:              "TouchpadOn",
	KeyTouchpadOff:             "TouchpadOff",
	KeyCameraZoomin:            "CameraZoomin",
	KeyCameraZoomout:           "CameraZoomout",
	KeyCameraUp:                "CameraUp",
	KeyCameraDown:              "CameraDown",
	KeyCameraLeft:              "CameraLeft",
	KeyCameraRight:             "CameraRight",
	KeyAttendantOn:             "AttendantOn",
	KeyAttendantOff:            "AttendantOff",
	KeyAttendantToggle:         "AttendantToggle",
	KeyLightsToggle:            "LightsToggle",
	BtnDpadUp:                  "BtnDpadUp",
	BtnDpadDown:                "BtnDpadDown",
	BtnDpadLeft:                "BtnDpadLeft",
	BtnDpadRight:               "BtnDpadRight",
	KeyAlsToggle:               "AlsToggle",
	KeyButtonconfig:            "Buttonconfig",
	KeyTaskmanager:             "Taskmanager",
	KeyJournal:                 "Journal",
	KeyControlpanel:            "Controlpanel",
	KeyAppselect:               "Appselect",
	KeyScreensaver:             "Screensaver",
	KeyVoicecommand:            "Voicecommand",
	KeyBrightnessMin:           "BrightnessMin",
	KeyBrightnessMax:           "BrightnessMax",
	KeyKbdinputassistPrev:      "KbdinputassistPrev",
	KeyKbdinputassistNext:      "KbdinputassistNext",
	KeyKbdinputassistPrevgroup: "KbdinputassistPrevgroup",
	KeyKbdinputassistNextgroup: "KbdinputassistNextgroup",
	KeyKbdinputassistAccept:    "KbdinputassistAccept",
	KeyKbdinputassistCancel:    "KbdinputassistCancel",
	KeyRightUp:                 "RightUp",
	KeyRightDown:               "RightDown",
	KeyLeftUp:                  "LeftUp",
	KeyLeftDown:                "LeftDown",
	KeyRootMenu:                "RootMenu",
	KeyMediaTopMenu:            "MediaTopMenu",
	KeyNumeric11:               "Numeric11",
	KeyNumeric12:               "Numeric12",
	KeyAudioDesc:               "AudioDesc",
	Key3dMode:                  "3dMode",
	KeyNextFavorite:            "NextFavorite",
	KeyStopRecord:              "StopRecord",
	KeyPauseRecord:             "PauseRecord",
	KeyVod:                     "Vod",
	KeyUnmute:                  "Unmute",
	KeyFastreverse:             "Fastreverse",
	KeySlowreverse:             "Slowreverse",
}
