<template>
<div class="small">
    <h1>Body.vue</h1>
    <bar-chart :chart-data="datacollection"></bar-chart>
    <radar-chart :chart-data="datacollection"></radar-chart>
</div>
</template>

<script>
import BarChart from './BarChart';
import RadarChart from './RadarChart';

export default {
  components: {
    BarChart,
    RadarChart,
  },
}
</script>

<style>
  .small {
    max-width: 50%;
    margin:  150px auto;
  }
</style>