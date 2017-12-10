# bufioReaderPool

This is a package that provide you a fast way to get bufio.Reader.

```
brp := bufioReaderPool.New()
reader := brp.Get()
defer brp.Put(reader)
// work with reader
```
