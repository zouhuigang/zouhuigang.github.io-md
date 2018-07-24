---
#标题
title: "syntax-highlighting"
#描述
description: ""
#创建日期
date: 2018-07-24
#修改日期
lastmod: 2018-07-24
#草稿
draft: false
#关键字
keywords: []
#标签
tags: [post,post]
#分类
categories: [post,post]
#作者
author: "邹慧刚"
---
More detail: [Syntax Highlighting | Hugo](https://gohugo.io/content-management/syntax-highlighting/)

```js
function helloWorld () {
  alert("Hello, World!")
}
```

<!--more-->

```java
public class HelloWorld {
  public static void main(String[] args) {
    System.out.println("Hello, World!");
  }
}
```

```kotlin
package hello

fun main(args: Array<String>) {
  println("Hello World!")
}
```

```c
#include <stdio.h>

/* Hello */
int main(void){
  printf("Hello, World!");
  return 0;
}
```

```cpp
// 'Hello World!' program

#include <iostream>

int main(){
  std::cout << "Hello World!" << std::endl;
  return 0;
}
```

```cs
using System;
class HelloWorld{
  public static void Main(){
    System.Console.WriteLine("Hello, World!");
  }
}
```

```html
<html>
<body>
  Hello, World!
</body>
</html>
```

```go
package main
import fmt "fmt"

func main()
{
   fmt.Printf("Hello, World!\n");
}
```

```scala
object HelloWorld with Application {
  Console.println("Hello, World!");
}
```

```php
<?php
  echo 'Hello, World!';
?>
```

```python
print("Hello, World!")
```

## no named code block

```
## this is a comment
$ echo this is a command
this is a command

## edit the file
$vi foo.md
+++
date = "2014-09-28"
title = "creating a new theme"
+++

bah and humbug
:wq

## show it
$ cat foo.md
+++
date = "2014-09-28"
title = "creating a new theme"
+++

bah and humbug
$
```


## highlight shortcode

example:

```shortcode
{{</* highlight go "linenos=table,hl_lines=8 15-17,linenostart=199" */>}}
// ... code
{{</* /highlight */>}}
```

result:

{{< highlight go "linenos=table,hl_lines=8 15-17,linenostart=199" >}}
// GetTitleFunc returns a func that can be used to transform a string to
// title case.
//
// The supported styles are
//
// - "Go" (strings.Title)
// - "AP" (see https://www.apstylebook.com/)
// - "Chicago" (see http://www.chicagomanualofstyle.org/home.html)
//
// If an unknown or empty style is provided, AP style is what you get.
func GetTitleFunc(style string) func(s string) string {
  switch strings.ToLower(style) {
  case "go":
    return strings.Title
  case "chicago":
    tc := transform.NewTitleConverter(transform.ChicagoStyle)
    return tc.Title
  default:
    tc := transform.NewTitleConverter(transform.APStyle)
    return tc.Title
  }
}
{{< /highlight >}}
