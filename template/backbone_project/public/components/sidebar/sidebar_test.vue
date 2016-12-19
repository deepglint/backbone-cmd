<script>

module.exports={
	el:"#show",
	data:{
	target:"sidebar"
	},
	ready:function(){
	console.log("The sidebar test main is loaded")
	},
	components:{
	'sidebarvue':require('./sidebar.vue')
	}
}
</script>

<template>
<div><h1>Test--{{target}} is below</h1></div>
<component is="sidebarvue" />
</template>