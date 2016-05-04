<script>

module.exports={
	el:"#show",
	data:{
	target:"helloworld"
	},
	ready:function(){
	console.log("The helloworld test main is loaded")
	},
	components:{
	'helloworldvue':require('./helloworld.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="helloworldvue" />
</template>