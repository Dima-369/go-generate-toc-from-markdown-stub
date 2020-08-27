Initially implemented for integration into `writeTOC()` in 
https://github.com/gomarkdown/markdown/blob/master/html/renderer.go, but 
I realized that I could just rewrite my markdown file to keep the 
headers in line, so I did not need this anymore.

The `gomarkdown` renderer generates `<ul><li><ul>` tags (which display as messy duplicated bullet points)
when the markdown headers are not correctly
stepped as here:

```
# First

### Third

## Second
```

The initial [Python-Markdown](https://github.com/Python-Markdown/markdown) implementation
I used did not exhibit this behavior, so I prematurely coded this basic implementation for the library 
before realizing the easy solution of just rewriting my markdown files.

## Unfinished implementation

This test case is not parsed correctly, so this should not be used anyway:

```md
#### ha
# hb
{
<ul>
	<li>ha</li>
	<li>hb</li>
</ul>
```

Compare to this Python script:

```python
import markdown

if __name__ == '__main__':
    inp = """
[TOC]
#### ha
# hb
    """
    print(markdown.markdown(inp, extensions=['toc'])
```

Which outputs:

```
<div class="toc">
<ul>
<li><a href="#ha">ha</a></li>
<li><a href="#hb">hb</a></li>
</ul>
</div>
<h4 id="ha">ha</h4>
<h1 id="hb">hb</h1>
```
