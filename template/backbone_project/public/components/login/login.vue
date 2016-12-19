<script>
/**
 * 登录
 * author: zhongjiewang
 * date: 2015-12-09
 */
module.exports = {
    replace: false,
    data: {
        logo: 'public/img/logo_welcome.png',
        userhead: 'public/img/icon_user_login.png',
        loading: false,
        username: '',
        password: '',
        msg: ''
    },
    methods: {
        onClick: function(e) {
            var self = this;
            e.preventDefault()
            if (self.username == '' || self.password == '') {
                self.msg = '用户名或密码不能为空.';
                self.loading = false;
                return;
            }
            self.loading = true;
            self.msg = "";
            bkb.ajax({
                url: "login",
                data: $("#login_form").serialize(),
                method: "POST"
            }).done(function(data){
                location.href = "./";
            }).fail(function(xhr, textStatus, data){
                self.msg = data.msg;
            }).always(function(){
                self.loading = false;
            });
        }
    }
}
</script>

<template>
<div class="bkb-login">
    <div class='logo-wrapper'>
        <img class="logo" v-bind="{'src': logo}">
    </div>
    <div class='form'>
        <img class="userhead"  v-bind="{'src': userhead}">
        <div class='login'>
            <validator name="validation1">
                <form novalidate id="login_form" action="login" method="POST" v-on:submit="onClick($event)">
                    <input name="username" :disabled="loading" v-validate:username="{ required: true }" class="input-text" v-bind:class="{error: ($validation1 && $validation1.username && $validation1.username.required && $validation1.username.touched)}" placeholder='用户名' type='text' v-model='username'>
                    <input name="password" :disabled="loading" v-validate:password="{ required: true }" class="input-text" v-bind:class="{error: ($validation1 && $validation1.password && $validation1.password.required && $validation1.password.touched)}" placeholder='密码' type='password' v-model='password'>
                     
                    <p class="err-msg">
                        <span v-show="$validation1.touched && $validation1.invalid || msg">
                            <span>{{msg}}</span>
                            <span v-if="$validation1.username.dirty &&$validation1.username.required">请输入用户名</span>
                            <span v-if=" $validation1.password.dirty && $validation1.password.required">请输入密码</span>
                        </span>
                    </p>

                    <button type="submit" class="login_btn" :disabled="$validation1.invalid || loading">登录<span v-show="loading">中</span></button>
                </form>
            </validator>
        </div>
    </div>
</div>
</template>

<style>
body, html {
    width: 100%;
    height: 100%;
    overflow: auto;
    position: relative;
    min-width: 430px;
    min-height: 460px;
    margin: 0;
    padding: 0;
}
body {
    background: #ffffff;
    font-family: 'Roboto', sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}

#MainBody {
    position: static;
}

.logo-wrapper,
.logo-wrapper .logo {
    width: 420px;
}

.bkb-login {
    position: absolute;
    text-align: center;
    top: 48%;
    left: 50%;
    -webkit-transform: translate3d(-50%, -50%, 0);
            transform: translate3d(-50%, -50%, 0);
}

.bkb-login .form {
    position: relative;
    width: 420px;
    height: 300px;
    margin: 80px auto 0;
    background: #f8f8f8;
    border: 1px solid #d2d2d2;
    border-radius: 10px;
    box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.15);
}
.bkb-login .userhead {
    position: absolute;
    top: -32px;
    left: 50%;
    width: 64px;
    margin-left: -32px;
}
.bkb-login .login {
    margin: 80px auto 0;
    width: 320px;
}
.input-text {
    outline: none;
    display: block;
    width: 320px;
    margin: 20px auto 10px;
    padding: 10px 15px;
    height: 40px;
    line-height: 20px;
    color: #ccc;

    border: 1px solid #d9d9d9;

    font-family: "Roboto";
    font-size: 14px;
    font-weight: 400;
    border-radius: 3px;
    box-sizing: border-box;

    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    -webkit-transition: all 0.3s linear 0s;
    transition: all 0.3s linear 0s;
}
.input-text:focus {
    color: #333333;
    border: 1px solid #33b5e5;
}
.input-text.error {
    color: #333333;
    border-color: #f00;
}
.login_btn {
    outline: 0;
    cursor: pointer;
    width: 220px;
    height: 40px;
    padding: 10px 15px;
    margin-top: 10px;
    margin-bottom: 10px;
    background: #0084ff;
    border: 1px solid #0075aa;
    color: #ffffff;

    font-family: "Roboto";
    font-size: 14px;
    border-radius: 22px;
    box-sizing: border-box;

    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;

    -webkit-transition: all 0.3s linear 0s;
    transition: all 0.3s linear 0s;
}
.login_btn:disabled {
    background: #ccc;
    cursor: default;
    border-color: #ccc;
    color: #999;
}
.err-msg {
    color: #ec0202;
    text-align: left;
    font-size: 12px;
    margin: 0;
    height: 20px;
}
.loading_form {
    position: absolute;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    background: rgba(0,0,0,0.5);
    z-index: 100;
}
/* Loading */
.loading-wait {
    position: absolute;
    left: 50%;
    top: 50%;
    -webkit-transform: translate3d(-50%, -50%, 0);
            transform: translate3d(-50%, -50%, 0);
    border-bottom: 6px solid rgba(0, 0, 0, .1);
    border-left: 6px solid rgba(0, 0, 0, .1);
    border-right: 6px solid rgba(0, 0, 0, .1);
    border-top: 6px solid rgba(0, 0, 0, .4);
    border-radius: 100%;
    height: 50px;
    width: 50px;
    -webkit-animation: rot .6s infinite linear;
            animation: rot .6s infinite linear;
}

@-webkit-keyframes rot {
    from {
        -webkit-transform: rotate(0deg);
                transform: rotate(0deg);
    }

    to {
        -webkit-transform: rotate(359deg);
                transform: rotate(359deg);
    }

}

@keyframes rot {
    from {
        -webkit-transform: rotate(0deg);
                transform: rotate(0deg);
    }

    to {
        -webkit-transform: rotate(359deg);
                transform: rotate(359deg);
    }

}
</style>