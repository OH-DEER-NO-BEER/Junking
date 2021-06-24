<template>
<!-- at first, do 'load()' function -->
<div class="small">
  <h1>Body.vue</h1>
  <bar-chart></bar-chart>
  <radar-chart></radar-chart>
</div>
</template>

<script>
// import BarChart from './BarChart';
// import RadarChart from './RadarChart';

export default {
  // components: {
  //   BarChart,
  //   RadarChart,
  // },
  mounted:{
  },
  methods:{
  }
}
</script>

<style>
  .small {
    max-width: 100%;
    height: auto;
    max-width: 50%;
    margin:  0 auto;
  }
</style>