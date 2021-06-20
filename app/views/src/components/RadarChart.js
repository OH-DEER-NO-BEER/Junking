import { Radar } from "vue-chartjs";
// data for pattern chart
const data = {
	labels: ["rock", "scissors", "paper"],
	datasets: [
		{
			label: "Your Data",
			data: [0.3, 0.4, 0.3],
			fill: true,
			backgroundColor: "rgba(255, 99, 132, 0.2)",
			borderColor: "rgb(255, 99, 132)",
			pointBackgroundColor: "rgb(255, 99, 132)",
			pointBorderColor: "#fff",
			pointHoverBackgroundColor: "#fff",
			pointHoverBorderColor: "rgb(255, 99, 132)",
		},
		{
			label: "Opponent Data",
			data: [0.5, 0.2, 0.3],
			fill: true,
			backgroundColor: "rgba(255, 99, 132, 0.2)",
			borderColor: "rgb(255, 99, 132)",
			pointBackgroundColor: "rgb(255, 99, 132)",
			pointBorderColor: "#fff",
			pointHoverBackgroundColor: "#fff",
			pointHoverBorderColor: "rgb(255, 99, 132)",
		},
	],
};
// option for pattern chart
var options = {
	type: "radar",
	data: data,
	options: {
		elements: {
			line: {
				borderWidth: 3,
			},
		},

		legend: {
			display: true,
			labels: {
				fontSize: 20,
			},
		},
		scale: {
			pointLabels: {
				fontSize: 13,
			},
			ticks: {
				max: 1,
				min: 0,
				stepSize: 0.1,
			},
		},
	},
};

export default {
	extends: Radar,
	name: "radar-chart",
	data() {
		return {
			data: data,
			options: options,
		};
	},
	mounted() {
		this.renderChart(this.data, this.options);
	},
};
