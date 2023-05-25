function changeview(num) {
    im = document.getElementById("img")
    a = document.getElementById("1")
    if (num == 1) {
        im.setAttribute("src", "1.png")
        a.innerText = "东决：绿军大胜热火总分1-3 塔图姆33+11+7破纪录巴特勒29分"
        a.setAttribute("href", "https://fans.sports.qq.com/post.htm?id=1766742676465516666&mid=69#1_allWithElite")
    } else if (num == 2) {
        im.setAttribute("src", "2.png")
        a.innerText = "辽宁4-0横扫浙江卫冕夺第三冠：赵继伟17+11 郭艾伦24+5"
        a.setAttribute("href", "https://fans.sports.qq.com/post.htm?id=1765967798106849402&mid=63#1_allWithElite")
    } else if (num == 3) {
        im.setAttribute("src", "3.png")
        a.innerText = "美售台战机延后交付真实原因曝光 专家：把台军作为军火生意提款机"
        a.setAttribute("href", "https://new.qq.com/rain/a/20230524A07D0900")
    }
}