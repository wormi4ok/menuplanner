<template>
  <div id="app">
    <Navbar/>
    <router-view/>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import ErrorHandler from '@/mixins/ErrorHandler';
import { mapActions } from 'vuex';

export default {
  name: 'App',
  components: {
    Navbar,
  },
  data: () => ({
    isLoading: false,
  }),
  mounted() {
    this.isLoading = true;
    this.fetchCurrentWeek();
    this.fetchRecipes();
    this.isLoading = false;
  },
  methods: {
    ...mapActions([
      'fetchCurrentWeek',
      'fetchRecipes',
    ]),
  },
  mixins: [
    ErrorHandler,
  ],
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
