<script>
/**
 * 修改密码
 * author: zhongjiewang
 * date: 2015-12-09
 */

// require('bower_components/parsleyjs/dist/parsley.js');
var map = {
    "success": "成功！",
    "info": "",
    "warning": "",
    "danger": "错误！"
};
function msgBuilder(msg, type){
    type = type || "danger"; // success info warning danger
    var text = map[type] || "";
    var tpl = "<div class='alert alert-{type}' role='alert'><strong>{text}</strong> {msg}</div>";
    return tpl.replace('{type}', type)
                .replace('{text}', text)
                .replace('{msg}', msg)
}

function validatMsg(type, msg){
    msg = msg ? msg : "";
    var ico = type !== 'success' ? '<span class="glyphicon glyphicon-remove"></span>' : '<span class="glyphicon glyphicon-ok"></span>'
    return '<span class="text-{type}">{msg}</span>'.replace('{msg}', ico + msg).replace('{type}', type);
}

module.exports = {
    replace: true,
    props: ['main'],
    data: function() {
        return {
            isLoading: false,
            oldname: window.username,
            oldpwd: '',
            newname: '',
            newpwd: '',
            confirmpwd: '',
            msg: "",
            oldpwdValidMsg: "",
            newpwdValidMsg: "",
            confpwdValidMsg: ""
        }
    },
    validators: {
        confirm: function(val){
            return this.newpwd == val
        }
    },
    methods: {
        saveChange: function(e) {
            var self = this;
            e.preventDefault();
            self.msg = "";
            self.oldpwdValidMsg = false;

            // if (self.oldpwd == '') {
            //     self.oldpwdValidMsg = validatMsg("danger", "旧密码不能为空.");
            //     return;
            // }else{
            //     self.oldpwdValidMsg = "";
            // }
            // if (self.newpwd == '') {
            //     self.newpwdValidMsg = validatMsg("danger", "新密码不能为空.");
            //     return;
            // }else{
            //     self.newpwdValidMsg = validatMsg("success");
            // }

            // if (!(/^[\S]{6,16}$/).exec(self.newpwd)){
            //     self.newpwdValidMsg = validatMsg("danger", "密码格式错误.");
            //     return
            // }else{
            //     self.newpwdValidMsg = validatMsg("success");
            // }

            // if (self.confirmpwd != self.newpwd) {
            //     self.confpwdValidMsg = validatMsg("danger", "密码输入不一致.");
            //     return;
            // }else{
            //     self.confpwdValidMsg = validatMsg("success");
            // }

            self.isLoading = true;

            bkb.ajax({
                method: "post",
                url: "changeauth",
                data: $("#changepwd_form").serialize()
            }).done(function(data){
                self.msg = msgBuilder("修改密码成功，将为您自动跳转到登录界面", "success");
                setTimeout(function(){
                    location.href = "login"
                }, 2000)
            }).fail(function(xhr, text, data){
                self.isLoading = false;
                if(data.code == 3){
                    $('[name="oldpwd"]', self.$el).focus()
                    self.oldpwdValidMsg = true;
                    self.oldpwd = "";
                }else{
                    self.msg = msgBuilder("很遗憾，密码修改失败，请重试", "danger");
                }
                
            }).always(function(){
                
            });
        },

    }
}
</script>

<template>
<div class="content m-change-pwd">
    <div class="err_form">{{{msg}}}</div>
    <validator name="validation1" >
    <form class="form-horizontal" novalidate id="changepwd_form" v-on:submit="saveChange($event)" action="/changeauth" method="post">
        <div class="form-group hidden">
            <label class="col-sm-2 control-label">用户名</label>
            <div class="col-sm-6">
                <input type="hidden" name="oldname" v-model='oldname'>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">旧密码</label>
            <div class="col-sm-6">
                <input type="password" class="form-control" v-validate:pwd="{ required: true }" name="oldpwd" placeholder="旧密码" v-model='oldpwd'>
            </div>
            <div class="col-sm-4">
                <span class="text-danger tip" v-if="$validation1.pwd.touched && $validation1.pwd.invalid"><span class="glyphicon glyphicon-remove"></span>旧密码不能为空</span>
                <span class="text-danger tip" v-if="oldpwdValidMsg"><span class="glyphicon glyphicon-remove"></span>密码输入错误</span>
            </div>
        </div>
        <div class="form-group">
            <div class="row">
                <label class="col-sm-2 control-label">新密码</label>
                <div class="col-sm-6">
                    <input type="password" class="form-control" name="newpwd" v-validate:newpwd="{ required: true, pattern: '\/\^\[\\S\]\{6,16\}\$\/'}" maxlength="16" placeholder="新密码" v-model='newpwd'>
                </div>

                <div class="col-sm-4">
                    <span class="text-danger tip" v-if="$validation1.newpwd.touched && $validation1.newpwd.required"><span class="glyphicon glyphicon-remove"></span>新密码不能为空</span>
                    <span class="text-danger tip" v-if="$validation1.newpwd.touched && !$validation1.newpwd.required && $validation1.newpwd.pattern"><span class="glyphicon glyphicon-remove"></span>密码格式错误</span>
                    <span class="text-success tip" v-if="$validation1.newpwd.touched && $validation1.newpwd.valid"><span class="glyphicon glyphicon-ok"></span></span>
                </div>
            </div>
            <div class="row">
                <p class="col-sm-6 col-sm-offset-2 pwd-tip">
                    密码可以为6-16位英文字母、数字或符号，字母区分大小写
                </p>
            </div>
        </div>
        <div class="form-group">
            <label class="col-sm-2 control-label">确认密码</label>
            <div class="col-sm-6">
                <input type="password" class="form-control confirmpwd" name="confirmpwd" v-validate:confirmpwd="{ confirm: true}" maxlength="16" v-model='confirmpwd' placeholder="确认密码">
            </div>

            <div class="col-sm-4">
                <span class="text-danger tip" v-if="$validation1.newpwd.valid && $validation1.confirmpwd.touched && $validation1.confirmpwd.invalid"><span class="glyphicon glyphicon-remove"></span>密码输入不一致</span>
                <span class="text-success tip" v-if="$validation1.newpwd.valid && $validation1.confirmpwd.touched && $validation1.confirmpwd.valid"><span class="glyphicon glyphicon-ok"></span></span>
            </div>
        </div>
        <div class="form-group">
            <div class="col-sm-offset-2 col-sm-10">
                <button type="submit"
                    v-bind="{disabled: $validation1.invalid || isLoading}" class="btn btn-default">保存</button>
            </div>
        </div>
    </form>
    </validator>
</div>
</template>
<style type="text/css">
.m-change-pwd .pwd-tip{
    margin-top: 10px;
    margin-bottom: 0;
    font-size: 13px;
}
.m-change-pwd .form-group .tip {
    margin-top: 7px;
    display: inline-block;
}
</style>