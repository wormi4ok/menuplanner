<template>
  <b-navbar-dropdown>
    <template #label>
      <strong>{{ displayName }}</strong>
    </template>
    <b-navbar-item @click="onLogout">
      Logout
    </b-navbar-item>
  </b-navbar-dropdown>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'NavbarUserMenu',
  computed: {
    ...mapGetters(['currentUser']),
    displayName() {
      return this.currentUser.name || this.currentUser.email;
    },
  },
  methods: {
    ...mapActions(['logOut']),
    onLogout() {
      this.logOut();
      this.$router.push({ name: 'Login' });
    },
  },
};
</script>

<style scoped>

</style>
