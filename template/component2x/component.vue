<script>
export default {
	data(){
		return {
			text: "helloworld",
			str: ['1', '2', '3']
		}
	},
	props:{
		text_field: {
			required: false,
            type: String,
            default: "text",
        }
	},
	methods: {
		sendurl(){
			setTimeout(function(){
				console.log("done")
			}, 100)
			this.testalert()
		},
		testalert(){
			console.log(this.text)
		}
	},
	computed: {
	},
	watch: {

	},
	events: {

	}
}
</script>
<template>
<div class="component-name">
	<div>{{text}}</div>
	<div class="row">
		<div class="col" v-for="i in str">
			{{i}}
		</div>
	</div>
	<button v-on:click="sendurl()">ClickMe</button>
</div>
</template>
<style lang="sass" scoped>
button{
	background-color: #000;
	color: #fff;
}
</style>