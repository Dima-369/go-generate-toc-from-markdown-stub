package main

import (
	"fmt"
	"os"
	"strings"
)

const testData = `
# ha
# hb
{
<ul>
	<li>ha</li>
	<li>hb</li>
</ul>

---

# ha
## hb
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
	</ul>
</ul>

---

# ha
## hb
# hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
	</ul>
	<li>hc</li>
</ul>

---

## ha
### hb
## hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
	</ul>
	<li>hc</li>
</ul>

---

# ha
## hb
### hc
# hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
		<ul>
			<li>hc</li>
		</ul>
	</ul>
	<li>hc</li>
</ul>

---

# ha
## hb
## hc
# hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
		<li>hc</li>
	</ul>
	<li>hc</li>
</ul>

---

# ha
## hb
### hc
# hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
		<ul>
			<li>hc</li>
		</ul>
	</ul>
	<li>hc</li>
</ul>

---

# ha
## hb
#### hc
# hc
{
<ul>
	<li>ha</li>
	<ul>
		<li>hb</li>
		<ul>
			<li>hc</li>
		</ul>
	</ul>
	<li>hc</li>
</ul>
`

func generateTableOfContentsHtmlTags(in string) string {
	// 0 is root
	lastHeadingLevel := 0
	currentUlLevel := 0
	output := "<ul>\n"

	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			continue
		}

		level := strings.Count(line, "#")
		text := strings.Split(line, " ")[1]

		// it's not important how far the diff is, just the direction if larger, smaller or equal
		levelDiff := level - lastHeadingLevel

		if levelDiff == 0 || lastHeadingLevel == 0 {
			output += strings.Repeat(" ", currentUlLevel*4) + "<li>" + text + "</li>\n"
		} else if levelDiff > 0 {
			output += strings.Repeat(" ", currentUlLevel*4) + "<ul>\n  <li>" + text + "</li>\n"
			currentUlLevel += 1
		} else {
			// levelDiff < 0
			for currentUlLevel >= 2 {
				output += "</ul>\n"
				currentUlLevel -= 1
			}
			currentUlLevel = 0
			output += strings.Repeat(" ", currentUlLevel*4) + "</ul>\n  <li>" + text + "</li>\n"
		}

		lastHeadingLevel = level
	}

	for currentUlLevel >= 1 {
		output += "</ul>\n"
		currentUlLevel -= 1
	}

	output += "</ul>"
	return output
}

func verify(in string, must string) {
	wi := withoutWhitespace(in)
	mu := withoutWhitespace(must)
	if wi != mu {
		println(fmt.Sprintf("Not equal!\nIn:\n%v\nMust:\n%v\n", in, must))
		os.Exit(1)
	}
}

// for easier HTML string comparison
func withoutWhitespace(in string) string {
	in = strings.TrimSpace(in)
	in = strings.ReplaceAll(in, " ", "")
	in = strings.ReplaceAll(in, "\t", "")
	return strings.ReplaceAll(in, "\n", "")
}

func main() {
	for i, t := range strings.Split(testData, "---") {
		parts := strings.Split(t, "{")
		generated := generateTableOfContentsHtmlTags(parts[0])
		verify(generated, parts[1])
		println(fmt.Sprintf("Passed test #%v\n", i))
	}
}
