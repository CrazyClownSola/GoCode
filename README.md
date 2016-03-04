# GoCode
> 大家好，我又来作死了，这段时间LZ一直在致力于转型的道路，在不断探索这各种好玩的有趣的东西，这一次我找到了一种新的开发语言**GO**，以免各位看官觉得我在忽悠各位，下面贴出[百科](https://zh.wikipedia.org/wiki/Go)，还附上GO的[官方地址](https://golang.org/)，当然具体这是一个怎样的语言，这些还请各位看官自行辨别。
> 补充说明一句，由于语言本身过于强大，有太多东西个人还无法理解和参透，这一次写的代码也只是对于Go的一种小尝试，有这样一篇文章[关于Go的5个阶段](http://www.techug.com/the-5-stages-of-learning-go-with-examples)，LZ还是依旧停留在菜鸟的这个阶段，大家看看就好，看看就好，不要扔东西！

## Sublime Text
> 工欲善其事必先利其器，在未接触Go之前，我并不知道Sublime是什么东西，然后在寻找IDE的过程中，我找到了他，并且喜欢上了这款编译工具，附上[官网地址](https://www.sublimetext.com/)，建议各位看官去尝试下这个工具。

## Go
> 有了IDE，就可以开始我们的代码编写过程，Go是个很适合处理高并发的请求的语言，有点类似[NodeJs](https://nodejs.org/en/)的定位，个人觉得这十分适合去作为一种中间件来实现webService，并且代码实现起来也非常的精简。
> 补充一嗓子，这边所展示的代码，上传的文件当中并没有，可以当作是一个教程。

```
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world!")
	}) // 绑定方法名和路径

	err := http.ListenAndServe("localhost:4000", nil)
	
	if err != nil {
		fmt.Println(err)
	}
}

```


## Mongodb





