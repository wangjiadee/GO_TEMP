/*
 * @Author: Ralph
 * @Type_file: go
 * @Date: 2020-07-28 20:17:10
 * @LastEditTime: 2020-07-28 20:17:36
 * @FilePath: \goele\19-面向对象的微博案例\weibo.go
 * @Effect: 自版
 */

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
	Attach(bFans FansInterfacer)
	// fans取关
	Derach(bFans FansInterfacer)
	// 微博发布通知fans
	Notify(wbid int)
}

// 定义针对粉丝的接口
type FansInterfacer interface {
	// 接收通知
	Update(bloggerI BloggerInterfacer, wbid int)
	// 评论微博
	Action(bloggerI BloggerInterfacer, wbid int)
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
	b.Notify(weibo.Id)
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
func (f *FriedFans) Update(bloggerI BloggerInterfacer, wbid int) {
	fmt.Println("Hello! %s 你所关注关注的博主发布了一条新的微博。", f.Name)
	f.Action(bloggerI, wbid)
}

// 定义粉丝对微博信息评论的方法
func (f *FriedFans) Action(bloggerI BloggerInterfacer, wbid int) {
	// 1: 获取博主发布的新微博
	/*
			以下写法错误,无法获取发布的微博数据，因为这里是创建了一个新的对象。
		     blogger:=new(Blogger)
		     weibo:=blogger.GetWeiBo(wbid)
		     fmt.Println(weibo)
	*/
	blogger, ok := bloggerI.(*Blogger)
	if ok {
		weibo := blogger.GetWeiBo(wbid)
		// 2: 进行评论。
		// 2.1 构建评论的内容
		// 判断微博类型。
		cType := weibo.Type
		message := "" // 评论内容
		switch cType {
		case 1: // 高兴的类型
			message = "非常好啊!!"
		case 2:
			message = "加油"

		}
		postComment := PostContent{0, message, time.Now(), cType, f.Name, blogger.Name}
		// 2.2 发布评论
		blogger.AddComment(postComment, wbid)
		//fmt.Println(weibo)
		// 展示粉丝发布的评论内容
		blogger.ShowComment(wbid)
	}

}

func (f *BadFans) Update(bloggerI BloggerInterfacer, wbid int) {
	fmt.Println("Hello! %s 你所关注关注的博主发布了一条新的微博。", f.Name)
}
func (f *BadFans) Action(bloggerI BloggerInterfacer, wbid int) {

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

//微博发布通知fans的方法：
func (b *Blogger) Notify(wbid int) {
	// 1.对与粉丝的切片进行遍历，获取粉丝的数据
	for _, fan := range b.Fans {
		fan.Update(b, wbid)
	}
}

// 定义粉丝获取微博的方法
func (b *Blogger) GetWeiBo(wbid int) *PostContent {
	for _, blog := range b.WeiBos {
		if blog.Id == wbid {
			return blog
		}
	}
	return nil
}

// 发布评论
func (b *Blogger) AddComment(comment PostContent, wbid int) {
	b.Comments[wbid] = append(b.Comments[wbid], &comment)
}

// 评论展示
func (b *Blogger) ShowComment(wbid int) {
	// 1: 根据传递过来的微博的编号，查询出对应的微博数据并且展示微博内容
	blog := b.GetWeiBo(wbid)
	fmt.Println("博主名称:", b.Name)
	fmt.Println("微博内容:", blog.Content)
	// 2: 展示出微博对应的评论内容
	for _, msg := range b.Comments[wbid] {
		fmt.Println("粉丝名称:", msg.PostMan)
		fmt.Println("评论内容:", msg.Content)
	}

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
	blg.Derach(friedFans) // 删除粉丝
	for _, v := range blg.Fans {
		fmt.Println(v)
	}
	blg.PostWeiBo("didi!小汉堡了喂！", 1)
}
