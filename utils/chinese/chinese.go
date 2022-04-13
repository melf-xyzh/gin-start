/**
 * @Time    :2022/3/21 9:10
 * @Author  :ZhangXiaoyu
 */

package chinese

var Chinese = new(chinese)

type chinese struct{}

// Len
/**
 *  @Description: 获取中文字符串长度
 *  @receiver c
 *  @param str
 *  @return int
 */
func (c chinese) Len(str string) int {
	rt := []rune(str)
	return len(rt)
}


// Cut
/**
 *  @Description: 截取中文字符串
 *  @receiver c
 *  @param str
 *  @param start
 *  @param end
 *  @return string
 */
func (c chinese) Cut(str string, start int, end int) string {
	rt := []rune(str)
	return string(rt[start:end])
}
