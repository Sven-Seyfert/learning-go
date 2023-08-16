package hotkey

import (
	"08-structured-types/logger"

	"github.com/eiannone/keyboard"
)

func GetPressedKey() rune {
	err := keyboard.Open()
	if err != nil {
		logger.LogFatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		logger.LogFatal(err)
	}

	return char
}
