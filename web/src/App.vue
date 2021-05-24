<template>
  <div id="app">
    <PseudoWindow @resize.passive="onResize"/>
    <Navbar v-if="isLoggedIn"/>
    <div class="container" :class="{ 'is-fluid': !this.$root.isMobile }">
      <router-view/>
    </div>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          Made with ♥️ by <a href="https://petrashov.ru" target="_blank">wormi4ok</a>
          © {{ (new Date).getFullYear() }}, Menuplanner {{ appVersion }}
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import ErrorHandler from '@/mixins/ErrorHandler';
import PseudoWindow from 'vue-pseudo-window';
import { mapGetters } from 'vuex';

export default {
  name: 'App',
  components: {
    PseudoWindow,
    Navbar,
  },
  data: () => ({
    appVersion: window.config.MP_VERSION || '',
  }),
  computed: {
    ...mapGetters([
      'isLoggedIn',
    ]),
  },
  methods: {
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
  margin-top: 30px;
}
</style>
