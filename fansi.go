package fansi

import (
	"fmt"
	"strconv"
)

const escSeq = "\033[%d;%d;%sm%s\033[0m"
const FG_DEFAULT = 39
const BG_DEFAULT = 49

const(
	BLACK = 30 + iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	LIGHT_GREY
)

const (
	DARK_GREY = 90 + iota
	LIGHT_RED
	LIGHT_GREEN
	LIGHT_YELLOW
	LIGHT_BLUE
	LIGHT_MAGENTA
	LIGHT_CYAN
	WHITE
)

//background colours
const (
	BG_BLACK = 40 + iota
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_MAGENTA
	BG_CYAN
	BG_LIGHT_GREY
)

const (
	BG_DARK_GREY = 100 + iota
	BG_LIGHT_RED
	BG_LIGHT_GREEN
	BG_LIGHT_YELLOW
	BG_LIGHT_BLUE
	BG_LIGHT_MAGENTA
	BG_LIGHT_CYAN
	BG_WHITE
)

func Pants(s string, c int, bg int, attrs ...int) (string, error) {

	fmt.Println(WHITE);

	if !(BLACK <= c && c <= LIGHT_GREY ||
	    DARK_GREY <= c && c <= WHITE) && c != 0 {

		return "", fmt.Errorf("Invalid foreground value: %d\n", c)
	} else if(c == 0) {
		c = FG_DEFAULT
	}

	if !(BG_BLACK <= bg && bg <= BG_LIGHT_GREY || 
	    BG_DARK_GREY <= bg && bg <= BG_WHITE) && bg != 0{

		return "", fmt.Errorf("Invalid background value: %d\n", bg)
	} else if(bg == 0) {
		bg = BG_DEFAULT
	}

	return fmt.Sprintf(escSeq, c, bg, strconv.Itoa(attrs[0]), s), nil
}
