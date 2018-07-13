There appears to be a bug in golang.org/x/image/tiff when creating compressed tiffs.

```bash
go build
./tiff_bug
```

Open the resulting `out-compressed.tiff` and `out-noncompressed.tiff` in Mac OS X Preview. The ~80MB uncompressed works fine, the [compressed](/out-compressed.tiff) gives an error:

> **The file “out-compressed.tiff” could not be opened.** It may be damaged or use a file format that Preview doesn’t recognize.
