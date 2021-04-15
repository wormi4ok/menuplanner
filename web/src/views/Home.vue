<template>
  <Week :menu="week.menu"/>
</template>

<script>
// @ is an alias to /src
import Week from '@/components/Week.vue';
import axios from 'axios';

export default {
  name: 'Home',
  components: {
    Week,
  },
  data() {
    return {
      week: {},
    };
  },
  created() {
    axios
      .get(process.env.VUE_APP_API_ADDRESS)
      .then((response) => {
        this.week = response.data;
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>
