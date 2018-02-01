package pin

import (
	"strconv"
)

// Pin Number
const (
	PIN01 = iota + 1
	PIN02
	PIN03
	PIN04
	PIN05
	PIN06
	PIN07
	PIN08
	PIN09
	PIN10
	PIN11
	PIN12
	PIN13
	PIN14
	PIN15
	PIN16
	PIN17
	PIN18
	PIN19
	PIN20
	PIN21
	PIN22
	PIN23
	PIN24
	PIN25
	PIN26
	PIN27
	PIN28
	PIN29
	PIN30
	PIN31
	PIN32
	PIN33
	PIN34
	PIN35
	PIN36
	PIN37
	PIN38
	PIN39
	PIN40
)

// Gobot raspi gpio mapping
var (
	IDSD   = strconv.Itoa(PIN27) // I²C ID EEPROM
	IDSC   = strconv.Itoa(PIN28) // I²C ID EEPROM
	GPIO02 = strconv.Itoa(PIN03) // SDA1, I²C
	GPIO03 = strconv.Itoa(PIN05) // SCL1, I²C
	GPIO04 = strconv.Itoa(PIN07) // GPIO_GCLK
	GPIO05 = strconv.Itoa(PIN29) //
	GPIO06 = strconv.Itoa(PIN31) //
	GPIO07 = strconv.Itoa(PIN26) // SPI_CE1_N
	GPIO08 = strconv.Itoa(PIN24) // SPI_CE0_N
	GPIO09 = strconv.Itoa(PIN21) // SPI_MISO
	GPIO10 = strconv.Itoa(PIN19) // SPI_MOSI
	GPIO11 = strconv.Itoa(PIN23) // SPI_CLK
	GPIO12 = strconv.Itoa(PIN32)
	GPIO13 = strconv.Itoa(PIN33)
	GPIO14 = strconv.Itoa(PIN08) // TXD0
	GPIO15 = strconv.Itoa(PIN10) // RXD0
	GPIO16 = strconv.Itoa(PIN36)
	GPIO17 = strconv.Itoa(PIN11) // GPIO_GEN0
	GPIO18 = strconv.Itoa(PIN12) // GPIO_GEN1
	GPIO19 = strconv.Itoa(PIN35)
	GPIO20 = strconv.Itoa(PIN38)
	GPIO21 = strconv.Itoa(PIN40)
	GPIO22 = strconv.Itoa(PIN15) // GPIO_GEN3
	GPIO23 = strconv.Itoa(PIN16) // GPIO_GEN4
	GPIO24 = strconv.Itoa(PIN18) // GPIO_GEN5
	GPIO25 = strconv.Itoa(PIN22) // GPIO_GEN6
	GPIO26 = strconv.Itoa(PIN37)
	GPIO27 = strconv.Itoa(PIN13) // GPIO_GEN2
)
