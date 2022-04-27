$(function (){

    var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVaWQiOjEsIkF2YXRhciI6Imh0dHBzOi8vcGljLnJtYi5iZHN0YXRpYy5jb20vYmpoL25ld3MvZjBmMTg3YTgwZGFiYmQ4NDQyNzgwNDk3YTI2YWI0MzUuanBlZyIsIlVzZXJuYW1lIjoiaXZhbnpoYW5nIiwiTmlja25hbWUiOiLlvKDlhqzlhqwiLCJFbWFpbCI6Im4ueXFlYkBxcS5jb20iLCJNb2JpbGUiOiIxOTg5NTExMTE3NCIsImV4cCI6MTY1MDM2MTczMiwiaXNzIjoieGlsaWFuZ3ppIn0.bvHLq-FiJTHLo8KKSgk1uH22IKg6uJROILb4SmXB7qo"
    console.log(getUrl()+"?token="+token)
    // 验证用户
    // checkUser()
    // 初始化websocket连接
    // var ws = new WebSocket(getUrl()+"?token="+token)
    var ws = new WebSocket(getUrl(), [token])
    // 获取连接状态
    console.log("是否连接成功："+ws.readyState)
    if (ws.readyState != 0) {
        alert("websocket连接失败")
        return
    }

    // 监听是否连接成功
    ws.onopen = function (){
        console.log('ws连接状态：'+ws.readyState)
        // 连接成功向服务器发送数据
        message = {
            "msg_id":getMessageId(),
            "sender":1,
            "recipient":1,
            "content": "connect"
        }
        ws.send(JSON.stringify(message))
    }
    // 接收服务器发送的数据
    ws.onmessage = function (data){

        console.log("接收到服务端消息：")
        console.log(data)
        // 解析接收到的数据并绑定显示
        BindMessage(getNowTime(), data.data)
    }
    // 监听连接关闭事件
    ws.onclose = function (){
        console.log('ws连接状态：'+ws.readyState)
    }
    // 监听并处理error事件
    ws.onerror = function (error){
        console.log(error)
    }
    // 处理客户端发送消息事件
    $("#button").click(function (){
        msg = $("#text-box").text().valueOf()
        if (msg == "") {
            alert("消息不能为空！")
            return
        }
        message = {
            "msg_id":getMessageId(),
            "sender":1,
            "recipient":1,
            "content": msg
        }
        ws.send(JSON.stringify(message))
        $("#text-box").text("")
        // console.log("message:"+JSON.stringify(message))
    })
});

// 生成消息id
function getMessageId(){
    rand = parseInt(Math.random()*(10000+1),10);
    timestamp = (new Date()).valueOf();
    return String(timestamp) + String(rand);
}

// 获取websocket地址
function getUrl(){
    var host = "ws://127.0.0.1:9001/ws"
    return host
}

// 获取当前时间
function getNowTime(){
    var newDate = new Date()
    nowTime = newDate.getFullYear()+"年"+(newDate.getMonth()+1)+"月"+newDate.getDate()+"日 "+newDate.getHours()+":"+newDate.getMinutes()+":"+newDate.getSeconds()
    return nowTime
}

// 解析接收到的信息并显示在页面
function BindMessage(time, data){
    $("#message-list-ul").append("<span>"+time+"</span>")
    $("#message-list-ul").append("<p>"+$.parseJSON(data).sender+"</p>")
    $("#message-list-ul").append("<li>"+$.parseJSON(data).content+"</li>")
}

// 验证用户
function checkUser(){
    //
}