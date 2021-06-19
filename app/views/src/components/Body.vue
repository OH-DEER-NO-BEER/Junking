<template>
<!-- at first, do 'load()' function -->
<div class="small" onload="load();">
  <h1>Body.vue</h1>
  <div>
    <v-text-field label="Room ID" placeholder="Put Room ID Here">
    </v-text-field>
    <v-btn dark rounded v-on:click="sendRoomID">submit</v-btn>
      <!-- <button class="button" v-on:click="onGet">get</button> -->
  </div>
  <bar-chart></bar-chart>
  <radar-chart></radar-chart>
</div>
</template>

<script>
import axios from 'axios';
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
      checkInRequest: "hoge",
    }
  },
  methods:{
    // when this component is loaded, do below function
    onLoad(){
      // delete former value if NOT null
      if(this.getRoomID() != null){
        this.removeRoomID();
        console.log("deleted");
      }
      // registar 'storageChange' as event listener
      window.addEventListener("storage", storageChange);
      },
    storageChange(event){
      // PROCESS FOR SENDING TO GOLANG
      this.sendRoomID();
    },
    removeRoomID(){
      window.localStorage.removeItem('roomID');
    },
    storeRoomID(){
      var inputRoomID = document.getElementById("input-room-ID");
      window.localStorage.setItem("roomID", inputRoomID.value);
    },
    getRoomID(){
      return window.localStorage.getItem("roomID");
    },
    sendRoomID(){
      // send roomID to Go-server via axios
      axios.post("/roomID", this.makeCheckInJson()).then(response =>{
        this.response = response.data;
      });
    },
    makeCheckInJson(){
      return {
        roomID: this.getRoomID(),
      }
    }
  }
}
</script>

<style>
  .small {
    max-width: 50%;
    margin:  150px auto;
  }
</style>