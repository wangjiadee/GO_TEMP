/*======================================================================
author：Ralph
time：2020-07-22
effect：面向对象综合应用案例-微博
======================================================================*/
package main

import (
	"fmt"
	"time"
)

// =========================start 定义类============================
// 定义博主类
type Blogger struct {
	Base
	WeiBos   []*PostContent         //微博内容结构体指针
	Comments map[int][]*PostContent //微博评论内容结构体指针,评论内能区分哪一个内容 map k:微博的编号 v：评论的内容
	Fans     []FansInterfacer       //存储粉丝的数据=存储粉丝的对象=粉丝类实现接口=可以赋值给接口变量 = 存储接口
}

// 针对微博内容的结构体
type PostContent struct {
	Id          int       //发布微博的编号
	Content     string    //内容
	CommentTime time.Time //发布微博的时间
	Type        int       //微博的类型
	PostMan     string    //微博发布人信息
	TO          string    //粉丝发布的人（博主的姓名）
}

// 定义粉丝类
type Fans struct {
	Base
}

// 粉丝子类
// 真心粉
type FriedFans struct {
	Fans
}

// 黑粉
type BadFans struct {
	Fans
}

// 博主粉丝父类
type Base struct {
	Id   int    //微博编号
	Name string //名字
}

// ========================= 定义类 end============================

// =========================start 定义接口=========================
// 定义针对博主类型的接口
type BloggerInterfacer interface {
	// fans关注
	Attach()
	// fans取关
	Derach()
	// 微博发布通知fans
	Notify()
}

// 定义针对粉丝的接口
type FansInterfacer interface {
	// 接收通知
	Update()
	// 评论微博
	Action()
}

// ========================= 定义接口 end============================

// ========================= start 定义方法==========================

// 定义发布微博的方法的实现
func (b *Blogger) PostWeiBo(content string, wbType int) {
	// 1：创建PostContent对象

	weibo := new(PostContent) //new 返回指针类型 追加到WeiBos的切片里面

	// 2：成员初始化
	weibo.Id = b.GetId() //微博编号 每次发布的微博标号不一样
	weibo.Content = content
	weibo.Type = wbType
	weibo.CommentTime = time.Now() //表示当前系统的时间
	weibo.PostMan = b.Name
	weibo.TO = "ALL" // 表示为微博的内容

	// 3：存储微博数据
	b.WeiBos = append(b.WeiBos, weibo)
	for _, v := range b.WeiBos {
		fmt.Println(*v)
	}
}

// 定义处理微博变化的方法
func (b *Blogger) GetId() int {
	//得到切片中存储的最后一条微博数据的ID值，然后+1，作为新的微博发布编号
	// 通过切片的下标开始获取ID值
	if len(b.WeiBos) == 0 {
		return 0
	} else {
		return b.WeiBos[len(b.WeiBos)-1].Id + 1
	}
}

// 定义粉丝关注博主Attach()的方法
func (b *Blogger) Attach(bFans FansInterfacer) {
	b.Fans = append(b.Fans, bFans)
	// fmt.Println(b.Fans)  //默认打印地址
}

// 真心粉和黑粉实现FansInterfacer接口的方法
func (f *FriedFans) Update() {

}
func (f *FriedFans) Action() {

}

func (f *BadFans) Update() {

}
func (f *BadFans) Action() {

}

//定义粉丝取消博主的关注的方法
func (b *Blogger) Derach(bFans FansInterfacer) {
	// 1:将要移除的粉丝的数据传递过来，然后在切片中查找
	for i := 0; i < len(b.Fans); i++ {
		if b.Fans[i] == bFans {
			b.Fans = append(b.Fans[:i], b.Fans[i+1:]...)
		}
	}
	// 2：从切片中移除粉丝的数据
}

// ========================= 定义方法 end============================

// ========================= start 定义普通函数======================
// 定义函数 完成博主类对象的创建
func NewBlogger(name string) *Blogger {
	// 1:创建对象
	blg := new(Blogger) //返回的是指针
	// 2: 初始化成员
	blg.Name = name
	blg.Comments = make(map[int][]*PostContent)
	blg.WeiBos = make([]*PostContent, 0)
	return blg
}

// ========================= 定义普通函数 end========================

// =========================main function============================
func main() {
	blg := NewBlogger("老八秘制小汉堡")
	friedFans := new(FriedFans)
	friedFans.Id = 1
	friedFans.Name = "真心粉丝1"
	blg.Attach(friedFans) // 添加粉丝
	// blg.Derach(friedFans) // 删除粉丝
	for _, v := range blg.Fans {
		fmt.Println(v)
	}
	blg.PostWeiBo("didi!小汉堡了喂！", 1)
}
