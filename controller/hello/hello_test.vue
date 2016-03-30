<script>

module.exports={
	el:"#show",
	data:{
	target:"hello"
	},
	ready:function(){
	console.log("The hello test main is loaded")
	},
	components:{
	'hellovue':require('./hello.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="hellovue" />
</template>