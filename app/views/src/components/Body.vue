<template>
<!-- at first, do 'load()' function -->
<div class="small">
  <h1>Body.vue</h1>
  <div>
    <v-text-field id="input-room-ID" label="Room ID" placeholder="Put Room ID Here">
    </v-text-field>
    <v-btn dark rounded>submit</v-btn>
      <!-- <button class="button" v-on:click="onGet">get</button> -->
  </div>
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
  data(){
    // data should be write down below
    return{
    }
  },
  mounted(){
  },
  methods:{
  }
}
</script>

<style>
  .small {
    max-width: 50%;
    margin:  150px auto;
  }
</style>