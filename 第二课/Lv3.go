package main

type Video struct {
	Author string
	Name string//视频名
	Praise int//点赞数
	Coin int//投币数
	Star int//收藏数
	Share string//转发
}
func (v *Video) Praises(){
	v.Praise++
}
func (v *Video) GiveStar() {
	v.Star++
}
func (v *Video) GiveCoin() {
	v.Coin++
}
func (v *Video) LongPress() {
	v.Praise++
	v.Star++
	v.Coin++
}
func (v *Video) all()   {
	v.Praises()
	v.GiveCoin()
	v.GiveStar()
}
func newVideo(author string, video string)Video{
	return Video{
		Author:author,
		Name:video,
	}
}
func main(){

}