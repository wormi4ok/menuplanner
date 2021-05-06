<template>
  <div id="app">
    <PseudoWindow @resize.passive="onResize"/>
    <Navbar/>
    <div class="container is-fluid">
      <router-view/>
    </div>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          Made with ♥️ by <a href="https://petrashov.ru" target="_blank">wormi4ok</a>.
        </p>
      </div>
    </footer>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import ErrorHandler from '@/mixins/ErrorHandler';
import PseudoWindow from 'vue-pseudo-window';
import { mapActions } from 'vuex';

export default {
  name: 'App',
  components: {
    PseudoWindow,
    Navbar,
  },
  data: () => ({
    isLoading: false,
  }),
  mounted() {
    this.isLoading = true;
    this.fetchCurrentWeek();
    this.fetchRecipes();
    this.fetchCourses();
    this.isLoading = false;
  },
  methods: {
    ...mapActions([
      'fetchCurrentWeek',
      'fetchRecipes',
      'fetchCourses',
    ]),
    onResize() {
      this.$root.isMobile = window.innerWidth <= 768;
    },
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
  margin-top: 60px;
}
</style>
