<script>

module.exports={
	el:"#show",
	data:{
	target:"login"
	},
	ready:function(){
	console.log("The login test main is loaded")
	},
	components:{
	'loginvue':require('./login.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="loginvue" />
</template>