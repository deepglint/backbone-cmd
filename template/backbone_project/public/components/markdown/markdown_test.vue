<script>

module.exports={
	el:"#show",
	data:{
	target:"markdown"
	},
	ready:function(){
	console.log("The markdown test main is loaded")
	},
	components:{
	'markdownvue':require('./markdown.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="markdownvue" />
</template>