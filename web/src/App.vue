<template>
  <div id="app">
    <Navbar/>
    <Week :menu="week.menu"/>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Week from '@/components/Week.vue';

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
    this.week = {
      menu: {
        0: {
          recipes: {
            0: {
              id: 1,
              name: 'Ovsyanoblin',
              calories: 300,
            },
            1: {
              id: 2,
              name: 'Pasta Carbonara',
              calories: 500,
            },
            2: {
              id: 3,
              name: 'Abendbrot',
              calories: 700,
            },
          },
        },
        1: {
          recipes: {
            0: {
              id: 1,
              name: 'Ovsyanoblin2',
              calories: 300,
            },
            1: {
              id: 2,
              name: 'Pasta Carbonara2',
              calories: 500,
            },
            2: {
              id: 3,
              name: 'Abendbrot2',
              calories: 700,
            },
          },
        },
        2: {
          recipes: {
            0: {
              id: 1,
              name: 'Ovsyanoblin2',
              calories: 300,
            },
            1: {
              id: 2,
              name: 'Pasta Carbonara2',
              calories: 500,
            },
            2: {
              id: 3,
              name: 'Abendbrot2',
              calories: 700,
            },
          },
        },
      },
    };
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
