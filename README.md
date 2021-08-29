# go-output-parser

One day I found a really big object in some logs. It was printed with `%v` using the fmt pkg.  
It was such a waste because we couldn't use it at all.  
So I made this.


### Limitations:
- converts all values to string (because there's no way to determine what they actually were just from the input)
- may not work for some cases (contributions are welcome)
