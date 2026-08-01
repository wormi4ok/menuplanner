<template>
  <div id="app">
    <Navbar v-if="isLoggedIn" />
    <div class="container" :class="{ 'is-fluid': !isMobile }">
      <router-view />
    </div>
    <footer class="footer">
      <div class="content has-text-centered">
        <p>
          Made with ♥️ by
          <a href="https://petrashov.com" target="_blank" rel="noopener noreferrer">wormi4ok</a>
          © {{ (new Date).getFullYear() }}, Menuplanner {{ appVersion }}
        </p>
      </div>
    </footer>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import ErrorHandler from '@/mixins/ErrorHandler';
import isMobile from '@/isMobile';
import { mapState } from 'pinia';
import { useUserStore } from '@/stores/user';

export default {
  name: 'App',
  components: {
    Navbar,
  },
  data: () => ({
    appVersion: window.config.MP_VERSION || '',
  }),
  computed: {
    ...mapState(useUserStore, [
      'isLoggedIn',
    ]),
    isMobile: () => isMobile.value,
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
