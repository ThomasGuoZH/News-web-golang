var im = document.getElementById("img")
var a = document.getElementById("t1")
var valuenode = document.getElementById("gzh love wqr")
var radiolist = document.getElementsByName("select")
function changeview(num) {
    valuenode.setAttribute("value", num)
    if (num === 1) {
        im.setAttribute("src", "image/1.png")
        a.innerText = "东决:绿军大胜热火总分1-3 塔图姆33+11+7破纪录巴特勒29分"
        a.setAttribute("href", "https://fans.sports.qq.com/post.htm?id=1766742676465516666&mid=69#1_allWithElite")
    } else if (num === 2) {
        im.setAttribute("src", "image/2.png")
        a.innerText = "辽宁4-0横扫浙江卫冕夺第三冠：赵继伟17+11 郭艾伦24+5"
        a.setAttribute("href", "https://fans.sports.qq.com/post.htm?id=1765967798106849402&mid=63#1_allWithElite")
    } else if (num === 3) {
        im.setAttribute("src", "image/3.png")
        a.innerText = "美售台战机延后交付真实原因曝光 专家：把台军作为军火生意提款机"
        a.setAttribute("href", "https://new.qq.com/rain/a/20230524A07D0900")
    } else if (num === 4) {
        im.setAttribute("src", "image/4.png")
        a.innerText = "#"
        a.setAttribute("href", "#")
    } else if (num === 5) {
        im.setAttribute("src", "image/5.png")
        a.innerText = "#"
        a.setAttribute("href", "#")
    }
}

function changeprev() {
    value = valuenode.getAttribute("value")
    if (value == 1) {
        changeview(5)
        radiolist[0].checked = false
        radiolist[4].checked = true
    }
    else {
        changeview(value - 1)
        console.log(value - 1)
        radiolist[value - 1].checked = false
        radiolist[value - 2].checked = true
    }
}
function changenext() {
    value = valuenode.getAttribute("value")
    if (value == 5) {
        changeview(1)
        radiolist[4].checked = false
        radiolist[0].checked = true
    }
    else {
        changeview(parseInt(value) + 1)
        radiolist[value - 1].checked = false
        radiolist[value].checked = true
    }
}