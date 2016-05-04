<script type="text/javascript">
	module.exports={
		data:function(){
			return {
				//hello:"world"
			}
		},
		props:[],
		/////////////Life Span/////////////////////
		// created:function(){
		// 	console.log("hello has been created");
		// },
		// beforeCompile:function(){
		// 	console.log("hello beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("hello has been compiled");
		// },
		// ready:function(){
		// 	console.log("hello has been ready");
		// },
		// attached:function(){
		// 	console.log("hello has been attached");
		// },
		// detached:function(){
		// 	console.log("hello has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before hello destoryed");
		// },
		// destoryed:function(){
		// 	console.log("hello has been destoryed");
		// },
		methods:{
			// sayHello:function(){
			// 	console.log("Hello,This is the hello component");
			// }
			sendurl:function(){
				var data={hello:"world"}
				$.ajax({
					type:'POST',
					url:"testhello",
					cache:false,
					data:JSON.stringify(data),
					contentType:"application/json",
					processData: false,
		            success:function(data){
		                console.log(data)
		              	alert(data)
		            },
		            error:function(data){
		                console.log("error");
		                console.log(data);
		               alert("error: "+data)
		            }
				});			
			}
		},
		computed:{
			// foo:function(){
			// 	return "bar"
			// }
		},
		events:{
			//////////Events 
			// 'hook:created':function(){
			// 	console.log('hello has created again');
			// },
			// greeting:function(msg){
			// 	console.log(msg);
			// },
			// hello:'sayHello'
			///////////Can be Triggered by
			//////vm.$emit('hello')
		},
		watch:{
			//////////Data Watcher
			// 'a':function(val,oldval){
			// 	console.log('new: %s, old: %s', val, oldVal);
			// }
		}

	}
</script>
<template>

<!-- <div>This is component hello</div>
 -->
 <button v-on:click="sendurl()">ClickMe</button>
 </template>
<style type="text/css"></style>