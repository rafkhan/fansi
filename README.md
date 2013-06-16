fansi
=====

This is a library for colourful terminal output using ANSI escape sequences.

##Bless me with your wisdom, oh colourful one.

It's so simple. All you do is call the only public function:

`fansi.Pants(s string, int fg, int bg, int attrs) string` 

It should be fairly self explanatory. The constants for fg and bg are in the source file.

###If it's not self explanatory

You pass in the string that you want to make colourful, and it returns the string with the
proper escape sequences. For example:

`fansi.Pants("TICKLE ME PINK", fansi.MAGENTA, fansi.BG_WHITE, fansi.SHIT_I_DIDNT_DO_THIS_YET)`



NOTE: I have yet to add the constants for text attributes (bold, underline, etc.) and
the default colours.

