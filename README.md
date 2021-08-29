# go-output-parser

One day I found a really big object printed in some logs using the fmt pkg.
It was such a waste- we couldn't use it at all.

Until I made this.

Limitations:
- converts all values to string (because there's no way to determine what they actually were just from the input)
- may not work for some cases (contributions are welcome)
