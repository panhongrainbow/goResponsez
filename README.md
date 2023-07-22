# goResponsez

Web pages and Markdown have the following differences

## H tag

Markdown and web page titles `can only be set up to size 6`.

<H1> ->title </H1>

<H2> ->title </H2>

<H3> ->title </H3>

<H4> ->title </H4>

<H5> ->title </H5>

<H6> ->title </H6>

<H7> ->nothing </H7>

## P tag

Web pages have P tags, but Markdown doesn’t because Markdown uses ` blank lines` to separate paragraphs.

<p> paragraph marked by p tag </p>

## hr tag

In Markdown, you can use "***" to generate a horizontal line, but you `cannot set its length`.
Even if you switch to `HTML format`, you `still cannot set its position`.

***

<hr width="300PX" align="right">
## Character table

|    Result    |            Description             | Entity Name | Entity Number |
| :----------: | :--------------------------------: | :---------: | :-----------: |
| non-breaking |               space                |   &nbsp;    |    &#160;     |
|      <       |             less than              |    &lt;     |     &#60;     |
|      >       |            greater than            |    &gt;     |     &#62;     |
|      &       |             ampersand              |    &amp;    |     &#38;     |
|      “       |       double quotation mark        |   &quot;    |     &#34;     |
|      ‘       | single quotation mark (apostrophe) |   &apos;    |     &#39;     |
|      ¢       |                cent                |   &cent;    |    &#162;     |
|      £       |               pound                |   &pound;   |    &#163;     |
|      ¥       |                yen                 |    &yen;    |    &#165;     |
|      €       |                euro                |   &euro;    |    &#8364;    |
|      ©       |             copyright              |   &copy;    |    &#169;     |
|      ®       |        registered trademark        |    &reg;    |    &#174;     |

## bold tag

**markdown's bold tag**

<b>html's bold tag</b>

## italic tag

*markdown's italic tag*

<i>html's italic tag</i>

## underline tag

<u>markdown's underline tag</u> (use the <u> tag in HTML to underline text in Markdown)

<u>html's underline tag</u>

## del tag

~~markdown's delete tag~~

<del>html's delete tag</del>

## combination tag

***markdown's combination tag*** (bold and italic)

<b><i>html's combination tag</i></b> (bold and italic)

## preformat tag

Both `preformat tag in html` and `code snippets in markdown` can be used to represent code. However, using `code snippets` is better.

<pre>
    func main() {
    	fmt.Println("hello world !")
    }
</pre>

```go
    func main() {
    	fmt.Println("hello world !")
    }
```

## font tag

markdown doesn’t support font color natively.

<font color=red size=50px> hello world </font>

## img tag

> As long as `one dimension is set`, the `ratio can be maintained` between the length and width.

The image in html

<img src="https://golang.google.cn/images/gophers/ladder.svg" alt="Gopher climbing a ladder">

The image in markdown

![image-20230712212812003](https://golang.google.cn/images/gophers/ladder.svg)

## embed tag

the major difference between the `<iframe>` and the `<embed> `is that

`<iframe>` is use to represent HTML content like nested HTML tags (another web page or HTML diocument) where as

`<embed>` is use to integrate external (mostly, non HTML) applications like flash, svg etc.



<iframe width="704" height="396" src="https://www.youtube.com/embed/nBxnNxrAQHc" title="YT in html" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>



[![YT in markdown](http://img.youtube.com/vi/nBxnNxrAQHc/0.jpg)]

[(https://www.youtube.com/watch?v=nBxnNxrAQHc "video")

## a herf tag

In markdown, using `markdown-formatted links` to ensure that Firefox opens the link correctly.

[Web site link in markdown](https://golang.google.cn)



As below, other web page features cannot be used.

<a herf=https://golang.google.cn>Web site link in html</a>

<a herf=https://golang.google.cn target=_blank>Open website link in new webpage</a>

<a herf=https://golang.google.cn target=_self> Website link in the same page</a>

If herf is **blank ("")**, it opens the current page. If target is not specified, **the default is target=_self**.



As below, you can also use `an image as a hyperlink`.

<a href=https://golang.google.cn target=_self> Website link in the same page <img src=http://img.youtube.com/vi/nBxnNxrAQHc/0.jpg /></a>

## anchor

Below are the anchors for markdown and html, respectively. Press the `CTL` key before using them.

[jump to the front markdown version](##H tag)

<a href="#H tag">jump to the front html version</a>

## anchor in anthor page

It will jump to `the anchor point of achor in anchor.html`

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title></title>title>
    </head>
    <body>
        <a herf="anchor.html#anchor">link</a>
    </body>
</html>
```

## unordered list

style is `square`

<ul style="list-style-type:square;">
    <li>option 1</li>
    <li>option 2</li>
    <li>option 3</li>
</ul>




style is `circle`

<ul style="list-style-type:circle;">
    <li>option 1</li>
    <li>option 2</li>
    <li>option 3</li>
</ul>




style is `disc`

<ul style="list-style-type:disc;">
    <li>option 1</li>
    <li>option 2</li>
    <li>option 3</li>
</ul>




style is `image`

<ul style="list-style-image: url(https://www.google.com/images/branding/googlelogo/1x/googlelogo_light_color_272x92dp.png);">
    <li>option 1</li>
    <li>option 2</li>
    <li>option 3</li>
</ul>

## ordered list

ordered list in markdown

1. option 1
2. option 2
3. option 3

ordered list in html

<ol>
    <li>option 1</li>
    <li>option 2</li>
    <li>option 3</li>
</ol>

## table tag

table in markdown

| Item      | Quantity | Price |
| --------- | -------- | ----- |
| Product A | 5        | $10   |
| Product B | 3        | $15   |

table in html

<table border="1px" width="300px" hight="300px" align="center" background="http://img.youtube.com/vi/nBxnNxrAQHc/0.jpg">
<tr bgcolor=yellow>
<th>Item</th>
<th>Quantity</th>
<th>Price</th>
</tr>
<tr>
<td>Product A</td>
<td align="center">5</td>
<td>$10</td>
</tr>
<tr>
<td>Product B</td>
<td>3</td>
<td>$15</td>
</tr>
</table>
merge cells

<table border="1">
  <tr>
    <td>Name</td> 
    <td>Gender</td>
    <td>Age</td>
  </tr>
  <tr>
    <td>John</td>
    <td>Male</td>  
    <td>18</td>
  </tr>
  <tr>
    <td>Mary</td>
    <td>Female</td>
    <td>20</td>
  </tr>
  <tr>
    <td colspan="3">Total: 2 people</td> 
  </tr>
</table>

## iframe tag

cat << EOF > index.html

```html
html
<html>
<head>
  <title>Iframe Example</title>
</head>

<body>

<h1>My Web Page</h1>

<p>This is the main content.</p>  

<iframe id="myframe" width="400" height="300" src="https://www.runoob.com"></iframe>

<button onclick="switchIframe('https://www.runoob.com')">Switch to Google</button>

<script>
function switchIframe(url) {
  document.getElementById('myframe').src = url; 
}
</script>

</body>
</html>
```

EOF

In this example, we can `change the page shown in the iframe by clicking the button`.

But now many websites `reject connections from iframes`.

This is because they worry about `clickjacking attacks`.

Clickjacking means hackers use iframes to secretly embed websites in their own pages.

Then they trick users to click and do bad things they did not want to do.



