$(function () {
    $('#login').click(function () {
        console.log("点击")
        // 获取用户名和密码
        var account = $('#account').val();
        var password = $('#password').val();
        console.log(account)

        // 对密码进行MD5加密
        var passwordMD5 = CryptoJS.MD5(password).toString();
        console.log(passwordMD5)
        alert(passwordMD5)

    })
})

