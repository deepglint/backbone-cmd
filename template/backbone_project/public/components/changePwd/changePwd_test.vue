<script>

module.exports={
	el:"#show",
	data:{
	target:"changePwd"
	},
	ready:function(){
	console.log("The changePwd test main is loaded")
	},
	components:{
	'changePwdvue':require('./changePwd.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="changePwdvue" />
</template>