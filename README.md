Initially implemented for integration into `writeTOC()` in 
https://github.com/gomarkdown/markdown/blob/master/html/renderer.go, but 
I realized that I could just rewrite my markdown file to keep the 
headers in line, so I did not need this anymore.

The `gomarkdown` renderer generates `<ul><li><ul>` tags when the markdown headers are not correctly
stepped as here:

```
# First

### Third

## Second
```

The initial [Python-Markdown](`https://github.com/Python-Markdown/markdown`) implementation
I used did not exhibit this behavior, so I prematurely coded this before realizing the easy solution
of just rewriting my markdown files.
