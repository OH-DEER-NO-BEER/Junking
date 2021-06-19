<script>
import axios from 'axios'

export default {
  data() {
    return {
      qu: "",
    }
  },

  methods:{
      getRequest: function() {
      this.axios
        .get("https://api.coindesk.com/v1/bpi/currentprice.json")
        .then(
          response => (
            (this.name = response.data.chartName),
            (this.rate = response.data.bpi.USD.rate)
          )
        );
    }
  },
};
</script>