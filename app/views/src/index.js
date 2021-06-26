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
		this.ws = new WebSocket('ws://' + window.location.host + '/ws'); //websocketアクセス

		this.ws.onmessage = function(msg){ //サーバーからのwebsocketからのデータ受信時に
			console.log(JSON.parse(msg.data)); //受信したデータを表示
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

vm.ws.onopen = function(){ //websocket開通時にroomIDを送っている。(mock用なので直接onopenの中に入れるのは今だけ)
	vm.sendRoomID();
}
