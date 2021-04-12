<template>
  <div id="app">
    <Navbar/>
    <Week :menu="week.menu"/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Week from '@/components/Week.vue';
import axios from 'axios';

export default {
  name: 'App',
  components: {
    Navbar,
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

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
