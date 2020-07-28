package main

import (
	"time"
	"fmt"
)

// 博主类
type Blogger struct {
	Base
	WeiBos []*PostContent
	// 评论内容，一定要区分属于哪个微博的，所以在这使用了map ,map的key表示的是微博的编号
	//  value 表示具体的评论内容
	Comments map[int][]*PostContent
	Fans     []FansInterfacer // 存储粉丝数据，实际上存储的就是粉丝对象（友好的和不友好），让友好的粉丝类实现接口，让不友好粉丝类实现接口
	// 那么友好粉丝类和不友好粉丝类创建出的对象可以赋值给接口变量。
}

// 定义一个函数，完成博主类对象的创建
func NewBlogger(name string) *Blogger {
	// 1: 创建对象
	blg := new(Blogger)
	// 2：初始化成员
	blg.Name = name
	blg.Comments = make(map[int][]*PostContent)
	blg.WeiBos = make([]*PostContent, 0)
	return blg

}

// 发布微博的方法的实现
// 第一个参数：表示微博的内容，第二个参数表示的是微博的类型
func (b *Blogger) PostWeiBo(content string, wbType int) {
	// 1: 创建PostContent对象
	weibo := new(PostContent)
	// 2: 成员进行初始化
	weibo.Id = b.GetId()
	weibo.Content = content
	weibo.Type = wbType
	weibo.CommentTime = time.Now() // 表示当前系统的时间
	weibo.PostMan = b.Name
	weibo.To = "All"
	// 3: 存储微博数据
	b.WeiBos = append(b.WeiBos, weibo)
	b.Notify(weibo.Id)

	//for _, value := range b.WeiBos {
	//	fmt.Println(*value)
	//}

}

// 获取微博的编号
func (b *Blogger) GetId() int {
	// 获取切片中存储的最后一条微博数据的ID值，然后让该值加1，作为新发布的微博的编号。
	if len(b.WeiBos) == 0 {
		return 0
	} else {
		return b.WeiBos[len(b.WeiBos)-1].Id + 1
	}
}

// 博主接口
type BloggerInterfacer interface {
	// 粉丝可以关注博主
	Attach(bFans FansInterfacer)
	// 粉丝可以取消对博主关注的方法
	Detach(bFans FansInterfacer)
	//微博发布后，通知粉丝
	Notify(wbid int)
}

// 当微博更新后，给粉丝发布通知。
func (b *Blogger) Notify(wbid int) {
	// 1: 对切片进行遍历，获取每个粉丝数据
	for _, fan := range b.Fans {
		fan.Update(b, wbid)
	}
	// 2: 发送通知
}

// 粉丝关注博主
func (b *Blogger) Attach(bFans FansInterfacer) {
	b.Fans = append(b.Fans, bFans)
}

// 粉丝取消对博主的关注
func (b *Blogger) Detach(bFans FansInterfacer) {
	// 1: 将要移除的粉丝数据传递过来，然后在切片中查询
	for i := 0; i < len(b.Fans); i++ {
		if b.Fans[i] == bFans {
			b.Fans = append(b.Fans[:i], b.Fans[i+1:]...)
		}
	}
	// 2: 从切片中移除粉丝数据
}

// 微博内容
// 评论内容信息
type PostContent struct {
	Id          int       // 编号
	Content     string    // 内容
	CommentTime time.Time // 时间
	Type        int       // 类型
	PostMan     string    // 发布人
	To          string    // 给谁发布的（博主姓名）
}

// 粉丝
type Fans struct {
	Base
}

// 粉丝操作的接口
type FansInterfacer interface {
	// 接收博主发出的通知
	Update(bloggerI BloggerInterfacer, wbid int)
	// 具体操作的方法，例如：发布评论
	Action(bloggerI BloggerInterfacer, wbid int)
}

// 获取博主新发布的微博
func (b *Blogger) GetWeiBo(wbid int) *PostContent {
	for _, blog := range b.WeiBos {
		if blog.Id == wbid {
			return blog
		}
	}
	return nil
}

// 友好的粉丝
type FriedFans struct {
	Fans
}

func (f *FriedFans) Update(bloggerI BloggerInterfacer, wbid int) {
	fmt.Printf("Hello:%s你所关注的博主发布了一个新的微博", f.Name)
	f.Action(bloggerI, wbid)

}
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
	for _, msg := range b.Comments[wbid]{
		fmt.Println("粉丝名称:",msg.PostMan)
		fmt.Println("评论内容:",msg.Content)
	}

}

// 不友好的粉丝
type BadFans struct {
	Fans
}

func (f *BadFans) Update(bloggerI BloggerInterfacer, wbid int) {
	fmt.Printf("Hello:%s你所关注的博主发布了一个新的微博", f.Name)
}
func (f *BadFans) Action(bloggerI BloggerInterfacer, wbid int) {

}

type Base struct {
	Id   int    // 编号
	Name string // 名字
}

func main() {
	blg := NewBlogger("张三")
	friedFans := new(FriedFans)
	friedFans.Id = 1
	friedFans.Name = "李四"
	blg.Attach(friedFans) // 添加粉丝

	//blg.Detach(friedFans) // 粉丝取消关注
	//for _,value:= range blg.Fans{
	//	fmt.Println(value)
	//}
	blg.PostWeiBo("今天天气很好", 1)

}
