import { Bar } from "vue-chartjs";
// data for win chart
var data = {
	labels: ["Total Games", "Total Win", "Total Lose"],
	datasets: [
		{
			label: "Bar Dataset",
			data: [50, 20, 30],
			backgroundColor: [
				"rgba(255, 99, 132, 0.2)",
				"rgba(54, 162, 235, 0.2)",
				"rgba(255, 206, 86, 0.2)",
			],
			borderColor: [
				"rgba(255, 99, 132, 1)",
				"rgba(54, 162, 235, 1)",
				"rgba(255, 206, 86, 1)",
			],
			borderWidth: 1,
		},
	],
};
// options for win chart
var options = {
	scales: {
		xAxes: [
			{
				scaleLabel: {
					display: true,
					labelString: "Status (Win)",
				},
			},
		],
		yAxes: [
			{
				ticks: {
					beginAtZero: true,
					stepSize: 10,
				},
			},
		],
	},
};

export default {
	extends: Bar,
	name: "bar-chart",
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
