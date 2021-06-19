import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify"; // path to vuetify export

let vm = new Vue({
	el: "#app",
	vuetify,
	data: () => {},
	methods: {},
	mounted: function() {},
	render: (h) => h(App),
});
