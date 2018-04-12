go-image2ascii
==============

Provides functions to convert image to ascii art.

Usage
-----
[sample](https://github.com/buchy/go-image2ascii/blob/master/_example/sample.go)

convert image to ascii
```
image2ascii.Convert(img)
```

change mapping between luminance values and ascii
```
image2ascii.SetConf([]image2ascii.LumAsciiMap{
    {0.0, 'a'},
    {0.1, 'b'},
    {0.2, 'c'},
    {0.3, 'd'},
    {0.4, 'e'},
    {0.5, 'f'},
    {0.6, 'g'},
    {0.7, 'h'},
    {0.8, 'i'},
    {0.9, 'j'},
    {1.0, 'k'},
})
```

Author
------
buchy