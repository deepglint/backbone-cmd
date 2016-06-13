<script>

module.exports={
	el:"#show",
	data:{
	target:"navbar"
	},
	ready:function(){
	console.log("The navbar test main is loaded")
	},
	components:{
	'navbarvue':require('./navbar.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="navbarvue" />
</template>