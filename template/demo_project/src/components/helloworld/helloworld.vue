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
		// 	console.log("helloworld has been created");
		// },
		// beforeCompile:function(){
		// 	console.log("helloworld beforeCompiled");
		// },
		// compiled:function(){
		// 	console.log("helloworld has been compiled");
		// },
		// ready:function(){
		// 	console.log("helloworld has been ready");
		// },
		// attached:function(){
		// 	console.log("helloworld has been attached");
		// },
		// detached:function(){
		// 	console.log("helloworld has been detached");
		// },
		// beforeDestory:function(){
		// 	console.log("before helloworld destoryed");
		// },
		// destoryed:function(){
		// 	console.log("helloworld has been destoryed");
		// },
		methods:{
			// sayHello:function(){
			// 	console.log("Hello,This is the helloworld component");
			// }
			sendurl:function(){
				var data={hello:"world"}
				$.ajax({
					type:'POST',
					url:"testhelloworld",
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
			// 	console.log('helloworld has created again');
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

<!-- <div>This is component helloworld</div>
 -->
 <button v-on:click="sendurl()">ClickMe</button>
 </template>
<style type="text/css"></style>