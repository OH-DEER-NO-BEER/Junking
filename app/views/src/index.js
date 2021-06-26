import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify"; // path to vuetify export

console.log("test");
var vm = new Vue({
	el: "#app",
	vuetify,
	data: {
		ws: null,
		temp: 0,
		roomID: "test room"
	},
	mounted: function() {
		console.log("created");
		this.ws = new WebSocket('ws://' + window.location.host + '/ws');		
		this.ws.onmessage = function(msg){
			console.log(JSON.parse(msg.data));
		}
	},
	methods: {
		sendRoomID: function (){
			console.log(JSON.stringify({roomId: this.roomID}));
			this.ws.send(
				JSON.stringify({
					roomId: this.roomID
				})
			);	
		},
	},
	render: (h) => h(App),
});

vm.ws.onopen = function(){
	vm.sendRoomID();
}