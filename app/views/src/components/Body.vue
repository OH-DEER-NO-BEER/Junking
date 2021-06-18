<template>
<div class="small">
    <h1>Body.vue</h1>
    <bar-chart></bar-chart>
    <radar-chart></radar-chart>
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