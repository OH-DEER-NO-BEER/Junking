<template>
<div class="small">
    <h1>Body.vue</h1>
    <line-chart :chart-data="datacollection"></line-chart>
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
    max-width: 600px;
    margin:  150px auto;
  }
</style>