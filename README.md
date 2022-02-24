# go-output-parser

Programmers depend on logs to find stacktraces and function inputs so they can reproduce bugs from production on their local machines. 
One day when trying to debug an error, I found a map with a lot of values in some logs.
It was printed with `%v` using the fmt pkg and so it was essentially unusable because there's no way to decode it.

This program converts these kinds of print statement outputs to json.


### Limitations:
- converts all values to string (because there's no way to determine what they actually were just from the input)
- may not work for some cases (contributions are welcome)
